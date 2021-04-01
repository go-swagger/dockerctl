// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new session API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for session API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Session(params *SessionParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  Session initializes interactive session

  Start a new interactive session with a server. Session allows server to call back to the client for advanced capabilities.

### Hijacking

This endpoint hijacks the HTTP connection to HTTP2 transport that allows the client to expose gPRC services on that connection.

For example, the client sends this request to upgrade the connection:

```
POST /session HTTP/1.1
Upgrade: h2c
Connection: Upgrade
```

The Docker daemon will respond with a `101 UPGRADED` response follow with the raw stream:

```
HTTP/1.1 101 UPGRADED
Connection: Upgrade
Upgrade: h2c
```

*/
func (a *Client) Session(params *SessionParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Session",
		Method:             "POST",
		PathPattern:        "/session",
		ProducesMediaTypes: []string{"application/vnd.docker.raw-stream"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
