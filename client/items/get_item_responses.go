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

// GetItemReader is a Reader for the GetItem structure.
type GetItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetItemOK creates a GetItemOK with default headers values
func NewGetItemOK() *GetItemOK {
	return &GetItemOK{}
}

/*GetItemOK handles this case with default header values.

Status Ok
*/
type GetItemOK struct {
	Payload models.Item
}

func (o *GetItemOK) Error() string {
	return fmt.Sprintf("[GET /items/{id}][%d] getItemOK  %+v", 200, o.Payload)
}

func (o *GetItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemNotFound creates a GetItemNotFound with default headers values
func NewGetItemNotFound() *GetItemNotFound {
	return &GetItemNotFound{}
}

/*GetItemNotFound handles this case with default header values.

No items found
*/
type GetItemNotFound struct {
}

func (o *GetItemNotFound) Error() string {
	return fmt.Sprintf("[GET /items/{id}][%d] getItemNotFound ", 404)
}

func (o *GetItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
