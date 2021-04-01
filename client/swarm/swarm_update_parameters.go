// Code generated by go-swagger; DO NOT EDIT.

package swarm

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
	"github.com/go-openapi/swag"

	"github.com/go-openapi/dockerctl/models"
)

// NewSwarmUpdateParams creates a new SwarmUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSwarmUpdateParams() *SwarmUpdateParams {
	return &SwarmUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSwarmUpdateParamsWithTimeout creates a new SwarmUpdateParams object
// with the ability to set a timeout on a request.
func NewSwarmUpdateParamsWithTimeout(timeout time.Duration) *SwarmUpdateParams {
	return &SwarmUpdateParams{
		timeout: timeout,
	}
}

// NewSwarmUpdateParamsWithContext creates a new SwarmUpdateParams object
// with the ability to set a context for a request.
func NewSwarmUpdateParamsWithContext(ctx context.Context) *SwarmUpdateParams {
	return &SwarmUpdateParams{
		Context: ctx,
	}
}

// NewSwarmUpdateParamsWithHTTPClient creates a new SwarmUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSwarmUpdateParamsWithHTTPClient(client *http.Client) *SwarmUpdateParams {
	return &SwarmUpdateParams{
		HTTPClient: client,
	}
}

/* SwarmUpdateParams contains all the parameters to send to the API endpoint
   for the swarm update operation.

   Typically these are written to a http.Request.
*/
type SwarmUpdateParams struct {

	// Body.
	Body *models.SwarmSpec

	/* RotateManagerToken.

	   Rotate the manager join token.
	*/
	RotateManagerToken *bool

	/* RotateManagerUnlockKey.

	   Rotate the manager unlock key.
	*/
	RotateManagerUnlockKey *bool

	/* RotateWorkerToken.

	   Rotate the worker join token.
	*/
	RotateWorkerToken *bool

	/* Version.

	   The version number of the swarm object being updated. This is required to avoid conflicting writes.

	   Format: int64
	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the swarm update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwarmUpdateParams) WithDefaults() *SwarmUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the swarm update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwarmUpdateParams) SetDefaults() {
	var (
		rotateManagerTokenDefault = bool(false)

		rotateManagerUnlockKeyDefault = bool(false)

		rotateWorkerTokenDefault = bool(false)
	)

	val := SwarmUpdateParams{
		RotateManagerToken:     &rotateManagerTokenDefault,
		RotateManagerUnlockKey: &rotateManagerUnlockKeyDefault,
		RotateWorkerToken:      &rotateWorkerTokenDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the swarm update params
func (o *SwarmUpdateParams) WithTimeout(timeout time.Duration) *SwarmUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the swarm update params
func (o *SwarmUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the swarm update params
func (o *SwarmUpdateParams) WithContext(ctx context.Context) *SwarmUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the swarm update params
func (o *SwarmUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the swarm update params
func (o *SwarmUpdateParams) WithHTTPClient(client *http.Client) *SwarmUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the swarm update params
func (o *SwarmUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the swarm update params
func (o *SwarmUpdateParams) WithBody(body *models.SwarmSpec) *SwarmUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the swarm update params
func (o *SwarmUpdateParams) SetBody(body *models.SwarmSpec) {
	o.Body = body
}

// WithRotateManagerToken adds the rotateManagerToken to the swarm update params
func (o *SwarmUpdateParams) WithRotateManagerToken(rotateManagerToken *bool) *SwarmUpdateParams {
	o.SetRotateManagerToken(rotateManagerToken)
	return o
}

// SetRotateManagerToken adds the rotateManagerToken to the swarm update params
func (o *SwarmUpdateParams) SetRotateManagerToken(rotateManagerToken *bool) {
	o.RotateManagerToken = rotateManagerToken
}

// WithRotateManagerUnlockKey adds the rotateManagerUnlockKey to the swarm update params
func (o *SwarmUpdateParams) WithRotateManagerUnlockKey(rotateManagerUnlockKey *bool) *SwarmUpdateParams {
	o.SetRotateManagerUnlockKey(rotateManagerUnlockKey)
	return o
}

// SetRotateManagerUnlockKey adds the rotateManagerUnlockKey to the swarm update params
func (o *SwarmUpdateParams) SetRotateManagerUnlockKey(rotateManagerUnlockKey *bool) {
	o.RotateManagerUnlockKey = rotateManagerUnlockKey
}

// WithRotateWorkerToken adds the rotateWorkerToken to the swarm update params
func (o *SwarmUpdateParams) WithRotateWorkerToken(rotateWorkerToken *bool) *SwarmUpdateParams {
	o.SetRotateWorkerToken(rotateWorkerToken)
	return o
}

// SetRotateWorkerToken adds the rotateWorkerToken to the swarm update params
func (o *SwarmUpdateParams) SetRotateWorkerToken(rotateWorkerToken *bool) {
	o.RotateWorkerToken = rotateWorkerToken
}

// WithVersion adds the version to the swarm update params
func (o *SwarmUpdateParams) WithVersion(version int64) *SwarmUpdateParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the swarm update params
func (o *SwarmUpdateParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SwarmUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.RotateManagerToken != nil {

		// query param rotateManagerToken
		var qrRotateManagerToken bool

		if o.RotateManagerToken != nil {
			qrRotateManagerToken = *o.RotateManagerToken
		}
		qRotateManagerToken := swag.FormatBool(qrRotateManagerToken)
		if qRotateManagerToken != "" {

			if err := r.SetQueryParam("rotateManagerToken", qRotateManagerToken); err != nil {
				return err
			}
		}
	}

	if o.RotateManagerUnlockKey != nil {

		// query param rotateManagerUnlockKey
		var qrRotateManagerUnlockKey bool

		if o.RotateManagerUnlockKey != nil {
			qrRotateManagerUnlockKey = *o.RotateManagerUnlockKey
		}
		qRotateManagerUnlockKey := swag.FormatBool(qrRotateManagerUnlockKey)
		if qRotateManagerUnlockKey != "" {

			if err := r.SetQueryParam("rotateManagerUnlockKey", qRotateManagerUnlockKey); err != nil {
				return err
			}
		}
	}

	if o.RotateWorkerToken != nil {

		// query param rotateWorkerToken
		var qrRotateWorkerToken bool

		if o.RotateWorkerToken != nil {
			qrRotateWorkerToken = *o.RotateWorkerToken
		}
		qRotateWorkerToken := swag.FormatBool(qrRotateWorkerToken)
		if qRotateWorkerToken != "" {

			if err := r.SetQueryParam("rotateWorkerToken", qRotateWorkerToken); err != nil {
				return err
			}
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {

		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
