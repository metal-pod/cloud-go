// Code generated by go-swagger; DO NOT EDIT.

package volume

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

// ClusterInfoReader is a Reader for the ClusterInfo structure.
type ClusterInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterInfoOK creates a ClusterInfoOK with default headers values
func NewClusterInfoOK() *ClusterInfoOK {
	return &ClusterInfoOK{}
}

/* ClusterInfoOK describes a response with status code 200, with default header values.

OK
*/
type ClusterInfoOK struct {
	Payload []models.V2ClusterInfo
}

func (o *ClusterInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/volume/clusterinfo][%d] clusterInfoOK  %+v", 200, o.Payload)
}
func (o *ClusterInfoOK) GetPayload() []models.V2ClusterInfo {
	return o.Payload
}

func (o *ClusterInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterInfoDefault creates a ClusterInfoDefault with default headers values
func NewClusterInfoDefault(code int) *ClusterInfoDefault {
	return &ClusterInfoDefault{
		_statusCode: code,
	}
}

/* ClusterInfoDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterInfoDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the cluster info default response
func (o *ClusterInfoDefault) Code() int {
	return o._statusCode
}

func (o *ClusterInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volume/clusterinfo][%d] clusterInfo default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterInfoDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ClusterInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
