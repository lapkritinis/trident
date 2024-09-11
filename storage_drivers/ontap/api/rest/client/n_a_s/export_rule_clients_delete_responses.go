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

// ExportRuleClientsDeleteReader is a Reader for the ExportRuleClientsDelete structure.
type ExportRuleClientsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleClientsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleClientsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleClientsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleClientsDeleteOK creates a ExportRuleClientsDeleteOK with default headers values
func NewExportRuleClientsDeleteOK() *ExportRuleClientsDeleteOK {
	return &ExportRuleClientsDeleteOK{}
}

/*
ExportRuleClientsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleClientsDeleteOK struct {
}

// IsSuccess returns true when this export rule clients delete o k response has a 2xx status code
func (o *ExportRuleClientsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export rule clients delete o k response has a 3xx status code
func (o *ExportRuleClientsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export rule clients delete o k response has a 4xx status code
func (o *ExportRuleClientsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export rule clients delete o k response has a 5xx status code
func (o *ExportRuleClientsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export rule clients delete o k response a status code equal to that given
func (o *ExportRuleClientsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export rule clients delete o k response
func (o *ExportRuleClientsDeleteOK) Code() int {
	return 200
}

func (o *ExportRuleClientsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients/{match}][%d] exportRuleClientsDeleteOK", 200)
}

func (o *ExportRuleClientsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients/{match}][%d] exportRuleClientsDeleteOK", 200)
}

func (o *ExportRuleClientsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportRuleClientsDeleteDefault creates a ExportRuleClientsDeleteDefault with default headers values
func NewExportRuleClientsDeleteDefault(code int) *ExportRuleClientsDeleteDefault {
	return &ExportRuleClientsDeleteDefault{
		_statusCode: code,
	}
}

/*
	ExportRuleClientsDeleteDefault describes a response with status code -1, with default header values.

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
| 1704050    | Invalid clientmatch: the clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704054    | Invalid clientmatch: invalid characters in netgroup name. Valid characters for a netgroup name are 0-9, A-Z, a-z, ".", "_" and "-" |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |
| 6691623    | User is not authorized |
*/
type ExportRuleClientsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export rule clients delete default response has a 2xx status code
func (o *ExportRuleClientsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export rule clients delete default response has a 3xx status code
func (o *ExportRuleClientsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export rule clients delete default response has a 4xx status code
func (o *ExportRuleClientsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export rule clients delete default response has a 5xx status code
func (o *ExportRuleClientsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export rule clients delete default response a status code equal to that given
func (o *ExportRuleClientsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export rule clients delete default response
func (o *ExportRuleClientsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleClientsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients/{match}][%d] export_rule_clients_delete default %s", o._statusCode, payload)
}

func (o *ExportRuleClientsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients/{match}][%d] export_rule_clients_delete default %s", o._statusCode, payload)
}

func (o *ExportRuleClientsDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleClientsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
