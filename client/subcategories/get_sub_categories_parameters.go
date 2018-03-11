// Code generated by go-swagger; DO NOT EDIT.

package subcategories

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

// NewGetSubCategoriesParams creates a new GetSubCategoriesParams object
// with the default values initialized.
func NewGetSubCategoriesParams() *GetSubCategoriesParams {

	return &GetSubCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubCategoriesParamsWithTimeout creates a new GetSubCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubCategoriesParamsWithTimeout(timeout time.Duration) *GetSubCategoriesParams {

	return &GetSubCategoriesParams{

		timeout: timeout,
	}
}

// NewGetSubCategoriesParamsWithContext creates a new GetSubCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubCategoriesParamsWithContext(ctx context.Context) *GetSubCategoriesParams {

	return &GetSubCategoriesParams{

		Context: ctx,
	}
}

// NewGetSubCategoriesParamsWithHTTPClient creates a new GetSubCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubCategoriesParamsWithHTTPClient(client *http.Client) *GetSubCategoriesParams {

	return &GetSubCategoriesParams{
		HTTPClient: client,
	}
}

/*GetSubCategoriesParams contains all the parameters to send to the API endpoint
for the get sub categories operation typically these are written to a http.Request
*/
type GetSubCategoriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sub categories params
func (o *GetSubCategoriesParams) WithTimeout(timeout time.Duration) *GetSubCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sub categories params
func (o *GetSubCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sub categories params
func (o *GetSubCategoriesParams) WithContext(ctx context.Context) *GetSubCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sub categories params
func (o *GetSubCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sub categories params
func (o *GetSubCategoriesParams) WithHTTPClient(client *http.Client) *GetSubCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sub categories params
func (o *GetSubCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
