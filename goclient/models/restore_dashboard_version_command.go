// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreDashboardVersionCommand restore dashboard version command
//
// swagger:model RestoreDashboardVersionCommand
type RestoreDashboardVersionCommand struct {

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this restore dashboard version command
func (m *RestoreDashboardVersionCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore dashboard version command based on context it is used
func (m *RestoreDashboardVersionCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreDashboardVersionCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreDashboardVersionCommand) UnmarshalBinary(b []byte) error {
	var res RestoreDashboardVersionCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
