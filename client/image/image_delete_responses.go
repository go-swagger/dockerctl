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

// ImageDeleteReader is a Reader for the ImageDelete structure.
type ImageDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageDeleteOK creates a ImageDeleteOK with default headers values
func NewImageDeleteOK() *ImageDeleteOK {
	return &ImageDeleteOK{}
}

/* ImageDeleteOK describes a response with status code 200, with default header values.

The image was deleted successfully
*/
type ImageDeleteOK struct {
	Payload []*models.ImageDeleteResponseItem
}

// IsSuccess returns true when this image delete o k response has a 2xx status code
func (o *ImageDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image delete o k response has a 3xx status code
func (o *ImageDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image delete o k response has a 4xx status code
func (o *ImageDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image delete o k response has a 5xx status code
func (o *ImageDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image delete o k response a status code equal to that given
func (o *ImageDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteOK  %+v", 200, o.Payload)
}

func (o *ImageDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteOK  %+v", 200, o.Payload)
}

func (o *ImageDeleteOK) GetPayload() []*models.ImageDeleteResponseItem {
	return o.Payload
}

func (o *ImageDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteNotFound creates a ImageDeleteNotFound with default headers values
func NewImageDeleteNotFound() *ImageDeleteNotFound {
	return &ImageDeleteNotFound{}
}

/* ImageDeleteNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageDeleteNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this image delete not found response has a 2xx status code
func (o *ImageDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image delete not found response has a 3xx status code
func (o *ImageDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image delete not found response has a 4xx status code
func (o *ImageDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image delete not found response has a 5xx status code
func (o *ImageDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image delete not found response a status code equal to that given
func (o *ImageDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImageDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ImageDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ImageDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteConflict creates a ImageDeleteConflict with default headers values
func NewImageDeleteConflict() *ImageDeleteConflict {
	return &ImageDeleteConflict{}
}

/* ImageDeleteConflict describes a response with status code 409, with default header values.

Conflict
*/
type ImageDeleteConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this image delete conflict response has a 2xx status code
func (o *ImageDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image delete conflict response has a 3xx status code
func (o *ImageDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image delete conflict response has a 4xx status code
func (o *ImageDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image delete conflict response has a 5xx status code
func (o *ImageDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image delete conflict response a status code equal to that given
func (o *ImageDeleteConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ImageDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteConflict  %+v", 409, o.Payload)
}

func (o *ImageDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteConflict  %+v", 409, o.Payload)
}

func (o *ImageDeleteConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageDeleteInternalServerError creates a ImageDeleteInternalServerError with default headers values
func NewImageDeleteInternalServerError() *ImageDeleteInternalServerError {
	return &ImageDeleteInternalServerError{}
}

/* ImageDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImageDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this image delete internal server error response has a 2xx status code
func (o *ImageDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image delete internal server error response has a 3xx status code
func (o *ImageDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image delete internal server error response has a 4xx status code
func (o *ImageDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image delete internal server error response has a 5xx status code
func (o *ImageDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image delete internal server error response a status code equal to that given
func (o *ImageDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /images/{name}][%d] imageDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
