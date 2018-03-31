// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Base base
// swagger:model Base

type Base struct {

	// category
	Category *Category `json:"category,omitempty"`

	// subcategory
	Subcategory *SubCategory `json:"subcategory,omitempty"`
}

/* polymorph Base category false */

/* polymorph Base subcategory false */

// Validate validates this base
func (m *Base) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubcategory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Base) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {

		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *Base) validateSubcategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Subcategory) { // not required
		return nil
	}

	if m.Subcategory != nil {

		if err := m.Subcategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subcategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Base) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Base) UnmarshalBinary(b []byte) error {
	var res Base
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
