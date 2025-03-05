// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *CreateQuoteReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Items == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *CreateQuoteReqItemsItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Quantity)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "quantity",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetOrderOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Order.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "order",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetQuoteOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Quote.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "quote",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ListOrdersOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Orders == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Orders {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "orders",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ListQuotesOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Quotes == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Quotes {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "quotes",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *Order) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Total)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "total",
			Error: err,
		})
	}
	if err := func() error {
		if s.Items == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OrderItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.OrderedQuantity)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "ordered_quantity",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.ShippedQuantity)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "shipped_quantity",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.UnitPrice)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "unit_price",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.TotalPrice)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "total_price",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.BackOrderedQuantity)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "back_ordered_quantity",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s OrderStatus) Validate() error {
	switch s {
	case "unspecified":
		return nil
	case "approved":
		return nil
	case "completed":
		return nil
	case "cancelled":
		return nil
	case "pending_approval":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *OrderSummary) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *Quote) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if s.Items == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *QuoteItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.OrderedQuantity)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "ordered_quantity",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.UnitPrice)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "unit_price",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.TotalPrice)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "total_price",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s QuoteStatus) Validate() error {
	switch s {
	case "unspecified":
		return nil
	case "pending_approval":
		return nil
	case "approved":
		return nil
	case "cancelled":
		return nil
	case "expired":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *QuoteSummary) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *SearchProductsOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Aggregations == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Aggregations {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "aggregations",
			Error: err,
		})
	}
	if err := func() error {
		if s.Products == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "products",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s SearchProductsOKAggregationsItem) Validate() error {
	switch s.Type {
	case TermsAggregationSearchProductsOKAggregationsItem:
		if err := s.TermsAggregation.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *SearchProductsReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Filters.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "filters",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s SearchProductsReqFilters) Validate() error {
	var failures []validate.FieldError
	for key, elem := range s {
		if err := func() error {
			if elem == nil {
				return errors.New("nil is invalid value")
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  key,
				Error: err,
			})
		}
	}

	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *TermsAggregation) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Buckets == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "buckets",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
