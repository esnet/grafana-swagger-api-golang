// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

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

// NewStarDashboardParams creates a new StarDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStarDashboardParams() *StarDashboardParams {
	return &StarDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStarDashboardParamsWithTimeout creates a new StarDashboardParams object
// with the ability to set a timeout on a request.
func NewStarDashboardParamsWithTimeout(timeout time.Duration) *StarDashboardParams {
	return &StarDashboardParams{
		timeout: timeout,
	}
}

// NewStarDashboardParamsWithContext creates a new StarDashboardParams object
// with the ability to set a context for a request.
func NewStarDashboardParamsWithContext(ctx context.Context) *StarDashboardParams {
	return &StarDashboardParams{
		Context: ctx,
	}
}

// NewStarDashboardParamsWithHTTPClient creates a new StarDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewStarDashboardParamsWithHTTPClient(client *http.Client) *StarDashboardParams {
	return &StarDashboardParams{
		HTTPClient: client,
	}
}

/*
StarDashboardParams contains all the parameters to send to the API endpoint

	for the star dashboard operation.

	Typically these are written to a http.Request.
*/
type StarDashboardParams struct {

	// DashboardID.
	DashboardID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the star dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StarDashboardParams) WithDefaults() *StarDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the star dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StarDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the star dashboard params
func (o *StarDashboardParams) WithTimeout(timeout time.Duration) *StarDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the star dashboard params
func (o *StarDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the star dashboard params
func (o *StarDashboardParams) WithContext(ctx context.Context) *StarDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the star dashboard params
func (o *StarDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the star dashboard params
func (o *StarDashboardParams) WithHTTPClient(client *http.Client) *StarDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the star dashboard params
func (o *StarDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the star dashboard params
func (o *StarDashboardParams) WithDashboardID(dashboardID string) *StarDashboardParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the star dashboard params
func (o *StarDashboardParams) SetDashboardID(dashboardID string) {
	o.DashboardID = dashboardID
}

// WriteToRequest writes these params to a swagger request
func (o *StarDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard_id
	if err := r.SetPathParam("dashboard_id", o.DashboardID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
