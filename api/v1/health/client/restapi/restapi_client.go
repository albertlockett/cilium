// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new restapi API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for restapi API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHealthz gets health of cilium node

Returns health and status information of the local node including
load and uptime, as well as the status of related components including
the Cilium daemon.

*/
func (a *Client) GetHealthz(params *GetHealthzParams) (*GetHealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthzParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHealthz",
		Method:             "GET",
		PathPattern:        "/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthzReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHealthzOK), nil

}

/*
GetHello says hello to cilium health

Returns a successful status code if this cilium-health instance is
reachable.

*/
func (a *Client) GetHello(params *GetHelloParams) (*GetHelloOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHelloParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHello",
		Method:             "GET",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHelloReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHelloOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
