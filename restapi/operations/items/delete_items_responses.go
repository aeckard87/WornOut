// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteItemsOKCode is the HTTP code returned for type DeleteItemsOK
const DeleteItemsOKCode int = 200

/*DeleteItemsOK Status Ok

swagger:response deleteItemsOK
*/
type DeleteItemsOK struct {
}

// NewDeleteItemsOK creates DeleteItemsOK with default headers values
func NewDeleteItemsOK() *DeleteItemsOK {
	return &DeleteItemsOK{}
}

// WriteResponse to the client
func (o *DeleteItemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteItemsNotFoundCode is the HTTP code returned for type DeleteItemsNotFound
const DeleteItemsNotFoundCode int = 404

/*DeleteItemsNotFound No items found

swagger:response deleteItemsNotFound
*/
type DeleteItemsNotFound struct {
}

// NewDeleteItemsNotFound creates DeleteItemsNotFound with default headers values
func NewDeleteItemsNotFound() *DeleteItemsNotFound {
	return &DeleteItemsNotFound{}
}

// WriteResponse to the client
func (o *DeleteItemsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
