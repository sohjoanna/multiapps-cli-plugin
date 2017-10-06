// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTasklistTaskLogsParams creates a new GetTasklistTaskLogsParams object
// with the default values initialized.
func NewGetTasklistTaskLogsParams() *GetTasklistTaskLogsParams {
	var ()
	return &GetTasklistTaskLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasklistTaskLogsParamsWithTimeout creates a new GetTasklistTaskLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTasklistTaskLogsParamsWithTimeout(timeout time.Duration) *GetTasklistTaskLogsParams {
	var ()
	return &GetTasklistTaskLogsParams{

		timeout: timeout,
	}
}

// NewGetTasklistTaskLogsParamsWithContext creates a new GetTasklistTaskLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTasklistTaskLogsParamsWithContext(ctx context.Context) *GetTasklistTaskLogsParams {
	var ()
	return &GetTasklistTaskLogsParams{

		Context: ctx,
	}
}

// NewGetTasklistTaskLogsParamsWithHTTPClient creates a new GetTasklistTaskLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTasklistTaskLogsParamsWithHTTPClient(client *http.Client) *GetTasklistTaskLogsParams {
	var ()
	return &GetTasklistTaskLogsParams{
		HTTPClient: client,
	}
}

/*GetTasklistTaskLogsParams contains all the parameters to send to the API endpoint
for the get tasklist task logs operation typically these are written to a http.Request
*/
type GetTasklistTaskLogsParams struct {

	/*TaskID*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) WithTimeout(timeout time.Duration) *GetTasklistTaskLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) WithContext(ctx context.Context) *GetTasklistTaskLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) WithHTTPClient(client *http.Client) *GetTasklistTaskLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskID adds the taskID to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) WithTaskID(taskID string) *GetTasklistTaskLogsParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get tasklist task logs params
func (o *GetTasklistTaskLogsParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasklistTaskLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
