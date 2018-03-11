// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDescriptorReader is a Reader for the DeleteDescriptor structure.
type DeleteDescriptorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDescriptorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDescriptorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDescriptorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteDescriptorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDescriptorOK creates a DeleteDescriptorOK with default headers values
func NewDeleteDescriptorOK() *DeleteDescriptorOK {
	return &DeleteDescriptorOK{}
}

/*DeleteDescriptorOK handles this case with default header values.

Status Ok
*/
type DeleteDescriptorOK struct {
}

func (o *DeleteDescriptorOK) Error() string {
	return fmt.Sprintf("[DELETE /Descriptors/{id}][%d] deleteDescriptorOK ", 200)
}

func (o *DeleteDescriptorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDescriptorBadRequest creates a DeleteDescriptorBadRequest with default headers values
func NewDeleteDescriptorBadRequest() *DeleteDescriptorBadRequest {
	return &DeleteDescriptorBadRequest{}
}

/*DeleteDescriptorBadRequest handles this case with default header values.

Invalid input
*/
type DeleteDescriptorBadRequest struct {
}

func (o *DeleteDescriptorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /Descriptors/{id}][%d] deleteDescriptorBadRequest ", 400)
}

func (o *DeleteDescriptorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDescriptorNotFound creates a DeleteDescriptorNotFound with default headers values
func NewDeleteDescriptorNotFound() *DeleteDescriptorNotFound {
	return &DeleteDescriptorNotFound{}
}

/*DeleteDescriptorNotFound handles this case with default header values.

No items found
*/
type DeleteDescriptorNotFound struct {
}

func (o *DeleteDescriptorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /Descriptors/{id}][%d] deleteDescriptorNotFound ", 404)
}

func (o *DeleteDescriptorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
