// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nspcc-dev/neofs-rest-gw/gen/models"
)

// DeleteStorageGroupOKCode is the HTTP code returned for type DeleteStorageGroupOK
const DeleteStorageGroupOKCode int = 200

/*DeleteStorageGroupOK Successful deletion.

swagger:response deleteStorageGroupOK
*/
type DeleteStorageGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewDeleteStorageGroupOK creates DeleteStorageGroupOK with default headers values
func NewDeleteStorageGroupOK() *DeleteStorageGroupOK {

	return &DeleteStorageGroupOK{}
}

// WithPayload adds the payload to the delete storage group o k response
func (o *DeleteStorageGroupOK) WithPayload(payload *models.SuccessResponse) *DeleteStorageGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage group o k response
func (o *DeleteStorageGroupOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageGroupBadRequestCode is the HTTP code returned for type DeleteStorageGroupBadRequest
const DeleteStorageGroupBadRequestCode int = 400

/*DeleteStorageGroupBadRequest Bad request.

swagger:response deleteStorageGroupBadRequest
*/
type DeleteStorageGroupBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteStorageGroupBadRequest creates DeleteStorageGroupBadRequest with default headers values
func NewDeleteStorageGroupBadRequest() *DeleteStorageGroupBadRequest {

	return &DeleteStorageGroupBadRequest{}
}

// WithPayload adds the payload to the delete storage group bad request response
func (o *DeleteStorageGroupBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteStorageGroupBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage group bad request response
func (o *DeleteStorageGroupBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageGroupBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
