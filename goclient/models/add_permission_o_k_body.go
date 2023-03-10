// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddPermissionOKBody add permission o k body
//
// swagger:model addPermissionOKBody
type AddPermissionOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// permission Id
	PermissionID int64 `json:"permissionId,omitempty"`
}

// Validate validates this add permission o k body
func (m *AddPermissionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add permission o k body based on context it is used
func (m *AddPermissionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddPermissionOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddPermissionOKBody) UnmarshalBinary(b []byte) error {
	var res AddPermissionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
