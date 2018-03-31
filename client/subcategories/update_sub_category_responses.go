// Code generated by go-swagger; DO NOT EDIT.

package subcategories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aeckard87/WornOut/models"
)

// UpdateSubCategoryReader is a Reader for the UpdateSubCategory structure.
type UpdateSubCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUpdateSubCategoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSubCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSubCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSubCategoryCreated creates a UpdateSubCategoryCreated with default headers values
func NewUpdateSubCategoryCreated() *UpdateSubCategoryCreated {
	return &UpdateSubCategoryCreated{}
}

/*UpdateSubCategoryCreated handles this case with default header values.

SubCategory Updated
*/
type UpdateSubCategoryCreated struct {
	Payload *models.SubCategory
}

func (o *UpdateSubCategoryCreated) Error() string {
	return fmt.Sprintf("[PUT /subcategories/{id}][%d] updateSubCategoryCreated  %+v", 201, o.Payload)
}

func (o *UpdateSubCategoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubCategoryBadRequest creates a UpdateSubCategoryBadRequest with default headers values
func NewUpdateSubCategoryBadRequest() *UpdateSubCategoryBadRequest {
	return &UpdateSubCategoryBadRequest{}
}

/*UpdateSubCategoryBadRequest handles this case with default header values.

Invalid input
*/
type UpdateSubCategoryBadRequest struct {
}

func (o *UpdateSubCategoryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subcategories/{id}][%d] updateSubCategoryBadRequest ", 400)
}

func (o *UpdateSubCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSubCategoryNotFound creates a UpdateSubCategoryNotFound with default headers values
func NewUpdateSubCategoryNotFound() *UpdateSubCategoryNotFound {
	return &UpdateSubCategoryNotFound{}
}

/*UpdateSubCategoryNotFound handles this case with default header values.

No items found
*/
type UpdateSubCategoryNotFound struct {
}

func (o *UpdateSubCategoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /subcategories/{id}][%d] updateSubCategoryNotFound ", 404)
}

func (o *UpdateSubCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}