// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new playlists API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for playlists API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePlaylist(params *CreatePlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePlaylistOK, error)

	DeletePlaylist(params *DeletePlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePlaylistOK, error)

	GetPlaylist(params *GetPlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlaylistOK, error)

	GetPlaylistDashboards(params *GetPlaylistDashboardsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlaylistDashboardsOK, error)

	GetPlaylistItems(params *GetPlaylistItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlaylistItemsOK, error)

	SearchPlaylists(params *SearchPlaylistsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchPlaylistsOK, error)

	UpdatePlaylist(params *UpdatePlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePlaylistOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePlaylist creates playlist
*/
func (a *Client) CreatePlaylist(params *CreatePlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePlaylistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePlaylistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPlaylist",
		Method:             "POST",
		PathPattern:        "/playlists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePlaylistReader{formats: a.formats},
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
	success, ok := result.(*CreatePlaylistOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPlaylist: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePlaylist deletes playlist
*/
func (a *Client) DeletePlaylist(params *DeletePlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePlaylistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePlaylistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePlaylist",
		Method:             "DELETE",
		PathPattern:        "/playlists/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePlaylistReader{formats: a.formats},
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
	success, ok := result.(*DeletePlaylistOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePlaylist: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPlaylist gets playlist
*/
func (a *Client) GetPlaylist(params *GetPlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlaylistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlaylistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPlaylist",
		Method:             "GET",
		PathPattern:        "/playlists/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlaylistReader{formats: a.formats},
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
	success, ok := result.(*GetPlaylistOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlaylist: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPlaylistDashboards gets playlist dashboards
*/
func (a *Client) GetPlaylistDashboards(params *GetPlaylistDashboardsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlaylistDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlaylistDashboardsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPlaylistDashboards",
		Method:             "GET",
		PathPattern:        "/playlists/{uid}/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlaylistDashboardsReader{formats: a.formats},
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
	success, ok := result.(*GetPlaylistDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlaylistDashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPlaylistItems gets playlist items
*/
func (a *Client) GetPlaylistItems(params *GetPlaylistItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPlaylistItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlaylistItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPlaylistItems",
		Method:             "GET",
		PathPattern:        "/playlists/{uid}/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlaylistItemsReader{formats: a.formats},
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
	success, ok := result.(*GetPlaylistItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlaylistItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchPlaylists gets playlists
*/
func (a *Client) SearchPlaylists(params *SearchPlaylistsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchPlaylistsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchPlaylistsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchPlaylists",
		Method:             "GET",
		PathPattern:        "/playlists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchPlaylistsReader{formats: a.formats},
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
	success, ok := result.(*SearchPlaylistsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchPlaylists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePlaylist updates playlist
*/
func (a *Client) UpdatePlaylist(params *UpdatePlaylistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePlaylistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePlaylistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePlaylist",
		Method:             "PUT",
		PathPattern:        "/playlists/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdatePlaylistReader{formats: a.formats},
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
	success, ok := result.(*UpdatePlaylistOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePlaylist: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}