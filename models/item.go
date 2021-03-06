// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Item item
// swagger:model Item

type Item struct {

	// descriptions
	Descriptions []Descriptor `json:"descriptions"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sub category ID
	SubCategoryID int64 `json:"subCategoryID,omitempty"`

	// user ID
	UserID int64 `json:"userID,omitempty"`
}

/* polymorph Item descriptions false */

/* polymorph Item id false */

/* polymorph Item name false */

/* polymorph Item subCategoryID false */

/* polymorph Item userID false */

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
