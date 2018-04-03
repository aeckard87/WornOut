// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteItemsBySubCategoryOKCode is the HTTP code returned for type DeleteItemsBySubCategoryOK
const DeleteItemsBySubCategoryOKCode int = 200

/*DeleteItemsBySubCategoryOK Status Ok

swagger:response deleteItemsBySubCategoryOK
*/
type DeleteItemsBySubCategoryOK struct {
}

// NewDeleteItemsBySubCategoryOK creates DeleteItemsBySubCategoryOK with default headers values
func NewDeleteItemsBySubCategoryOK() *DeleteItemsBySubCategoryOK {
	return &DeleteItemsBySubCategoryOK{}
}

// WriteResponse to the client
func (o *DeleteItemsBySubCategoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteItemsBySubCategoryNotFoundCode is the HTTP code returned for type DeleteItemsBySubCategoryNotFound
const DeleteItemsBySubCategoryNotFoundCode int = 404

/*DeleteItemsBySubCategoryNotFound No items found

swagger:response deleteItemsBySubCategoryNotFound
*/
type DeleteItemsBySubCategoryNotFound struct {
}

// NewDeleteItemsBySubCategoryNotFound creates DeleteItemsBySubCategoryNotFound with default headers values
func NewDeleteItemsBySubCategoryNotFound() *DeleteItemsBySubCategoryNotFound {
	return &DeleteItemsBySubCategoryNotFound{}
}

// WriteResponse to the client
func (o *DeleteItemsBySubCategoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
