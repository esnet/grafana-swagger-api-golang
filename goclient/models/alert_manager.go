// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertManager AlertManager models a configured Alert Manager.
//
// swagger:model AlertManager
type AlertManager struct {

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this alert manager
func (m *AlertManager) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alert manager based on context it is used
func (m *AlertManager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertManager) UnmarshalBinary(b []byte) error {
	var res AlertManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
