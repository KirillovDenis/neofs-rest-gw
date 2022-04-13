// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nspcc-dev/neofs-rest-gw/gen/models"
)

// ListContainersOKCode is the HTTP code returned for type ListContainersOK
const ListContainersOKCode int = 200

/*ListContainersOK Containers info

swagger:response listContainersOK
*/
type ListContainersOK struct {

	/*
	  In: Body
	*/
	Payload *models.ContainerList `json:"body,omitempty"`
}

// NewListContainersOK creates ListContainersOK with default headers values
func NewListContainersOK() *ListContainersOK {

	return &ListContainersOK{}
}

// WithPayload adds the payload to the list containers o k response
func (o *ListContainersOK) WithPayload(payload *models.ContainerList) *ListContainersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list containers o k response
func (o *ListContainersOK) SetPayload(payload *models.ContainerList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListContainersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListContainersBadRequestCode is the HTTP code returned for type ListContainersBadRequest
const ListContainersBadRequestCode int = 400

/*ListContainersBadRequest Bad request

swagger:response listContainersBadRequest
*/
type ListContainersBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewListContainersBadRequest creates ListContainersBadRequest with default headers values
func NewListContainersBadRequest() *ListContainersBadRequest {

	return &ListContainersBadRequest{}
}

// WithPayload adds the payload to the list containers bad request response
func (o *ListContainersBadRequest) WithPayload(payload models.Error) *ListContainersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list containers bad request response
func (o *ListContainersBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListContainersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
