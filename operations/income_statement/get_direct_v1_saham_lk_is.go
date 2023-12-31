// Code generated by go-swagger; DO NOT EDIT.

package income_statement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDirectV1SahamLkIsHandlerFunc turns a function with the right signature into a get direct v1 saham lk is handler
type GetDirectV1SahamLkIsHandlerFunc func(GetDirectV1SahamLkIsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDirectV1SahamLkIsHandlerFunc) Handle(params GetDirectV1SahamLkIsParams) middleware.Responder {
	return fn(params)
}

// GetDirectV1SahamLkIsHandler interface for that can handle valid get direct v1 saham lk is params
type GetDirectV1SahamLkIsHandler interface {
	Handle(GetDirectV1SahamLkIsParams) middleware.Responder
}

// NewGetDirectV1SahamLkIs creates a new http.Handler for the get direct v1 saham lk is operation
func NewGetDirectV1SahamLkIs(ctx *middleware.Context, handler GetDirectV1SahamLkIsHandler) *GetDirectV1SahamLkIs {
	return &GetDirectV1SahamLkIs{Context: ctx, Handler: handler}
}

/* GetDirectV1SahamLkIs swagger:route GET /direct/v1/saham/lk/is/ Income Statement getDirectV1SahamLkIs

Get Income Statement Data

Get Rasio Keuangan Data based on provided parameters.

*/
type GetDirectV1SahamLkIs struct {
	Context *middleware.Context
	Handler GetDirectV1SahamLkIsHandler
}

func (o *GetDirectV1SahamLkIs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDirectV1SahamLkIsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
