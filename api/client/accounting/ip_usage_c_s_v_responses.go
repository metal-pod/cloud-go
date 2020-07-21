// Code generated by go-swagger; DO NOT EDIT.

package accounting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fi-ts/cloud-go/api/models"
)

// IPUsageCSVReader is a Reader for the IPUsageCSV structure.
type IPUsageCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPUsageCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPUsageCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewIPUsageCSVDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPUsageCSVOK creates a IPUsageCSVOK with default headers values
func NewIPUsageCSVOK() *IPUsageCSVOK {
	return &IPUsageCSVOK{}
}

/*IPUsageCSVOK handles this case with default header values.

OK
*/
type IPUsageCSVOK struct {
	Payload string
}

func (o *IPUsageCSVOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/ip-usage-csv][%d] ipUsageCSVOK  %+v", 200, o.Payload)
}

func (o *IPUsageCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPUsageCSVDefault creates a IPUsageCSVDefault with default headers values
func NewIPUsageCSVDefault(code int) *IPUsageCSVDefault {
	return &IPUsageCSVDefault{
		_statusCode: code,
	}
}

/*IPUsageCSVDefault handles this case with default header values.

Error
*/
type IPUsageCSVDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the ip usage c s v default response
func (o *IPUsageCSVDefault) Code() int {
	return o._statusCode
}

func (o *IPUsageCSVDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/ip-usage-csv][%d] ipUsageCSV default  %+v", o._statusCode, o.Payload)
}

func (o *IPUsageCSVDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
