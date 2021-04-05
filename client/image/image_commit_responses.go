// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// ImageCommitReader is a Reader for the ImageCommit structure.
type ImageCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImageCommitCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageCommitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageCommitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageCommitCreated creates a ImageCommitCreated with default headers values
func NewImageCommitCreated() *ImageCommitCreated {
	return &ImageCommitCreated{}
}

/* ImageCommitCreated describes a response with status code 201, with default header values.

no error
*/
type ImageCommitCreated struct {
	Payload *models.IDResponse
}

func (o *ImageCommitCreated) Error() string {
	return fmt.Sprintf("[POST /commit][%d] imageCommitCreated  %+v", 201, o.Payload)
}
func (o *ImageCommitCreated) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *ImageCommitCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageCommitNotFound creates a ImageCommitNotFound with default headers values
func NewImageCommitNotFound() *ImageCommitNotFound {
	return &ImageCommitNotFound{}
}

/* ImageCommitNotFound describes a response with status code 404, with default header values.

no such container
*/
type ImageCommitNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ImageCommitNotFound) Error() string {
	return fmt.Sprintf("[POST /commit][%d] imageCommitNotFound  %+v", 404, o.Payload)
}
func (o *ImageCommitNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageCommitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageCommitInternalServerError creates a ImageCommitInternalServerError with default headers values
func NewImageCommitInternalServerError() *ImageCommitInternalServerError {
	return &ImageCommitInternalServerError{}
}

/* ImageCommitInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ImageCommitInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageCommitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /commit][%d] imageCommitInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageCommitInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageCommitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
