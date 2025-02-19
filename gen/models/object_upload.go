// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectUpload Request body to create object.
// Example: {"attributes":[{"key":"User-Attribute","value":"some-value"}],"containerId":"5HZTn5qkRnmgSz9gSrw22CEdPPk6nQhkwf2Mgzyvkikv","fileName":"myFile.txt","payload":"Y29udGVudCBvZiBmaWxl"}
//
// swagger:model ObjectUpload
type ObjectUpload struct {

	// attributes
	Attributes []*Attribute `json:"attributes"`

	// container Id
	// Required: true
	ContainerID *string `json:"containerId"`

	// file name
	// Required: true
	FileName *string `json:"fileName"`

	// payload
	Payload string `json:"payload,omitempty"`
}

// Validate validates this object upload
func (m *ObjectUpload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectUpload) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectUpload) validateContainerID(formats strfmt.Registry) error {

	if err := validate.Required("containerId", "body", m.ContainerID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectUpload) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("fileName", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this object upload based on the context it is used
func (m *ObjectUpload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectUpload) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectUpload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectUpload) UnmarshalBinary(b []byte) error {
	var res ObjectUpload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
