// Code generated by go-swagger; DO NOT EDIT.

package exec

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

// NewExecInspectParams creates a new ExecInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecInspectParams() *ExecInspectParams {
	return &ExecInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecInspectParamsWithTimeout creates a new ExecInspectParams object
// with the ability to set a timeout on a request.
func NewExecInspectParamsWithTimeout(timeout time.Duration) *ExecInspectParams {
	return &ExecInspectParams{
		timeout: timeout,
	}
}

// NewExecInspectParamsWithContext creates a new ExecInspectParams object
// with the ability to set a context for a request.
func NewExecInspectParamsWithContext(ctx context.Context) *ExecInspectParams {
	return &ExecInspectParams{
		Context: ctx,
	}
}

// NewExecInspectParamsWithHTTPClient creates a new ExecInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecInspectParamsWithHTTPClient(client *http.Client) *ExecInspectParams {
	return &ExecInspectParams{
		HTTPClient: client,
	}
}

/* ExecInspectParams contains all the parameters to send to the API endpoint
   for the exec inspect operation.

   Typically these are written to a http.Request.
*/
type ExecInspectParams struct {

	/* ID.

	   Exec instance ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the exec inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecInspectParams) WithDefaults() *ExecInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the exec inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the exec inspect params
func (o *ExecInspectParams) WithTimeout(timeout time.Duration) *ExecInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exec inspect params
func (o *ExecInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exec inspect params
func (o *ExecInspectParams) WithContext(ctx context.Context) *ExecInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exec inspect params
func (o *ExecInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exec inspect params
func (o *ExecInspectParams) WithHTTPClient(client *http.Client) *ExecInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exec inspect params
func (o *ExecInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the exec inspect params
func (o *ExecInspectParams) WithID(id string) *ExecInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the exec inspect params
func (o *ExecInspectParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExecInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
