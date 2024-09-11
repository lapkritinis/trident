// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsDestinationCreateReader is a Reader for the EmsDestinationCreate structure.
type EmsDestinationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsDestinationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmsDestinationCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsDestinationCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsDestinationCreateCreated creates a EmsDestinationCreateCreated with default headers values
func NewEmsDestinationCreateCreated() *EmsDestinationCreateCreated {
	return &EmsDestinationCreateCreated{}
}

/*
EmsDestinationCreateCreated describes a response with status code 201, with default header values.

Created
*/
type EmsDestinationCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.EmsDestinationResponse
}

// IsSuccess returns true when this ems destination create created response has a 2xx status code
func (o *EmsDestinationCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems destination create created response has a 3xx status code
func (o *EmsDestinationCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems destination create created response has a 4xx status code
func (o *EmsDestinationCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems destination create created response has a 5xx status code
func (o *EmsDestinationCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ems destination create created response a status code equal to that given
func (o *EmsDestinationCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ems destination create created response
func (o *EmsDestinationCreateCreated) Code() int {
	return 201
}

func (o *EmsDestinationCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/destinations][%d] emsDestinationCreateCreated %s", 201, payload)
}

func (o *EmsDestinationCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/destinations][%d] emsDestinationCreateCreated %s", 201, payload)
}

func (o *EmsDestinationCreateCreated) GetPayload() *models.EmsDestinationResponse {
	return o.Payload
}

func (o *EmsDestinationCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.EmsDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsDestinationCreateDefault creates a EmsDestinationCreateDefault with default headers values
func NewEmsDestinationCreateDefault(code int) *EmsDestinationCreateDefault {
	return &EmsDestinationCreateDefault{
		_statusCode: code,
	}
}

/*
	EmsDestinationCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983104 | The syslog destination provided is invalid |
| 983107 | A provided filter does not exist |
| 983116 | The number of notifications has reached maximum capacity |
| 983117 | The number of destinations has reached maximum capacity |
| 983129 | The rest-api destination provided must contain a valid scheme, such as http// or https// |
| 983130 | The rest-api destination provided contains an invalid URL |
| 983131 | The rest-api destination provided contains an invalid IPv6 URL |
| 983144 | The security certificate information provided is incomplete. Provide the certificate and serial number |
| 983145 | The rest-api destination provided has an 'http://' scheme. It is invalid to provide certificate information |
| 983149 | New SNMP destinations cannot be created |
| 983152 | Default destinations cannot be modified or removed |
| 983153 | The security certificate provided does not exist |
| 983154 | The necessary private key is not installed on the system |
| 983170 | Connectivity check is not supported for the specified destination type on the node |
| 983171 | Connectivity check failed for the specified destination on the node |
| 983176 | Connectivity check is only supported for TCP-based syslog destinations |
| 983179 | The value for the destination field cannot be empty |
| 983180 | The destination name provided cannot be empty |
| 983181 | The destination name provided cannot contain spaces |
| 983182 | The destination name provided is invalid. The destination name must contain between 2 and 64 characters and start and end with an alphanumeric symbol or _(underscore). The allowed special characters are _(underscore) and -(hyphen) |
| 983184 | A provided property cannot be configured on the requested destination type |
| 983200 | Access control role compatibility issue with provided filters |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsDestinationCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems destination create default response has a 2xx status code
func (o *EmsDestinationCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems destination create default response has a 3xx status code
func (o *EmsDestinationCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems destination create default response has a 4xx status code
func (o *EmsDestinationCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems destination create default response has a 5xx status code
func (o *EmsDestinationCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems destination create default response a status code equal to that given
func (o *EmsDestinationCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems destination create default response
func (o *EmsDestinationCreateDefault) Code() int {
	return o._statusCode
}

func (o *EmsDestinationCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/destinations][%d] ems_destination_create default %s", o._statusCode, payload)
}

func (o *EmsDestinationCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/destinations][%d] ems_destination_create default %s", o._statusCode, payload)
}

func (o *EmsDestinationCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsDestinationCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
