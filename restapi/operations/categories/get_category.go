// Code generated by go-swagger; DO NOT EDIT.

package categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCategoryHandlerFunc turns a function with the right signature into a get category handler
type GetCategoryHandlerFunc func(GetCategoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCategoryHandlerFunc) Handle(params GetCategoryParams) middleware.Responder {
	return fn(params)
}

// GetCategoryHandler interface for that can handle valid get category params
type GetCategoryHandler interface {
	Handle(GetCategoryParams) middleware.Responder
}

// NewGetCategory creates a new http.Handler for the get category operation
func NewGetCategory(ctx *middleware.Context, handler GetCategoryHandler) *GetCategory {
	return &GetCategory{Context: ctx, Handler: handler}
}

/*GetCategory swagger:route GET /categories/{id} Categories getCategory

Get a specific Category

*/
type GetCategory struct {
	Context *middleware.Context
	Handler GetCategoryHandler
}

func (o *GetCategory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCategoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}