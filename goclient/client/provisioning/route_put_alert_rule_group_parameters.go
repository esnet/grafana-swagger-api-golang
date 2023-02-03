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

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// NewRoutePutAlertRuleGroupParams creates a new RoutePutAlertRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRoutePutAlertRuleGroupParams() *RoutePutAlertRuleGroupParams {
	return &RoutePutAlertRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRoutePutAlertRuleGroupParamsWithTimeout creates a new RoutePutAlertRuleGroupParams object
// with the ability to set a timeout on a request.
func NewRoutePutAlertRuleGroupParamsWithTimeout(timeout time.Duration) *RoutePutAlertRuleGroupParams {
	return &RoutePutAlertRuleGroupParams{
		timeout: timeout,
	}
}

// NewRoutePutAlertRuleGroupParamsWithContext creates a new RoutePutAlertRuleGroupParams object
// with the ability to set a context for a request.
func NewRoutePutAlertRuleGroupParamsWithContext(ctx context.Context) *RoutePutAlertRuleGroupParams {
	return &RoutePutAlertRuleGroupParams{
		Context: ctx,
	}
}

// NewRoutePutAlertRuleGroupParamsWithHTTPClient creates a new RoutePutAlertRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewRoutePutAlertRuleGroupParamsWithHTTPClient(client *http.Client) *RoutePutAlertRuleGroupParams {
	return &RoutePutAlertRuleGroupParams{
		HTTPClient: client,
	}
}

/*
RoutePutAlertRuleGroupParams contains all the parameters to send to the API endpoint

	for the route put alert rule group operation.

	Typically these are written to a http.Request.
*/
type RoutePutAlertRuleGroupParams struct {

	// Body.
	Body *models.AlertRuleGroup

	// FolderUID.
	FolderUID string

	// Group.
	Group string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route put alert rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoutePutAlertRuleGroupParams) WithDefaults() *RoutePutAlertRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route put alert rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoutePutAlertRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) WithTimeout(timeout time.Duration) *RoutePutAlertRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) WithContext(ctx context.Context) *RoutePutAlertRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) WithHTTPClient(client *http.Client) *RoutePutAlertRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) WithBody(body *models.AlertRuleGroup) *RoutePutAlertRuleGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) SetBody(body *models.AlertRuleGroup) {
	o.Body = body
}

// WithFolderUID adds the folderUID to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) WithFolderUID(folderUID string) *RoutePutAlertRuleGroupParams {
	o.SetFolderUID(folderUID)
	return o
}

// SetFolderUID adds the folderUid to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) SetFolderUID(folderUID string) {
	o.FolderUID = folderUID
}

// WithGroup adds the group to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) WithGroup(group string) *RoutePutAlertRuleGroupParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the route put alert rule group params
func (o *RoutePutAlertRuleGroupParams) SetGroup(group string) {
	o.Group = group
}

// WriteToRequest writes these params to a swagger request
func (o *RoutePutAlertRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param FolderUID
	if err := r.SetPathParam("FolderUID", o.FolderUID); err != nil {
		return err
	}

	// path param Group
	if err := r.SetPathParam("Group", o.Group); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}