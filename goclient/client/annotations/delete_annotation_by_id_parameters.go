// Code generated by go-swagger; DO NOT EDIT.

package annotations

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

// NewDeleteAnnotationByIDParams creates a new DeleteAnnotationByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAnnotationByIDParams() *DeleteAnnotationByIDParams {
	return &DeleteAnnotationByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAnnotationByIDParamsWithTimeout creates a new DeleteAnnotationByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAnnotationByIDParamsWithTimeout(timeout time.Duration) *DeleteAnnotationByIDParams {
	return &DeleteAnnotationByIDParams{
		timeout: timeout,
	}
}

// NewDeleteAnnotationByIDParamsWithContext creates a new DeleteAnnotationByIDParams object
// with the ability to set a context for a request.
func NewDeleteAnnotationByIDParamsWithContext(ctx context.Context) *DeleteAnnotationByIDParams {
	return &DeleteAnnotationByIDParams{
		Context: ctx,
	}
}

// NewDeleteAnnotationByIDParamsWithHTTPClient creates a new DeleteAnnotationByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAnnotationByIDParamsWithHTTPClient(client *http.Client) *DeleteAnnotationByIDParams {
	return &DeleteAnnotationByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAnnotationByIDParams contains all the parameters to send to the API endpoint

	for the delete annotation by ID operation.

	Typically these are written to a http.Request.
*/
type DeleteAnnotationByIDParams struct {

	// AnnotationID.
	AnnotationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete annotation by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAnnotationByIDParams) WithDefaults() *DeleteAnnotationByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete annotation by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAnnotationByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) WithTimeout(timeout time.Duration) *DeleteAnnotationByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) WithContext(ctx context.Context) *DeleteAnnotationByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) WithHTTPClient(client *http.Client) *DeleteAnnotationByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnnotationID adds the annotationID to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) WithAnnotationID(annotationID string) *DeleteAnnotationByIDParams {
	o.SetAnnotationID(annotationID)
	return o
}

// SetAnnotationID adds the annotationId to the delete annotation by ID params
func (o *DeleteAnnotationByIDParams) SetAnnotationID(annotationID string) {
	o.AnnotationID = annotationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAnnotationByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param annotation_id
	if err := r.SetPathParam("annotation_id", o.AnnotationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
