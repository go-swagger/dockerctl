// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new node API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NodeDelete(params *NodeDeleteParams, opts ...ClientOption) (*NodeDeleteOK, error)

	NodeInspect(params *NodeInspectParams, opts ...ClientOption) (*NodeInspectOK, error)

	NodeList(params *NodeListParams, opts ...ClientOption) (*NodeListOK, error)

	NodeUpdate(params *NodeUpdateParams, opts ...ClientOption) (*NodeUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  NodeDelete deletes a node
*/
func (a *Client) NodeDelete(params *NodeDeleteParams, opts ...ClientOption) (*NodeDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeDelete",
		Method:             "DELETE",
		PathPattern:        "/nodes/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodeDeleteReader{formats: a.formats},
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
	success, ok := result.(*NodeDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NodeDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NodeInspect inspects a node
*/
func (a *Client) NodeInspect(params *NodeInspectParams, opts ...ClientOption) (*NodeInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeInspect",
		Method:             "GET",
		PathPattern:        "/nodes/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodeInspectReader{formats: a.formats},
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
	success, ok := result.(*NodeInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NodeInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NodeList lists nodes
*/
func (a *Client) NodeList(params *NodeListParams, opts ...ClientOption) (*NodeListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeList",
		Method:             "GET",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodeListReader{formats: a.formats},
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
	success, ok := result.(*NodeListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NodeList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NodeUpdate updates a node
*/
func (a *Client) NodeUpdate(params *NodeUpdateParams, opts ...ClientOption) (*NodeUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeUpdate",
		Method:             "POST",
		PathPattern:        "/nodes/{id}/update",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodeUpdateReader{formats: a.formats},
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
	success, ok := result.(*NodeUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NodeUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
