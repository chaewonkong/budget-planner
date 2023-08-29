package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("expense", func() {
	Title("Expense Service")
	Description("Inputs and calculations on expenses")
	Server("expense", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:18080")
		})
	})
})

var _ = Service("expense", func() {
	Description("Expense handler")

	Method("AddExpense", func() {
		Payload(func() {
			Field(1, "amount", Float64, "amount spent")
			Field(3, "category", String, "category")
			Field(4, "memo", String, "memo")
			Required("amount", "category")
		})

		Result(func() {
			Field(1, "message", String, "message")
		})

		HTTP(func() {
			POST("/ledger/expenses")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
