// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateItemHandlerFunc turns a function with the right signature into a create item handler
type CreateItemHandlerFunc func(CreateItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateItemHandlerFunc) Handle(params CreateItemParams) middleware.Responder {
	return fn(params)
}

// CreateItemHandler interface for that can handle valid create item params
type CreateItemHandler interface {
	Handle(CreateItemParams) middleware.Responder
}

// NewCreateItem creates a new http.Handler for the create item operation
func NewCreateItem(ctx *middleware.Context, handler CreateItemHandler) *CreateItem {
	return &CreateItem{Context: ctx, Handler: handler}
}

/*CreateItem swagger:route POST /items Items createItem

Create an Item

*/
type CreateItem struct {
	Context *middleware.Context
	Handler CreateItemHandler
}

func (o *CreateItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
