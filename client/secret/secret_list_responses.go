// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// SecretListReader is a Reader for the SecretList structure.
type SecretListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSecretListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSecretListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretListOK creates a SecretListOK with default headers values
func NewSecretListOK() *SecretListOK {
	return &SecretListOK{}
}

/* SecretListOK describes a response with status code 200, with default header values.

no error
*/
type SecretListOK struct {
	Payload []*models.Secret
}

func (o *SecretListOK) Error() string {
	return fmt.Sprintf("[GET /secrets][%d] secretListOK  %+v", 200, o.Payload)
}
func (o *SecretListOK) GetPayload() []*models.Secret {
	return o.Payload
}

func (o *SecretListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretListInternalServerError creates a SecretListInternalServerError with default headers values
func NewSecretListInternalServerError() *SecretListInternalServerError {
	return &SecretListInternalServerError{}
}

/* SecretListInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SecretListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SecretListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /secrets][%d] secretListInternalServerError  %+v", 500, o.Payload)
}
func (o *SecretListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecretListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretListServiceUnavailable creates a SecretListServiceUnavailable with default headers values
func NewSecretListServiceUnavailable() *SecretListServiceUnavailable {
	return &SecretListServiceUnavailable{}
}

/* SecretListServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type SecretListServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *SecretListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /secrets][%d] secretListServiceUnavailable  %+v", 503, o.Payload)
}
func (o *SecretListServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecretListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
