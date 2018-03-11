// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetItemHandlerFunc turns a function with the right signature into a get item handler
type GetItemHandlerFunc func(GetItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetItemHandlerFunc) Handle(params GetItemParams) middleware.Responder {
	return fn(params)
}

// GetItemHandler interface for that can handle valid get item params
type GetItemHandler interface {
	Handle(GetItemParams) middleware.Responder
}

// NewGetItem creates a new http.Handler for the get item operation
func NewGetItem(ctx *middleware.Context, handler GetItemHandler) *GetItem {
	return &GetItem{Context: ctx, Handler: handler}
}

/*GetItem swagger:route GET /items/{id} Items getItem

Get specific Item

*/
type GetItem struct {
	Context *middleware.Context
	Handler GetItemHandler
}

func (o *GetItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
