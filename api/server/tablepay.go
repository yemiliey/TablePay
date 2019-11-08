package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/goadesign/goa"
)

// TablepayController implements the tablepay resource.
type TablepayController struct {
	*goa.Controller
	posDataAPIClient servicepb.DataAPIClient
}

// NewTablepayController creates a tablepay controller.
func NewTablepayController(service *goa.Service, posDataAPIClient servicepb.DataAPIClient) *TablepayController {
	return &TablepayController{
		Controller:       service.NewController("TablepayController"),
		posDataAPIClient: posDataAPIClient,
	}
}

// GetEncriptionKeyByMerchantforeignAndApikey runs the get_encription_key_by_merchantforeign_and_apikey action.
func (c *TablepayController) GetEncriptionKeyByMerchantforeignAndApikey(ctx *app.GetEncriptionKeyByMerchantforeignAndApikeyTablepayContext) error {

	mfid := ctx.Mfid
	apiKey := ctx.Apikey

	if mfid == "" || apiKey == "" {
		return ctx.BadRequest(errors.New("Missing merchant id or api key"))
	}

	urlStr := fmt.Sprintf("https://apisandbox.dev.clover.com/v2/merchant/%s/pay/key?access_token=%s",
		mfid, apiKey)
	resp, err := http.Get(urlStr)
	if err != nil {
		return ctx.BadRequest(errors.New("Bad response"))
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ctx.BadRequest(errors.New("Problem reading"))
		}
		return ctx.OK(bodyBytes)
	}

	return ctx.BadRequest(errors.New("Bad status"))
}

// GetOrderByLocationAndTable runs the get_order_by_location_and_table action.
func (c *TablepayController) GetOrderByLocationAndTable(ctx *app.GetOrderByLocationAndTableTablepayContext) error {
	// GET endpoint that takes table name submit from frontend
	// look up API key by merchant id from POS Data API
	// get order from Clover (filter by table name which is the title field, sort DESC by modified time, take the top one)
	// return order line items + order id + api key + merchant foreign id

	// get merchant from pos data api
	cloverVendor := pospb.POSVendor_CLOVER_VENDOR
	request := servicepb.GetMerchantsByLocationIDsRequest{
		VendorName:  cloverVendor,
		LocationIds: []string{ctx.Lid},
	}
	posCtx1, cancel1 := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel1()
	merchantResponse, err := c.posDataAPIClient.GetMerchantsByLocationIDs(posCtx1, &request)
	if err != nil {
		ctx.Service.LogError(
			"an error occurred while getting merchant by locationid from POS api.",
			"location", ctx.Lid,
			"request", request,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("error while getting merchant by locationid from POS api")))
	}
	if len(merchantResponse.Merchants) == 0 {
		ctx.Service.LogError(
			"no merchants found from pos data api",
			"location", ctx.Lid,
			"request", request,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("no merchants found from pos data api")))
	}
	merchantID := merchantResponse.Merchants[0].MerchantId
	merchantForeignID := merchantResponse.Merchants[0].ForeignId

	posCtx2, cancel2 := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel2()
	authResponse, err := c.posDataAPIClient.GetAuth(posCtx2, &servicepb.GetAuthRequest{MerchantId: merchantID})
	if err != nil {
		ctx.Service.LogError(
			"an error occurred while getting auth from POS api.",
			"location", ctx.Lid,
			"merchantid", merchantID,
			"request", request,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("error while getting auth from POS api")))
	}
	apiToken := authResponse.ApiToken

	// HACK until we get POS Data API working
	// if ctx.Lid != "5c671cd0719df80001d4656c" {
	// 	return ctx.NotFound()
	// }
	// apiToken := "d8ca4666-6295-ec3f-9e14-d147fd2baa80"
	// merchantForeignID := "F5YXVF6JV7W86"

	orderResponse, err := http.Get(
		fmt.Sprintf(
			"https://apisandbox.dev.clover.com/v3/merchants/%s/orders?expand=lineItems&access_token=%s&orderBy=modifiedTime%%20DESC&filter=title=%s",
			merchantForeignID,
			apiToken,
			url.PathEscape(ctx.Table),
		),
	)
	if err != nil {
		ctx.Service.LogError(
			"error while getting order from clover api.",
			"merchantForeignID", merchantForeignID,
			"table", ctx.Table,
			"apitoken", apiToken,
			"error", err,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("error while unmarshaling clover api response")))
	}
	defer orderResponse.Body.Close()

	if orderResponse.StatusCode == http.StatusNotFound {
		ctx.Service.LogInfo(
			"order not found",
			"merchantForeignID", merchantForeignID,
			"table", ctx.Table,
			"apitoken", apiToken,
			"statuscode", orderResponse.StatusCode,
		)
		return ctx.NotFound()
	}
	if orderResponse.StatusCode != http.StatusOK {
		ctx.Service.LogError(
			"unexpected status code while getting order from clover api",
			"merchantForeignID", merchantForeignID,
			"table", ctx.Table,
			"apitoken", apiToken,
			"statuscode", orderResponse.StatusCode,
			"error", errors.New("unexpected status code"),
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("unexpected status code while getting order from clover api")))
	}

	body, err := ioutil.ReadAll(orderResponse.Body)
	if err != nil {
		ctx.Service.LogError(
			"error while doing ioutil.ReadAll on clover api response body.",
			"body", orderResponse.Body,
			"error", err,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("error while unmarshaling clover api response")))
	}

	type LineItemElement struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	type LineItems struct {
		Elements []LineItemElement `json:"elements"`
	}

	type CloverOrder struct {
		ID        string    `json:"id"`
		LineItems LineItems `json:"lineItems"`
	}

	type CloverOrders struct {
		Elements []CloverOrder `json:"elements,omitempty"`
	}

	var data CloverOrders
	err = json.Unmarshal(body, &data)
	if err != nil {
		ctx.Service.LogError(
			"error unmarshaling clover api response.",
			"body", body,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("error while unmarshaling clover api response")))
	}

	orderID := ""
	lineItems := []LineItemElement{}

	if len(data.Elements) == 0 {
		ctx.Service.LogInfo(
			"no orders in response",
		)
		return ctx.NotFound()
	}

	// take first (latest) order
	orderID = data.Elements[0].ID
	lineItems = data.Elements[0].LineItems.Elements

	type TableEntryResponse struct {
		OrderID    string            `json:"order_id,omitempty"`
		MerchantID string            `json:"merchant_id,omitempty"`
		APIKey     string            `json:"api_key,omitempty"`
		LineItems  []LineItemElement `json:"line_items,omitempty"`
	}

	ter := TableEntryResponse{
		OrderID:    orderID,
		MerchantID: merchantForeignID,
		APIKey:     apiToken,
		LineItems:  lineItems,
	}

	response, err := json.Marshal(ter)
	if err != nil {
		ctx.Service.LogError(
			"error marshalling response",
			"body", orderResponse.Body,
			"error", err,
		)
		return ctx.InternalServerError(goa.ErrInternal(errors.New("error marshalling response")))
	}

	// TablepayController_GetOrderByLocationAndTable: end_implement
	return ctx.OK(response)
}

