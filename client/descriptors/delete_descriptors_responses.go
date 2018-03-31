// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDescriptorsReader is a Reader for the DeleteDescriptors structure.
type DeleteDescriptorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDescriptorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDescriptorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDescriptorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteDescriptorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDescriptorsOK creates a DeleteDescriptorsOK with default headers values
func NewDeleteDescriptorsOK() *DeleteDescriptorsOK {
	return &DeleteDescriptorsOK{}
}

/*DeleteDescriptorsOK handles this case with default header values.

Status Ok
*/
type DeleteDescriptorsOK struct {
}

func (o *DeleteDescriptorsOK) Error() string {
	return fmt.Sprintf("[DELETE /descriptors][%d] deleteDescriptorsOK ", 200)
}

func (o *DeleteDescriptorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDescriptorsBadRequest creates a DeleteDescriptorsBadRequest with default headers values
func NewDeleteDescriptorsBadRequest() *DeleteDescriptorsBadRequest {
	return &DeleteDescriptorsBadRequest{}
}

/*DeleteDescriptorsBadRequest handles this case with default header values.

Invalid input
*/
type DeleteDescriptorsBadRequest struct {
}

func (o *DeleteDescriptorsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /descriptors][%d] deleteDescriptorsBadRequest ", 400)
}

func (o *DeleteDescriptorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDescriptorsNotFound creates a DeleteDescriptorsNotFound with default headers values
func NewDeleteDescriptorsNotFound() *DeleteDescriptorsNotFound {
	return &DeleteDescriptorsNotFound{}
}

/*DeleteDescriptorsNotFound handles this case with default header values.

No items found
*/
type DeleteDescriptorsNotFound struct {
}

func (o *DeleteDescriptorsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /descriptors][%d] deleteDescriptorsNotFound ", 404)
}

func (o *DeleteDescriptorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}