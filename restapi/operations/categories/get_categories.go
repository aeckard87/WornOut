// Code generated by go-swagger; DO NOT EDIT.

package categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCategoriesHandlerFunc turns a function with the right signature into a get categories handler
type GetCategoriesHandlerFunc func(GetCategoriesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCategoriesHandlerFunc) Handle(params GetCategoriesParams) middleware.Responder {
	return fn(params)
}

// GetCategoriesHandler interface for that can handle valid get categories params
type GetCategoriesHandler interface {
	Handle(GetCategoriesParams) middleware.Responder
}

// NewGetCategories creates a new http.Handler for the get categories operation
func NewGetCategories(ctx *middleware.Context, handler GetCategoriesHandler) *GetCategories {
	return &GetCategories{Context: ctx, Handler: handler}
}

/*GetCategories swagger:route GET /categories Categories getCategories

Get all categories

*/
type GetCategories struct {
	Context *middleware.Context
	Handler GetCategoriesHandler
}

func (o *GetCategories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCategoriesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
