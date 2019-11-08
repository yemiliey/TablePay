package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

// CloverPaymentType is the cloverpayment type
var CloverPaymentType = dsl.Type("business_owner", func() {
	dsl.Description("Clover payment info.")
	dsl.Attribute("apikey", goa.String, func() {
		dsl.Description("API key for Clover.")
		dsl.Example("3e5c0000-1111-1111-2222-0000c7809931")
	})
	dsl.Attribute("amount", goa.String, func() {
		dsl.Description("Amount to pay in cents.")
		dsl.Example("399")
	})
	dsl.Attribute("exp_month", goa.String, func() {
		dsl.Description("Card's expiry month.")
		dsl.Example("1")
	})
	dsl.Attribute("exp_year", goa.String, func() {
		dsl.Description("Card's expiry year.")
		dsl.Example("2020")
	})
	dsl.Attribute("last4", goa.String, func() {
		dsl.Description("Last 4 digits of the card's number.")
		dsl.Example("1234")
	})
	dsl.Attribute("first6", goa.String, func() {
		dsl.Description("First 6 digits of the card's number.")
		dsl.Example("123456")
	})
	dsl.Attribute("card_encrypted", goa.String, func() {
		dsl.Description("Card's encrypted token.")
	})
})
