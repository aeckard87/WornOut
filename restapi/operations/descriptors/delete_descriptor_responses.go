// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDescriptorOKCode is the HTTP code returned for type DeleteDescriptorOK
const DeleteDescriptorOKCode int = 200

/*DeleteDescriptorOK Status Ok

swagger:response deleteDescriptorOK
*/
type DeleteDescriptorOK struct {
}

// NewDeleteDescriptorOK creates DeleteDescriptorOK with default headers values
func NewDeleteDescriptorOK() *DeleteDescriptorOK {
	return &DeleteDescriptorOK{}
}

// WriteResponse to the client
func (o *DeleteDescriptorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteDescriptorBadRequestCode is the HTTP code returned for type DeleteDescriptorBadRequest
const DeleteDescriptorBadRequestCode int = 400

/*DeleteDescriptorBadRequest Invalid input

swagger:response deleteDescriptorBadRequest
*/
type DeleteDescriptorBadRequest struct {
}

// NewDeleteDescriptorBadRequest creates DeleteDescriptorBadRequest with default headers values
func NewDeleteDescriptorBadRequest() *DeleteDescriptorBadRequest {
	return &DeleteDescriptorBadRequest{}
}

// WriteResponse to the client
func (o *DeleteDescriptorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// DeleteDescriptorNotFoundCode is the HTTP code returned for type DeleteDescriptorNotFound
const DeleteDescriptorNotFoundCode int = 404

/*DeleteDescriptorNotFound No items found

swagger:response deleteDescriptorNotFound
*/
type DeleteDescriptorNotFound struct {
}

// NewDeleteDescriptorNotFound creates DeleteDescriptorNotFound with default headers values
func NewDeleteDescriptorNotFound() *DeleteDescriptorNotFound {
	return &DeleteDescriptorNotFound{}
}

// WriteResponse to the client
func (o *DeleteDescriptorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
