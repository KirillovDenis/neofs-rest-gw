// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/TrueCloudLab/frostfs-rest-gw/gen/models"
)

// PutContainerEACLHandlerFunc turns a function with the right signature into a put container e ACL handler
type PutContainerEACLHandlerFunc func(PutContainerEACLParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutContainerEACLHandlerFunc) Handle(params PutContainerEACLParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutContainerEACLHandler interface for that can handle valid put container e ACL params
type PutContainerEACLHandler interface {
	Handle(PutContainerEACLParams, *models.Principal) middleware.Responder
}

// NewPutContainerEACL creates a new http.Handler for the put container e ACL operation
func NewPutContainerEACL(ctx *middleware.Context, handler PutContainerEACLHandler) *PutContainerEACL {
	return &PutContainerEACL{Context: ctx, Handler: handler}
}

/*
	PutContainerEACL swagger:route PUT /containers/{containerId}/eacl putContainerEAcl

Set container EACL by id
*/
type PutContainerEACL struct {
	Context *middleware.Context
	Handler PutContainerEACLHandler
}

func (o *PutContainerEACL) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutContainerEACLParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
