// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// NewRouteGetContactpointsParams creates a new RouteGetContactpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteGetContactpointsParams() *RouteGetContactpointsParams {
	return &RouteGetContactpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteGetContactpointsParamsWithTimeout creates a new RouteGetContactpointsParams object
// with the ability to set a timeout on a request.
func NewRouteGetContactpointsParamsWithTimeout(timeout time.Duration) *RouteGetContactpointsParams {
	return &RouteGetContactpointsParams{
		timeout: timeout,
	}
}

// NewRouteGetContactpointsParamsWithContext creates a new RouteGetContactpointsParams object
// with the ability to set a context for a request.
func NewRouteGetContactpointsParamsWithContext(ctx context.Context) *RouteGetContactpointsParams {
	return &RouteGetContactpointsParams{
		Context: ctx,
	}
}

// NewRouteGetContactpointsParamsWithHTTPClient creates a new RouteGetContactpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteGetContactpointsParamsWithHTTPClient(client *http.Client) *RouteGetContactpointsParams {
	return &RouteGetContactpointsParams{
		HTTPClient: client,
	}
}

/*
RouteGetContactpointsParams contains all the parameters to send to the API endpoint

	for the route get contactpoints operation.

	Typically these are written to a http.Request.
*/
type RouteGetContactpointsParams struct {

	/* Name.

	   Filter by name
	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route get contactpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetContactpointsParams) WithDefaults() *RouteGetContactpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route get contactpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetContactpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route get contactpoints params
func (o *RouteGetContactpointsParams) WithTimeout(timeout time.Duration) *RouteGetContactpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route get contactpoints params
func (o *RouteGetContactpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route get contactpoints params
func (o *RouteGetContactpointsParams) WithContext(ctx context.Context) *RouteGetContactpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route get contactpoints params
func (o *RouteGetContactpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route get contactpoints params
func (o *RouteGetContactpointsParams) WithHTTPClient(client *http.Client) *RouteGetContactpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route get contactpoints params
func (o *RouteGetContactpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the route get contactpoints params
func (o *RouteGetContactpointsParams) WithName(name *string) *RouteGetContactpointsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the route get contactpoints params
func (o *RouteGetContactpointsParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RouteGetContactpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
