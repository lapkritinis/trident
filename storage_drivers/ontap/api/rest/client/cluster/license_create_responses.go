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

// LicenseCreateReader is a Reader for the LicenseCreate structure.
type LicenseCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LicenseCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLicenseCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLicenseCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLicenseCreateCreated creates a LicenseCreateCreated with default headers values
func NewLicenseCreateCreated() *LicenseCreateCreated {
	return &LicenseCreateCreated{}
}

/*
LicenseCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LicenseCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.LicensePackageResponse
}

// IsSuccess returns true when this license create created response has a 2xx status code
func (o *LicenseCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this license create created response has a 3xx status code
func (o *LicenseCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this license create created response has a 4xx status code
func (o *LicenseCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this license create created response has a 5xx status code
func (o *LicenseCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this license create created response a status code equal to that given
func (o *LicenseCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the license create created response
func (o *LicenseCreateCreated) Code() int {
	return 201
}

func (o *LicenseCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/licensing/licenses][%d] licenseCreateCreated %s", 201, payload)
}

func (o *LicenseCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/licensing/licenses][%d] licenseCreateCreated %s", 201, payload)
}

func (o *LicenseCreateCreated) GetPayload() *models.LicensePackageResponse {
	return o.Payload
}

func (o *LicenseCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.LicensePackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLicenseCreateDefault creates a LicenseCreateDefault with default headers values
func NewLicenseCreateDefault(code int) *LicenseCreateDefault {
	return &LicenseCreateDefault{
		_statusCode: code,
	}
}

/*
	LicenseCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115117 | Generic licensing error |
| 1115122 | No cluster serial number found |
| 1115124 | No node serial number found |
| 1115130 | No license code was provided |
| 1115131 | Installation of the license failed |
| 1115132 | License already exists on system |
| 1115134 | Serial number does not belong to node |
| 1115141 | License data is invalid |
| 1115142 | License signature is invalid |
| 1115143 | Internal error applying the requested license |
| 1115152 | License does not apply to the platform |
| 1115154 | Unable to retrieve cluster ID |
| 1115155 | Invalid cluster ID found |
| 1115159 | License is not in an acceptable format |
| 1115160 | License has already expired |
| 1115164 | Minimum ONTAP version requirements not met |
| 1115165 | Minimum ONTAP version requirements are not met for license type enabled |
| 1115166 | Minimum ONTAP version requirements are not met for license protocol SEC-COMP-BNDL-ENBLD |
| 1115179 | FlexCache is not supported on this system |
| 1115180 | FlexCache is not supported on cloud systems |
| 1115407 | Capacity pool licenses cannot be installed directly |
| 1115427 | License is incompatible with capacity pools licensing mode |
| 1115562 | One or more errors occurred when installing a NLFv2 license |
| 1115563 | Package details and serial number of license contained within the NLFv2 failure |
| 1115564 | Package cannot be deleted individually as it is part of a bundle |
| 1115565 | NLFv2 install failed as the license serial number is already in use |
| 1115616 | Package details and serial number of license included in the install conflict |
| 1115617 | NLFv2 license install failed with summary of conflicting licenses |
| 1115618 | NLFv2 license install failed as a license with newer timestamp already exists |
| 5375355 | The cluster has more nodes than are supported by All SAN Array. |
| 5375366 | The cluster has one or more nodes that do not support All SAN Array. |
| 66846818 | Failed to interpret FlexCache license information |
| 66846821 | FlexCache is not supported on cloud systems |
| 66846822 | Invalid FlexCache capacity information provided |
| 655294464 | Failed to extract license contents |
| 655294465 | License key is invalid |
| 655294466 | Serial number is invalid |
| 655294467 | Version number is invalid |
| 655294468 | Expired license |
| 655294469 | License does not apply to the platform |
| 655294470 | License does not apply to the product |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LicenseCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponses
}

// IsSuccess returns true when this license create default response has a 2xx status code
func (o *LicenseCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this license create default response has a 3xx status code
func (o *LicenseCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this license create default response has a 4xx status code
func (o *LicenseCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this license create default response has a 5xx status code
func (o *LicenseCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this license create default response a status code equal to that given
func (o *LicenseCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the license create default response
func (o *LicenseCreateDefault) Code() int {
	return o._statusCode
}

func (o *LicenseCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/licensing/licenses][%d] license_create default %s", o._statusCode, payload)
}

func (o *LicenseCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/licensing/licenses][%d] license_create default %s", o._statusCode, payload)
}

func (o *LicenseCreateDefault) GetPayload() *models.ErrorResponses {
	return o.Payload
}

func (o *LicenseCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
