// Code generated by go-swagger; DO NOT EDIT.

package dashboard_permissions

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
	"github.com/go-openapi/swag"
)

// NewGetDashboardPermissionsListByIDParams creates a new GetDashboardPermissionsListByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardPermissionsListByIDParams() *GetDashboardPermissionsListByIDParams {
	return &GetDashboardPermissionsListByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardPermissionsListByIDParamsWithTimeout creates a new GetDashboardPermissionsListByIDParams object
// with the ability to set a timeout on a request.
func NewGetDashboardPermissionsListByIDParamsWithTimeout(timeout time.Duration) *GetDashboardPermissionsListByIDParams {
	return &GetDashboardPermissionsListByIDParams{
		timeout: timeout,
	}
}

// NewGetDashboardPermissionsListByIDParamsWithContext creates a new GetDashboardPermissionsListByIDParams object
// with the ability to set a context for a request.
func NewGetDashboardPermissionsListByIDParamsWithContext(ctx context.Context) *GetDashboardPermissionsListByIDParams {
	return &GetDashboardPermissionsListByIDParams{
		Context: ctx,
	}
}

// NewGetDashboardPermissionsListByIDParamsWithHTTPClient creates a new GetDashboardPermissionsListByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardPermissionsListByIDParamsWithHTTPClient(client *http.Client) *GetDashboardPermissionsListByIDParams {
	return &GetDashboardPermissionsListByIDParams{
		HTTPClient: client,
	}
}

/*
GetDashboardPermissionsListByIDParams contains all the parameters to send to the API endpoint

	for the get dashboard permissions list by ID operation.

	Typically these are written to a http.Request.
*/
type GetDashboardPermissionsListByIDParams struct {

	// DashboardID.
	//
	// Format: int64
	DashboardID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard permissions list by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardPermissionsListByIDParams) WithDefaults() *GetDashboardPermissionsListByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard permissions list by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardPermissionsListByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) WithTimeout(timeout time.Duration) *GetDashboardPermissionsListByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) WithContext(ctx context.Context) *GetDashboardPermissionsListByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) WithHTTPClient(client *http.Client) *GetDashboardPermissionsListByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) WithDashboardID(dashboardID int64) *GetDashboardPermissionsListByIDParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the get dashboard permissions list by ID params
func (o *GetDashboardPermissionsListByIDParams) SetDashboardID(dashboardID int64) {
	o.DashboardID = dashboardID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardPermissionsListByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param DashboardID
	if err := r.SetPathParam("DashboardID", swag.FormatInt64(o.DashboardID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
