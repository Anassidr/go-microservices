// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-microservices/models"
)

// CreateProductReader is a Reader for the CreateProduct structure.
type CreateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewCreateProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProductOK creates a CreateProductOK with default headers values
func NewCreateProductOK() *CreateProductOK {
	return &CreateProductOK{}
}

/* CreateProductOK describes a response with status code 200, with default header values.

Data structure representing a single product
*/
type CreateProductOK struct {
	Payload *models.Product
}

// IsSuccess returns true when this create product o k response has a 2xx status code
func (o *CreateProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create product o k response has a 3xx status code
func (o *CreateProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create product o k response has a 4xx status code
func (o *CreateProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create product o k response has a 5xx status code
func (o *CreateProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create product o k response a status code equal to that given
func (o *CreateProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create product o k response
func (o *CreateProductOK) Code() int {
	return 200
}

func (o *CreateProductOK) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductOK  %+v", 200, o.Payload)
}

func (o *CreateProductOK) String() string {
	return fmt.Sprintf("[POST /products][%d] createProductOK  %+v", 200, o.Payload)
}

func (o *CreateProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *CreateProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductUnprocessableEntity creates a CreateProductUnprocessableEntity with default headers values
func NewCreateProductUnprocessableEntity() *CreateProductUnprocessableEntity {
	return &CreateProductUnprocessableEntity{}
}

/* CreateProductUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors defined as an array of strings
*/
type CreateProductUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this create product unprocessable entity response has a 2xx status code
func (o *CreateProductUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create product unprocessable entity response has a 3xx status code
func (o *CreateProductUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create product unprocessable entity response has a 4xx status code
func (o *CreateProductUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create product unprocessable entity response has a 5xx status code
func (o *CreateProductUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create product unprocessable entity response a status code equal to that given
func (o *CreateProductUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create product unprocessable entity response
func (o *CreateProductUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateProductUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /products][%d] createProductUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateProductUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductNotImplemented creates a CreateProductNotImplemented with default headers values
func NewCreateProductNotImplemented() *CreateProductNotImplemented {
	return &CreateProductNotImplemented{}
}

/* CreateProductNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type CreateProductNotImplemented struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this create product not implemented response has a 2xx status code
func (o *CreateProductNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create product not implemented response has a 3xx status code
func (o *CreateProductNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create product not implemented response has a 4xx status code
func (o *CreateProductNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this create product not implemented response has a 5xx status code
func (o *CreateProductNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this create product not implemented response a status code equal to that given
func (o *CreateProductNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the create product not implemented response
func (o *CreateProductNotImplemented) Code() int {
	return 501
}

func (o *CreateProductNotImplemented) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductNotImplemented  %+v", 501, o.Payload)
}

func (o *CreateProductNotImplemented) String() string {
	return fmt.Sprintf("[POST /products][%d] createProductNotImplemented  %+v", 501, o.Payload)
}

func (o *CreateProductNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
