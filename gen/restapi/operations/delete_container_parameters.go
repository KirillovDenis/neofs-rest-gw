// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteContainerParams creates a new DeleteContainerParams object
// with the default values initialized.
func NewDeleteContainerParams() DeleteContainerParams {

	var (
		// initialize parameters with default values

		walletConnectDefault = bool(false)
	)

	return DeleteContainerParams{
		WalletConnect: &walletConnectDefault,
	}
}

// DeleteContainerParams contains all the bound params for the delete container operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteContainer
type DeleteContainerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Base64 encoded signature for bearer token.
	  In: header
	*/
	XBearerSignature *string
	/*Hex encoded the public part of the key that signed the bearer token.
	  In: header
	*/
	XBearerSignatureKey *string
	/*Base58 encoded container id.
	  Required: true
	  In: path
	*/
	ContainerID string
	/*Use wallet connect signature scheme or native NeoFS signature.
	  In: query
	  Default: false
	*/
	WalletConnect *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteContainerParams() beforehand.
func (o *DeleteContainerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXBearerSignature(r.Header[http.CanonicalHeaderKey("X-Bearer-Signature")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXBearerSignatureKey(r.Header[http.CanonicalHeaderKey("X-Bearer-Signature-Key")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rContainerID, rhkContainerID, _ := route.Params.GetOK("containerId")
	if err := o.bindContainerID(rContainerID, rhkContainerID, route.Formats); err != nil {
		res = append(res, err)
	}

	qWalletConnect, qhkWalletConnect, _ := qs.GetOK("walletConnect")
	if err := o.bindWalletConnect(qWalletConnect, qhkWalletConnect, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXBearerSignature binds and validates parameter XBearerSignature from header.
func (o *DeleteContainerParams) bindXBearerSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XBearerSignature = &raw

	return nil
}

// bindXBearerSignatureKey binds and validates parameter XBearerSignatureKey from header.
func (o *DeleteContainerParams) bindXBearerSignatureKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XBearerSignatureKey = &raw

	return nil
}

// bindContainerID binds and validates parameter ContainerID from path.
func (o *DeleteContainerParams) bindContainerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ContainerID = raw

	return nil
}

// bindWalletConnect binds and validates parameter WalletConnect from query.
func (o *DeleteContainerParams) bindWalletConnect(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewDeleteContainerParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("walletConnect", "query", "bool", raw)
	}
	o.WalletConnect = &value

	return nil
}
