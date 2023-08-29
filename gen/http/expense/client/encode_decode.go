// Code generated by goa v3.12.4, DO NOT EDIT.
//
// expense HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/chaewonkong/budget-planner/design/expense

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	expense "github.com/chaewonkong/budget-planner/gen/expense"
	goahttp "goa.design/goa/v3/http"
)

// BuildAddExpenseRequest instantiates a HTTP request object with method and
// path set to call the "expense" service "AddExpense" endpoint
func (c *Client) BuildAddExpenseRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddExpenseExpensePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("expense", "AddExpense", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddExpenseRequest returns an encoder for requests sent to the expense
// AddExpense server.
func EncodeAddExpenseRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*expense.AddExpensePayload)
		if !ok {
			return goahttp.ErrInvalidType("expense", "AddExpense", "*expense.AddExpensePayload", v)
		}
		body := NewAddExpenseRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("expense", "AddExpense", err)
		}
		return nil
	}
}

// DecodeAddExpenseResponse returns a decoder for responses returned by the
// expense AddExpense endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeAddExpenseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AddExpenseResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("expense", "AddExpense", err)
			}
			res := NewAddExpenseResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("expense", "AddExpense", resp.StatusCode, string(body))
		}
	}
}