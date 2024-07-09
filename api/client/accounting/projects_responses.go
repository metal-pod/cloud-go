// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

// ProjectsReader is a Reader for the Projects structure.
type ProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectsOK creates a ProjectsOK with default headers values
func NewProjectsOK() *ProjectsOK {
	return &ProjectsOK{}
}

/*
ProjectsOK describes a response with status code 200, with default header values.

OK
*/
type ProjectsOK struct {
	Payload []*models.V1ProjectInfoResponse
}

// IsSuccess returns true when this projects o k response has a 2xx status code
func (o *ProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects o k response has a 3xx status code
func (o *ProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects o k response has a 4xx status code
func (o *ProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects o k response has a 5xx status code
func (o *ProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects o k response a status code equal to that given
func (o *ProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the projects o k response
func (o *ProjectsOK) Code() int {
	return 200
}

func (o *ProjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/projects][%d] projectsOK %s", 200, payload)
}

func (o *ProjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/projects][%d] projectsOK %s", 200, payload)
}

func (o *ProjectsOK) GetPayload() []*models.V1ProjectInfoResponse {
	return o.Payload
}

func (o *ProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDefault creates a ProjectsDefault with default headers values
func NewProjectsDefault(code int) *ProjectsDefault {
	return &ProjectsDefault{
		_statusCode: code,
	}
}

/*
ProjectsDefault describes a response with status code -1, with default header values.

Error
*/
type ProjectsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this projects default response has a 2xx status code
func (o *ProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this projects default response has a 3xx status code
func (o *ProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this projects default response has a 4xx status code
func (o *ProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this projects default response has a 5xx status code
func (o *ProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this projects default response a status code equal to that given
func (o *ProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the projects default response
func (o *ProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/projects][%d] projects default %s", o._statusCode, payload)
}

func (o *ProjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/projects][%d] projects default %s", o._statusCode, payload)
}

func (o *ProjectsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
