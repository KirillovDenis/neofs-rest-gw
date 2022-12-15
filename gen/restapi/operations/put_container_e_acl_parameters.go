// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/TrueCloudLab/frostfs-rest-gw/gen/models"
)

// NewPutContainerEACLParams creates a new PutContainerEACLParams object
// with the default values initialized.
func NewPutContainerEACLParams() PutContainerEACLParams {

	var (
		// initialize parameters with default values

		walletConnectDefault = bool(false)
	)

	return PutContainerEACLParams{
		WalletConnect: &walletConnectDefault,
	}
}

// PutContainerEACLParams contains all the bound params for the put container e ACL operation
// typically these are obtained from a http.Request
//
// swagger:parameters putContainerEACL
type PutContainerEACLParams struct {

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
	/*EACL for container.
	  Required: true
	  In: body
	*/
	Eacl *models.Eacl
	/*Use wallet connect signature scheme or native NeoFS signature.
	  In: query
	  Default: false
	*/
	WalletConnect *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutContainerEACLParams() beforehand.
func (o *PutContainerEACLParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Eacl
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("eacl", "body", ""))
			} else {
				res = append(res, errors.NewParseError("eacl", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Eacl = &body
			}
		}
	} else {
		res = append(res, errors.Required("eacl", "body", ""))
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
func (o *PutContainerEACLParams) bindXBearerSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PutContainerEACLParams) bindXBearerSignatureKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PutContainerEACLParams) bindContainerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PutContainerEACLParams) bindWalletConnect(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPutContainerEACLParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("walletConnect", "query", "bool", raw)
	}
	o.WalletConnect = &value

	return nil
}
