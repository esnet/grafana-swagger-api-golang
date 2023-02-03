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

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// NewCreateQueryParams creates a new CreateQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateQueryParams() *CreateQueryParams {
	return &CreateQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateQueryParamsWithTimeout creates a new CreateQueryParams object
// with the ability to set a timeout on a request.
func NewCreateQueryParamsWithTimeout(timeout time.Duration) *CreateQueryParams {
	return &CreateQueryParams{
		timeout: timeout,
	}
}

// NewCreateQueryParamsWithContext creates a new CreateQueryParams object
// with the ability to set a context for a request.
func NewCreateQueryParamsWithContext(ctx context.Context) *CreateQueryParams {
	return &CreateQueryParams{
		Context: ctx,
	}
}

// NewCreateQueryParamsWithHTTPClient creates a new CreateQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateQueryParamsWithHTTPClient(client *http.Client) *CreateQueryParams {
	return &CreateQueryParams{
		HTTPClient: client,
	}
}

/*
CreateQueryParams contains all the parameters to send to the API endpoint

	for the create query operation.

	Typically these are written to a http.Request.
*/
type CreateQueryParams struct {

	// Body.
	Body *models.CreateQueryInQueryHistoryCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateQueryParams) WithDefaults() *CreateQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create query params
func (o *CreateQueryParams) WithTimeout(timeout time.Duration) *CreateQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create query params
func (o *CreateQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create query params
func (o *CreateQueryParams) WithContext(ctx context.Context) *CreateQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create query params
func (o *CreateQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create query params
func (o *CreateQueryParams) WithHTTPClient(client *http.Client) *CreateQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create query params
func (o *CreateQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create query params
func (o *CreateQueryParams) WithBody(body *models.CreateQueryInQueryHistoryCommand) *CreateQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create query params
func (o *CreateQueryParams) SetBody(body *models.CreateQueryInQueryHistoryCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
