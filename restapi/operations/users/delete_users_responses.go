// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteUsersOKCode is the HTTP code returned for type DeleteUsersOK
const DeleteUsersOKCode int = 200

/*DeleteUsersOK Status Ok

swagger:response deleteUsersOK
*/
type DeleteUsersOK struct {
}

// NewDeleteUsersOK creates DeleteUsersOK with default headers values
func NewDeleteUsersOK() *DeleteUsersOK {
	return &DeleteUsersOK{}
}

// WriteResponse to the client
func (o *DeleteUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// DeleteUsersNotFoundCode is the HTTP code returned for type DeleteUsersNotFound
const DeleteUsersNotFoundCode int = 404

/*DeleteUsersNotFound No items found

swagger:response deleteUsersNotFound
*/
type DeleteUsersNotFound struct {
}

// NewDeleteUsersNotFound creates DeleteUsersNotFound with default headers values
func NewDeleteUsersNotFound() *DeleteUsersNotFound {
	return &DeleteUsersNotFound{}
}

// WriteResponse to the client
func (o *DeleteUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
