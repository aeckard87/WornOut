// Code generated by go-swagger; DO NOT EDIT.

package subcategories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteSubCategoriesOKCode is the HTTP code returned for type DeleteSubCategoriesOK
const DeleteSubCategoriesOKCode int = 200

/*DeleteSubCategoriesOK Status Ok

swagger:response deleteSubCategoriesOK
*/
type DeleteSubCategoriesOK struct {
}

// NewDeleteSubCategoriesOK creates DeleteSubCategoriesOK with default headers values
func NewDeleteSubCategoriesOK() *DeleteSubCategoriesOK {
	return &DeleteSubCategoriesOK{}
}

// WriteResponse to the client
func (o *DeleteSubCategoriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteSubCategoriesNotFoundCode is the HTTP code returned for type DeleteSubCategoriesNotFound
const DeleteSubCategoriesNotFoundCode int = 404

/*DeleteSubCategoriesNotFound No items found

swagger:response deleteSubCategoriesNotFound
*/
type DeleteSubCategoriesNotFound struct {
}

// NewDeleteSubCategoriesNotFound creates DeleteSubCategoriesNotFound with default headers values
func NewDeleteSubCategoriesNotFound() *DeleteSubCategoriesNotFound {
	return &DeleteSubCategoriesNotFound{}
}

// WriteResponse to the client
func (o *DeleteSubCategoriesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
