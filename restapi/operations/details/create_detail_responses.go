// Code generated by go-swagger; DO NOT EDIT.

package details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aeckard87/WornOut/models"
)

// CreateDetailCreatedCode is the HTTP code returned for type CreateDetailCreated
const CreateDetailCreatedCode int = 201

/*CreateDetailCreated Category Created

swagger:response createDetailCreated
*/
type CreateDetailCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Detail `json:"body,omitempty"`
}

// NewCreateDetailCreated creates CreateDetailCreated with default headers values
func NewCreateDetailCreated() *CreateDetailCreated {
	return &CreateDetailCreated{}
}

// WithPayload adds the payload to the create detail created response
func (o *CreateDetailCreated) WithPayload(payload *models.Detail) *CreateDetailCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create detail created response
func (o *CreateDetailCreated) SetPayload(payload *models.Detail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDetailCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDetailBadRequestCode is the HTTP code returned for type CreateDetailBadRequest
const CreateDetailBadRequestCode int = 400

/*CreateDetailBadRequest Invalid input

swagger:response createDetailBadRequest
*/
type CreateDetailBadRequest struct {
}

// NewCreateDetailBadRequest creates CreateDetailBadRequest with default headers values
func NewCreateDetailBadRequest() *CreateDetailBadRequest {
	return &CreateDetailBadRequest{}
}

// WriteResponse to the client
func (o *CreateDetailBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}
