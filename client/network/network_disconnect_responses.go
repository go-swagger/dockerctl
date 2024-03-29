// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/dockerctl/models"
)

// NetworkDisconnectReader is a Reader for the NetworkDisconnect structure.
type NetworkDisconnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkDisconnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkDisconnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNetworkDisconnectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNetworkDisconnectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkDisconnectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkDisconnectOK creates a NetworkDisconnectOK with default headers values
func NewNetworkDisconnectOK() *NetworkDisconnectOK {
	return &NetworkDisconnectOK{}
}

/* NetworkDisconnectOK describes a response with status code 200, with default header values.

No error
*/
type NetworkDisconnectOK struct {
}

// IsSuccess returns true when this network disconnect o k response has a 2xx status code
func (o *NetworkDisconnectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network disconnect o k response has a 3xx status code
func (o *NetworkDisconnectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network disconnect o k response has a 4xx status code
func (o *NetworkDisconnectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network disconnect o k response has a 5xx status code
func (o *NetworkDisconnectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network disconnect o k response a status code equal to that given
func (o *NetworkDisconnectOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkDisconnectOK) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectOK ", 200)
}

func (o *NetworkDisconnectOK) String() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectOK ", 200)
}

func (o *NetworkDisconnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkDisconnectForbidden creates a NetworkDisconnectForbidden with default headers values
func NewNetworkDisconnectForbidden() *NetworkDisconnectForbidden {
	return &NetworkDisconnectForbidden{}
}

/* NetworkDisconnectForbidden describes a response with status code 403, with default header values.

Operation not supported for swarm scoped networks
*/
type NetworkDisconnectForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network disconnect forbidden response has a 2xx status code
func (o *NetworkDisconnectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network disconnect forbidden response has a 3xx status code
func (o *NetworkDisconnectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network disconnect forbidden response has a 4xx status code
func (o *NetworkDisconnectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this network disconnect forbidden response has a 5xx status code
func (o *NetworkDisconnectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this network disconnect forbidden response a status code equal to that given
func (o *NetworkDisconnectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NetworkDisconnectForbidden) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectForbidden  %+v", 403, o.Payload)
}

func (o *NetworkDisconnectForbidden) String() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectForbidden  %+v", 403, o.Payload)
}

func (o *NetworkDisconnectForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkDisconnectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkDisconnectNotFound creates a NetworkDisconnectNotFound with default headers values
func NewNetworkDisconnectNotFound() *NetworkDisconnectNotFound {
	return &NetworkDisconnectNotFound{}
}

/* NetworkDisconnectNotFound describes a response with status code 404, with default header values.

Network or container not found
*/
type NetworkDisconnectNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network disconnect not found response has a 2xx status code
func (o *NetworkDisconnectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network disconnect not found response has a 3xx status code
func (o *NetworkDisconnectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network disconnect not found response has a 4xx status code
func (o *NetworkDisconnectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this network disconnect not found response has a 5xx status code
func (o *NetworkDisconnectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this network disconnect not found response a status code equal to that given
func (o *NetworkDisconnectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NetworkDisconnectNotFound) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectNotFound  %+v", 404, o.Payload)
}

func (o *NetworkDisconnectNotFound) String() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectNotFound  %+v", 404, o.Payload)
}

func (o *NetworkDisconnectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkDisconnectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkDisconnectInternalServerError creates a NetworkDisconnectInternalServerError with default headers values
func NewNetworkDisconnectInternalServerError() *NetworkDisconnectInternalServerError {
	return &NetworkDisconnectInternalServerError{}
}

/* NetworkDisconnectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type NetworkDisconnectInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network disconnect internal server error response has a 2xx status code
func (o *NetworkDisconnectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network disconnect internal server error response has a 3xx status code
func (o *NetworkDisconnectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network disconnect internal server error response has a 4xx status code
func (o *NetworkDisconnectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network disconnect internal server error response has a 5xx status code
func (o *NetworkDisconnectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network disconnect internal server error response a status code equal to that given
func (o *NetworkDisconnectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkDisconnectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkDisconnectInternalServerError) String() string {
	return fmt.Sprintf("[POST /networks/{id}/disconnect][%d] networkDisconnectInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkDisconnectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkDisconnectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NetworkDisconnectBody network disconnect body
swagger:model NetworkDisconnectBody
*/
type NetworkDisconnectBody struct {

	// The ID or name of the container to disconnect from the network.
	Container string `json:"Container,omitempty"`

	// Force the container to disconnect from the network.
	Force bool `json:"Force,omitempty"`
}

// Validate validates this network disconnect body
func (o *NetworkDisconnectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network disconnect body based on context it is used
func (o *NetworkDisconnectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDisconnectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDisconnectBody) UnmarshalBinary(b []byte) error {
	var res NetworkDisconnectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
