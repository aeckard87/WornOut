// Code generated by go-swagger; DO NOT EDIT.

package subcategories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aeckard87/test/models"
)

// GetSubCategoriesReader is a Reader for the GetSubCategories structure.
type GetSubCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSubCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSubCategoriesOK creates a GetSubCategoriesOK with default headers values
func NewGetSubCategoriesOK() *GetSubCategoriesOK {
	return &GetSubCategoriesOK{}
}

/*GetSubCategoriesOK handles this case with default header values.

Status Ok
*/
type GetSubCategoriesOK struct {
	Payload models.SubCategories
}

func (o *GetSubCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /subcategories][%d] getSubCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetSubCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubCategoriesNotFound creates a GetSubCategoriesNotFound with default headers values
func NewGetSubCategoriesNotFound() *GetSubCategoriesNotFound {
	return &GetSubCategoriesNotFound{}
}

/*GetSubCategoriesNotFound handles this case with default header values.

No items found
*/
type GetSubCategoriesNotFound struct {
}

func (o *GetSubCategoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /subcategories][%d] getSubCategoriesNotFound ", 404)
}

func (o *GetSubCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
