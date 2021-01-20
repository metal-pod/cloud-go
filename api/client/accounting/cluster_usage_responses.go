// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

// ClusterUsageReader is a Reader for the ClusterUsage structure.
type ClusterUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterUsageOK creates a ClusterUsageOK with default headers values
func NewClusterUsageOK() *ClusterUsageOK {
	return &ClusterUsageOK{}
}

/* ClusterUsageOK describes a response with status code 200, with default header values.

OK
*/
type ClusterUsageOK struct {
	Payload *models.V1ClusterUsageResponse
}

func (o *ClusterUsageOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/cluster-usage][%d] clusterUsageOK  %+v", 200, o.Payload)
}
func (o *ClusterUsageOK) GetPayload() *models.V1ClusterUsageResponse {
	return o.Payload
}

func (o *ClusterUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterUsageDefault creates a ClusterUsageDefault with default headers values
func NewClusterUsageDefault(code int) *ClusterUsageDefault {
	return &ClusterUsageDefault{
		_statusCode: code,
	}
}

/* ClusterUsageDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterUsageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the cluster usage default response
func (o *ClusterUsageDefault) Code() int {
	return o._statusCode
}

func (o *ClusterUsageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/cluster-usage][%d] clusterUsage default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterUsageDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ClusterUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
