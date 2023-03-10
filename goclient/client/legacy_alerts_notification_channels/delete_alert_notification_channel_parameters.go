// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteAlertNotificationChannelParams creates a new DeleteAlertNotificationChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAlertNotificationChannelParams() *DeleteAlertNotificationChannelParams {
	return &DeleteAlertNotificationChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlertNotificationChannelParamsWithTimeout creates a new DeleteAlertNotificationChannelParams object
// with the ability to set a timeout on a request.
func NewDeleteAlertNotificationChannelParamsWithTimeout(timeout time.Duration) *DeleteAlertNotificationChannelParams {
	return &DeleteAlertNotificationChannelParams{
		timeout: timeout,
	}
}

// NewDeleteAlertNotificationChannelParamsWithContext creates a new DeleteAlertNotificationChannelParams object
// with the ability to set a context for a request.
func NewDeleteAlertNotificationChannelParamsWithContext(ctx context.Context) *DeleteAlertNotificationChannelParams {
	return &DeleteAlertNotificationChannelParams{
		Context: ctx,
	}
}

// NewDeleteAlertNotificationChannelParamsWithHTTPClient creates a new DeleteAlertNotificationChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAlertNotificationChannelParamsWithHTTPClient(client *http.Client) *DeleteAlertNotificationChannelParams {
	return &DeleteAlertNotificationChannelParams{
		HTTPClient: client,
	}
}

/*
DeleteAlertNotificationChannelParams contains all the parameters to send to the API endpoint

	for the delete alert notification channel operation.

	Typically these are written to a http.Request.
*/
type DeleteAlertNotificationChannelParams struct {

	// NotificationChannelID.
	//
	// Format: int64
	NotificationChannelID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete alert notification channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertNotificationChannelParams) WithDefaults() *DeleteAlertNotificationChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete alert notification channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertNotificationChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) WithTimeout(timeout time.Duration) *DeleteAlertNotificationChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) WithContext(ctx context.Context) *DeleteAlertNotificationChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) WithHTTPClient(client *http.Client) *DeleteAlertNotificationChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationChannelID adds the notificationChannelID to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) WithNotificationChannelID(notificationChannelID int64) *DeleteAlertNotificationChannelParams {
	o.SetNotificationChannelID(notificationChannelID)
	return o
}

// SetNotificationChannelID adds the notificationChannelId to the delete alert notification channel params
func (o *DeleteAlertNotificationChannelParams) SetNotificationChannelID(notificationChannelID int64) {
	o.NotificationChannelID = notificationChannelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlertNotificationChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param notification_channel_id
	if err := r.SetPathParam("notification_channel_id", swag.FormatInt64(o.NotificationChannelID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
