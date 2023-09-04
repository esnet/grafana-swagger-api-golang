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

// DashboardACLUpdateItem dashboard ACL update item
//
// swagger:model DashboardACLUpdateItem
type DashboardACLUpdateItem struct {

	// permission
	Permission PermissionType `json:"permission,omitempty"`

	// role
	// Enum: [None Viewer Editor Admin]
	Role string `json:"role,omitempty"`

	// team Id
	TeamID int64 `json:"teamId,omitempty"`

	// user Id
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this dashboard ACL update item
func (m *DashboardACLUpdateItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardACLUpdateItem) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permission")
		}
		return err
	}

	return nil
}

var dashboardAclUpdateItemTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dashboardAclUpdateItemTypeRolePropEnum = append(dashboardAclUpdateItemTypeRolePropEnum, v)
	}
}

const (

	// DashboardACLUpdateItemRoleNone captures enum value "None"
	DashboardACLUpdateItemRoleNone string = "None"

	// DashboardACLUpdateItemRoleViewer captures enum value "Viewer"
	DashboardACLUpdateItemRoleViewer string = "Viewer"

	// DashboardACLUpdateItemRoleEditor captures enum value "Editor"
	DashboardACLUpdateItemRoleEditor string = "Editor"

	// DashboardACLUpdateItemRoleAdmin captures enum value "Admin"
	DashboardACLUpdateItemRoleAdmin string = "Admin"
)

// prop value enum
func (m *DashboardACLUpdateItem) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dashboardAclUpdateItemTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DashboardACLUpdateItem) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dashboard ACL update item based on the context it is used
func (m *DashboardACLUpdateItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardACLUpdateItem) contextValidatePermission(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Permission.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permission")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardACLUpdateItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardACLUpdateItem) UnmarshalBinary(b []byte) error {
	var res DashboardACLUpdateItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
