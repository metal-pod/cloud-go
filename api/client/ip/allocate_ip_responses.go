// Code generated by go-swagger; DO NOT EDIT.

package ip

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

// AllocateIPReader is a Reader for the AllocateIP structure.
type AllocateIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocateIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAllocateIPCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAllocateIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllocateIPCreated creates a AllocateIPCreated with default headers values
func NewAllocateIPCreated() *AllocateIPCreated {
	return &AllocateIPCreated{}
}

/*
AllocateIPCreated describes a response with status code 201, with default header values.

Created
*/
type AllocateIPCreated struct {
	Payload *models.ModelsV1IPResponse
}

// IsSuccess returns true when this allocate Ip created response has a 2xx status code
func (o *AllocateIPCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this allocate Ip created response has a 3xx status code
func (o *AllocateIPCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allocate Ip created response has a 4xx status code
func (o *AllocateIPCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this allocate Ip created response has a 5xx status code
func (o *AllocateIPCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this allocate Ip created response a status code equal to that given
func (o *AllocateIPCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the allocate Ip created response
func (o *AllocateIPCreated) Code() int {
	return 201
}

func (o *AllocateIPCreated) Error() string {
	return fmt.Sprintf("[POST /v1/ip/allocate][%d] allocateIpCreated  %+v", 201, o.Payload)
}

func (o *AllocateIPCreated) String() string {
	return fmt.Sprintf("[POST /v1/ip/allocate][%d] allocateIpCreated  %+v", 201, o.Payload)
}

func (o *AllocateIPCreated) GetPayload() *models.ModelsV1IPResponse {
	return o.Payload
}

func (o *AllocateIPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsV1IPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateIPDefault creates a AllocateIPDefault with default headers values
func NewAllocateIPDefault(code int) *AllocateIPDefault {
	return &AllocateIPDefault{
		_statusCode: code,
	}
}

/*
AllocateIPDefault describes a response with status code -1, with default header values.

Error
*/
type AllocateIPDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this allocate IP default response has a 2xx status code
func (o *AllocateIPDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this allocate IP default response has a 3xx status code
func (o *AllocateIPDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this allocate IP default response has a 4xx status code
func (o *AllocateIPDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this allocate IP default response has a 5xx status code
func (o *AllocateIPDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this allocate IP default response a status code equal to that given
func (o *AllocateIPDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the allocate IP default response
func (o *AllocateIPDefault) Code() int {
	return o._statusCode
}

func (o *AllocateIPDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ip/allocate][%d] allocateIP default  %+v", o._statusCode, o.Payload)
}

func (o *AllocateIPDefault) String() string {
	return fmt.Sprintf("[POST /v1/ip/allocate][%d] allocateIP default  %+v", o._statusCode, o.Payload)
}

func (o *AllocateIPDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AllocateIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
