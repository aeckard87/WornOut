// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteItemsReader is a Reader for the DeleteItems structure.
type DeleteItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteItemsOK creates a DeleteItemsOK with default headers values
func NewDeleteItemsOK() *DeleteItemsOK {
	return &DeleteItemsOK{}
}

/*DeleteItemsOK handles this case with default header values.

Status Ok
*/
type DeleteItemsOK struct {
}

func (o *DeleteItemsOK) Error() string {
	return fmt.Sprintf("[DELETE /items][%d] deleteItemsOK ", 200)
}

func (o *DeleteItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteItemsNotFound creates a DeleteItemsNotFound with default headers values
func NewDeleteItemsNotFound() *DeleteItemsNotFound {
	return &DeleteItemsNotFound{}
}

/*DeleteItemsNotFound handles this case with default header values.

No items found
*/
type DeleteItemsNotFound struct {
}

func (o *DeleteItemsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /items][%d] deleteItemsNotFound ", 404)
}

func (o *DeleteItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}