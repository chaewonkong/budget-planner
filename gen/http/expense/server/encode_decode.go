// Code generated by goa v3.12.4, DO NOT EDIT.
//
// expense HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/chaewonkong/budget-planner/design/expense

package server

import (
	"context"
	"io"
	"net/http"

	expense "github.com/chaewonkong/budget-planner/gen/expense"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddExpenseResponse returns an encoder for responses returned by the
// expense AddExpense endpoint.
func EncodeAddExpenseResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*expense.AddExpenseResult)
		enc := encoder(ctx, w)
		body := NewAddExpenseResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddExpenseRequest returns a decoder for requests sent to the expense
// AddExpense endpoint.
func DecodeAddExpenseRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddExpenseRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddExpenseRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddExpensePayload(&body)

		return payload, nil
	}
}