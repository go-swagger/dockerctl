// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// SwarmLeaveReader is a Reader for the SwarmLeave structure.
type SwarmLeaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwarmLeaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwarmLeaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSwarmLeaveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSwarmLeaveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSwarmLeaveOK creates a SwarmLeaveOK with default headers values
func NewSwarmLeaveOK() *SwarmLeaveOK {
	return &SwarmLeaveOK{}
}

/* SwarmLeaveOK describes a response with status code 200, with default header values.

no error
*/
type SwarmLeaveOK struct {
}

// IsSuccess returns true when this swarm leave o k response has a 2xx status code
func (o *SwarmLeaveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this swarm leave o k response has a 3xx status code
func (o *SwarmLeaveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swarm leave o k response has a 4xx status code
func (o *SwarmLeaveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this swarm leave o k response has a 5xx status code
func (o *SwarmLeaveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this swarm leave o k response a status code equal to that given
func (o *SwarmLeaveOK) IsCode(code int) bool {
	return code == 200
}

func (o *SwarmLeaveOK) Error() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveOK ", 200)
}

func (o *SwarmLeaveOK) String() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveOK ", 200)
}

func (o *SwarmLeaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSwarmLeaveInternalServerError creates a SwarmLeaveInternalServerError with default headers values
func NewSwarmLeaveInternalServerError() *SwarmLeaveInternalServerError {
	return &SwarmLeaveInternalServerError{}
}

/* SwarmLeaveInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SwarmLeaveInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this swarm leave internal server error response has a 2xx status code
func (o *SwarmLeaveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this swarm leave internal server error response has a 3xx status code
func (o *SwarmLeaveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swarm leave internal server error response has a 4xx status code
func (o *SwarmLeaveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this swarm leave internal server error response has a 5xx status code
func (o *SwarmLeaveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this swarm leave internal server error response a status code equal to that given
func (o *SwarmLeaveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SwarmLeaveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveInternalServerError  %+v", 500, o.Payload)
}

func (o *SwarmLeaveInternalServerError) String() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveInternalServerError  %+v", 500, o.Payload)
}

func (o *SwarmLeaveInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwarmLeaveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwarmLeaveServiceUnavailable creates a SwarmLeaveServiceUnavailable with default headers values
func NewSwarmLeaveServiceUnavailable() *SwarmLeaveServiceUnavailable {
	return &SwarmLeaveServiceUnavailable{}
}

/* SwarmLeaveServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type SwarmLeaveServiceUnavailable struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this swarm leave service unavailable response has a 2xx status code
func (o *SwarmLeaveServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this swarm leave service unavailable response has a 3xx status code
func (o *SwarmLeaveServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swarm leave service unavailable response has a 4xx status code
func (o *SwarmLeaveServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this swarm leave service unavailable response has a 5xx status code
func (o *SwarmLeaveServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this swarm leave service unavailable response a status code equal to that given
func (o *SwarmLeaveServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SwarmLeaveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SwarmLeaveServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /swarm/leave][%d] swarmLeaveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SwarmLeaveServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwarmLeaveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
