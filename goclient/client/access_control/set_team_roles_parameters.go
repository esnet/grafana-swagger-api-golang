// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// NewSetTeamRolesParams creates a new SetTeamRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetTeamRolesParams() *SetTeamRolesParams {
	return &SetTeamRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetTeamRolesParamsWithTimeout creates a new SetTeamRolesParams object
// with the ability to set a timeout on a request.
func NewSetTeamRolesParamsWithTimeout(timeout time.Duration) *SetTeamRolesParams {
	return &SetTeamRolesParams{
		timeout: timeout,
	}
}

// NewSetTeamRolesParamsWithContext creates a new SetTeamRolesParams object
// with the ability to set a context for a request.
func NewSetTeamRolesParamsWithContext(ctx context.Context) *SetTeamRolesParams {
	return &SetTeamRolesParams{
		Context: ctx,
	}
}

// NewSetTeamRolesParamsWithHTTPClient creates a new SetTeamRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetTeamRolesParamsWithHTTPClient(client *http.Client) *SetTeamRolesParams {
	return &SetTeamRolesParams{
		HTTPClient: client,
	}
}

/*
SetTeamRolesParams contains all the parameters to send to the API endpoint

	for the set team roles operation.

	Typically these are written to a http.Request.
*/
type SetTeamRolesParams struct {

	// TeamID.
	//
	// Format: int64
	TeamID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set team roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTeamRolesParams) WithDefaults() *SetTeamRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set team roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTeamRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set team roles params
func (o *SetTeamRolesParams) WithTimeout(timeout time.Duration) *SetTeamRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set team roles params
func (o *SetTeamRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set team roles params
func (o *SetTeamRolesParams) WithContext(ctx context.Context) *SetTeamRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set team roles params
func (o *SetTeamRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set team roles params
func (o *SetTeamRolesParams) WithHTTPClient(client *http.Client) *SetTeamRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set team roles params
func (o *SetTeamRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the set team roles params
func (o *SetTeamRolesParams) WithTeamID(teamID int64) *SetTeamRolesParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the set team roles params
func (o *SetTeamRolesParams) SetTeamID(teamID int64) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *SetTeamRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param teamId
	if err := r.SetPathParam("teamId", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
