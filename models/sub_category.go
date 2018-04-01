// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubCategory sub category
// swagger:model SubCategory

type SubCategory struct {

	// id
	ID int64 `gorm:"primary_key;AUTO_INCEMENT;not null" json:"id,omitempty"`

	// subcategory
	Subcategory string `gorm:"unique; not null" json:"subcategory,omitempty"`

	//Category Foreign Key
	CategoryID int64 `gorm:"foreignkey:CategoryID;not null" json:"category_id,omitempty"`
}

/* polymorph SubCategory id false */

/* polymorph SubCategory subcategory false */

// Validate validates this sub category
func (m *SubCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SubCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubCategory) UnmarshalBinary(b []byte) error {
	var res SubCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
