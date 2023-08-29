package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("ledger", func() {
	Description("Record cash inflow and outflow")

	Method("RecordExpense", func() {
		Payload(func() {
			Field(1, "userID", Int64, "user id")
			Field(2, "amount", Float64, "amount spent")
			Field(3, "category", String, "category")
			Field(4, "memo", String, "memo")
			Required("amount", "category")
		})

		Result(Empty)

		HTTP(func() {
			POST("/user/{userID}/expenses")
			Body(func() {
				Attribute("amount")
				Attribute("category")
				Attribute("memo")
				Required("amount", "category")
			})
			Response(StatusCreated)
			Response(StatusBadRequest)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
