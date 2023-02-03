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

// Threshold Threshold a single step on the threshold list
//
// swagger:model Threshold
type Threshold struct {

	// color
	Color string `json:"color,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// value
	Value ConfFloat64 `json:"value,omitempty"`
}

// Validate validates this threshold
func (m *Threshold) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Threshold) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("value")
		}
		return err
	}

	return nil
}

// ContextValidate validate this threshold based on the context it is used
func (m *Threshold) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Threshold) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Value.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Threshold) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Threshold) UnmarshalBinary(b []byte) error {
	var res Threshold
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}