// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

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

// NewGetTeamGroupsAPIParams creates a new GetTeamGroupsAPIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeamGroupsAPIParams() *GetTeamGroupsAPIParams {
	return &GetTeamGroupsAPIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamGroupsAPIParamsWithTimeout creates a new GetTeamGroupsAPIParams object
// with the ability to set a timeout on a request.
func NewGetTeamGroupsAPIParamsWithTimeout(timeout time.Duration) *GetTeamGroupsAPIParams {
	return &GetTeamGroupsAPIParams{
		timeout: timeout,
	}
}

// NewGetTeamGroupsAPIParamsWithContext creates a new GetTeamGroupsAPIParams object
// with the ability to set a context for a request.
func NewGetTeamGroupsAPIParamsWithContext(ctx context.Context) *GetTeamGroupsAPIParams {
	return &GetTeamGroupsAPIParams{
		Context: ctx,
	}
}

// NewGetTeamGroupsAPIParamsWithHTTPClient creates a new GetTeamGroupsAPIParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeamGroupsAPIParamsWithHTTPClient(client *http.Client) *GetTeamGroupsAPIParams {
	return &GetTeamGroupsAPIParams{
		HTTPClient: client,
	}
}

/*
GetTeamGroupsAPIParams contains all the parameters to send to the API endpoint

	for the get team groups Api operation.

	Typically these are written to a http.Request.
*/
type GetTeamGroupsAPIParams struct {

	// TeamID.
	//
	// Format: int64
	TeamID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get team groups Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamGroupsAPIParams) WithDefaults() *GetTeamGroupsAPIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get team groups Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamGroupsAPIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get team groups Api params
func (o *GetTeamGroupsAPIParams) WithTimeout(timeout time.Duration) *GetTeamGroupsAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team groups Api params
func (o *GetTeamGroupsAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team groups Api params
func (o *GetTeamGroupsAPIParams) WithContext(ctx context.Context) *GetTeamGroupsAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team groups Api params
func (o *GetTeamGroupsAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team groups Api params
func (o *GetTeamGroupsAPIParams) WithHTTPClient(client *http.Client) *GetTeamGroupsAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team groups Api params
func (o *GetTeamGroupsAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the get team groups Api params
func (o *GetTeamGroupsAPIParams) WithTeamID(teamID int64) *GetTeamGroupsAPIParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get team groups Api params
func (o *GetTeamGroupsAPIParams) SetTeamID(teamID int64) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamGroupsAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
