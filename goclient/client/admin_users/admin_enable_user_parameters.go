// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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

// NewAdminEnableUserParams creates a new AdminEnableUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminEnableUserParams() *AdminEnableUserParams {
	return &AdminEnableUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminEnableUserParamsWithTimeout creates a new AdminEnableUserParams object
// with the ability to set a timeout on a request.
func NewAdminEnableUserParamsWithTimeout(timeout time.Duration) *AdminEnableUserParams {
	return &AdminEnableUserParams{
		timeout: timeout,
	}
}

// NewAdminEnableUserParamsWithContext creates a new AdminEnableUserParams object
// with the ability to set a context for a request.
func NewAdminEnableUserParamsWithContext(ctx context.Context) *AdminEnableUserParams {
	return &AdminEnableUserParams{
		Context: ctx,
	}
}

// NewAdminEnableUserParamsWithHTTPClient creates a new AdminEnableUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminEnableUserParamsWithHTTPClient(client *http.Client) *AdminEnableUserParams {
	return &AdminEnableUserParams{
		HTTPClient: client,
	}
}

/*
AdminEnableUserParams contains all the parameters to send to the API endpoint

	for the admin enable user operation.

	Typically these are written to a http.Request.
*/
type AdminEnableUserParams struct {

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin enable user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminEnableUserParams) WithDefaults() *AdminEnableUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin enable user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminEnableUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin enable user params
func (o *AdminEnableUserParams) WithTimeout(timeout time.Duration) *AdminEnableUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin enable user params
func (o *AdminEnableUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin enable user params
func (o *AdminEnableUserParams) WithContext(ctx context.Context) *AdminEnableUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin enable user params
func (o *AdminEnableUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin enable user params
func (o *AdminEnableUserParams) WithHTTPClient(client *http.Client) *AdminEnableUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin enable user params
func (o *AdminEnableUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the admin enable user params
func (o *AdminEnableUserParams) WithUserID(userID int64) *AdminEnableUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin enable user params
func (o *AdminEnableUserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminEnableUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