// PostCloverPayByLocationAndTable runs the post_clover_pay_by_location_and_table action.
func (c *TablepayController) PostCloverPayByLocationAndTable(ctx *app.PostCloverPayByLocationAndTableTablepayContext) error {
	payURL := fmt.Sprintf(
		"https://apisandbox.dev.clover.com/v2/merchant/%s/pay?access_token=%s",
		ctx.Mfid,
		ctx.Payload.Apikey,
	)

	type CloverPayPayload struct {
		OrderID         string `json:"orderId"`
		Currency        string `json:"currency"`
		Amount          string `json:"amount"`
		First6          string `json:"first6"`
		Last4           string `json:"last4"`
		ExpirationMonth string `json:"expMonth"`
		ExpirationYear  string `json:"expYear"`
		CVV             string `json:"cvv"`
		CardEncrypted   string `json:"cardEncrypted"`
		// taxAmount - seems to be optional
		// zip - seems to be optional
	}
	payPayload := CloverPayPayload{
		OrderID:         ctx.Orderid,
		Currency:        "usd",
		Amount:          ctx.Payload.Amount,
		First6:          ctx.Payload.First6,
		Last4:           ctx.Payload.Last4,
		ExpirationMonth: ctx.Payload.ExpMonth,
		ExpirationYear:  ctx.Payload.ExpYear,
		CVV:             "None",
		CardEncrypted:   ctx.Payload.CardEncrypted,
	}

	// create request
	bodyBytes, err := json.Marshal(payPayload)
	if err != nil {
		c.Service.LogError(
			"error marshaling pay payload json",
			"error", err,
		)
		return ctx.InternalServerError(err)
	}
	body := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(http.MethodPost, payURL, body)
	if err != nil {
		c.Service.LogError(
			"error creating pay request",
			"error", err,
		)
		return ctx.InternalServerError(err)
	}

	// send request
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		c.Service.LogError(
			"error making pay request to clover",
			"error", err,
		)
		return ctx.InternalServerError(err)
	}
	defer resp.Body.Close()

	// handle response
	if resp.StatusCode != http.StatusOK {
		c.Service.LogError(
			"got unexpected response from clover when making payment",
			"statuscode", resp.StatusCode,
			"payURL", payURL,
			"payPayload", payPayload,
			"error", err,
		)
		return ctx.InternalServerError(err)
	}

	c.Service.LogInfo(
		"successfully paid",
	)

	return ctx.OK(nil)
}
