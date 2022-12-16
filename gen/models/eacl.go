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

// Eacl EACL FrostFS table.
// Example: {"containerId":"5HZTn5qkRnmgSz9gSrw22CEdPPk6nQhkwf2Mgzyvkikv","records":[{"action":"GET","filters":[{"headerType":"OBJECT","key":"FileName","matchType":"STRING_EQUAL","value":"myfile"}],"operation":"ALLOW","targets":[{"role":"OTHERS"}]}]}
//
// swagger:model Eacl
type Eacl struct {

	// container Id
	// Read Only: true
	ContainerID string `json:"containerId,omitempty"`

	// records
	// Required: true
	Records []*Record `json:"records"`
}

// Validate validates this eacl
func (m *Eacl) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Eacl) validateRecords(formats strfmt.Registry) error {

	if err := validate.Required("records", "body", m.Records); err != nil {
		return err
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this eacl based on the context it is used
func (m *Eacl) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Eacl) contextValidateContainerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "containerId", "body", string(m.ContainerID)); err != nil {
		return err
	}

	return nil
}

func (m *Eacl) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Eacl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Eacl) UnmarshalBinary(b []byte) error {
	var res Eacl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
