// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateDescriptorHandlerFunc turns a function with the right signature into a update descriptor handler
type UpdateDescriptorHandlerFunc func(UpdateDescriptorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateDescriptorHandlerFunc) Handle(params UpdateDescriptorParams) middleware.Responder {
	return fn(params)
}

// UpdateDescriptorHandler interface for that can handle valid update descriptor params
type UpdateDescriptorHandler interface {
	Handle(UpdateDescriptorParams) middleware.Responder
}

// NewUpdateDescriptor creates a new http.Handler for the update descriptor operation
func NewUpdateDescriptor(ctx *middleware.Context, handler UpdateDescriptorHandler) *UpdateDescriptor {
	return &UpdateDescriptor{Context: ctx, Handler: handler}
}

/*UpdateDescriptor swagger:route PUT /descriptors/{id} Descriptors updateDescriptor

Update a specific Descriptor

*/
type UpdateDescriptor struct {
	Context *middleware.Context
	Handler UpdateDescriptorHandler
}

func (o *UpdateDescriptor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateDescriptorParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
