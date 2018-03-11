// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteDescriptorHandlerFunc turns a function with the right signature into a delete descriptor handler
type DeleteDescriptorHandlerFunc func(DeleteDescriptorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDescriptorHandlerFunc) Handle(params DeleteDescriptorParams) middleware.Responder {
	return fn(params)
}

// DeleteDescriptorHandler interface for that can handle valid delete descriptor params
type DeleteDescriptorHandler interface {
	Handle(DeleteDescriptorParams) middleware.Responder
}

// NewDeleteDescriptor creates a new http.Handler for the delete descriptor operation
func NewDeleteDescriptor(ctx *middleware.Context, handler DeleteDescriptorHandler) *DeleteDescriptor {
	return &DeleteDescriptor{Context: ctx, Handler: handler}
}

/*DeleteDescriptor swagger:route DELETE /Descriptors/{id} Descriptors deleteDescriptor

Delete Descriptor

*/
type DeleteDescriptor struct {
	Context *middleware.Context
	Handler DeleteDescriptorHandler
}

func (o *DeleteDescriptor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteDescriptorParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
