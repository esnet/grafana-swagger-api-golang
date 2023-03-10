// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// NewRemoveOrgUserParams creates a new RemoveOrgUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveOrgUserParams() *RemoveOrgUserParams {
	return &RemoveOrgUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveOrgUserParamsWithTimeout creates a new RemoveOrgUserParams object
// with the ability to set a timeout on a request.
func NewRemoveOrgUserParamsWithTimeout(timeout time.Duration) *RemoveOrgUserParams {
	return &RemoveOrgUserParams{
		timeout: timeout,
	}
}

// NewRemoveOrgUserParamsWithContext creates a new RemoveOrgUserParams object
// with the ability to set a context for a request.
func NewRemoveOrgUserParamsWithContext(ctx context.Context) *RemoveOrgUserParams {
	return &RemoveOrgUserParams{
		Context: ctx,
	}
}

// NewRemoveOrgUserParamsWithHTTPClient creates a new RemoveOrgUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveOrgUserParamsWithHTTPClient(client *http.Client) *RemoveOrgUserParams {
	return &RemoveOrgUserParams{
		HTTPClient: client,
	}
}

/*
RemoveOrgUserParams contains all the parameters to send to the API endpoint

	for the remove org user operation.

	Typically these are written to a http.Request.
*/
type RemoveOrgUserParams struct {

	// OrgID.
	//
	// Format: int64
	OrgID int64

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove org user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveOrgUserParams) WithDefaults() *RemoveOrgUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove org user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveOrgUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove org user params
func (o *RemoveOrgUserParams) WithTimeout(timeout time.Duration) *RemoveOrgUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove org user params
func (o *RemoveOrgUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove org user params
func (o *RemoveOrgUserParams) WithContext(ctx context.Context) *RemoveOrgUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove org user params
func (o *RemoveOrgUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove org user params
func (o *RemoveOrgUserParams) WithHTTPClient(client *http.Client) *RemoveOrgUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove org user params
func (o *RemoveOrgUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgID adds the orgID to the remove org user params
func (o *RemoveOrgUserParams) WithOrgID(orgID int64) *RemoveOrgUserParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the remove org user params
func (o *RemoveOrgUserParams) SetOrgID(orgID int64) {
	o.OrgID = orgID
}

// WithUserID adds the userID to the remove org user params
func (o *RemoveOrgUserParams) WithUserID(userID int64) *RemoveOrgUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the remove org user params
func (o *RemoveOrgUserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveOrgUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org_id
	if err := r.SetPathParam("org_id", swag.FormatInt64(o.OrgID)); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
