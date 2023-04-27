// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// ListConstraintsReader is a Reader for the ListConstraints structure.
type ListConstraintsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConstraintsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConstraintsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListConstraintsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListConstraintsOK creates a ListConstraintsOK with default headers values
func NewListConstraintsOK() *ListConstraintsOK {
	return &ListConstraintsOK{}
}

/*
ListConstraintsOK describes a response with status code 200, with default header values.

OK
*/
type ListConstraintsOK struct {
	Payload *models.V1ShootConstraints
}

// IsSuccess returns true when this list constraints o k response has a 2xx status code
func (o *ListConstraintsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list constraints o k response has a 3xx status code
func (o *ListConstraintsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraints o k response has a 4xx status code
func (o *ListConstraintsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list constraints o k response has a 5xx status code
func (o *ListConstraintsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraints o k response a status code equal to that given
func (o *ListConstraintsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list constraints o k response
func (o *ListConstraintsOK) Code() int {
	return 200
}

func (o *ListConstraintsOK) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/constraints][%d] listConstraintsOK  %+v", 200, o.Payload)
}

func (o *ListConstraintsOK) String() string {
	return fmt.Sprintf("[GET /v1/cluster/constraints][%d] listConstraintsOK  %+v", 200, o.Payload)
}

func (o *ListConstraintsOK) GetPayload() *models.V1ShootConstraints {
	return o.Payload
}

func (o *ListConstraintsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ShootConstraints)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConstraintsDefault creates a ListConstraintsDefault with default headers values
func NewListConstraintsDefault(code int) *ListConstraintsDefault {
	return &ListConstraintsDefault{
		_statusCode: code,
	}
}

/*
ListConstraintsDefault describes a response with status code -1, with default header values.

Error
*/
type ListConstraintsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this list constraints default response has a 2xx status code
func (o *ListConstraintsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list constraints default response has a 3xx status code
func (o *ListConstraintsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list constraints default response has a 4xx status code
func (o *ListConstraintsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list constraints default response has a 5xx status code
func (o *ListConstraintsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list constraints default response a status code equal to that given
func (o *ListConstraintsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list constraints default response
func (o *ListConstraintsDefault) Code() int {
	return o._statusCode
}

func (o *ListConstraintsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/constraints][%d] listConstraints default  %+v", o._statusCode, o.Payload)
}

func (o *ListConstraintsDefault) String() string {
	return fmt.Sprintf("[GET /v1/cluster/constraints][%d] listConstraints default  %+v", o._statusCode, o.Payload)
}

func (o *ListConstraintsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListConstraintsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
