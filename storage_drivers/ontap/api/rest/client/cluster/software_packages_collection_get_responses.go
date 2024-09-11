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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SoftwarePackagesCollectionGetReader is a Reader for the SoftwarePackagesCollectionGet structure.
type SoftwarePackagesCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwarePackagesCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwarePackagesCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwarePackagesCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwarePackagesCollectionGetOK creates a SoftwarePackagesCollectionGetOK with default headers values
func NewSoftwarePackagesCollectionGetOK() *SoftwarePackagesCollectionGetOK {
	return &SoftwarePackagesCollectionGetOK{}
}

/*
SoftwarePackagesCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SoftwarePackagesCollectionGetOK struct {
	Payload *models.SoftwarePackageResponse
}

// IsSuccess returns true when this software packages collection get o k response has a 2xx status code
func (o *SoftwarePackagesCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this software packages collection get o k response has a 3xx status code
func (o *SoftwarePackagesCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this software packages collection get o k response has a 4xx status code
func (o *SoftwarePackagesCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this software packages collection get o k response has a 5xx status code
func (o *SoftwarePackagesCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this software packages collection get o k response a status code equal to that given
func (o *SoftwarePackagesCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the software packages collection get o k response
func (o *SoftwarePackagesCollectionGetOK) Code() int {
	return 200
}

func (o *SoftwarePackagesCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/packages][%d] softwarePackagesCollectionGetOK %s", 200, payload)
}

func (o *SoftwarePackagesCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/packages][%d] softwarePackagesCollectionGetOK %s", 200, payload)
}

func (o *SoftwarePackagesCollectionGetOK) GetPayload() *models.SoftwarePackageResponse {
	return o.Payload
}

func (o *SoftwarePackagesCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwarePackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwarePackagesCollectionGetDefault creates a SoftwarePackagesCollectionGetDefault with default headers values
func NewSoftwarePackagesCollectionGetDefault(code int) *SoftwarePackagesCollectionGetDefault {
	return &SoftwarePackagesCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SoftwarePackagesCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SoftwarePackagesCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this software packages collection get default response has a 2xx status code
func (o *SoftwarePackagesCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this software packages collection get default response has a 3xx status code
func (o *SoftwarePackagesCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this software packages collection get default response has a 4xx status code
func (o *SoftwarePackagesCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this software packages collection get default response has a 5xx status code
func (o *SoftwarePackagesCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this software packages collection get default response a status code equal to that given
func (o *SoftwarePackagesCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the software packages collection get default response
func (o *SoftwarePackagesCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SoftwarePackagesCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/packages][%d] software_packages_collection_get default %s", o._statusCode, payload)
}

func (o *SoftwarePackagesCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/packages][%d] software_packages_collection_get default %s", o._statusCode, payload)
}

func (o *SoftwarePackagesCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SoftwarePackagesCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
