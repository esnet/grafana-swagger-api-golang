// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Sample Sample is a single sample belonging to a metric.
//
// swagger:model Sample
type Sample struct {

	// h
	H *FloatHistogram `json:"H,omitempty"`

	// metric
	Metric Labels `json:"Metric,omitempty"`

	// t
	T int64 `json:"T,omitempty"`

	// v
	V float64 `json:"V,omitempty"`
}

// Validate validates this sample
func (m *Sample) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateH(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sample) validateH(formats strfmt.Registry) error {
	if swag.IsZero(m.H) { // not required
		return nil
	}

	if m.H != nil {
		if err := m.H.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("H")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("H")
			}
			return err
		}
	}

	return nil
}

func (m *Sample) validateMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	if err := m.Metric.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Metric")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Metric")
		}
		return err
	}

	return nil
}

// ContextValidate validate this sample based on the context it is used
func (m *Sample) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateH(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sample) contextValidateH(ctx context.Context, formats strfmt.Registry) error {

	if m.H != nil {
		if err := m.H.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("H")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("H")
			}
			return err
		}
	}

	return nil
}

func (m *Sample) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Metric.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Metric")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Metric")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Sample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sample) UnmarshalBinary(b []byte) error {
	var res Sample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
