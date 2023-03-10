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

// GetHomeDashboardResponse Get home dashboard response.
//
// swagger:model GetHomeDashboardResponse
type GetHomeDashboardResponse struct {
	GetHomeDashboardResponseAllOf0

	GetHomeDashboardResponseAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetHomeDashboardResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetHomeDashboardResponseAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetHomeDashboardResponseAllOf0 = aO0

	// AO1
	var aO1 GetHomeDashboardResponseAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetHomeDashboardResponseAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetHomeDashboardResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.GetHomeDashboardResponseAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetHomeDashboardResponseAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get home dashboard response
func (m *GetHomeDashboardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetHomeDashboardResponseAllOf0
	if err := m.GetHomeDashboardResponseAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with GetHomeDashboardResponseAllOf1
	if err := m.GetHomeDashboardResponseAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get home dashboard response based on the context it is used
func (m *GetHomeDashboardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetHomeDashboardResponseAllOf0
	if err := m.GetHomeDashboardResponseAllOf0.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with GetHomeDashboardResponseAllOf1
	if err := m.GetHomeDashboardResponseAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetHomeDashboardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetHomeDashboardResponse) UnmarshalBinary(b []byte) error {
	var res GetHomeDashboardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
