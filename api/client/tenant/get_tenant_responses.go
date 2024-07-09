// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
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

/*
GetTenantOK describes a response with status code 200, with default header values.

OK
*/
type GetTenantOK struct {
	Payload *models.V1TenantResponse
}

// IsSuccess returns true when this get tenant o k response has a 2xx status code
func (o *GetTenantOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tenant o k response has a 3xx status code
func (o *GetTenantOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant o k response has a 4xx status code
func (o *GetTenantOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tenant o k response has a 5xx status code
func (o *GetTenantOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant o k response a status code equal to that given
func (o *GetTenantOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tenant o k response
func (o *GetTenantOK) Code() int {
	return 200
}

func (o *GetTenantOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/tenant/{id}][%d] getTenantOK %s", 200, payload)
}

func (o *GetTenantOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/tenant/{id}][%d] getTenantOK %s", 200, payload)
}

func (o *GetTenantOK) GetPayload() *models.V1TenantResponse {
	return o.Payload
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

/*
GetTenantDefault describes a response with status code -1, with default header values.

Error
*/
type GetTenantDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get tenant default response has a 2xx status code
func (o *GetTenantDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get tenant default response has a 3xx status code
func (o *GetTenantDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get tenant default response has a 4xx status code
func (o *GetTenantDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get tenant default response has a 5xx status code
func (o *GetTenantDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get tenant default response a status code equal to that given
func (o *GetTenantDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get tenant default response
func (o *GetTenantDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/tenant/{id}][%d] getTenant default %s", o._statusCode, payload)
}

func (o *GetTenantDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/tenant/{id}][%d] getTenant default %s", o._statusCode, payload)
}

func (o *GetTenantDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetTenantDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
