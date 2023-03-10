// Code generated by go-swagger; DO NOT EDIT.

package query_history

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

// NewStarQueryParams creates a new StarQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStarQueryParams() *StarQueryParams {
	return &StarQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStarQueryParamsWithTimeout creates a new StarQueryParams object
// with the ability to set a timeout on a request.
func NewStarQueryParamsWithTimeout(timeout time.Duration) *StarQueryParams {
	return &StarQueryParams{
		timeout: timeout,
	}
}

// NewStarQueryParamsWithContext creates a new StarQueryParams object
// with the ability to set a context for a request.
func NewStarQueryParamsWithContext(ctx context.Context) *StarQueryParams {
	return &StarQueryParams{
		Context: ctx,
	}
}

// NewStarQueryParamsWithHTTPClient creates a new StarQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewStarQueryParamsWithHTTPClient(client *http.Client) *StarQueryParams {
	return &StarQueryParams{
		HTTPClient: client,
	}
}

/*
StarQueryParams contains all the parameters to send to the API endpoint

	for the star query operation.

	Typically these are written to a http.Request.
*/
type StarQueryParams struct {

	// QueryHistoryUID.
	QueryHistoryUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the star query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StarQueryParams) WithDefaults() *StarQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the star query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StarQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the star query params
func (o *StarQueryParams) WithTimeout(timeout time.Duration) *StarQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the star query params
func (o *StarQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the star query params
func (o *StarQueryParams) WithContext(ctx context.Context) *StarQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the star query params
func (o *StarQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the star query params
func (o *StarQueryParams) WithHTTPClient(client *http.Client) *StarQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the star query params
func (o *StarQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryHistoryUID adds the queryHistoryUID to the star query params
func (o *StarQueryParams) WithQueryHistoryUID(queryHistoryUID string) *StarQueryParams {
	o.SetQueryHistoryUID(queryHistoryUID)
	return o
}

// SetQueryHistoryUID adds the queryHistoryUid to the star query params
func (o *StarQueryParams) SetQueryHistoryUID(queryHistoryUID string) {
	o.QueryHistoryUID = queryHistoryUID
}

// WriteToRequest writes these params to a swagger request
func (o *StarQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_history_uid
	if err := r.SetPathParam("query_history_uid", o.QueryHistoryUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
