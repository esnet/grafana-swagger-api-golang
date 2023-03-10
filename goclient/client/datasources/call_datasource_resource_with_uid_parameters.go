// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

// NewCallDatasourceResourceWithUIDParams creates a new CallDatasourceResourceWithUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCallDatasourceResourceWithUIDParams() *CallDatasourceResourceWithUIDParams {
	return &CallDatasourceResourceWithUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCallDatasourceResourceWithUIDParamsWithTimeout creates a new CallDatasourceResourceWithUIDParams object
// with the ability to set a timeout on a request.
func NewCallDatasourceResourceWithUIDParamsWithTimeout(timeout time.Duration) *CallDatasourceResourceWithUIDParams {
	return &CallDatasourceResourceWithUIDParams{
		timeout: timeout,
	}
}

// NewCallDatasourceResourceWithUIDParamsWithContext creates a new CallDatasourceResourceWithUIDParams object
// with the ability to set a context for a request.
func NewCallDatasourceResourceWithUIDParamsWithContext(ctx context.Context) *CallDatasourceResourceWithUIDParams {
	return &CallDatasourceResourceWithUIDParams{
		Context: ctx,
	}
}

// NewCallDatasourceResourceWithUIDParamsWithHTTPClient creates a new CallDatasourceResourceWithUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCallDatasourceResourceWithUIDParamsWithHTTPClient(client *http.Client) *CallDatasourceResourceWithUIDParams {
	return &CallDatasourceResourceWithUIDParams{
		HTTPClient: client,
	}
}

/*
CallDatasourceResourceWithUIDParams contains all the parameters to send to the API endpoint

	for the call datasource resource with UID operation.

	Typically these are written to a http.Request.
*/
type CallDatasourceResourceWithUIDParams struct {

	// DatasourceProxyRoute.
	DatasourceProxyRoute string

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the call datasource resource with UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CallDatasourceResourceWithUIDParams) WithDefaults() *CallDatasourceResourceWithUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the call datasource resource with UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CallDatasourceResourceWithUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) WithTimeout(timeout time.Duration) *CallDatasourceResourceWithUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) WithContext(ctx context.Context) *CallDatasourceResourceWithUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) WithHTTPClient(client *http.Client) *CallDatasourceResourceWithUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasourceProxyRoute adds the datasourceProxyRoute to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) WithDatasourceProxyRoute(datasourceProxyRoute string) *CallDatasourceResourceWithUIDParams {
	o.SetDatasourceProxyRoute(datasourceProxyRoute)
	return o
}

// SetDatasourceProxyRoute adds the datasourceProxyRoute to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) SetDatasourceProxyRoute(datasourceProxyRoute string) {
	o.DatasourceProxyRoute = datasourceProxyRoute
}

// WithUID adds the uid to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) WithUID(uid string) *CallDatasourceResourceWithUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the call datasource resource with UID params
func (o *CallDatasourceResourceWithUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *CallDatasourceResourceWithUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasource_proxy_route
	if err := r.SetPathParam("datasource_proxy_route", o.DatasourceProxyRoute); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
