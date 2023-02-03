// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// NewUpdateUserQuotaParams creates a new UpdateUserQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserQuotaParams() *UpdateUserQuotaParams {
	return &UpdateUserQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserQuotaParamsWithTimeout creates a new UpdateUserQuotaParams object
// with the ability to set a timeout on a request.
func NewUpdateUserQuotaParamsWithTimeout(timeout time.Duration) *UpdateUserQuotaParams {
	return &UpdateUserQuotaParams{
		timeout: timeout,
	}
}

// NewUpdateUserQuotaParamsWithContext creates a new UpdateUserQuotaParams object
// with the ability to set a context for a request.
func NewUpdateUserQuotaParamsWithContext(ctx context.Context) *UpdateUserQuotaParams {
	return &UpdateUserQuotaParams{
		Context: ctx,
	}
}

// NewUpdateUserQuotaParamsWithHTTPClient creates a new UpdateUserQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserQuotaParamsWithHTTPClient(client *http.Client) *UpdateUserQuotaParams {
	return &UpdateUserQuotaParams{
		HTTPClient: client,
	}
}

/*
UpdateUserQuotaParams contains all the parameters to send to the API endpoint

	for the update user quota operation.

	Typically these are written to a http.Request.
*/
type UpdateUserQuotaParams struct {

	// Body.
	Body *models.UpdateUserQuotaCmd

	// QuotaTarget.
	QuotaTarget string

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserQuotaParams) WithDefaults() *UpdateUserQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user quota params
func (o *UpdateUserQuotaParams) WithTimeout(timeout time.Duration) *UpdateUserQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user quota params
func (o *UpdateUserQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user quota params
func (o *UpdateUserQuotaParams) WithContext(ctx context.Context) *UpdateUserQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user quota params
func (o *UpdateUserQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user quota params
func (o *UpdateUserQuotaParams) WithHTTPClient(client *http.Client) *UpdateUserQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user quota params
func (o *UpdateUserQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user quota params
func (o *UpdateUserQuotaParams) WithBody(body *models.UpdateUserQuotaCmd) *UpdateUserQuotaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user quota params
func (o *UpdateUserQuotaParams) SetBody(body *models.UpdateUserQuotaCmd) {
	o.Body = body
}

// WithQuotaTarget adds the quotaTarget to the update user quota params
func (o *UpdateUserQuotaParams) WithQuotaTarget(quotaTarget string) *UpdateUserQuotaParams {
	o.SetQuotaTarget(quotaTarget)
	return o
}

// SetQuotaTarget adds the quotaTarget to the update user quota params
func (o *UpdateUserQuotaParams) SetQuotaTarget(quotaTarget string) {
	o.QuotaTarget = quotaTarget
}

// WithUserID adds the userID to the update user quota params
func (o *UpdateUserQuotaParams) WithUserID(userID int64) *UpdateUserQuotaParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user quota params
func (o *UpdateUserQuotaParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param quota_target
	if err := r.SetPathParam("quota_target", o.QuotaTarget); err != nil {
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
