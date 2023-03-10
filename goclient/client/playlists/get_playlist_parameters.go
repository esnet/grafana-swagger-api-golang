// Code generated by go-swagger; DO NOT EDIT.

package playlists

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

// NewGetPlaylistParams creates a new GetPlaylistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlaylistParams() *GetPlaylistParams {
	return &GetPlaylistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlaylistParamsWithTimeout creates a new GetPlaylistParams object
// with the ability to set a timeout on a request.
func NewGetPlaylistParamsWithTimeout(timeout time.Duration) *GetPlaylistParams {
	return &GetPlaylistParams{
		timeout: timeout,
	}
}

// NewGetPlaylistParamsWithContext creates a new GetPlaylistParams object
// with the ability to set a context for a request.
func NewGetPlaylistParamsWithContext(ctx context.Context) *GetPlaylistParams {
	return &GetPlaylistParams{
		Context: ctx,
	}
}

// NewGetPlaylistParamsWithHTTPClient creates a new GetPlaylistParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlaylistParamsWithHTTPClient(client *http.Client) *GetPlaylistParams {
	return &GetPlaylistParams{
		HTTPClient: client,
	}
}

/*
GetPlaylistParams contains all the parameters to send to the API endpoint

	for the get playlist operation.

	Typically these are written to a http.Request.
*/
type GetPlaylistParams struct {

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get playlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaylistParams) WithDefaults() *GetPlaylistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get playlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaylistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get playlist params
func (o *GetPlaylistParams) WithTimeout(timeout time.Duration) *GetPlaylistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get playlist params
func (o *GetPlaylistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get playlist params
func (o *GetPlaylistParams) WithContext(ctx context.Context) *GetPlaylistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get playlist params
func (o *GetPlaylistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get playlist params
func (o *GetPlaylistParams) WithHTTPClient(client *http.Client) *GetPlaylistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get playlist params
func (o *GetPlaylistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the get playlist params
func (o *GetPlaylistParams) WithUID(uid string) *GetPlaylistParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get playlist params
func (o *GetPlaylistParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlaylistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
