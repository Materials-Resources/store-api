// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
)

func encodeContactUsResponse(response *ContactUsOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeCreateQuoteResponse(response CreateQuoteRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreateQuoteCreated:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateQuoteUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *CreateQuoteUnprocessableEntity:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422)
		span.SetStatus(codes.Error, http.StatusText(422))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetActiveBranchResponse(response GetActiveBranchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetActiveBranchOK:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetActiveBranchUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetInvoiceReportResponse(response GetInvoiceReportRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetInvoiceReportOK:
		w.Header().Set("Content-Type", "application/pdf")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetInvoiceReportUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetInvoiceReportNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetOrderResponse(response GetOrderRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetOrderOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetOrderUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetOrderNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetPackingListReportResponse(response GetPackingListReportRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetPackingListReportOK:
		w.Header().Set("Content-Type", "application/pdf")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetPackingListReportUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetPackingListReportNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetProductResponse(response GetProductRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetProductOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetProductNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetQuoteResponse(response GetQuoteRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetQuoteOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetQuoteUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetQuoteNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetRecentPurchasesResponse(response GetRecentPurchasesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetRecentPurchasesOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetRecentPurchasesUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *GetRecentPurchasesNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListCustomerBranchesResponse(response ListCustomerBranchesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListCustomerBranchesOK:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ListCustomerBranchesUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListInvoicesResponse(response ListInvoicesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListInvoicesOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ListInvoicesUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListOrdersResponse(response ListOrdersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListOrdersOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ListOrdersUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListQuotesResponse(response ListQuotesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListQuotesOK:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ListQuotesUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeSearchProductsResponse(response *SearchProductsOK, w http.ResponseWriter, span trace.Span) error {
	if err := func() error {
		if err := response.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "validate")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSetActiveBranchResponse(response SetActiveBranchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SetActiveBranchOK:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *SetActiveBranchBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *SetActiveBranchForbidden:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeErrorResponse(response *ErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	e := new(jx.Encoder)
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
