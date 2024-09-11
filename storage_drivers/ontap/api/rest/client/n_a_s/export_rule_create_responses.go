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

// ExportRuleCreateReader is a Reader for the ExportRuleCreate structure.
type ExportRuleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExportRuleCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleCreateCreated creates a ExportRuleCreateCreated with default headers values
func NewExportRuleCreateCreated() *ExportRuleCreateCreated {
	return &ExportRuleCreateCreated{}
}

/*
ExportRuleCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ExportRuleCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.ExportRuleResponse
}

// IsSuccess returns true when this export rule create created response has a 2xx status code
func (o *ExportRuleCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export rule create created response has a 3xx status code
func (o *ExportRuleCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export rule create created response has a 4xx status code
func (o *ExportRuleCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this export rule create created response has a 5xx status code
func (o *ExportRuleCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this export rule create created response a status code equal to that given
func (o *ExportRuleCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the export rule create created response
func (o *ExportRuleCreateCreated) Code() int {
	return 201
}

func (o *ExportRuleCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules][%d] exportRuleCreateCreated %s", 201, payload)
}

func (o *ExportRuleCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules][%d] exportRuleCreateCreated %s", 201, payload)
}

func (o *ExportRuleCreateCreated) GetPayload() *models.ExportRuleResponse {
	return o.Payload
}

func (o *ExportRuleCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.ExportRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleCreateDefault creates a ExportRuleCreateDefault with default headers values
func NewExportRuleCreateDefault(code int) *ExportRuleCreateDefault {
	return &ExportRuleCreateDefault{
		_statusCode: code,
	}
}

/*
	ExportRuleCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
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
| 1704050    | Invalid clientmatch: clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704054    | Invalid clientmatch: invalid characters in netgroup name. Valid characters for a netgroup name are 0-9, A-Z, a-z, ".", "_" and "-" |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |
| 1704070    | Export policy rule already exists. The export policy rule was not created because an export policy rule with the same values already exists. |
| 3277000    | Upgrade all nodes to ONTAP 9.0.0 or above to use krb5p as a security flavor in export-policy rules |
| 3277083    | User ID is not valid. Enter a value for User ID from 0 to 4294967295 |
| 3277162    | The specified \"index\", 0, is invalid. Valid values are values from 1 to 4294967295 |
| 3277163    | The system cannot automatically specify an "index" for this rule, because a rule with "index" 4294967295 exists. Either specify an unused "index", or update the existing rules so that "index" 4294967295 is not used. |
| 3277149    | The "Anon" field cannot be an empty string |
*/
type ExportRuleCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export rule create default response has a 2xx status code
func (o *ExportRuleCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export rule create default response has a 3xx status code
func (o *ExportRuleCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export rule create default response has a 4xx status code
func (o *ExportRuleCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export rule create default response has a 5xx status code
func (o *ExportRuleCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export rule create default response a status code equal to that given
func (o *ExportRuleCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export rule create default response
func (o *ExportRuleCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules][%d] export_rule_create default %s", o._statusCode, payload)
}

func (o *ExportRuleCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules][%d] export_rule_create default %s", o._statusCode, payload)
}

func (o *ExportRuleCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
