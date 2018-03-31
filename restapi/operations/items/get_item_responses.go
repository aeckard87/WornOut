// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aeckard87/WornOut/models"
)

// GetItemOKCode is the HTTP code returned for type GetItemOK
const GetItemOKCode int = 200

/*GetItemOK Status Ok

swagger:response getItemOK
*/
type GetItemOK struct {

	/*
	  In: Body
	*/
	Payload models.Items `json:"body,omitempty"`
}

// NewGetItemOK creates GetItemOK with default headers values
func NewGetItemOK() *GetItemOK {
	return &GetItemOK{}
}

// WithPayload adds the payload to the get item o k response
func (o *GetItemOK) WithPayload(payload models.Items) *GetItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get item o k response
func (o *GetItemOK) SetPayload(payload models.Items) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Items, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetItemNotFoundCode is the HTTP code returned for type GetItemNotFound
const GetItemNotFoundCode int = 404

/*GetItemNotFound No items found

swagger:response getItemNotFound
*/
type GetItemNotFound struct {
}

// NewGetItemNotFound creates GetItemNotFound with default headers values
func NewGetItemNotFound() *GetItemNotFound {
	return &GetItemNotFound{}
}

// WriteResponse to the client
func (o *GetItemNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}