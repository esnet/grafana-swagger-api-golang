// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddUserRoleCommand add user role command
//
// swagger:model AddUserRoleCommand
type AddUserRoleCommand struct {

	// global
	Global bool `json:"global,omitempty"`

	// role Uid
	RoleUID string `json:"roleUid,omitempty"`
}

// Validate validates this add user role command
func (m *AddUserRoleCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add user role command based on context it is used
func (m *AddUserRoleCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddUserRoleCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddUserRoleCommand) UnmarshalBinary(b []byte) error {
	var res AddUserRoleCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
