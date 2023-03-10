// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// NewGetSharingOptionsParams creates a new GetSharingOptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSharingOptionsParams() *GetSharingOptionsParams {
	return &GetSharingOptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSharingOptionsParamsWithTimeout creates a new GetSharingOptionsParams object
// with the ability to set a timeout on a request.
func NewGetSharingOptionsParamsWithTimeout(timeout time.Duration) *GetSharingOptionsParams {
	return &GetSharingOptionsParams{
		timeout: timeout,
	}
}

// NewGetSharingOptionsParamsWithContext creates a new GetSharingOptionsParams object
// with the ability to set a context for a request.
func NewGetSharingOptionsParamsWithContext(ctx context.Context) *GetSharingOptionsParams {
	return &GetSharingOptionsParams{
		Context: ctx,
	}
}

// NewGetSharingOptionsParamsWithHTTPClient creates a new GetSharingOptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSharingOptionsParamsWithHTTPClient(client *http.Client) *GetSharingOptionsParams {
	return &GetSharingOptionsParams{
		HTTPClient: client,
	}
}

/*
GetSharingOptionsParams contains all the parameters to send to the API endpoint

	for the get sharing options operation.

	Typically these are written to a http.Request.
*/
type GetSharingOptionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sharing options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSharingOptionsParams) WithDefaults() *GetSharingOptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sharing options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSharingOptionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sharing options params
func (o *GetSharingOptionsParams) WithTimeout(timeout time.Duration) *GetSharingOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sharing options params
func (o *GetSharingOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sharing options params
func (o *GetSharingOptionsParams) WithContext(ctx context.Context) *GetSharingOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sharing options params
func (o *GetSharingOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sharing options params
func (o *GetSharingOptionsParams) WithHTTPClient(client *http.Client) *GetSharingOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sharing options params
func (o *GetSharingOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSharingOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
