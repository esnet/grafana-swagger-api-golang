// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserIDDTO user Id d t o
//
// swagger:model UserIdDTO
type UserIDDTO struct {

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this user Id d t o
func (m *UserIDDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user Id d t o based on context it is used
func (m *UserIDDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserIDDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIDDTO) UnmarshalBinary(b []byte) error {
	var res UserIDDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
