// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nspcc-dev/neofs-rest-gw/gen/models"
)

// GetContainerEACLOKCode is the HTTP code returned for type GetContainerEACLOK
const GetContainerEACLOKCode int = 200

/*GetContainerEACLOK Container EACL information.

swagger:response getContainerEAclOK
*/
type GetContainerEACLOK struct {
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload *models.Eacl `json:"body,omitempty"`
}

// NewGetContainerEACLOK creates GetContainerEACLOK with default headers values
func NewGetContainerEACLOK() *GetContainerEACLOK {

	return &GetContainerEACLOK{}
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the get container e Acl o k response
func (o *GetContainerEACLOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *GetContainerEACLOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the get container e Acl o k response
func (o *GetContainerEACLOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the get container e Acl o k response
func (o *GetContainerEACLOK) WithPayload(payload *models.Eacl) *GetContainerEACLOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get container e Acl o k response
func (o *GetContainerEACLOK) SetPayload(payload *models.Eacl) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContainerEACLOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetContainerEACLBadRequestCode is the HTTP code returned for type GetContainerEACLBadRequest
const GetContainerEACLBadRequestCode int = 400

/*GetContainerEACLBadRequest Bad request.

swagger:response getContainerEAclBadRequest
*/
type GetContainerEACLBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetContainerEACLBadRequest creates GetContainerEACLBadRequest with default headers values
func NewGetContainerEACLBadRequest() *GetContainerEACLBadRequest {

	return &GetContainerEACLBadRequest{}
}

// WithPayload adds the payload to the get container e Acl bad request response
func (o *GetContainerEACLBadRequest) WithPayload(payload *models.ErrorResponse) *GetContainerEACLBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get container e Acl bad request response
func (o *GetContainerEACLBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContainerEACLBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
