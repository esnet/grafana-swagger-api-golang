// Code generated by go-swagger; DO NOT EDIT.

package org

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

// NewUpdateCurrentOrgParams creates a new UpdateCurrentOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCurrentOrgParams() *UpdateCurrentOrgParams {
	return &UpdateCurrentOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCurrentOrgParamsWithTimeout creates a new UpdateCurrentOrgParams object
// with the ability to set a timeout on a request.
func NewUpdateCurrentOrgParamsWithTimeout(timeout time.Duration) *UpdateCurrentOrgParams {
	return &UpdateCurrentOrgParams{
		timeout: timeout,
	}
}

// NewUpdateCurrentOrgParamsWithContext creates a new UpdateCurrentOrgParams object
// with the ability to set a context for a request.
func NewUpdateCurrentOrgParamsWithContext(ctx context.Context) *UpdateCurrentOrgParams {
	return &UpdateCurrentOrgParams{
		Context: ctx,
	}
}

// NewUpdateCurrentOrgParamsWithHTTPClient creates a new UpdateCurrentOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCurrentOrgParamsWithHTTPClient(client *http.Client) *UpdateCurrentOrgParams {
	return &UpdateCurrentOrgParams{
		HTTPClient: client,
	}
}

/*
UpdateCurrentOrgParams contains all the parameters to send to the API endpoint

	for the update current org operation.

	Typically these are written to a http.Request.
*/
type UpdateCurrentOrgParams struct {

	// Body.
	Body *models.UpdateOrgForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update current org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentOrgParams) WithDefaults() *UpdateCurrentOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update current org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update current org params
func (o *UpdateCurrentOrgParams) WithTimeout(timeout time.Duration) *UpdateCurrentOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update current org params
func (o *UpdateCurrentOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update current org params
func (o *UpdateCurrentOrgParams) WithContext(ctx context.Context) *UpdateCurrentOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update current org params
func (o *UpdateCurrentOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update current org params
func (o *UpdateCurrentOrgParams) WithHTTPClient(client *http.Client) *UpdateCurrentOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update current org params
func (o *UpdateCurrentOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update current org params
func (o *UpdateCurrentOrgParams) WithBody(body *models.UpdateOrgForm) *UpdateCurrentOrgParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update current org params
func (o *UpdateCurrentOrgParams) SetBody(body *models.UpdateOrgForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCurrentOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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