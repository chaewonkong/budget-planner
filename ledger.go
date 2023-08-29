package ledgerapi

import (
	"context"
	"log"

	ledger "github.com/chaewonkong/budget-planner/gen/ledger"
)

// ledger service example implementation.
// The example methods log the requests and return zero values.
type ledgersrvc struct {
	logger *log.Logger
}

// NewLedger returns the ledger service implementation.
func NewLedger(logger *log.Logger) ledger.Service {
	return &ledgersrvc{logger}
}

// RecordExpense implements RecordExpense.
func (s *ledgersrvc) RecordExpense(ctx context.Context, p *ledger.RecordExpensePayload) (err error) {
	s.logger.Print("ledger.RecordExpense")
	return
}
