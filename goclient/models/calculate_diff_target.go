// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CalculateDiffTarget calculate diff target
//
// swagger:model CalculateDiffTarget
type CalculateDiffTarget struct {

	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// unsaved dashboard
	UnsavedDashboard JSON `json:"unsavedDashboard,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this calculate diff target
func (m *CalculateDiffTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this calculate diff target based on context it is used
func (m *CalculateDiffTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CalculateDiffTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculateDiffTarget) UnmarshalBinary(b []byte) error {
	var res CalculateDiffTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
