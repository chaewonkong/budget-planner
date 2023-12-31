// Code generated by goa v3.12.4, DO NOT EDIT.
//
// ledger client
//
// Command:
// $ goa gen github.com/chaewonkong/budget-planner/design/ledger --output
// gen/ledger

package ledger

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ledger" service client.
type Client struct {
	RecordExpenseEndpoint goa.Endpoint
}

// NewClient initializes a "ledger" service client given the endpoints.
func NewClient(recordExpense goa.Endpoint) *Client {
	return &Client{
		RecordExpenseEndpoint: recordExpense,
	}
}

// RecordExpense calls the "RecordExpense" endpoint of the "ledger" service.
func (c *Client) RecordExpense(ctx context.Context, p *RecordExpensePayload) (err error) {
	_, err = c.RecordExpenseEndpoint(ctx, p)
	return
}
