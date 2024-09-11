// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NvmeServiceCollectionGetReader is a Reader for the NvmeServiceCollectionGet structure.
type NvmeServiceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeServiceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeServiceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeServiceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeServiceCollectionGetOK creates a NvmeServiceCollectionGetOK with default headers values
func NewNvmeServiceCollectionGetOK() *NvmeServiceCollectionGetOK {
	return &NvmeServiceCollectionGetOK{}
}

/*
NvmeServiceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeServiceCollectionGetOK struct {
	Payload *models.NvmeServiceResponse
}

// IsSuccess returns true when this nvme service collection get o k response has a 2xx status code
func (o *NvmeServiceCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme service collection get o k response has a 3xx status code
func (o *NvmeServiceCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme service collection get o k response has a 4xx status code
func (o *NvmeServiceCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme service collection get o k response has a 5xx status code
func (o *NvmeServiceCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme service collection get o k response a status code equal to that given
func (o *NvmeServiceCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme service collection get o k response
func (o *NvmeServiceCollectionGetOK) Code() int {
	return 200
}

func (o *NvmeServiceCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services][%d] nvmeServiceCollectionGetOK %s", 200, payload)
}

func (o *NvmeServiceCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services][%d] nvmeServiceCollectionGetOK %s", 200, payload)
}

func (o *NvmeServiceCollectionGetOK) GetPayload() *models.NvmeServiceResponse {
	return o.Payload
}

func (o *NvmeServiceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeServiceCollectionGetDefault creates a NvmeServiceCollectionGetDefault with default headers values
func NewNvmeServiceCollectionGetDefault(code int) *NvmeServiceCollectionGetDefault {
	return &NvmeServiceCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NvmeServiceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NvmeServiceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme service collection get default response has a 2xx status code
func (o *NvmeServiceCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme service collection get default response has a 3xx status code
func (o *NvmeServiceCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme service collection get default response has a 4xx status code
func (o *NvmeServiceCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme service collection get default response has a 5xx status code
func (o *NvmeServiceCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme service collection get default response a status code equal to that given
func (o *NvmeServiceCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme service collection get default response
func (o *NvmeServiceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeServiceCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services][%d] nvme_service_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeServiceCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services][%d] nvme_service_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeServiceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeServiceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
