// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemDTO item d t o
//
// swagger:model ItemDTO
type ItemDTO struct {

	// alert Id
	AlertID int64 `json:"alertId,omitempty"`

	// alert name
	AlertName string `json:"alertName,omitempty"`

	// avatar Url
	AvatarURL string `json:"avatarUrl,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// dashboard UID
	DashboardUID string `json:"dashboardUID,omitempty"`

	// data
	Data JSON `json:"data,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// new state
	NewState string `json:"newState,omitempty"`

	// panel Id
	PanelID int64 `json:"panelId,omitempty"`

	// prev state
	PrevState string `json:"prevState,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// text
	Text string `json:"text,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// time end
	TimeEnd int64 `json:"timeEnd,omitempty"`

	// updated
	Updated int64 `json:"updated,omitempty"`

	// user Id
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this item d t o
func (m *ItemDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this item d t o based on context it is used
func (m *ItemDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemDTO) UnmarshalBinary(b []byte) error {
	var res ItemDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
