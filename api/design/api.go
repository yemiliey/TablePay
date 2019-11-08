package design

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

func init() {
	dsl.API("Table Pay", func() {
		dsl.Origin("*", func() {
			dsl.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
			dsl.Headers("Accept", "Content-Type")
			dsl.Expose("Content-Type", "Origin")
			dsl.MaxAge(600)
			dsl.Credentials()
		})
		dsl.Title("Table Pay API")
		dsl.Description("API for table pay")
		dsl.Scheme("http")
		dsl.Host("localhost:8081")
		dsl.Version("1")
		dsl.Trait("LocationIdPathParam", func() {
			dsl.Param("lid", goa.String, func() {
				dsl.Description("ID of location.")
				dsl.Pattern("[0-9a-fA-F]{24}")
				dsl.Example("507f191e810c19729de860ea")
			})
			dsl.Required("lid")
		})
		dsl.Trait("TableNamePathParam", func() {
			dsl.Param("table", goa.String, func() {
				dsl.Description("Table Name.")
			})
			dsl.Required("table")
		})
		dsl.Trait("MerchantForeignIdPathParam", func() {
			dsl.Param("mfid", goa.String, func() {
				dsl.Description("Merchant Foreign Id.")
				dsl.Example("EJ9AWMETAWFC0")
			})
			dsl.Required("table")
		})
		dsl.Trait("ApiKeyPathParam", func() {
			dsl.Param("apikey", goa.String, func() {
				dsl.Description("Api key.")
			})
			dsl.Required("apikey")
		})
		dsl.Trait("OrderIdPathParam", func() {
			dsl.Param("orderid", goa.String, func() {
				dsl.Description("Order Id.")
				dsl.Example("EJ9AWMETAWFC0")
			})
			dsl.Required("orderid")
		})
	})
	dsl.Resource("tablepay", TablePayResource)
	dsl.Resource("assets", AssetsResource)
}
