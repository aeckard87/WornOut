// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aeckard87/WornOut/models"
)

// UpdateItemReader is a Reader for the UpdateItem structure.
type UpdateItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUpdateItemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateItemCreated creates a UpdateItemCreated with default headers values
func NewUpdateItemCreated() *UpdateItemCreated {
	return &UpdateItemCreated{}
}

/*UpdateItemCreated handles this case with default header values.

Category Created
*/
type UpdateItemCreated struct {
	Payload *models.Item
}

func (o *UpdateItemCreated) Error() string {
	return fmt.Sprintf("[PUT /items/{id}][%d] updateItemCreated  %+v", 201, o.Payload)
}

func (o *UpdateItemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateItemBadRequest creates a UpdateItemBadRequest with default headers values
func NewUpdateItemBadRequest() *UpdateItemBadRequest {
	return &UpdateItemBadRequest{}
}

/*UpdateItemBadRequest handles this case with default header values.

Invalid input
*/
type UpdateItemBadRequest struct {
}

func (o *UpdateItemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /items/{id}][%d] updateItemBadRequest ", 400)
}

func (o *UpdateItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
