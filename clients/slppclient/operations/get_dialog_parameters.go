package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDialogParams creates a new GetDialogParams object
// with the default values initialized.
func NewGetDialogParams() *GetDialogParams {
	var ()
	return &GetDialogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDialogParamsWithTimeout creates a new GetDialogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDialogParamsWithTimeout(timeout time.Duration) *GetDialogParams {
	var ()
	return &GetDialogParams{

		timeout: timeout,
	}
}

// NewGetDialogParamsWithContext creates a new GetDialogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDialogParamsWithContext(ctx context.Context) *GetDialogParams {
	var ()
	return &GetDialogParams{

		Context: ctx,
	}
}

/*GetDialogParams contains all the parameters to send to the API endpoint
for the get dialog operation typically these are written to a http.Request
*/
type GetDialogParams struct {

	/*DialogID*/
	DialogID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get dialog params
func (o *GetDialogParams) WithTimeout(timeout time.Duration) *GetDialogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dialog params
func (o *GetDialogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dialog params
func (o *GetDialogParams) WithContext(ctx context.Context) *GetDialogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dialog params
func (o *GetDialogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDialogID adds the dialogID to the get dialog params
func (o *GetDialogParams) WithDialogID(dialogID string) *GetDialogParams {
	o.SetDialogID(dialogID)
	return o
}

// SetDialogID adds the dialogId to the get dialog params
func (o *GetDialogParams) SetDialogID(dialogID string) {
	o.DialogID = dialogID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDialogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param dialogId
	if err := r.SetPathParam("dialogId", o.DialogID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
