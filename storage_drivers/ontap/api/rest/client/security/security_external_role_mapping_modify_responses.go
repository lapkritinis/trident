// Code generated by go-swagger; DO NOT EDIT.

package security

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

// SecurityExternalRoleMappingModifyReader is a Reader for the SecurityExternalRoleMappingModify structure.
type SecurityExternalRoleMappingModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityExternalRoleMappingModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityExternalRoleMappingModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityExternalRoleMappingModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityExternalRoleMappingModifyOK creates a SecurityExternalRoleMappingModifyOK with default headers values
func NewSecurityExternalRoleMappingModifyOK() *SecurityExternalRoleMappingModifyOK {
	return &SecurityExternalRoleMappingModifyOK{}
}

/*
SecurityExternalRoleMappingModifyOK describes a response with status code 200, with default header values.

OK
*/
type SecurityExternalRoleMappingModifyOK struct {
	Payload *models.SecurityExternalRoleMapping
}

// IsSuccess returns true when this security external role mapping modify o k response has a 2xx status code
func (o *SecurityExternalRoleMappingModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security external role mapping modify o k response has a 3xx status code
func (o *SecurityExternalRoleMappingModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security external role mapping modify o k response has a 4xx status code
func (o *SecurityExternalRoleMappingModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security external role mapping modify o k response has a 5xx status code
func (o *SecurityExternalRoleMappingModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security external role mapping modify o k response a status code equal to that given
func (o *SecurityExternalRoleMappingModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security external role mapping modify o k response
func (o *SecurityExternalRoleMappingModifyOK) Code() int {
	return 200
}

func (o *SecurityExternalRoleMappingModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/external-role-mappings/{external_role}/{provider}][%d] securityExternalRoleMappingModifyOK %s", 200, payload)
}

func (o *SecurityExternalRoleMappingModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/external-role-mappings/{external_role}/{provider}][%d] securityExternalRoleMappingModifyOK %s", 200, payload)
}

func (o *SecurityExternalRoleMappingModifyOK) GetPayload() *models.SecurityExternalRoleMapping {
	return o.Payload
}

func (o *SecurityExternalRoleMappingModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityExternalRoleMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityExternalRoleMappingModifyDefault creates a SecurityExternalRoleMappingModifyDefault with default headers values
func NewSecurityExternalRoleMappingModifyDefault(code int) *SecurityExternalRoleMappingModifyDefault {
	return &SecurityExternalRoleMappingModifyDefault{
		_statusCode: code,
	}
}

/*
	SecurityExternalRoleMappingModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5636243 | Provided ONTAP role is not configured. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityExternalRoleMappingModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security external role mapping modify default response has a 2xx status code
func (o *SecurityExternalRoleMappingModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security external role mapping modify default response has a 3xx status code
func (o *SecurityExternalRoleMappingModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security external role mapping modify default response has a 4xx status code
func (o *SecurityExternalRoleMappingModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security external role mapping modify default response has a 5xx status code
func (o *SecurityExternalRoleMappingModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security external role mapping modify default response a status code equal to that given
func (o *SecurityExternalRoleMappingModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security external role mapping modify default response
func (o *SecurityExternalRoleMappingModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityExternalRoleMappingModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/external-role-mappings/{external_role}/{provider}][%d] security_external_role_mapping_modify default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/external-role-mappings/{external_role}/{provider}][%d] security_external_role_mapping_modify default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityExternalRoleMappingModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
