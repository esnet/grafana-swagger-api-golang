// Code generated by go-swagger; DO NOT EDIT.

package annotations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new annotations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for annotations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAnnotationByID(params *DeleteAnnotationByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAnnotationByIDOK, error)

	GetAnnotationByID(params *GetAnnotationByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAnnotationByIDOK, error)

	GetAnnotationTags(params *GetAnnotationTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAnnotationTagsOK, error)

	GetAnnotations(params *GetAnnotationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAnnotationsOK, error)

	MassDeleteAnnotations(params *MassDeleteAnnotationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MassDeleteAnnotationsOK, error)

	PatchAnnotation(params *PatchAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAnnotationOK, error)

	PostAnnotation(params *PostAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAnnotationOK, error)

	PostGraphiteAnnotation(params *PostGraphiteAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostGraphiteAnnotationOK, error)

	UpdateAnnotation(params *UpdateAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAnnotationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAnnotationByID deletes annotation by ID

Deletes the annotation that matches the specified ID.
*/
func (a *Client) DeleteAnnotationByID(params *DeleteAnnotationByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAnnotationByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAnnotationByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAnnotationByID",
		Method:             "DELETE",
		PathPattern:        "/annotations/{annotation_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAnnotationByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAnnotationByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAnnotationByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAnnotationByID gets annotation by ID
*/
func (a *Client) GetAnnotationByID(params *GetAnnotationByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAnnotationByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAnnotationByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAnnotationByID",
		Method:             "GET",
		PathPattern:        "/annotations/{annotation_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAnnotationByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAnnotationByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAnnotationByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAnnotationTags finds annotations tags

Find all the event tags created in the annotations.
*/
func (a *Client) GetAnnotationTags(params *GetAnnotationTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAnnotationTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAnnotationTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAnnotationTags",
		Method:             "GET",
		PathPattern:        "/annotations/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAnnotationTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAnnotationTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAnnotationTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAnnotations finds annotations

Starting in Grafana v6.4 regions annotations are now returned in one entity that now includes the timeEnd property.
*/
func (a *Client) GetAnnotations(params *GetAnnotationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAnnotationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAnnotationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAnnotations",
		Method:             "GET",
		PathPattern:        "/annotations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAnnotationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAnnotationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAnnotations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
MassDeleteAnnotations deletes multiple annotations
*/
func (a *Client) MassDeleteAnnotations(params *MassDeleteAnnotationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MassDeleteAnnotationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMassDeleteAnnotationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "massDeleteAnnotations",
		Method:             "POST",
		PathPattern:        "/annotations/mass-delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MassDeleteAnnotationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MassDeleteAnnotationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for massDeleteAnnotations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PatchAnnotation patches annotation

	Updates one or more properties of an annotation that matches the specified ID.

This operation currently supports updating of the `text`, `tags`, `time` and `timeEnd` properties.
This is available in Grafana 6.0.0-beta2 and above.
*/
func (a *Client) PatchAnnotation(params *PatchAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAnnotationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAnnotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchAnnotation",
		Method:             "PATCH",
		PathPattern:        "/annotations/{annotation_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchAnnotationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAnnotationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchAnnotation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostAnnotation creates annotation

	Creates an annotation in the Grafana database. The dashboardId and panelId fields are optional. If they are not specified then an organization annotation is created and can be queried in any dashboard that adds the Grafana annotations data source. When creating a region annotation include the timeEnd property.

The format for `time` and `timeEnd` should be epoch numbers in millisecond resolution.
The response for this HTTP request is slightly different in versions prior to v6.4. In prior versions you would also get an endId if you where creating a region. But in 6.4 regions are represented using a single event with time and timeEnd properties.
*/
func (a *Client) PostAnnotation(params *PostAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAnnotationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAnnotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postAnnotation",
		Method:             "POST",
		PathPattern:        "/annotations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostAnnotationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAnnotationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postAnnotation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostGraphiteAnnotation creates annotation in graphite format

Creates an annotation by using Graphite-compatible event format. The `when` and `data` fields are optional. If `when` is not specified then the current time will be used as annotation’s timestamp. The `tags` field can also be in prior to Graphite `0.10.0` format (string with multiple tags being separated by a space).
*/
func (a *Client) PostGraphiteAnnotation(params *PostGraphiteAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostGraphiteAnnotationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostGraphiteAnnotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postGraphiteAnnotation",
		Method:             "POST",
		PathPattern:        "/annotations/graphite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostGraphiteAnnotationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostGraphiteAnnotationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postGraphiteAnnotation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAnnotation updates annotation

Updates all properties of an annotation that matches the specified id. To only update certain property, consider using the Patch Annotation operation.
*/
func (a *Client) UpdateAnnotation(params *UpdateAnnotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAnnotationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAnnotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAnnotation",
		Method:             "PUT",
		PathPattern:        "/annotations/{annotation_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAnnotationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAnnotationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAnnotation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
