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
	"github.com/go-openapi/swag"
)

// NewRouteGetAlertRuleGroupExportParams creates a new RouteGetAlertRuleGroupExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteGetAlertRuleGroupExportParams() *RouteGetAlertRuleGroupExportParams {
	return &RouteGetAlertRuleGroupExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteGetAlertRuleGroupExportParamsWithTimeout creates a new RouteGetAlertRuleGroupExportParams object
// with the ability to set a timeout on a request.
func NewRouteGetAlertRuleGroupExportParamsWithTimeout(timeout time.Duration) *RouteGetAlertRuleGroupExportParams {
	return &RouteGetAlertRuleGroupExportParams{
		timeout: timeout,
	}
}

// NewRouteGetAlertRuleGroupExportParamsWithContext creates a new RouteGetAlertRuleGroupExportParams object
// with the ability to set a context for a request.
func NewRouteGetAlertRuleGroupExportParamsWithContext(ctx context.Context) *RouteGetAlertRuleGroupExportParams {
	return &RouteGetAlertRuleGroupExportParams{
		Context: ctx,
	}
}

// NewRouteGetAlertRuleGroupExportParamsWithHTTPClient creates a new RouteGetAlertRuleGroupExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteGetAlertRuleGroupExportParamsWithHTTPClient(client *http.Client) *RouteGetAlertRuleGroupExportParams {
	return &RouteGetAlertRuleGroupExportParams{
		HTTPClient: client,
	}
}

/*
RouteGetAlertRuleGroupExportParams contains all the parameters to send to the API endpoint

	for the route get alert rule group export operation.

	Typically these are written to a http.Request.
*/
type RouteGetAlertRuleGroupExportParams struct {

	// FolderUID.
	FolderUID string

	// Group.
	Group string

	/* Download.

	   Whether to initiate a download of the file or not.
	*/
	Download *bool

	/* Format.

	   Format of the downloaded file, either yaml or json. Accept header can also be used, but the query parameter will take precedence.

	   Default: "yaml"
	*/
	Format *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route get alert rule group export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetAlertRuleGroupExportParams) WithDefaults() *RouteGetAlertRuleGroupExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route get alert rule group export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetAlertRuleGroupExportParams) SetDefaults() {
	var (
		downloadDefault = bool(false)

		formatDefault = string("yaml")
	)

	val := RouteGetAlertRuleGroupExportParams{
		Download: &downloadDefault,
		Format:   &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithTimeout(timeout time.Duration) *RouteGetAlertRuleGroupExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithContext(ctx context.Context) *RouteGetAlertRuleGroupExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithHTTPClient(client *http.Client) *RouteGetAlertRuleGroupExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderUID adds the folderUID to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithFolderUID(folderUID string) *RouteGetAlertRuleGroupExportParams {
	o.SetFolderUID(folderUID)
	return o
}

// SetFolderUID adds the folderUid to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetFolderUID(folderUID string) {
	o.FolderUID = folderUID
}

// WithGroup adds the group to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithGroup(group string) *RouteGetAlertRuleGroupExportParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetGroup(group string) {
	o.Group = group
}

// WithDownload adds the download to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithDownload(download *bool) *RouteGetAlertRuleGroupExportParams {
	o.SetDownload(download)
	return o
}

// SetDownload adds the download to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetDownload(download *bool) {
	o.Download = download
}

// WithFormat adds the format to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) WithFormat(format *string) *RouteGetAlertRuleGroupExportParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the route get alert rule group export params
func (o *RouteGetAlertRuleGroupExportParams) SetFormat(format *string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *RouteGetAlertRuleGroupExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FolderUID
	if err := r.SetPathParam("FolderUID", o.FolderUID); err != nil {
		return err
	}

	// path param Group
	if err := r.SetPathParam("Group", o.Group); err != nil {
		return err
	}

	if o.Download != nil {

		// query param download
		var qrDownload bool

		if o.Download != nil {
			qrDownload = *o.Download
		}
		qDownload := swag.FormatBool(qrDownload)
		if qDownload != "" {

			if err := r.SetQueryParam("download", qDownload); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
