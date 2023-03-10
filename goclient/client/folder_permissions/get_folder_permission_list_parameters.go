// Code generated by go-swagger; DO NOT EDIT.

package folder_permissions

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

// NewGetFolderPermissionListParams creates a new GetFolderPermissionListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFolderPermissionListParams() *GetFolderPermissionListParams {
	return &GetFolderPermissionListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFolderPermissionListParamsWithTimeout creates a new GetFolderPermissionListParams object
// with the ability to set a timeout on a request.
func NewGetFolderPermissionListParamsWithTimeout(timeout time.Duration) *GetFolderPermissionListParams {
	return &GetFolderPermissionListParams{
		timeout: timeout,
	}
}

// NewGetFolderPermissionListParamsWithContext creates a new GetFolderPermissionListParams object
// with the ability to set a context for a request.
func NewGetFolderPermissionListParamsWithContext(ctx context.Context) *GetFolderPermissionListParams {
	return &GetFolderPermissionListParams{
		Context: ctx,
	}
}

// NewGetFolderPermissionListParamsWithHTTPClient creates a new GetFolderPermissionListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFolderPermissionListParamsWithHTTPClient(client *http.Client) *GetFolderPermissionListParams {
	return &GetFolderPermissionListParams{
		HTTPClient: client,
	}
}

/*
GetFolderPermissionListParams contains all the parameters to send to the API endpoint

	for the get folder permission list operation.

	Typically these are written to a http.Request.
*/
type GetFolderPermissionListParams struct {

	// FolderUID.
	FolderUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get folder permission list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFolderPermissionListParams) WithDefaults() *GetFolderPermissionListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get folder permission list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFolderPermissionListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get folder permission list params
func (o *GetFolderPermissionListParams) WithTimeout(timeout time.Duration) *GetFolderPermissionListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get folder permission list params
func (o *GetFolderPermissionListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get folder permission list params
func (o *GetFolderPermissionListParams) WithContext(ctx context.Context) *GetFolderPermissionListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get folder permission list params
func (o *GetFolderPermissionListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get folder permission list params
func (o *GetFolderPermissionListParams) WithHTTPClient(client *http.Client) *GetFolderPermissionListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get folder permission list params
func (o *GetFolderPermissionListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderUID adds the folderUID to the get folder permission list params
func (o *GetFolderPermissionListParams) WithFolderUID(folderUID string) *GetFolderPermissionListParams {
	o.SetFolderUID(folderUID)
	return o
}

// SetFolderUID adds the folderUid to the get folder permission list params
func (o *GetFolderPermissionListParams) SetFolderUID(folderUID string) {
	o.FolderUID = folderUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFolderPermissionListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param folder_uid
	if err := r.SetPathParam("folder_uid", o.FolderUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
