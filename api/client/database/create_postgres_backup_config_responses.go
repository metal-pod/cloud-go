// Code generated by go-swagger; DO NOT EDIT.

package database

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

// CreatePostgresBackupConfigReader is a Reader for the CreatePostgresBackupConfig structure.
type CreatePostgresBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePostgresBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePostgresBackupConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePostgresBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePostgresBackupConfigCreated creates a CreatePostgresBackupConfigCreated with default headers values
func NewCreatePostgresBackupConfigCreated() *CreatePostgresBackupConfigCreated {
	return &CreatePostgresBackupConfigCreated{}
}

/*
CreatePostgresBackupConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreatePostgresBackupConfigCreated struct {
	Payload *models.V1PostgresBackupConfigResponse
}

// IsSuccess returns true when this create postgres backup config created response has a 2xx status code
func (o *CreatePostgresBackupConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create postgres backup config created response has a 3xx status code
func (o *CreatePostgresBackupConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create postgres backup config created response has a 4xx status code
func (o *CreatePostgresBackupConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create postgres backup config created response has a 5xx status code
func (o *CreatePostgresBackupConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create postgres backup config created response a status code equal to that given
func (o *CreatePostgresBackupConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create postgres backup config created response
func (o *CreatePostgresBackupConfigCreated) Code() int {
	return 201
}

func (o *CreatePostgresBackupConfigCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/database/postgres/backup-config][%d] createPostgresBackupConfigCreated %s", 201, payload)
}

func (o *CreatePostgresBackupConfigCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/database/postgres/backup-config][%d] createPostgresBackupConfigCreated %s", 201, payload)
}

func (o *CreatePostgresBackupConfigCreated) GetPayload() *models.V1PostgresBackupConfigResponse {
	return o.Payload
}

func (o *CreatePostgresBackupConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PostgresBackupConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePostgresBackupConfigDefault creates a CreatePostgresBackupConfigDefault with default headers values
func NewCreatePostgresBackupConfigDefault(code int) *CreatePostgresBackupConfigDefault {
	return &CreatePostgresBackupConfigDefault{
		_statusCode: code,
	}
}

/*
CreatePostgresBackupConfigDefault describes a response with status code -1, with default header values.

Error
*/
type CreatePostgresBackupConfigDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this create postgres backup config default response has a 2xx status code
func (o *CreatePostgresBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create postgres backup config default response has a 3xx status code
func (o *CreatePostgresBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create postgres backup config default response has a 4xx status code
func (o *CreatePostgresBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create postgres backup config default response has a 5xx status code
func (o *CreatePostgresBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create postgres backup config default response a status code equal to that given
func (o *CreatePostgresBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create postgres backup config default response
func (o *CreatePostgresBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreatePostgresBackupConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/database/postgres/backup-config][%d] createPostgresBackupConfig default %s", o._statusCode, payload)
}

func (o *CreatePostgresBackupConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/database/postgres/backup-config][%d] createPostgresBackupConfig default %s", o._statusCode, payload)
}

func (o *CreatePostgresBackupConfigDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CreatePostgresBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
