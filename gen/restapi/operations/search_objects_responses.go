// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nspcc-dev/neofs-rest-gw/gen/models"
)

// SearchObjectsOKCode is the HTTP code returned for type SearchObjectsOK
const SearchObjectsOKCode int = 200

/*SearchObjectsOK List of objects

swagger:response searchObjectsOK
*/
type SearchObjectsOK struct {
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload *models.ObjectList `json:"body,omitempty"`
}

// NewSearchObjectsOK creates SearchObjectsOK with default headers values
func NewSearchObjectsOK() *SearchObjectsOK {

	return &SearchObjectsOK{}
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the search objects o k response
func (o *SearchObjectsOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *SearchObjectsOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the search objects o k response
func (o *SearchObjectsOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the search objects o k response
func (o *SearchObjectsOK) WithPayload(payload *models.ObjectList) *SearchObjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search objects o k response
func (o *SearchObjectsOK) SetPayload(payload *models.ObjectList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchObjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// SearchObjectsBadRequestCode is the HTTP code returned for type SearchObjectsBadRequest
const SearchObjectsBadRequestCode int = 400

/*SearchObjectsBadRequest Bad request

swagger:response searchObjectsBadRequest
*/
type SearchObjectsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSearchObjectsBadRequest creates SearchObjectsBadRequest with default headers values
func NewSearchObjectsBadRequest() *SearchObjectsBadRequest {

	return &SearchObjectsBadRequest{}
}

// WithPayload adds the payload to the search objects bad request response
func (o *SearchObjectsBadRequest) WithPayload(payload *models.ErrorResponse) *SearchObjectsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search objects bad request response
func (o *SearchObjectsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchObjectsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
