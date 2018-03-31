// Code generated by go-swagger; DO NOT EDIT.

package subcategories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteSubCategoryOKCode is the HTTP code returned for type DeleteSubCategoryOK
const DeleteSubCategoryOKCode int = 200

/*DeleteSubCategoryOK Status Ok

swagger:response deleteSubCategoryOK
*/
type DeleteSubCategoryOK struct {
}

// NewDeleteSubCategoryOK creates DeleteSubCategoryOK with default headers values
func NewDeleteSubCategoryOK() *DeleteSubCategoryOK {
	return &DeleteSubCategoryOK{}
}

// WriteResponse to the client
func (o *DeleteSubCategoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteSubCategoryBadRequestCode is the HTTP code returned for type DeleteSubCategoryBadRequest
const DeleteSubCategoryBadRequestCode int = 400

/*DeleteSubCategoryBadRequest Invalid input

swagger:response deleteSubCategoryBadRequest
*/
type DeleteSubCategoryBadRequest struct {
}

// NewDeleteSubCategoryBadRequest creates DeleteSubCategoryBadRequest with default headers values
func NewDeleteSubCategoryBadRequest() *DeleteSubCategoryBadRequest {
	return &DeleteSubCategoryBadRequest{}
}

// WriteResponse to the client
func (o *DeleteSubCategoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// DeleteSubCategoryNotFoundCode is the HTTP code returned for type DeleteSubCategoryNotFound
const DeleteSubCategoryNotFoundCode int = 404

/*DeleteSubCategoryNotFound No items found

swagger:response deleteSubCategoryNotFound
*/
type DeleteSubCategoryNotFound struct {
}

// NewDeleteSubCategoryNotFound creates DeleteSubCategoryNotFound with default headers values
func NewDeleteSubCategoryNotFound() *DeleteSubCategoryNotFound {
	return &DeleteSubCategoryNotFound{}
}

// WriteResponse to the client
func (o *DeleteSubCategoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}