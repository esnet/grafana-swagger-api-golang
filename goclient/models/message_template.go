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

// MessageTemplate message template
//
// swagger:model MessageTemplate
type MessageTemplate struct {

	// name
	Name string `json:"name,omitempty"`

	// provenance
	Provenance Provenance `json:"provenance,omitempty"`

	// template
	Template string `json:"template,omitempty"`
}

// Validate validates this message template
func (m *MessageTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvenance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageTemplate) validateProvenance(formats strfmt.Registry) error {
	if swag.IsZero(m.Provenance) { // not required
		return nil
	}

	if err := m.Provenance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

// ContextValidate validate this message template based on the context it is used
func (m *MessageTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageTemplate) contextValidateProvenance(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Provenance.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageTemplate) UnmarshalBinary(b []byte) error {
	var res MessageTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}