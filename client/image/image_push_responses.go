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

// ImagePushReader is a Reader for the ImagePush structure.
type ImagePushReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagePushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagePushOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImagePushNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagePushInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagePushOK creates a ImagePushOK with default headers values
func NewImagePushOK() *ImagePushOK {
	return &ImagePushOK{}
}

/* ImagePushOK describes a response with status code 200, with default header values.

No error
*/
type ImagePushOK struct {
}

// IsSuccess returns true when this image push o k response has a 2xx status code
func (o *ImagePushOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image push o k response has a 3xx status code
func (o *ImagePushOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image push o k response has a 4xx status code
func (o *ImagePushOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image push o k response has a 5xx status code
func (o *ImagePushOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image push o k response a status code equal to that given
func (o *ImagePushOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagePushOK) Error() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushOK ", 200)
}

func (o *ImagePushOK) String() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushOK ", 200)
}

func (o *ImagePushOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImagePushNotFound creates a ImagePushNotFound with default headers values
func NewImagePushNotFound() *ImagePushNotFound {
	return &ImagePushNotFound{}
}

/* ImagePushNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImagePushNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this image push not found response has a 2xx status code
func (o *ImagePushNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image push not found response has a 3xx status code
func (o *ImagePushNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image push not found response has a 4xx status code
func (o *ImagePushNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image push not found response has a 5xx status code
func (o *ImagePushNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image push not found response a status code equal to that given
func (o *ImagePushNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagePushNotFound) Error() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushNotFound  %+v", 404, o.Payload)
}

func (o *ImagePushNotFound) String() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushNotFound  %+v", 404, o.Payload)
}

func (o *ImagePushNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImagePushNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushInternalServerError creates a ImagePushInternalServerError with default headers values
func NewImagePushInternalServerError() *ImagePushInternalServerError {
	return &ImagePushInternalServerError{}
}

/* ImagePushInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImagePushInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this image push internal server error response has a 2xx status code
func (o *ImagePushInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image push internal server error response has a 3xx status code
func (o *ImagePushInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image push internal server error response has a 4xx status code
func (o *ImagePushInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image push internal server error response has a 5xx status code
func (o *ImagePushInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image push internal server error response a status code equal to that given
func (o *ImagePushInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagePushInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushInternalServerError  %+v", 500, o.Payload)
}

func (o *ImagePushInternalServerError) String() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushInternalServerError  %+v", 500, o.Payload)
}

func (o *ImagePushInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImagePushInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
