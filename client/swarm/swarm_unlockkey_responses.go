// Code generated by go-swagger; DO NOT EDIT.

package swarm

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

// SwarmUnlockkeyReader is a Reader for the SwarmUnlockkey structure.
type SwarmUnlockkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwarmUnlockkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwarmUnlockkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSwarmUnlockkeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSwarmUnlockkeyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSwarmUnlockkeyOK creates a SwarmUnlockkeyOK with default headers values
func NewSwarmUnlockkeyOK() *SwarmUnlockkeyOK {
	return &SwarmUnlockkeyOK{}
}

/* SwarmUnlockkeyOK describes a response with status code 200, with default header values.

no error
*/
type SwarmUnlockkeyOK struct {
	Payload *SwarmUnlockkeyOKBody
}

// IsSuccess returns true when this swarm unlockkey o k response has a 2xx status code
func (o *SwarmUnlockkeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this swarm unlockkey o k response has a 3xx status code
func (o *SwarmUnlockkeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swarm unlockkey o k response has a 4xx status code
func (o *SwarmUnlockkeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this swarm unlockkey o k response has a 5xx status code
func (o *SwarmUnlockkeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this swarm unlockkey o k response a status code equal to that given
func (o *SwarmUnlockkeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *SwarmUnlockkeyOK) Error() string {
	return fmt.Sprintf("[GET /swarm/unlockkey][%d] swarmUnlockkeyOK  %+v", 200, o.Payload)
}

func (o *SwarmUnlockkeyOK) String() string {
	return fmt.Sprintf("[GET /swarm/unlockkey][%d] swarmUnlockkeyOK  %+v", 200, o.Payload)
}

func (o *SwarmUnlockkeyOK) GetPayload() *SwarmUnlockkeyOKBody {
	return o.Payload
}

func (o *SwarmUnlockkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SwarmUnlockkeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwarmUnlockkeyInternalServerError creates a SwarmUnlockkeyInternalServerError with default headers values
func NewSwarmUnlockkeyInternalServerError() *SwarmUnlockkeyInternalServerError {
	return &SwarmUnlockkeyInternalServerError{}
}

/* SwarmUnlockkeyInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SwarmUnlockkeyInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this swarm unlockkey internal server error response has a 2xx status code
func (o *SwarmUnlockkeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this swarm unlockkey internal server error response has a 3xx status code
func (o *SwarmUnlockkeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swarm unlockkey internal server error response has a 4xx status code
func (o *SwarmUnlockkeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this swarm unlockkey internal server error response has a 5xx status code
func (o *SwarmUnlockkeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this swarm unlockkey internal server error response a status code equal to that given
func (o *SwarmUnlockkeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SwarmUnlockkeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /swarm/unlockkey][%d] swarmUnlockkeyInternalServerError  %+v", 500, o.Payload)
}

func (o *SwarmUnlockkeyInternalServerError) String() string {
	return fmt.Sprintf("[GET /swarm/unlockkey][%d] swarmUnlockkeyInternalServerError  %+v", 500, o.Payload)
}

func (o *SwarmUnlockkeyInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwarmUnlockkeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwarmUnlockkeyServiceUnavailable creates a SwarmUnlockkeyServiceUnavailable with default headers values
func NewSwarmUnlockkeyServiceUnavailable() *SwarmUnlockkeyServiceUnavailable {
	return &SwarmUnlockkeyServiceUnavailable{}
}

/* SwarmUnlockkeyServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type SwarmUnlockkeyServiceUnavailable struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this swarm unlockkey service unavailable response has a 2xx status code
func (o *SwarmUnlockkeyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this swarm unlockkey service unavailable response has a 3xx status code
func (o *SwarmUnlockkeyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swarm unlockkey service unavailable response has a 4xx status code
func (o *SwarmUnlockkeyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this swarm unlockkey service unavailable response has a 5xx status code
func (o *SwarmUnlockkeyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this swarm unlockkey service unavailable response a status code equal to that given
func (o *SwarmUnlockkeyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SwarmUnlockkeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /swarm/unlockkey][%d] swarmUnlockkeyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SwarmUnlockkeyServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /swarm/unlockkey][%d] swarmUnlockkeyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SwarmUnlockkeyServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwarmUnlockkeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SwarmUnlockkeyOKBody UnlockKeyResponse
// Example: {"UnlockKey":"SWMKEY-1-7c37Cc8654o6p38HnroywCi19pllOnGtbdZEgtKxZu8"}
swagger:model SwarmUnlockkeyOKBody
*/
type SwarmUnlockkeyOKBody struct {

	// The swarm's unlock key.
	UnlockKey string `json:"UnlockKey,omitempty"`
}

// Validate validates this swarm unlockkey o k body
func (o *SwarmUnlockkeyOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swarm unlockkey o k body based on context it is used
func (o *SwarmUnlockkeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SwarmUnlockkeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwarmUnlockkeyOKBody) UnmarshalBinary(b []byte) error {
	var res SwarmUnlockkeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
