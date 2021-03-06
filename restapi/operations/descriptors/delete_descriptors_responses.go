// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDescriptorsOKCode is the HTTP code returned for type DeleteDescriptorsOK
const DeleteDescriptorsOKCode int = 200

/*DeleteDescriptorsOK Status Ok

swagger:response deleteDescriptorsOK
*/
type DeleteDescriptorsOK struct {
}

// NewDeleteDescriptorsOK creates DeleteDescriptorsOK with default headers values
func NewDeleteDescriptorsOK() *DeleteDescriptorsOK {
	return &DeleteDescriptorsOK{}
}

// WriteResponse to the client
func (o *DeleteDescriptorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteDescriptorsBadRequestCode is the HTTP code returned for type DeleteDescriptorsBadRequest
const DeleteDescriptorsBadRequestCode int = 400

/*DeleteDescriptorsBadRequest Invalid input

swagger:response deleteDescriptorsBadRequest
*/
type DeleteDescriptorsBadRequest struct {
}

// NewDeleteDescriptorsBadRequest creates DeleteDescriptorsBadRequest with default headers values
func NewDeleteDescriptorsBadRequest() *DeleteDescriptorsBadRequest {
	return &DeleteDescriptorsBadRequest{}
}

// WriteResponse to the client
func (o *DeleteDescriptorsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// DeleteDescriptorsNotFoundCode is the HTTP code returned for type DeleteDescriptorsNotFound
const DeleteDescriptorsNotFoundCode int = 404

/*DeleteDescriptorsNotFound No items found

swagger:response deleteDescriptorsNotFound
*/
type DeleteDescriptorsNotFound struct {
}

// NewDeleteDescriptorsNotFound creates DeleteDescriptorsNotFound with default headers values
func NewDeleteDescriptorsNotFound() *DeleteDescriptorsNotFound {
	return &DeleteDescriptorsNotFound{}
}

// WriteResponse to the client
func (o *DeleteDescriptorsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
