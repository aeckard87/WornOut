// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDescriptorHandlerFunc turns a function with the right signature into a get descriptor handler
type GetDescriptorHandlerFunc func(GetDescriptorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDescriptorHandlerFunc) Handle(params GetDescriptorParams) middleware.Responder {
	return fn(params)
}

// GetDescriptorHandler interface for that can handle valid get descriptor params
type GetDescriptorHandler interface {
	Handle(GetDescriptorParams) middleware.Responder
}

// NewGetDescriptor creates a new http.Handler for the get descriptor operation
func NewGetDescriptor(ctx *middleware.Context, handler GetDescriptorHandler) *GetDescriptor {
	return &GetDescriptor{Context: ctx, Handler: handler}
}

/*GetDescriptor swagger:route GET /Descriptors/{id} Descriptors getDescriptor

Get a specific Descriptor

*/
type GetDescriptor struct {
	Context *middleware.Context
	Handler GetDescriptorHandler
}

func (o *GetDescriptor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDescriptorParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
