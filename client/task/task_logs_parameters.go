// Code generated by go-swagger; DO NOT EDIT.

package task

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
)

// NewTaskLogsParams creates a new TaskLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTaskLogsParams() *TaskLogsParams {
	return &TaskLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTaskLogsParamsWithTimeout creates a new TaskLogsParams object
// with the ability to set a timeout on a request.
func NewTaskLogsParamsWithTimeout(timeout time.Duration) *TaskLogsParams {
	return &TaskLogsParams{
		timeout: timeout,
	}
}

// NewTaskLogsParamsWithContext creates a new TaskLogsParams object
// with the ability to set a context for a request.
func NewTaskLogsParamsWithContext(ctx context.Context) *TaskLogsParams {
	return &TaskLogsParams{
		Context: ctx,
	}
}

// NewTaskLogsParamsWithHTTPClient creates a new TaskLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTaskLogsParamsWithHTTPClient(client *http.Client) *TaskLogsParams {
	return &TaskLogsParams{
		HTTPClient: client,
	}
}

/* TaskLogsParams contains all the parameters to send to the API endpoint
   for the task logs operation.

   Typically these are written to a http.Request.
*/
type TaskLogsParams struct {

	/* Details.

	   Show task context and extra details provided to logs.
	*/
	Details *bool

	/* Follow.

	   Keep connection after returning logs.
	*/
	Follow *bool

	/* ID.

	   ID of the task
	*/
	ID string

	/* Since.

	   Only return logs since this time, as a UNIX timestamp
	*/
	Since *int64

	/* Stderr.

	   Return logs from `stderr`
	*/
	Stderr *bool

	/* Stdout.

	   Return logs from `stdout`
	*/
	Stdout *bool

	/* Tail.

	   Only return this number of log lines from the end of the logs. Specify as an integer or `all` to output all log lines.

	   Default: "all"
	*/
	Tail *string

	/* Timestamps.

	   Add timestamps to every log line
	*/
	Timestamps *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the task logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaskLogsParams) WithDefaults() *TaskLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the task logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TaskLogsParams) SetDefaults() {
	var (
		detailsDefault = bool(false)

		followDefault = bool(false)

		sinceDefault = int64(0)

		stderrDefault = bool(false)

		stdoutDefault = bool(false)

		tailDefault = string("all")

		timestampsDefault = bool(false)
	)

	val := TaskLogsParams{
		Details:    &detailsDefault,
		Follow:     &followDefault,
		Since:      &sinceDefault,
		Stderr:     &stderrDefault,
		Stdout:     &stdoutDefault,
		Tail:       &tailDefault,
		Timestamps: &timestampsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the task logs params
func (o *TaskLogsParams) WithTimeout(timeout time.Duration) *TaskLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the task logs params
func (o *TaskLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the task logs params
func (o *TaskLogsParams) WithContext(ctx context.Context) *TaskLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the task logs params
func (o *TaskLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the task logs params
func (o *TaskLogsParams) WithHTTPClient(client *http.Client) *TaskLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the task logs params
func (o *TaskLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetails adds the details to the task logs params
func (o *TaskLogsParams) WithDetails(details *bool) *TaskLogsParams {
	o.SetDetails(details)
	return o
}

// SetDetails adds the details to the task logs params
func (o *TaskLogsParams) SetDetails(details *bool) {
	o.Details = details
}

// WithFollow adds the follow to the task logs params
func (o *TaskLogsParams) WithFollow(follow *bool) *TaskLogsParams {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the task logs params
func (o *TaskLogsParams) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithID adds the id to the task logs params
func (o *TaskLogsParams) WithID(id string) *TaskLogsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the task logs params
func (o *TaskLogsParams) SetID(id string) {
	o.ID = id
}

// WithSince adds the since to the task logs params
func (o *TaskLogsParams) WithSince(since *int64) *TaskLogsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the task logs params
func (o *TaskLogsParams) SetSince(since *int64) {
	o.Since = since
}

// WithStderr adds the stderr to the task logs params
func (o *TaskLogsParams) WithStderr(stderr *bool) *TaskLogsParams {
	o.SetStderr(stderr)
	return o
}

// SetStderr adds the stderr to the task logs params
func (o *TaskLogsParams) SetStderr(stderr *bool) {
	o.Stderr = stderr
}

// WithStdout adds the stdout to the task logs params
func (o *TaskLogsParams) WithStdout(stdout *bool) *TaskLogsParams {
	o.SetStdout(stdout)
	return o
}

// SetStdout adds the stdout to the task logs params
func (o *TaskLogsParams) SetStdout(stdout *bool) {
	o.Stdout = stdout
}

// WithTail adds the tail to the task logs params
func (o *TaskLogsParams) WithTail(tail *string) *TaskLogsParams {
	o.SetTail(tail)
	return o
}

// SetTail adds the tail to the task logs params
func (o *TaskLogsParams) SetTail(tail *string) {
	o.Tail = tail
}

// WithTimestamps adds the timestamps to the task logs params
func (o *TaskLogsParams) WithTimestamps(timestamps *bool) *TaskLogsParams {
	o.SetTimestamps(timestamps)
	return o
}

// SetTimestamps adds the timestamps to the task logs params
func (o *TaskLogsParams) SetTimestamps(timestamps *bool) {
	o.Timestamps = timestamps
}

// WriteToRequest writes these params to a swagger request
func (o *TaskLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Details != nil {

		// query param details
		var qrDetails bool

		if o.Details != nil {
			qrDetails = *o.Details
		}
		qDetails := swag.FormatBool(qrDetails)
		if qDetails != "" {

			if err := r.SetQueryParam("details", qDetails); err != nil {
				return err
			}
		}
	}

	if o.Follow != nil {

		// query param follow
		var qrFollow bool

		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {

			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Since != nil {

		// query param since
		var qrSince int64

		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {

			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}
	}

	if o.Stderr != nil {

		// query param stderr
		var qrStderr bool

		if o.Stderr != nil {
			qrStderr = *o.Stderr
		}
		qStderr := swag.FormatBool(qrStderr)
		if qStderr != "" {

			if err := r.SetQueryParam("stderr", qStderr); err != nil {
				return err
			}
		}
	}

	if o.Stdout != nil {

		// query param stdout
		var qrStdout bool

		if o.Stdout != nil {
			qrStdout = *o.Stdout
		}
		qStdout := swag.FormatBool(qrStdout)
		if qStdout != "" {

			if err := r.SetQueryParam("stdout", qStdout); err != nil {
				return err
			}
		}
	}

	if o.Tail != nil {

		// query param tail
		var qrTail string

		if o.Tail != nil {
			qrTail = *o.Tail
		}
		qTail := qrTail
		if qTail != "" {

			if err := r.SetQueryParam("tail", qTail); err != nil {
				return err
			}
		}
	}

	if o.Timestamps != nil {

		// query param timestamps
		var qrTimestamps bool

		if o.Timestamps != nil {
			qrTimestamps = *o.Timestamps
		}
		qTimestamps := swag.FormatBool(qrTimestamps)
		if qTimestamps != "" {

			if err := r.SetQueryParam("timestamps", qTimestamps); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
