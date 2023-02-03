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

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// NewPostGraphiteAnnotationParams creates a new PostGraphiteAnnotationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGraphiteAnnotationParams() *PostGraphiteAnnotationParams {
	return &PostGraphiteAnnotationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGraphiteAnnotationParamsWithTimeout creates a new PostGraphiteAnnotationParams object
// with the ability to set a timeout on a request.
func NewPostGraphiteAnnotationParamsWithTimeout(timeout time.Duration) *PostGraphiteAnnotationParams {
	return &PostGraphiteAnnotationParams{
		timeout: timeout,
	}
}

// NewPostGraphiteAnnotationParamsWithContext creates a new PostGraphiteAnnotationParams object
// with the ability to set a context for a request.
func NewPostGraphiteAnnotationParamsWithContext(ctx context.Context) *PostGraphiteAnnotationParams {
	return &PostGraphiteAnnotationParams{
		Context: ctx,
	}
}

// NewPostGraphiteAnnotationParamsWithHTTPClient creates a new PostGraphiteAnnotationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGraphiteAnnotationParamsWithHTTPClient(client *http.Client) *PostGraphiteAnnotationParams {
	return &PostGraphiteAnnotationParams{
		HTTPClient: client,
	}
}

/*
PostGraphiteAnnotationParams contains all the parameters to send to the API endpoint

	for the post graphite annotation operation.

	Typically these are written to a http.Request.
*/
type PostGraphiteAnnotationParams struct {

	// Body.
	Body *models.PostGraphiteAnnotationsCmd

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post graphite annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGraphiteAnnotationParams) WithDefaults() *PostGraphiteAnnotationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post graphite annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGraphiteAnnotationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) WithTimeout(timeout time.Duration) *PostGraphiteAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) WithContext(ctx context.Context) *PostGraphiteAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) WithHTTPClient(client *http.Client) *PostGraphiteAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) WithBody(body *models.PostGraphiteAnnotationsCmd) *PostGraphiteAnnotationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post graphite annotation params
func (o *PostGraphiteAnnotationParams) SetBody(body *models.PostGraphiteAnnotationsCmd) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostGraphiteAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
