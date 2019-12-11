// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/cloud-go/api/models"
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

/*UpdateIPOK handles this case with default header values.

OK
*/
type UpdateIPOK struct {
	Payload *models.ModelsV1IPResponse
}

func (o *UpdateIPOK) Error() string {
	return fmt.Sprintf("[POST /v1/ip][%d] updateIpOK  %+v", 200, o.Payload)
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

/*UpdateIPDefault handles this case with default header values.

Error
*/
type UpdateIPDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the update IP default response
func (o *UpdateIPDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIPDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ip][%d] updateIP default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
