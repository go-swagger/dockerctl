// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SystemAuth(params *SystemAuthParams, opts ...ClientOption) (*SystemAuthOK, *SystemAuthNoContent, error)

	SystemDataUsage(params *SystemDataUsageParams, opts ...ClientOption) (*SystemDataUsageOK, error)

	SystemEvents(params *SystemEventsParams, opts ...ClientOption) (*SystemEventsOK, error)

	SystemInfo(params *SystemInfoParams, opts ...ClientOption) (*SystemInfoOK, error)

	SystemPing(params *SystemPingParams, opts ...ClientOption) (*SystemPingOK, error)

	SystemPingHead(params *SystemPingHeadParams, opts ...ClientOption) (*SystemPingHeadOK, error)

	SystemVersion(params *SystemVersionParams, opts ...ClientOption) (*SystemVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SystemAuth checks auth configuration

  Validate credentials for a registry and, if available, get an identity token for accessing the registry without password.
*/
func (a *Client) SystemAuth(params *SystemAuthParams, opts ...ClientOption) (*SystemAuthOK, *SystemAuthNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemAuth",
		Method:             "POST",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SystemAuthOK:
		return value, nil, nil
	case *SystemAuthNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemDataUsage gets data usage information
*/
func (a *Client) SystemDataUsage(params *SystemDataUsageParams, opts ...ClientOption) (*SystemDataUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemDataUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemDataUsage",
		Method:             "GET",
		PathPattern:        "/system/df",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemDataUsageReader{formats: a.formats},
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
	success, ok := result.(*SystemDataUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemDataUsage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemEvents monitors events

  Stream real-time events from the server.

Various objects within Docker report events when something happens to them.

Containers report these events: `attach`, `commit`, `copy`, `create`, `destroy`, `detach`, `die`, `exec_create`, `exec_detach`, `exec_start`, `exec_die`, `export`, `health_status`, `kill`, `oom`, `pause`, `rename`, `resize`, `restart`, `start`, `stop`, `top`, `unpause`, and `update`

Images report these events: `delete`, `import`, `load`, `pull`, `push`, `save`, `tag`, and `untag`

Volumes report these events: `create`, `mount`, `unmount`, and `destroy`

Networks report these events: `create`, `connect`, `disconnect`, `destroy`, `update`, and `remove`

The Docker daemon reports these events: `reload`

Services report these events: `create`, `update`, and `remove`

Nodes report these events: `create`, `update`, and `remove`

Secrets report these events: `create`, `update`, and `remove`

Configs report these events: `create`, `update`, and `remove`

*/
func (a *Client) SystemEvents(params *SystemEventsParams, opts ...ClientOption) (*SystemEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemEvents",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemEventsReader{formats: a.formats},
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
	success, ok := result.(*SystemEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemInfo gets system information
*/
func (a *Client) SystemInfo(params *SystemInfoParams, opts ...ClientOption) (*SystemInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemInfo",
		Method:             "GET",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemInfoReader{formats: a.formats},
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
	success, ok := result.(*SystemInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemPing pings

  This is a dummy endpoint you can use to test if the server is accessible.
*/
func (a *Client) SystemPing(params *SystemPingParams, opts ...ClientOption) (*SystemPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemPing",
		Method:             "GET",
		PathPattern:        "/_ping",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemPingReader{formats: a.formats},
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
	success, ok := result.(*SystemPingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemPing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemPingHead pings

  This is a dummy endpoint you can use to test if the server is accessible.
*/
func (a *Client) SystemPingHead(params *SystemPingHeadParams, opts ...ClientOption) (*SystemPingHeadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemPingHeadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemPingHead",
		Method:             "HEAD",
		PathPattern:        "/_ping",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemPingHeadReader{formats: a.formats},
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
	success, ok := result.(*SystemPingHeadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemPingHead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemVersion gets version

  Returns the version of Docker that is running and various information about the system that Docker is running on.
*/
func (a *Client) SystemVersion(params *SystemVersionParams, opts ...ClientOption) (*SystemVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemVersionReader{formats: a.formats},
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
	success, ok := result.(*SystemVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
