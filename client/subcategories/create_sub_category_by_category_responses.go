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

// CreateSubCategoryByCategoryReader is a Reader for the CreateSubCategoryByCategory structure.
type CreateSubCategoryByCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubCategoryByCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSubCategoryByCategoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSubCategoryByCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSubCategoryByCategoryCreated creates a CreateSubCategoryByCategoryCreated with default headers values
func NewCreateSubCategoryByCategoryCreated() *CreateSubCategoryByCategoryCreated {
	return &CreateSubCategoryByCategoryCreated{}
}

/*CreateSubCategoryByCategoryCreated handles this case with default header values.

SubCategory Created
*/
type CreateSubCategoryByCategoryCreated struct {
	Payload *models.SubCategory
}

func (o *CreateSubCategoryByCategoryCreated) Error() string {
	return fmt.Sprintf("[POST /catagories/{id}/subcategories][%d] createSubCategoryByCategoryCreated  %+v", 201, o.Payload)
}

func (o *CreateSubCategoryByCategoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubCategoryByCategoryBadRequest creates a CreateSubCategoryByCategoryBadRequest with default headers values
func NewCreateSubCategoryByCategoryBadRequest() *CreateSubCategoryByCategoryBadRequest {
	return &CreateSubCategoryByCategoryBadRequest{}
}

/*CreateSubCategoryByCategoryBadRequest handles this case with default header values.

Invalid input
*/
type CreateSubCategoryByCategoryBadRequest struct {
}

func (o *CreateSubCategoryByCategoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /catagories/{id}/subcategories][%d] createSubCategoryByCategoryBadRequest ", 400)
}

func (o *CreateSubCategoryByCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}