// Code generated by go-swagger; DO NOT EDIT.

package gateway

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

// DescribeReader is a Reader for the Describe structure.
type DescribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeOK creates a DescribeOK with default headers values
func NewDescribeOK() *DescribeOK {
	return &DescribeOK{}
}

/* DescribeOK describes a response with status code 200, with default header values.

gateway info fetched
*/
type DescribeOK struct {
	Payload *models.V1GatewayResponse
}

func (o *DescribeOK) Error() string {
	return fmt.Sprintf("[GET /v1/gateway/{id}][%d] describeOK  %+v", 200, o.Payload)
}
func (o *DescribeOK) GetPayload() *models.V1GatewayResponse {
	return o.Payload
}

func (o *DescribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GatewayResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDefault creates a DescribeDefault with default headers values
func NewDescribeDefault(code int) *DescribeDefault {
	return &DescribeDefault{
		_statusCode: code,
	}
}

/* DescribeDefault describes a response with status code -1, with default header values.

error
*/
type DescribeDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the describe default response
func (o *DescribeDefault) Code() int {
	return o._statusCode
}

func (o *DescribeDefault) Error() string {
	return fmt.Sprintf("[GET /v1/gateway/{id}][%d] describe default  %+v", o._statusCode, o.Payload)
}
func (o *DescribeDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DescribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
