package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("ledger", func() {
	Title("Ledger API")
	Description("Record cash inflow and outflow")
	Server("ledger", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:18080")
		})
	})
})
