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

// NewRemoveTeamRoleParams creates a new RemoveTeamRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTeamRoleParams() *RemoveTeamRoleParams {
	return &RemoveTeamRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTeamRoleParamsWithTimeout creates a new RemoveTeamRoleParams object
// with the ability to set a timeout on a request.
func NewRemoveTeamRoleParamsWithTimeout(timeout time.Duration) *RemoveTeamRoleParams {
	return &RemoveTeamRoleParams{
		timeout: timeout,
	}
}

// NewRemoveTeamRoleParamsWithContext creates a new RemoveTeamRoleParams object
// with the ability to set a context for a request.
func NewRemoveTeamRoleParamsWithContext(ctx context.Context) *RemoveTeamRoleParams {
	return &RemoveTeamRoleParams{
		Context: ctx,
	}
}

// NewRemoveTeamRoleParamsWithHTTPClient creates a new RemoveTeamRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTeamRoleParamsWithHTTPClient(client *http.Client) *RemoveTeamRoleParams {
	return &RemoveTeamRoleParams{
		HTTPClient: client,
	}
}

/*
RemoveTeamRoleParams contains all the parameters to send to the API endpoint

	for the remove team role operation.

	Typically these are written to a http.Request.
*/
type RemoveTeamRoleParams struct {

	// RoleUID.
	RoleUID string

	// TeamID.
	//
	// Format: int64
	TeamID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove team role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTeamRoleParams) WithDefaults() *RemoveTeamRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove team role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTeamRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove team role params
func (o *RemoveTeamRoleParams) WithTimeout(timeout time.Duration) *RemoveTeamRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove team role params
func (o *RemoveTeamRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove team role params
func (o *RemoveTeamRoleParams) WithContext(ctx context.Context) *RemoveTeamRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove team role params
func (o *RemoveTeamRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove team role params
func (o *RemoveTeamRoleParams) WithHTTPClient(client *http.Client) *RemoveTeamRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove team role params
func (o *RemoveTeamRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleUID adds the roleUID to the remove team role params
func (o *RemoveTeamRoleParams) WithRoleUID(roleUID string) *RemoveTeamRoleParams {
	o.SetRoleUID(roleUID)
	return o
}

// SetRoleUID adds the roleUid to the remove team role params
func (o *RemoveTeamRoleParams) SetRoleUID(roleUID string) {
	o.RoleUID = roleUID
}

// WithTeamID adds the teamID to the remove team role params
func (o *RemoveTeamRoleParams) WithTeamID(teamID int64) *RemoveTeamRoleParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the remove team role params
func (o *RemoveTeamRoleParams) SetTeamID(teamID int64) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTeamRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleUID
	if err := r.SetPathParam("roleUID", o.RoleUID); err != nil {
		return err
	}

	// path param teamId
	if err := r.SetPathParam("teamId", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}