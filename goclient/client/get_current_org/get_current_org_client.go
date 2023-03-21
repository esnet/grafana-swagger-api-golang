// Code generated by go-swagger; DO NOT EDIT.

package get_current_org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new get current org API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for get current org API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCurrentOrgQuota(params *GetCurrentOrgQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentOrgQuotaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetCurrentOrgQuota fetches organization quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `orgs.quotas:read` and scope `org:id:1` (orgIDScope).
*/
func (a *Client) GetCurrentOrgQuota(params *GetCurrentOrgQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentOrgQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentOrgQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCurrentOrgQuota",
		Method:             "GET",
		PathPattern:        "/org/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCurrentOrgQuotaReader{formats: a.formats},
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
	success, ok := result.(*GetCurrentOrgQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentOrgQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
