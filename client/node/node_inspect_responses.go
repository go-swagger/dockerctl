// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/dockerctl/models"
)

// NodeInspectReader is a Reader for the NodeInspect structure.
type NodeInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewNodeInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodeInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNodeInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodeInspectOK creates a NodeInspectOK with default headers values
func NewNodeInspectOK() *NodeInspectOK {
	return &NodeInspectOK{}
}

/* NodeInspectOK describes a response with status code 200, with default header values.

no error
*/
type NodeInspectOK struct {
	Payload *models.Node
}

func (o *NodeInspectOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] nodeInspectOK  %+v", 200, o.Payload)
}
func (o *NodeInspectOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *NodeInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeInspectNotFound creates a NodeInspectNotFound with default headers values
func NewNodeInspectNotFound() *NodeInspectNotFound {
	return &NodeInspectNotFound{}
}

/* NodeInspectNotFound describes a response with status code 404, with default header values.

no such node
*/
type NodeInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *NodeInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] nodeInspectNotFound  %+v", 404, o.Payload)
}
func (o *NodeInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeInspectInternalServerError creates a NodeInspectInternalServerError with default headers values
func NewNodeInspectInternalServerError() *NodeInspectInternalServerError {
	return &NodeInspectInternalServerError{}
}

/* NodeInspectInternalServerError describes a response with status code 500, with default header values.

server error
*/
type NodeInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *NodeInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] nodeInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *NodeInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeInspectServiceUnavailable creates a NodeInspectServiceUnavailable with default headers values
func NewNodeInspectServiceUnavailable() *NodeInspectServiceUnavailable {
	return &NodeInspectServiceUnavailable{}
}

/* NodeInspectServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type NodeInspectServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *NodeInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /nodes/{id}][%d] nodeInspectServiceUnavailable  %+v", 503, o.Payload)
}
func (o *NodeInspectServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
