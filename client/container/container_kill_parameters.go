// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewContainerKillParams creates a new ContainerKillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerKillParams() *ContainerKillParams {
	return &ContainerKillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerKillParamsWithTimeout creates a new ContainerKillParams object
// with the ability to set a timeout on a request.
func NewContainerKillParamsWithTimeout(timeout time.Duration) *ContainerKillParams {
	return &ContainerKillParams{
		timeout: timeout,
	}
}

// NewContainerKillParamsWithContext creates a new ContainerKillParams object
// with the ability to set a context for a request.
func NewContainerKillParamsWithContext(ctx context.Context) *ContainerKillParams {
	return &ContainerKillParams{
		Context: ctx,
	}
}

// NewContainerKillParamsWithHTTPClient creates a new ContainerKillParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerKillParamsWithHTTPClient(client *http.Client) *ContainerKillParams {
	return &ContainerKillParams{
		HTTPClient: client,
	}
}

/* ContainerKillParams contains all the parameters to send to the API endpoint
   for the container kill operation.

   Typically these are written to a http.Request.
*/
type ContainerKillParams struct {

	/* ID.

	   ID or name of the container
	*/
	ID string

	/* Signal.

	   Signal to send to the container as an integer or string (e.g. `SIGINT`)

	   Default: "SIGKILL"
	*/
	Signal *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container kill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerKillParams) WithDefaults() *ContainerKillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container kill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerKillParams) SetDefaults() {
	var (
		signalDefault = string("SIGKILL")
	)

	val := ContainerKillParams{
		Signal: &signalDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container kill params
func (o *ContainerKillParams) WithTimeout(timeout time.Duration) *ContainerKillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container kill params
func (o *ContainerKillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container kill params
func (o *ContainerKillParams) WithContext(ctx context.Context) *ContainerKillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container kill params
func (o *ContainerKillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container kill params
func (o *ContainerKillParams) WithHTTPClient(client *http.Client) *ContainerKillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container kill params
func (o *ContainerKillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the container kill params
func (o *ContainerKillParams) WithID(id string) *ContainerKillParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container kill params
func (o *ContainerKillParams) SetID(id string) {
	o.ID = id
}

// WithSignal adds the signal to the container kill params
func (o *ContainerKillParams) WithSignal(signal *string) *ContainerKillParams {
	o.SetSignal(signal)
	return o
}

// SetSignal adds the signal to the container kill params
func (o *ContainerKillParams) SetSignal(signal *string) {
	o.Signal = signal
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerKillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Signal != nil {

		// query param signal
		var qrSignal string

		if o.Signal != nil {
			qrSignal = *o.Signal
		}
		qSignal := qrSignal
		if qSignal != "" {

			if err := r.SetQueryParam("signal", qSignal); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
