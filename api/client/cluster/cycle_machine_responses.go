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

// CycleMachineReader is a Reader for the CycleMachine structure.
type CycleMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CycleMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCycleMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCycleMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCycleMachineOK creates a CycleMachineOK with default headers values
func NewCycleMachineOK() *CycleMachineOK {
	return &CycleMachineOK{}
}

/* CycleMachineOK describes a response with status code 200, with default header values.

OK
*/
type CycleMachineOK struct {
	Payload *models.V1ClusterResponse
}

func (o *CycleMachineOK) Error() string {
	return fmt.Sprintf("[POST /v1/cluster/{id}/cyclemachine][%d] cycleMachineOK  %+v", 200, o.Payload)
}
func (o *CycleMachineOK) GetPayload() *models.V1ClusterResponse {
	return o.Payload
}

func (o *CycleMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCycleMachineDefault creates a CycleMachineDefault with default headers values
func NewCycleMachineDefault(code int) *CycleMachineDefault {
	return &CycleMachineDefault{
		_statusCode: code,
	}
}

/* CycleMachineDefault describes a response with status code -1, with default header values.

Error
*/
type CycleMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the cycle machine default response
func (o *CycleMachineDefault) Code() int {
	return o._statusCode
}

func (o *CycleMachineDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cluster/{id}/cyclemachine][%d] cycleMachine default  %+v", o._statusCode, o.Payload)
}
func (o *CycleMachineDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CycleMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
