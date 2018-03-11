// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDescriptorsByDetailOKCode is the HTTP code returned for type DeleteDescriptorsByDetailOK
const DeleteDescriptorsByDetailOKCode int = 200

/*DeleteDescriptorsByDetailOK Status Ok

swagger:response deleteDescriptorsByDetailOK
*/
type DeleteDescriptorsByDetailOK struct {
}

// NewDeleteDescriptorsByDetailOK creates DeleteDescriptorsByDetailOK with default headers values
func NewDeleteDescriptorsByDetailOK() *DeleteDescriptorsByDetailOK {
	return &DeleteDescriptorsByDetailOK{}
}

// WriteResponse to the client
func (o *DeleteDescriptorsByDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteDescriptorsByDetailNotFoundCode is the HTTP code returned for type DeleteDescriptorsByDetailNotFound
const DeleteDescriptorsByDetailNotFoundCode int = 404

/*DeleteDescriptorsByDetailNotFound No items found

swagger:response deleteDescriptorsByDetailNotFound
*/
type DeleteDescriptorsByDetailNotFound struct {
}

// NewDeleteDescriptorsByDetailNotFound creates DeleteDescriptorsByDetailNotFound with default headers values
func NewDeleteDescriptorsByDetailNotFound() *DeleteDescriptorsByDetailNotFound {
	return &DeleteDescriptorsByDetailNotFound{}
}

// WriteResponse to the client
func (o *DeleteDescriptorsByDetailNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
