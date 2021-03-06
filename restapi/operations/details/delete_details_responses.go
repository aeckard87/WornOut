// Code generated by go-swagger; DO NOT EDIT.

package details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDetailsOKCode is the HTTP code returned for type DeleteDetailsOK
const DeleteDetailsOKCode int = 200

/*DeleteDetailsOK Status Ok

swagger:response deleteDetailsOK
*/
type DeleteDetailsOK struct {
}

// NewDeleteDetailsOK creates DeleteDetailsOK with default headers values
func NewDeleteDetailsOK() *DeleteDetailsOK {
	return &DeleteDetailsOK{}
}

// WriteResponse to the client
func (o *DeleteDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteDetailsNotFoundCode is the HTTP code returned for type DeleteDetailsNotFound
const DeleteDetailsNotFoundCode int = 404

/*DeleteDetailsNotFound No items found

swagger:response deleteDetailsNotFound
*/
type DeleteDetailsNotFound struct {
}

// NewDeleteDetailsNotFound creates DeleteDetailsNotFound with default headers values
func NewDeleteDetailsNotFound() *DeleteDetailsNotFound {
	return &DeleteDetailsNotFound{}
}

// WriteResponse to the client
func (o *DeleteDetailsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
