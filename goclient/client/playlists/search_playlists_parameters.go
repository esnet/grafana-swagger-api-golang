// Code generated by go-swagger; DO NOT EDIT.

package playlists

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

// NewSearchPlaylistsParams creates a new SearchPlaylistsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchPlaylistsParams() *SearchPlaylistsParams {
	return &SearchPlaylistsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchPlaylistsParamsWithTimeout creates a new SearchPlaylistsParams object
// with the ability to set a timeout on a request.
func NewSearchPlaylistsParamsWithTimeout(timeout time.Duration) *SearchPlaylistsParams {
	return &SearchPlaylistsParams{
		timeout: timeout,
	}
}

// NewSearchPlaylistsParamsWithContext creates a new SearchPlaylistsParams object
// with the ability to set a context for a request.
func NewSearchPlaylistsParamsWithContext(ctx context.Context) *SearchPlaylistsParams {
	return &SearchPlaylistsParams{
		Context: ctx,
	}
}

// NewSearchPlaylistsParamsWithHTTPClient creates a new SearchPlaylistsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchPlaylistsParamsWithHTTPClient(client *http.Client) *SearchPlaylistsParams {
	return &SearchPlaylistsParams{
		HTTPClient: client,
	}
}

/*
SearchPlaylistsParams contains all the parameters to send to the API endpoint

	for the search playlists operation.

	Typically these are written to a http.Request.
*/
type SearchPlaylistsParams struct {

	/* Limit.

	   in:limit

	   Format: int64
	*/
	Limit *int64

	// Query.
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search playlists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchPlaylistsParams) WithDefaults() *SearchPlaylistsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search playlists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchPlaylistsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search playlists params
func (o *SearchPlaylistsParams) WithTimeout(timeout time.Duration) *SearchPlaylistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search playlists params
func (o *SearchPlaylistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search playlists params
func (o *SearchPlaylistsParams) WithContext(ctx context.Context) *SearchPlaylistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search playlists params
func (o *SearchPlaylistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search playlists params
func (o *SearchPlaylistsParams) WithHTTPClient(client *http.Client) *SearchPlaylistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search playlists params
func (o *SearchPlaylistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the search playlists params
func (o *SearchPlaylistsParams) WithLimit(limit *int64) *SearchPlaylistsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search playlists params
func (o *SearchPlaylistsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithQuery adds the query to the search playlists params
func (o *SearchPlaylistsParams) WithQuery(query *string) *SearchPlaylistsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search playlists params
func (o *SearchPlaylistsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SearchPlaylistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
