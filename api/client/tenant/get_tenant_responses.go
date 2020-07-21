// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fi-ts/cloud-go/api/models"
)

// GetTenantReader is a Reader for the GetTenant structure.
type GetTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTenantDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantOK creates a GetTenantOK with default headers values
func NewGetTenantOK() *GetTenantOK {
	return &GetTenantOK{}
}

/*GetTenantOK handles this case with default header values.

OK
*/
type GetTenantOK struct {
	Payload *models.V1TenantResponse
}

func (o *GetTenantOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{id}][%d] getTenantOK  %+v", 200, o.Payload)
}

func (o *GetTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TenantResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantDefault creates a GetTenantDefault with default headers values
func NewGetTenantDefault(code int) *GetTenantDefault {
	return &GetTenantDefault{
		_statusCode: code,
	}
}

/*GetTenantDefault handles this case with default header values.

Error
*/
type GetTenantDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the get tenant default response
func (o *GetTenantDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantDefault) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{id}][%d] getTenant default  %+v", o._statusCode, o.Payload)
}

func (o *GetTenantDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
