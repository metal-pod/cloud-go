// Code generated by go-swagger; DO NOT EDIT.

package masterdata

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

// GetMasterdataReader is a Reader for the GetMasterdata structure.
type GetMasterdataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMasterdataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMasterdataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetMasterdataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMasterdataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMasterdataOK creates a GetMasterdataOK with default headers values
func NewGetMasterdataOK() *GetMasterdataOK {
	return &GetMasterdataOK{}
}

/*
GetMasterdataOK describes a response with status code 200, with default header values.

Ok
*/
type GetMasterdataOK struct {
	Payload *models.V1MasterdataLookupResponse
}

// IsSuccess returns true when this get masterdata o k response has a 2xx status code
func (o *GetMasterdataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get masterdata o k response has a 3xx status code
func (o *GetMasterdataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get masterdata o k response has a 4xx status code
func (o *GetMasterdataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get masterdata o k response has a 5xx status code
func (o *GetMasterdataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get masterdata o k response a status code equal to that given
func (o *GetMasterdataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get masterdata o k response
func (o *GetMasterdataOK) Code() int {
	return 200
}

func (o *GetMasterdataOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdataOK %s", 200, payload)
}

func (o *GetMasterdataOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdataOK %s", 200, payload)
}

func (o *GetMasterdataOK) GetPayload() *models.V1MasterdataLookupResponse {
	return o.Payload
}

func (o *GetMasterdataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MasterdataLookupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMasterdataNotFound creates a GetMasterdataNotFound with default headers values
func NewGetMasterdataNotFound() *GetMasterdataNotFound {
	return &GetMasterdataNotFound{}
}

/*
GetMasterdataNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetMasterdataNotFound struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get masterdata not found response has a 2xx status code
func (o *GetMasterdataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get masterdata not found response has a 3xx status code
func (o *GetMasterdataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get masterdata not found response has a 4xx status code
func (o *GetMasterdataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get masterdata not found response has a 5xx status code
func (o *GetMasterdataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get masterdata not found response a status code equal to that given
func (o *GetMasterdataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get masterdata not found response
func (o *GetMasterdataNotFound) Code() int {
	return 404
}

func (o *GetMasterdataNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdataNotFound %s", 404, payload)
}

func (o *GetMasterdataNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdataNotFound %s", 404, payload)
}

func (o *GetMasterdataNotFound) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetMasterdataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMasterdataDefault creates a GetMasterdataDefault with default headers values
func NewGetMasterdataDefault(code int) *GetMasterdataDefault {
	return &GetMasterdataDefault{
		_statusCode: code,
	}
}

/*
GetMasterdataDefault describes a response with status code -1, with default header values.

Error
*/
type GetMasterdataDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get masterdata default response has a 2xx status code
func (o *GetMasterdataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get masterdata default response has a 3xx status code
func (o *GetMasterdataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get masterdata default response has a 4xx status code
func (o *GetMasterdataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get masterdata default response has a 5xx status code
func (o *GetMasterdataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get masterdata default response a status code equal to that given
func (o *GetMasterdataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get masterdata default response
func (o *GetMasterdataDefault) Code() int {
	return o._statusCode
}

func (o *GetMasterdataDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdata default %s", o._statusCode, payload)
}

func (o *GetMasterdataDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdata default %s", o._statusCode, payload)
}

func (o *GetMasterdataDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetMasterdataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
