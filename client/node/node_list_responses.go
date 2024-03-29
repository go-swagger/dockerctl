// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// NodeListReader is a Reader for the NodeList structure.
type NodeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewNodeListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNodeListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodeListOK creates a NodeListOK with default headers values
func NewNodeListOK() *NodeListOK {
	return &NodeListOK{}
}

/* NodeListOK describes a response with status code 200, with default header values.

no error
*/
type NodeListOK struct {
	Payload []*models.Node
}

// IsSuccess returns true when this node list o k response has a 2xx status code
func (o *NodeListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node list o k response has a 3xx status code
func (o *NodeListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node list o k response has a 4xx status code
func (o *NodeListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node list o k response has a 5xx status code
func (o *NodeListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node list o k response a status code equal to that given
func (o *NodeListOK) IsCode(code int) bool {
	return code == 200
}

func (o *NodeListOK) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListOK  %+v", 200, o.Payload)
}

func (o *NodeListOK) String() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListOK  %+v", 200, o.Payload)
}

func (o *NodeListOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *NodeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeListInternalServerError creates a NodeListInternalServerError with default headers values
func NewNodeListInternalServerError() *NodeListInternalServerError {
	return &NodeListInternalServerError{}
}

/* NodeListInternalServerError describes a response with status code 500, with default header values.

server error
*/
type NodeListInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node list internal server error response has a 2xx status code
func (o *NodeListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node list internal server error response has a 3xx status code
func (o *NodeListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node list internal server error response has a 4xx status code
func (o *NodeListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this node list internal server error response has a 5xx status code
func (o *NodeListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this node list internal server error response a status code equal to that given
func (o *NodeListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NodeListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListInternalServerError  %+v", 500, o.Payload)
}

func (o *NodeListInternalServerError) String() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListInternalServerError  %+v", 500, o.Payload)
}

func (o *NodeListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeListServiceUnavailable creates a NodeListServiceUnavailable with default headers values
func NewNodeListServiceUnavailable() *NodeListServiceUnavailable {
	return &NodeListServiceUnavailable{}
}

/* NodeListServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type NodeListServiceUnavailable struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node list service unavailable response has a 2xx status code
func (o *NodeListServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node list service unavailable response has a 3xx status code
func (o *NodeListServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node list service unavailable response has a 4xx status code
func (o *NodeListServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this node list service unavailable response has a 5xx status code
func (o *NodeListServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this node list service unavailable response a status code equal to that given
func (o *NodeListServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *NodeListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListServiceUnavailable  %+v", 503, o.Payload)
}

func (o *NodeListServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListServiceUnavailable  %+v", 503, o.Payload)
}

func (o *NodeListServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
