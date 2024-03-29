// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "Table Pay": tablepay TestHelpers
//
// Command:
// $ goagen
// --design=github.com/zenreach/platypus/platform/tablepay/api/design
// --out=$(GOPATH)/src/github.com/zenreach/platypus/platform/tablepay/api
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/zenreach/platypus/platform/tablepay/api/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// GetEncriptionKeyByMerchantforeignAndApikeyTablepayBadRequest runs the method GetEncriptionKeyByMerchantforeignAndApikey of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEncriptionKeyByMerchantforeignAndApikeyTablepayBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, apikey string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/%v", mfid, apikey),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["apikey"] = []string{fmt.Sprintf("%v", apikey)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getEncriptionKeyByMerchantforeignAndApikeyCtx, _err := app.NewGetEncriptionKeyByMerchantforeignAndApikeyTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetEncriptionKeyByMerchantforeignAndApikey(getEncriptionKeyByMerchantforeignAndApikeyCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetEncriptionKeyByMerchantforeignAndApikeyTablepayInternalServerError runs the method GetEncriptionKeyByMerchantforeignAndApikey of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEncriptionKeyByMerchantforeignAndApikeyTablepayInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, apikey string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/%v", mfid, apikey),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["apikey"] = []string{fmt.Sprintf("%v", apikey)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getEncriptionKeyByMerchantforeignAndApikeyCtx, _err := app.NewGetEncriptionKeyByMerchantforeignAndApikeyTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetEncriptionKeyByMerchantforeignAndApikey(getEncriptionKeyByMerchantforeignAndApikeyCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetEncriptionKeyByMerchantforeignAndApikeyTablepayNotFound runs the method GetEncriptionKeyByMerchantforeignAndApikey of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEncriptionKeyByMerchantforeignAndApikeyTablepayNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, apikey string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/%v", mfid, apikey),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["apikey"] = []string{fmt.Sprintf("%v", apikey)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getEncriptionKeyByMerchantforeignAndApikeyCtx, _err := app.NewGetEncriptionKeyByMerchantforeignAndApikeyTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetEncriptionKeyByMerchantforeignAndApikey(getEncriptionKeyByMerchantforeignAndApikeyCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GetEncriptionKeyByMerchantforeignAndApikeyTablepayOK runs the method GetEncriptionKeyByMerchantforeignAndApikey of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEncriptionKeyByMerchantforeignAndApikeyTablepayOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, apikey string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/%v", mfid, apikey),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["apikey"] = []string{fmt.Sprintf("%v", apikey)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getEncriptionKeyByMerchantforeignAndApikeyCtx, _err := app.NewGetEncriptionKeyByMerchantforeignAndApikeyTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetEncriptionKeyByMerchantforeignAndApikey(getEncriptionKeyByMerchantforeignAndApikeyCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// GetOrderByLocationAndTableTablepayBadRequest runs the method GetOrderByLocationAndTable of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetOrderByLocationAndTableTablepayBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, lid string, table string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/order/%v", lid, table),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["lid"] = []string{fmt.Sprintf("%v", lid)}
	prms["table"] = []string{fmt.Sprintf("%v", table)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getOrderByLocationAndTableCtx, _err := app.NewGetOrderByLocationAndTableTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetOrderByLocationAndTable(getOrderByLocationAndTableCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetOrderByLocationAndTableTablepayInternalServerError runs the method GetOrderByLocationAndTable of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetOrderByLocationAndTableTablepayInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, lid string, table string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/order/%v", lid, table),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["lid"] = []string{fmt.Sprintf("%v", lid)}
	prms["table"] = []string{fmt.Sprintf("%v", table)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getOrderByLocationAndTableCtx, _err := app.NewGetOrderByLocationAndTableTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetOrderByLocationAndTable(getOrderByLocationAndTableCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetOrderByLocationAndTableTablepayNotFound runs the method GetOrderByLocationAndTable of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetOrderByLocationAndTableTablepayNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, lid string, table string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/order/%v", lid, table),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["lid"] = []string{fmt.Sprintf("%v", lid)}
	prms["table"] = []string{fmt.Sprintf("%v", table)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getOrderByLocationAndTableCtx, _err := app.NewGetOrderByLocationAndTableTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetOrderByLocationAndTable(getOrderByLocationAndTableCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GetOrderByLocationAndTableTablepayOK runs the method GetOrderByLocationAndTable of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetOrderByLocationAndTableTablepayOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, lid string, table string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tablepay/%v/order/%v", lid, table),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["lid"] = []string{fmt.Sprintf("%v", lid)}
	prms["table"] = []string{fmt.Sprintf("%v", table)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	getOrderByLocationAndTableCtx, _err := app.NewGetOrderByLocationAndTableTablepayContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetOrderByLocationAndTable(getOrderByLocationAndTableCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// PostCloverPayByLocationAndTableTablepayBadRequest runs the method PostCloverPayByLocationAndTable of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PostCloverPayByLocationAndTableTablepayBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, orderid string, payload *app.CloverPaymentPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/pay/%v/%v", mfid, orderid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["orderid"] = []string{fmt.Sprintf("%v", orderid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	postCloverPayByLocationAndTableCtx, __err := app.NewPostCloverPayByLocationAndTableTablepayContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	postCloverPayByLocationAndTableCtx.Payload = payload

	// Perform action
	__err = ctrl.PostCloverPayByLocationAndTable(postCloverPayByLocationAndTableCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// PostCloverPayByLocationAndTableTablepayInternalServerError runs the method PostCloverPayByLocationAndTable of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PostCloverPayByLocationAndTableTablepayInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, orderid string, payload *app.CloverPaymentPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/pay/%v/%v", mfid, orderid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["orderid"] = []string{fmt.Sprintf("%v", orderid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	postCloverPayByLocationAndTableCtx, __err := app.NewPostCloverPayByLocationAndTableTablepayContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	postCloverPayByLocationAndTableCtx.Payload = payload

	// Perform action
	__err = ctrl.PostCloverPayByLocationAndTable(postCloverPayByLocationAndTableCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// PostCloverPayByLocationAndTableTablepayNotFound runs the method PostCloverPayByLocationAndTable of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PostCloverPayByLocationAndTableTablepayNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, orderid string, payload *app.CloverPaymentPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/pay/%v/%v", mfid, orderid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["orderid"] = []string{fmt.Sprintf("%v", orderid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	postCloverPayByLocationAndTableCtx, __err := app.NewPostCloverPayByLocationAndTableTablepayContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	postCloverPayByLocationAndTableCtx.Payload = payload

	// Perform action
	__err = ctrl.PostCloverPayByLocationAndTable(postCloverPayByLocationAndTableCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// PostCloverPayByLocationAndTableTablepayOK runs the method PostCloverPayByLocationAndTable of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PostCloverPayByLocationAndTableTablepayOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TablepayController, mfid string, orderid string, payload *app.CloverPaymentPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/pay/%v/%v", mfid, orderid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["mfid"] = []string{fmt.Sprintf("%v", mfid)}
	prms["orderid"] = []string{fmt.Sprintf("%v", orderid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TablepayTest"), rw, req, prms)
	postCloverPayByLocationAndTableCtx, __err := app.NewPostCloverPayByLocationAndTableTablepayContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	postCloverPayByLocationAndTableCtx.Payload = payload

	// Perform action
	__err = ctrl.PostCloverPayByLocationAndTable(postCloverPayByLocationAndTableCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
