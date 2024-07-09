// Code generated by go-swagger; DO NOT EDIT.

package ip

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

// UpdateIPReader is a Reader for the UpdateIP structure.
type UpdateIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIPOK creates a UpdateIPOK with default headers values
func NewUpdateIPOK() *UpdateIPOK {
	return &UpdateIPOK{}
}

/*
UpdateIPOK describes a response with status code 200, with default header values.

OK
*/
type UpdateIPOK struct {
	Payload *models.ModelsV1IPResponse
}

// IsSuccess returns true when this update Ip o k response has a 2xx status code
func (o *UpdateIPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Ip o k response has a 3xx status code
func (o *UpdateIPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Ip o k response has a 4xx status code
func (o *UpdateIPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Ip o k response has a 5xx status code
func (o *UpdateIPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update Ip o k response a status code equal to that given
func (o *UpdateIPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update Ip o k response
func (o *UpdateIPOK) Code() int {
	return 200
}

func (o *UpdateIPOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ip][%d] updateIpOK %s", 200, payload)
}

func (o *UpdateIPOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ip][%d] updateIpOK %s", 200, payload)
}

func (o *UpdateIPOK) GetPayload() *models.ModelsV1IPResponse {
	return o.Payload
}

func (o *UpdateIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsV1IPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPDefault creates a UpdateIPDefault with default headers values
func NewUpdateIPDefault(code int) *UpdateIPDefault {
	return &UpdateIPDefault{
		_statusCode: code,
	}
}

/*
UpdateIPDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateIPDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update IP default response has a 2xx status code
func (o *UpdateIPDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update IP default response has a 3xx status code
func (o *UpdateIPDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update IP default response has a 4xx status code
func (o *UpdateIPDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update IP default response has a 5xx status code
func (o *UpdateIPDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update IP default response a status code equal to that given
func (o *UpdateIPDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update IP default response
func (o *UpdateIPDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIPDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ip][%d] updateIP default %s", o._statusCode, payload)
}

func (o *UpdateIPDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ip][%d] updateIP default %s", o._statusCode, payload)
}

func (o *UpdateIPDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
