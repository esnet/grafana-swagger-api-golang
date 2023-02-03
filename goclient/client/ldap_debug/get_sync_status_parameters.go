// Code generated by go-swagger; DO NOT EDIT.

package ldap_debug

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
)

// NewGetSyncStatusParams creates a new GetSyncStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSyncStatusParams() *GetSyncStatusParams {
	return &GetSyncStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncStatusParamsWithTimeout creates a new GetSyncStatusParams object
// with the ability to set a timeout on a request.
func NewGetSyncStatusParamsWithTimeout(timeout time.Duration) *GetSyncStatusParams {
	return &GetSyncStatusParams{
		timeout: timeout,
	}
}

// NewGetSyncStatusParamsWithContext creates a new GetSyncStatusParams object
// with the ability to set a context for a request.
func NewGetSyncStatusParamsWithContext(ctx context.Context) *GetSyncStatusParams {
	return &GetSyncStatusParams{
		Context: ctx,
	}
}

// NewGetSyncStatusParamsWithHTTPClient creates a new GetSyncStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSyncStatusParamsWithHTTPClient(client *http.Client) *GetSyncStatusParams {
	return &GetSyncStatusParams{
		HTTPClient: client,
	}
}

/*
GetSyncStatusParams contains all the parameters to send to the API endpoint

	for the get sync status operation.

	Typically these are written to a http.Request.
*/
type GetSyncStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sync status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyncStatusParams) WithDefaults() *GetSyncStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sync status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyncStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sync status params
func (o *GetSyncStatusParams) WithTimeout(timeout time.Duration) *GetSyncStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync status params
func (o *GetSyncStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync status params
func (o *GetSyncStatusParams) WithContext(ctx context.Context) *GetSyncStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync status params
func (o *GetSyncStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync status params
func (o *GetSyncStatusParams) WithHTTPClient(client *http.Client) *GetSyncStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync status params
func (o *GetSyncStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}