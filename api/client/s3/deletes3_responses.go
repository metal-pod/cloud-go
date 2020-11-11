// Code generated by go-swagger; DO NOT EDIT.

package s3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
)

// Deletes3Reader is a Reader for the Deletes3 structure.
type Deletes3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Deletes3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletes3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletes3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletes3OK creates a Deletes3OK with default headers values
func NewDeletes3OK() *Deletes3OK {
	return &Deletes3OK{}
}

/*Deletes3OK handles this case with default header values.

OK
*/
type Deletes3OK struct {
	Payload *models.V1S3Response
}

func (o *Deletes3OK) Error() string {
	return fmt.Sprintf("[DELETE /v1/s3][%d] deletes3OK  %+v", 200, o.Payload)
}

func (o *Deletes3OK) GetPayload() *models.V1S3Response {
	return o.Payload
}

func (o *Deletes3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1S3Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletes3Default creates a Deletes3Default with default headers values
func NewDeletes3Default(code int) *Deletes3Default {
	return &Deletes3Default{
		_statusCode: code,
	}
}

/*Deletes3Default handles this case with default header values.

Error
*/
type Deletes3Default struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the deletes3 default response
func (o *Deletes3Default) Code() int {
	return o._statusCode
}

func (o *Deletes3Default) Error() string {
	return fmt.Sprintf("[DELETE /v1/s3][%d] deletes3 default  %+v", o._statusCode, o.Payload)
}

func (o *Deletes3Default) GetPayload() *models.HttperrorsHTTPErrorResponse {
	return o.Payload
}

func (o *Deletes3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
