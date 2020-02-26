// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ZoneIdentityByName ZoneIdentityByName
// swagger:model ZoneIdentityByName
type ZoneIdentityByName struct {

	// The name for this zone
	// Required: true
	// Max Length: 63
	// Min Length: 1
	// Pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$
	Name *string `json:"name"`
}

// Validate validates this zone identity by name
func (m *ZoneIdentityByName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneIdentityByName) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 63); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneIdentityByName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneIdentityByName) UnmarshalBinary(b []byte) error {
	var res ZoneIdentityByName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}