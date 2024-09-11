// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// ExportPolicyCreateReader is a Reader for the ExportPolicyCreate structure.
type ExportPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExportPolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportPolicyCreateCreated creates a ExportPolicyCreateCreated with default headers values
func NewExportPolicyCreateCreated() *ExportPolicyCreateCreated {
	return &ExportPolicyCreateCreated{}
}

/*
ExportPolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ExportPolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.ExportPolicyResponse
}

// IsSuccess returns true when this export policy create created response has a 2xx status code
func (o *ExportPolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export policy create created response has a 3xx status code
func (o *ExportPolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export policy create created response has a 4xx status code
func (o *ExportPolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this export policy create created response has a 5xx status code
func (o *ExportPolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this export policy create created response a status code equal to that given
func (o *ExportPolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the export policy create created response
func (o *ExportPolicyCreateCreated) Code() int {
	return 201
}

func (o *ExportPolicyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies][%d] exportPolicyCreateCreated %s", 201, payload)
}

func (o *ExportPolicyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies][%d] exportPolicyCreateCreated %s", 201, payload)
}

func (o *ExportPolicyCreateCreated) GetPayload() *models.ExportPolicyResponse {
	return o.Payload
}

func (o *ExportPolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.ExportPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportPolicyCreateDefault creates a ExportPolicyCreateDefault with default headers values
func NewExportPolicyCreateDefault(code int) *ExportPolicyCreateDefault {
	return &ExportPolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	ExportPolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1703952    | Invalid ruleset name provided. No spaces allowed in a ruleset name|
| 1703954    | Export policy does not exist |
| 1704036    | Invalid clientmatch:  missing domain name |
| 1704037    | Invalid clientmatch:  missing network name |
| 1704038    | Invalid clientmatch:  missing netgroup name |
| 1704039    | Invalid clientmatch |
| 1704040    | Invalid clientmatch: address bytes masked out by netmask are non-zero |
| 1704041    | Invalid clientmatch: address bytes masked to zero by netmask |
| 1704042    | Invalid clientmatch: too many bits in netmask |
| 1704043    | Invalid clientmatch: invalid netmask |
| 1704044    | Invalid clientmatch: invalid characters in host name |
| 1704045    | Invalid clientmatch: invalid characters in domain name |
| 1704047    | The export policy name cannot be longer than 256 characters |
| 1704049    | Invalid clientmatch: clientmatch lists require an effective cluster version of Data ONTAP 9.0 or later. Upgrade all nodes to Data ONTAP 9.0 or above to use features that operate on lists of clientmatch strings in export-policy rules |
| 1704050    | Invalid clientmatch: clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704054    | Invalid clientmatch: invalid characters in netgroup name. Valid characters for a netgroup name are 0-9, A-Z, a-z, ".", "_" and "-" |
| 1704055    | Export policies are only supported for data Vservers |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |
| 1704071    | The export policy was not created because of duplicate export policy rules present in the request |
| 2621462    | The specified SVM name does not exist |
| 2621519    | SVM name is invalid. The SVM name must begin with a letter or an underscore. If the SVM is of type \"sync-source\", the maximum supported length is 41. Otherwise, the maximum supported length is 47 |
| 2621643    | The specified SVM name is too long |
| 2621685    | SVM name length cannot be zero |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 3277000    | Upgrade all nodes to ONTAP 9.0.0 or above to use krb5p as a security flavor in export-policy rules |
| 3277083    | User ID is not valid. Enter a value for User ID from 0 to 4294967295 |
| 3277149    | The "Anon" field cannot be an empty string |
| 6691623    | User is not authorized |
*/
type ExportPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export policy create default response has a 2xx status code
func (o *ExportPolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export policy create default response has a 3xx status code
func (o *ExportPolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export policy create default response has a 4xx status code
func (o *ExportPolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export policy create default response has a 5xx status code
func (o *ExportPolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export policy create default response a status code equal to that given
func (o *ExportPolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export policy create default response
func (o *ExportPolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExportPolicyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies][%d] export_policy_create default %s", o._statusCode, payload)
}

func (o *ExportPolicyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies][%d] export_policy_create default %s", o._statusCode, payload)
}

func (o *ExportPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
