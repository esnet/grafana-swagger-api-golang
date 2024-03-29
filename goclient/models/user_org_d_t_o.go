// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserOrgDTO user org d t o
//
// swagger:model UserOrgDTO
type UserOrgDTO struct {

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// role
	// Enum: [None Viewer Editor Admin]
	Role string `json:"role,omitempty"`
}

// Validate validates this user org d t o
func (m *UserOrgDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userOrgDTOTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userOrgDTOTypeRolePropEnum = append(userOrgDTOTypeRolePropEnum, v)
	}
}

const (

	// UserOrgDTORoleNone captures enum value "None"
	UserOrgDTORoleNone string = "None"

	// UserOrgDTORoleViewer captures enum value "Viewer"
	UserOrgDTORoleViewer string = "Viewer"

	// UserOrgDTORoleEditor captures enum value "Editor"
	UserOrgDTORoleEditor string = "Editor"

	// UserOrgDTORoleAdmin captures enum value "Admin"
	UserOrgDTORoleAdmin string = "Admin"
)

// prop value enum
func (m *UserOrgDTO) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userOrgDTOTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserOrgDTO) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user org d t o based on context it is used
func (m *UserOrgDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserOrgDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserOrgDTO) UnmarshalBinary(b []byte) error {
	var res UserOrgDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
