package expenseapi

import (
	"context"
	"log"

	expense "github.com/chaewonkong/budget-planner/gen/expense"
)

// expense service example implementation.
// The example methods log the requests and return zero values.
type expensesrvc struct {
	logger *log.Logger
}

// NewExpense returns the expense service implementation.
func NewExpense(logger *log.Logger) expense.Service {
	return &expensesrvc{logger}
}

// AddExpense implements AddExpense.
func (s *expensesrvc) AddExpense(ctx context.Context, p *expense.AddExpensePayload) (res *expense.AddExpenseResult, err error) {
	res = &expense.AddExpenseResult{}
	resultMsg := "ok"
	res.Message = &resultMsg
	s.logger.Print("expense.AddExpense")
	return
}
