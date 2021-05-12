// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// HealthHandlerFunc turns a function with the right signature into a health handler
type HealthHandlerFunc func(HealthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HealthHandlerFunc) Handle(params HealthParams) middleware.Responder {
	return fn(params)
}

// HealthHandler interface for that can handle valid health params
type HealthHandler interface {
	Handle(HealthParams) middleware.Responder
}

// NewHealth creates a new http.Handler for the health operation
func NewHealth(ctx *middleware.Context, handler HealthHandler) *Health {
	return &Health{Context: ctx, Handler: handler}
}

/*Health swagger:route GET /health health

Healtcheck function

*/
type Health struct {
	Context *middleware.Context
	Handler HealthHandler
}

func (o *Health) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHealthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
