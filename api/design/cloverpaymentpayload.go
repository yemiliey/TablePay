package design

import (
	dsl "github.com/goadesign/goa/design/apidsl"
)

// CloverPaymentPayload is the cloverpayment payload type
var CloverPaymentPayload = dsl.Type("clover_payment_payload", func() {
	dsl.Reference(CloverPaymentType)
	dsl.Attribute("apikey")
	dsl.Attribute("amount")
	dsl.Attribute("exp_month")
	dsl.Attribute("exp_year")
	dsl.Attribute("last4")
	dsl.Attribute("first6")
	dsl.Attribute("card_encrypted")

	dsl.Required("apikey")
	dsl.Required("amount")
	dsl.Required("exp_month")
	dsl.Required("exp_year")
	dsl.Required("last4")
	dsl.Required("first6")
	dsl.Required("card_encrypted")
})
