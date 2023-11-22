// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationTemplateContent notification template content
//
// swagger:model NotificationTemplateContent
type NotificationTemplateContent struct {

	// template
	Template string `json:"template,omitempty"`
}

// Validate validates this notification template content
func (m *NotificationTemplateContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notification template content based on context it is used
func (m *NotificationTemplateContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationTemplateContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationTemplateContent) UnmarshalBinary(b []byte) error {
	var res NotificationTemplateContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}