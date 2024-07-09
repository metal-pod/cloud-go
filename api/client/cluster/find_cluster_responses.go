// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// FindClusterReader is a Reader for the FindCluster structure.
type FindClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindClusterOK creates a FindClusterOK with default headers values
func NewFindClusterOK() *FindClusterOK {
	return &FindClusterOK{}
}

/*
FindClusterOK describes a response with status code 200, with default header values.

OK
*/
type FindClusterOK struct {
	Payload *models.V1ClusterResponse
}

// IsSuccess returns true when this find cluster o k response has a 2xx status code
func (o *FindClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find cluster o k response has a 3xx status code
func (o *FindClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find cluster o k response has a 4xx status code
func (o *FindClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find cluster o k response has a 5xx status code
func (o *FindClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find cluster o k response a status code equal to that given
func (o *FindClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find cluster o k response
func (o *FindClusterOK) Code() int {
	return 200
}

func (o *FindClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}][%d] findClusterOK %s", 200, payload)
}

func (o *FindClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}][%d] findClusterOK %s", 200, payload)
}

func (o *FindClusterOK) GetPayload() *models.V1ClusterResponse {
	return o.Payload
}

func (o *FindClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindClusterDefault creates a FindClusterDefault with default headers values
func NewFindClusterDefault(code int) *FindClusterDefault {
	return &FindClusterDefault{
		_statusCode: code,
	}
}

/*
FindClusterDefault describes a response with status code -1, with default header values.

Error
*/
type FindClusterDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this find cluster default response has a 2xx status code
func (o *FindClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find cluster default response has a 3xx status code
func (o *FindClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find cluster default response has a 4xx status code
func (o *FindClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find cluster default response has a 5xx status code
func (o *FindClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find cluster default response a status code equal to that given
func (o *FindClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the find cluster default response
func (o *FindClusterDefault) Code() int {
	return o._statusCode
}

func (o *FindClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}][%d] findCluster default %s", o._statusCode, payload)
}

func (o *FindClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/cluster/{id}][%d] findCluster default %s", o._statusCode, payload)
}

func (o *FindClusterDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
