// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anassidr/go-microservices/product-api/sdk/models"
)

// ListProductsReader is a Reader for the ListProducts structure.
type ListProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProductsOK creates a ListProductsOK with default headers values
func NewListProductsOK() *ListProductsOK {
	return &ListProductsOK{}
}

/* ListProductsOK describes a response with status code 200, with default header values.

A list of products
*/
type ListProductsOK struct {
	Payload []*models.Product
}

// IsSuccess returns true when this list products o k response has a 2xx status code
func (o *ListProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list products o k response has a 3xx status code
func (o *ListProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list products o k response has a 4xx status code
func (o *ListProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list products o k response has a 5xx status code
func (o *ListProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list products o k response a status code equal to that given
func (o *ListProductsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list products o k response
func (o *ListProductsOK) Code() int {
	return 200
}

func (o *ListProductsOK) Error() string {
	return fmt.Sprintf("[GET /products][%d] listProductsOK  %+v", 200, o.Payload)
}

func (o *ListProductsOK) String() string {
	return fmt.Sprintf("[GET /products][%d] listProductsOK  %+v", 200, o.Payload)
}

func (o *ListProductsOK) GetPayload() []*models.Product {
	return o.Payload
}

func (o *ListProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
