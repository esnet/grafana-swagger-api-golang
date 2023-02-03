// Code generated by go-swagger; DO NOT EDIT.

package recording_rules

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

// NewTestCreateRecordingRuleParams creates a new TestCreateRecordingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestCreateRecordingRuleParams() *TestCreateRecordingRuleParams {
	return &TestCreateRecordingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestCreateRecordingRuleParamsWithTimeout creates a new TestCreateRecordingRuleParams object
// with the ability to set a timeout on a request.
func NewTestCreateRecordingRuleParamsWithTimeout(timeout time.Duration) *TestCreateRecordingRuleParams {
	return &TestCreateRecordingRuleParams{
		timeout: timeout,
	}
}

// NewTestCreateRecordingRuleParamsWithContext creates a new TestCreateRecordingRuleParams object
// with the ability to set a context for a request.
func NewTestCreateRecordingRuleParamsWithContext(ctx context.Context) *TestCreateRecordingRuleParams {
	return &TestCreateRecordingRuleParams{
		Context: ctx,
	}
}

// NewTestCreateRecordingRuleParamsWithHTTPClient creates a new TestCreateRecordingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestCreateRecordingRuleParamsWithHTTPClient(client *http.Client) *TestCreateRecordingRuleParams {
	return &TestCreateRecordingRuleParams{
		HTTPClient: client,
	}
}

/*
TestCreateRecordingRuleParams contains all the parameters to send to the API endpoint

	for the test create recording rule operation.

	Typically these are written to a http.Request.
*/
type TestCreateRecordingRuleParams struct {

	// Body.
	Body *models.RecordingRuleJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test create recording rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestCreateRecordingRuleParams) WithDefaults() *TestCreateRecordingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test create recording rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestCreateRecordingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test create recording rule params
func (o *TestCreateRecordingRuleParams) WithTimeout(timeout time.Duration) *TestCreateRecordingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test create recording rule params
func (o *TestCreateRecordingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test create recording rule params
func (o *TestCreateRecordingRuleParams) WithContext(ctx context.Context) *TestCreateRecordingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test create recording rule params
func (o *TestCreateRecordingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test create recording rule params
func (o *TestCreateRecordingRuleParams) WithHTTPClient(client *http.Client) *TestCreateRecordingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test create recording rule params
func (o *TestCreateRecordingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test create recording rule params
func (o *TestCreateRecordingRuleParams) WithBody(body *models.RecordingRuleJSON) *TestCreateRecordingRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test create recording rule params
func (o *TestCreateRecordingRuleParams) SetBody(body *models.RecordingRuleJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestCreateRecordingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
