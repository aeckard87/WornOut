// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aeckard87/WornOut/models"
)

// GetDescriptorOKCode is the HTTP code returned for type GetDescriptorOK
const GetDescriptorOKCode int = 200

/*GetDescriptorOK Status Ok

swagger:response getDescriptorOK
*/
type GetDescriptorOK struct {

	/*
	  In: Body
	*/
	Payload *models.Descriptor `json:"body,omitempty"`
}

// NewGetDescriptorOK creates GetDescriptorOK with default headers values
func NewGetDescriptorOK() *GetDescriptorOK {
	return &GetDescriptorOK{}
}

// WithPayload adds the payload to the get descriptor o k response
func (o *GetDescriptorOK) WithPayload(payload *models.Descriptor) *GetDescriptorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get descriptor o k response
func (o *GetDescriptorOK) SetPayload(payload *models.Descriptor) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDescriptorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDescriptorBadRequestCode is the HTTP code returned for type GetDescriptorBadRequest
const GetDescriptorBadRequestCode int = 400

/*GetDescriptorBadRequest Invalid input

swagger:response getDescriptorBadRequest
*/
type GetDescriptorBadRequest struct {
}

// NewGetDescriptorBadRequest creates GetDescriptorBadRequest with default headers values
func NewGetDescriptorBadRequest() *GetDescriptorBadRequest {
	return &GetDescriptorBadRequest{}
}

// WriteResponse to the client
func (o *GetDescriptorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// GetDescriptorNotFoundCode is the HTTP code returned for type GetDescriptorNotFound
const GetDescriptorNotFoundCode int = 404

/*GetDescriptorNotFound No items found

swagger:response getDescriptorNotFound
*/
type GetDescriptorNotFound struct {
}

// NewGetDescriptorNotFound creates GetDescriptorNotFound with default headers values
func NewGetDescriptorNotFound() *GetDescriptorNotFound {
	return &GetDescriptorNotFound{}
}

// WriteResponse to the client
func (o *GetDescriptorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}