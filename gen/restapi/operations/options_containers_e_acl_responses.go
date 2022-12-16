// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// OptionsContainersEACLOKCode is the HTTP code returned for type OptionsContainersEACLOK
const OptionsContainersEACLOKCode int = 200

/*OptionsContainersEACLOK CORS

swagger:response optionsContainersEAclOK
*/
type OptionsContainersEACLOK struct {
	/*

	 */
	AccessControlAllowHeaders string `json:"Access-Control-Allow-Headers"`
	/*

	 */
	AccessControlAllowMethods string `json:"Access-Control-Allow-Methods"`
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}

// NewOptionsContainersEACLOK creates OptionsContainersEACLOK with default headers values
func NewOptionsContainersEACLOK() *OptionsContainersEACLOK {

	return &OptionsContainersEACLOK{}
}

// WithAccessControlAllowHeaders adds the accessControlAllowHeaders to the options containers e Acl o k response
func (o *OptionsContainersEACLOK) WithAccessControlAllowHeaders(accessControlAllowHeaders string) *OptionsContainersEACLOK {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
	return o
}

// SetAccessControlAllowHeaders sets the accessControlAllowHeaders to the options containers e Acl o k response
func (o *OptionsContainersEACLOK) SetAccessControlAllowHeaders(accessControlAllowHeaders string) {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
}

// WithAccessControlAllowMethods adds the accessControlAllowMethods to the options containers e Acl o k response
func (o *OptionsContainersEACLOK) WithAccessControlAllowMethods(accessControlAllowMethods string) *OptionsContainersEACLOK {
	o.AccessControlAllowMethods = accessControlAllowMethods
	return o
}

// SetAccessControlAllowMethods sets the accessControlAllowMethods to the options containers e Acl o k response
func (o *OptionsContainersEACLOK) SetAccessControlAllowMethods(accessControlAllowMethods string) {
	o.AccessControlAllowMethods = accessControlAllowMethods
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the options containers e Acl o k response
func (o *OptionsContainersEACLOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *OptionsContainersEACLOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the options containers e Acl o k response
func (o *OptionsContainersEACLOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WriteResponse to the client
func (o *OptionsContainersEACLOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Headers

	accessControlAllowHeaders := o.AccessControlAllowHeaders
	if accessControlAllowHeaders != "" {
		rw.Header().Set("Access-Control-Allow-Headers", accessControlAllowHeaders)
	}

	// response header Access-Control-Allow-Methods

	accessControlAllowMethods := o.AccessControlAllowMethods
	if accessControlAllowMethods != "" {
		rw.Header().Set("Access-Control-Allow-Methods", accessControlAllowMethods)
	}

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
