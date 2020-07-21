// Code generated by go-swagger; DO NOT EDIT.

package s3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fi-ts/cloud-go/api/models"
)

// Gets3Reader is a Reader for the Gets3 structure.
type Gets3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Gets3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGets3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGets3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGets3OK creates a Gets3OK with default headers values
func NewGets3OK() *Gets3OK {
	return &Gets3OK{}
}

/*Gets3OK handles this case with default header values.

OK
*/
type Gets3OK struct {
	Payload *models.V1S3CredentialsResponse
}

func (o *Gets3OK) Error() string {
	return fmt.Sprintf("[GET /v1/s3][%d] gets3OK  %+v", 200, o.Payload)
}

func (o *Gets3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1S3CredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGets3Default creates a Gets3Default with default headers values
func NewGets3Default(code int) *Gets3Default {
	return &Gets3Default{
		_statusCode: code,
	}
}

/*Gets3Default handles this case with default header values.

Error
*/
type Gets3Default struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the gets3 default response
func (o *Gets3Default) Code() int {
	return o._statusCode
}

func (o *Gets3Default) Error() string {
	return fmt.Sprintf("[GET /v1/s3][%d] gets3 default  %+v", o._statusCode, o.Payload)
}

func (o *Gets3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
