// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aeckard87/WornOut/models"
)

// GetItemsBySubCategoryOKCode is the HTTP code returned for type GetItemsBySubCategoryOK
const GetItemsBySubCategoryOKCode int = 200

/*GetItemsBySubCategoryOK Status Ok

swagger:response getItemsBySubCategoryOK
*/
type GetItemsBySubCategoryOK struct {

	/*
	  In: Body
	*/
	Payload models.Items `json:"body,omitempty"`
}

// NewGetItemsBySubCategoryOK creates GetItemsBySubCategoryOK with default headers values
func NewGetItemsBySubCategoryOK() *GetItemsBySubCategoryOK {
	return &GetItemsBySubCategoryOK{}
}

// WithPayload adds the payload to the get items by sub category o k response
func (o *GetItemsBySubCategoryOK) WithPayload(payload models.Items) *GetItemsBySubCategoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get items by sub category o k response
func (o *GetItemsBySubCategoryOK) SetPayload(payload models.Items) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemsBySubCategoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Items, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetItemsBySubCategoryBadRequestCode is the HTTP code returned for type GetItemsBySubCategoryBadRequest
const GetItemsBySubCategoryBadRequestCode int = 400

/*GetItemsBySubCategoryBadRequest Invalid input

swagger:response getItemsBySubCategoryBadRequest
*/
type GetItemsBySubCategoryBadRequest struct {
}

// NewGetItemsBySubCategoryBadRequest creates GetItemsBySubCategoryBadRequest with default headers values
func NewGetItemsBySubCategoryBadRequest() *GetItemsBySubCategoryBadRequest {
	return &GetItemsBySubCategoryBadRequest{}
}

// WriteResponse to the client
func (o *GetItemsBySubCategoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// GetItemsBySubCategoryNotFoundCode is the HTTP code returned for type GetItemsBySubCategoryNotFound
const GetItemsBySubCategoryNotFoundCode int = 404

/*GetItemsBySubCategoryNotFound No items found

swagger:response getItemsBySubCategoryNotFound
*/
type GetItemsBySubCategoryNotFound struct {
}

// NewGetItemsBySubCategoryNotFound creates GetItemsBySubCategoryNotFound with default headers values
func NewGetItemsBySubCategoryNotFound() *GetItemsBySubCategoryNotFound {
	return &GetItemsBySubCategoryNotFound{}
}

// WriteResponse to the client
func (o *GetItemsBySubCategoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}