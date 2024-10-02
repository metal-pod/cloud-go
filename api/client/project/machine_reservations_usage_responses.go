// Code generated by go-swagger; DO NOT EDIT.

package project

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

// MachineReservationsUsageReader is a Reader for the MachineReservationsUsage structure.
type MachineReservationsUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineReservationsUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachineReservationsUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMachineReservationsUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMachineReservationsUsageOK creates a MachineReservationsUsageOK with default headers values
func NewMachineReservationsUsageOK() *MachineReservationsUsageOK {
	return &MachineReservationsUsageOK{}
}

/*
MachineReservationsUsageOK describes a response with status code 200, with default header values.

OK
*/
type MachineReservationsUsageOK struct {
	Payload []*models.V1MachineReservationUsageResponse
}

// IsSuccess returns true when this machine reservations usage o k response has a 2xx status code
func (o *MachineReservationsUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machine reservations usage o k response has a 3xx status code
func (o *MachineReservationsUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machine reservations usage o k response has a 4xx status code
func (o *MachineReservationsUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machine reservations usage o k response has a 5xx status code
func (o *MachineReservationsUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machine reservations usage o k response a status code equal to that given
func (o *MachineReservationsUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the machine reservations usage o k response
func (o *MachineReservationsUsageOK) Code() int {
	return 200
}

func (o *MachineReservationsUsageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine/usage][%d] machineReservationsUsageOK %s", 200, payload)
}

func (o *MachineReservationsUsageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine/usage][%d] machineReservationsUsageOK %s", 200, payload)
}

func (o *MachineReservationsUsageOK) GetPayload() []*models.V1MachineReservationUsageResponse {
	return o.Payload
}

func (o *MachineReservationsUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachineReservationsUsageDefault creates a MachineReservationsUsageDefault with default headers values
func NewMachineReservationsUsageDefault(code int) *MachineReservationsUsageDefault {
	return &MachineReservationsUsageDefault{
		_statusCode: code,
	}
}

/*
MachineReservationsUsageDefault describes a response with status code -1, with default header values.

Error
*/
type MachineReservationsUsageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this machine reservations usage default response has a 2xx status code
func (o *MachineReservationsUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this machine reservations usage default response has a 3xx status code
func (o *MachineReservationsUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this machine reservations usage default response has a 4xx status code
func (o *MachineReservationsUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this machine reservations usage default response has a 5xx status code
func (o *MachineReservationsUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this machine reservations usage default response a status code equal to that given
func (o *MachineReservationsUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the machine reservations usage default response
func (o *MachineReservationsUsageDefault) Code() int {
	return o._statusCode
}

func (o *MachineReservationsUsageDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine/usage][%d] machineReservationsUsage default %s", o._statusCode, payload)
}

func (o *MachineReservationsUsageDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine/usage][%d] machineReservationsUsage default %s", o._statusCode, payload)
}

func (o *MachineReservationsUsageDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *MachineReservationsUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
