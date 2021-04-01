// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/dockerctl/models"
)

// ExecResizeReader is a Reader for the ExecResize structure.
type ExecResizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecResizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExecResizeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewExecResizeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecResizeCreated creates a ExecResizeCreated with default headers values
func NewExecResizeCreated() *ExecResizeCreated {
	return &ExecResizeCreated{}
}

/* ExecResizeCreated describes a response with status code 201, with default header values.

No error
*/
type ExecResizeCreated struct {
}

func (o *ExecResizeCreated) Error() string {
	return fmt.Sprintf("[POST /exec/{id}/resize][%d] execResizeCreated ", 201)
}

func (o *ExecResizeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecResizeNotFound creates a ExecResizeNotFound with default headers values
func NewExecResizeNotFound() *ExecResizeNotFound {
	return &ExecResizeNotFound{}
}

/* ExecResizeNotFound describes a response with status code 404, with default header values.

No such exec instance
*/
type ExecResizeNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ExecResizeNotFound) Error() string {
	return fmt.Sprintf("[POST /exec/{id}/resize][%d] execResizeNotFound  %+v", 404, o.Payload)
}
func (o *ExecResizeNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecResizeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
