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

// PlaylistItem PlaylistItem is the Go representation of a playlist.Item.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
//
// swagger:model PlaylistItem
type PlaylistItem struct {

	// Title is an unused property -- it will be removed in the future
	Title string `json:"title,omitempty"`

	// type
	Type PlaylistItemType `json:"type,omitempty"`

	// Value depends on type and describes the playlist item.
	//
	// dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
	// is not portable as the numerical identifier is non-deterministic between different instances.
	// Will be replaced by dashboard_by_uid in the future. (deprecated)
	// dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
	// dashboards behind the tag will be added to the playlist.
	// dashboard_by_uid: The value is the dashboard UID
	Value string `json:"value,omitempty"`
}

// Validate validates this playlist item
func (m *PlaylistItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlaylistItem) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this playlist item based on the context it is used
func (m *PlaylistItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlaylistItem) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlaylistItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlaylistItem) UnmarshalBinary(b []byte) error {
	var res PlaylistItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}