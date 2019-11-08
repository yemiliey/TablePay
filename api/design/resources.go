package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

// TablePayResource is the table pay resource.
var TablePayResource = func() {
	dsl.Origin("*", func() {
		dsl.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		dsl.Headers("Accept", "Content-Type")
		dsl.Expose("Content-Type", "Origin")
		dsl.MaxAge(600)
		dsl.Credentials()
	})
	dsl.Action("get_order_by_location_and_table", func() {
		dsl.Description("Get order for the given table at the given location.")
		dsl.Routing(dsl.GET("/tablepay/:lid/order/:table"))
		dsl.Params(func() {
			dsl.UseTrait("LocationIdPathParam")
			dsl.UseTrait("TableNamePathParam")

			dsl.Required("LocationIdPathParam")
			dsl.Required("TableNamePathParam")
		})
		dsl.Response(goa.OK)
		dsl.Response(goa.NotFound)
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
		dsl.Response(goa.InternalServerError, goa.ErrorMedia)
	})
	dsl.Action("get_encription_key_by_merchantforeign_and_apikey", func() {
		dsl.Description("Get order for the given table at the given location.")
		dsl.Routing(dsl.GET("/tablepay/:mfid/:apikey"))
		dsl.Params(func() {
			dsl.UseTrait("MerchantForeignIdPathParam")
			dsl.UseTrait("ApiKeyPathParam")

			dsl.Required("MerchantForeignIdPathParam")
			dsl.Required("ApiKeyPathParam")
		})
		dsl.Response(goa.OK)
		dsl.Response(goa.NotFound)
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
		dsl.Response(goa.InternalServerError, goa.ErrorMedia)
	})
	dsl.Action("post_clover_pay_by_location_and_table", func() {
		dsl.Description("Call Clover pay endpoint.")
		dsl.Routing(dsl.POST("/pay/:mfid/:orderid"))
		dsl.Params(func() {
			dsl.UseTrait("MerchantForeignIdPathParam")
			dsl.UseTrait("OrderIdPathParam")

			dsl.Required("MerchantForeignIdPathParam")
			dsl.Required("OrderIdPathParam")
		})
		dsl.Payload(CloverPaymentPayload)
		dsl.Response(goa.OK)
		dsl.Response(goa.NotFound)
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
		dsl.Response(goa.InternalServerError, goa.ErrorMedia)
	})
}

var AssetsResource = func() {
	dsl.Origin("*", func() {
		dsl.Methods("GET")
	})
	dsl.Files("/*filepath", "platform/tablepay/frontend/build")
}
