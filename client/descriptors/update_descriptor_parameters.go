// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aeckard87/WornOut/models"
)

// NewUpdateDescriptorParams creates a new UpdateDescriptorParams object
// with the default values initialized.
func NewUpdateDescriptorParams() *UpdateDescriptorParams {
	var ()
	return &UpdateDescriptorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDescriptorParamsWithTimeout creates a new UpdateDescriptorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDescriptorParamsWithTimeout(timeout time.Duration) *UpdateDescriptorParams {
	var ()
	return &UpdateDescriptorParams{

		timeout: timeout,
	}
}

// NewUpdateDescriptorParamsWithContext creates a new UpdateDescriptorParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDescriptorParamsWithContext(ctx context.Context) *UpdateDescriptorParams {
	var ()
	return &UpdateDescriptorParams{

		Context: ctx,
	}
}

// NewUpdateDescriptorParamsWithHTTPClient creates a new UpdateDescriptorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDescriptorParamsWithHTTPClient(client *http.Client) *UpdateDescriptorParams {
	var ()
	return &UpdateDescriptorParams{
		HTTPClient: client,
	}
}

/*UpdateDescriptorParams contains all the parameters to send to the API endpoint
for the update descriptor operation typically these are written to a http.Request
*/
type UpdateDescriptorParams struct {

	/*Body
	  a descriptor that will be updated

	*/
	Body *models.Descriptor
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update descriptor params
func (o *UpdateDescriptorParams) WithTimeout(timeout time.Duration) *UpdateDescriptorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update descriptor params
func (o *UpdateDescriptorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update descriptor params
func (o *UpdateDescriptorParams) WithContext(ctx context.Context) *UpdateDescriptorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update descriptor params
func (o *UpdateDescriptorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update descriptor params
func (o *UpdateDescriptorParams) WithHTTPClient(client *http.Client) *UpdateDescriptorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update descriptor params
func (o *UpdateDescriptorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update descriptor params
func (o *UpdateDescriptorParams) WithBody(body *models.Descriptor) *UpdateDescriptorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update descriptor params
func (o *UpdateDescriptorParams) SetBody(body *models.Descriptor) {
	o.Body = body
}

// WithID adds the id to the update descriptor params
func (o *UpdateDescriptorParams) WithID(id int64) *UpdateDescriptorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update descriptor params
func (o *UpdateDescriptorParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDescriptorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.Descriptor)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
