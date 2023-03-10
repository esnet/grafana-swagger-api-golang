// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetUserRolesCommand set user roles command
//
// swagger:model SetUserRolesCommand
type SetUserRolesCommand struct {

	// global
	Global bool `json:"global,omitempty"`

	// include hidden
	IncludeHidden bool `json:"includeHidden,omitempty"`

	// role uids
	RoleUids []string `json:"roleUids"`
}

// Validate validates this set user roles command
func (m *SetUserRolesCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set user roles command based on context it is used
func (m *SetUserRolesCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetUserRolesCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetUserRolesCommand) UnmarshalBinary(b []byte) error {
	var res SetUserRolesCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
