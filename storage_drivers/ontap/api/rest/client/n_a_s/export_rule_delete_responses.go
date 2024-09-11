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

// ExportRuleDeleteReader is a Reader for the ExportRuleDelete structure.
type ExportRuleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleDeleteOK creates a ExportRuleDeleteOK with default headers values
func NewExportRuleDeleteOK() *ExportRuleDeleteOK {
	return &ExportRuleDeleteOK{}
}

/*
ExportRuleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleDeleteOK struct {
}

// IsSuccess returns true when this export rule delete o k response has a 2xx status code
func (o *ExportRuleDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export rule delete o k response has a 3xx status code
func (o *ExportRuleDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export rule delete o k response has a 4xx status code
func (o *ExportRuleDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export rule delete o k response has a 5xx status code
func (o *ExportRuleDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export rule delete o k response a status code equal to that given
func (o *ExportRuleDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export rule delete o k response
func (o *ExportRuleDeleteOK) Code() int {
	return 200
}

func (o *ExportRuleDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] exportRuleDeleteOK", 200)
}

func (o *ExportRuleDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] exportRuleDeleteOK", 200)
}

func (o *ExportRuleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportRuleDeleteDefault creates a ExportRuleDeleteDefault with default headers values
func NewExportRuleDeleteDefault(code int) *ExportRuleDeleteDefault {
	return &ExportRuleDeleteDefault{
		_statusCode: code,
	}
}

/*
	ExportRuleDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1703945    | Ruleset is in use by a volume.  It cannot be deleted until all volumes that refer to it are first deleted|
| 1703946    | Cannot determine if the ruleset is in use by a volume.  It cannot be deleted until all volumes that refer to it are first deleted|
| 1703954    | Export policy does not exist |
*/
type ExportRuleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export rule delete default response has a 2xx status code
func (o *ExportRuleDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export rule delete default response has a 3xx status code
func (o *ExportRuleDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export rule delete default response has a 4xx status code
func (o *ExportRuleDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export rule delete default response has a 5xx status code
func (o *ExportRuleDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export rule delete default response a status code equal to that given
func (o *ExportRuleDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export rule delete default response
func (o *ExportRuleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] export_rule_delete default %s", o._statusCode, payload)
}

func (o *ExportRuleDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] export_rule_delete default %s", o._statusCode, payload)
}

func (o *ExportRuleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
