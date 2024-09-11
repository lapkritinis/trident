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

// GroupPolicyObjectCentralAccessRuleGetReader is a Reader for the GroupPolicyObjectCentralAccessRuleGet structure.
type GroupPolicyObjectCentralAccessRuleGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPolicyObjectCentralAccessRuleGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPolicyObjectCentralAccessRuleGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPolicyObjectCentralAccessRuleGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPolicyObjectCentralAccessRuleGetOK creates a GroupPolicyObjectCentralAccessRuleGetOK with default headers values
func NewGroupPolicyObjectCentralAccessRuleGetOK() *GroupPolicyObjectCentralAccessRuleGetOK {
	return &GroupPolicyObjectCentralAccessRuleGetOK{}
}

/*
GroupPolicyObjectCentralAccessRuleGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupPolicyObjectCentralAccessRuleGetOK struct {
	Payload *models.GroupPolicyObjectCentralAccessRule
}

// IsSuccess returns true when this group policy object central access rule get o k response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group policy object central access rule get o k response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group policy object central access rule get o k response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group policy object central access rule get o k response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group policy object central access rule get o k response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessRuleGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group policy object central access rule get o k response
func (o *GroupPolicyObjectCentralAccessRuleGetOK) Code() int {
	return 200
}

func (o *GroupPolicyObjectCentralAccessRuleGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules/{name}][%d] groupPolicyObjectCentralAccessRuleGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules/{name}][%d] groupPolicyObjectCentralAccessRuleGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleGetOK) GetPayload() *models.GroupPolicyObjectCentralAccessRule {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessRuleGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPolicyObjectCentralAccessRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPolicyObjectCentralAccessRuleGetDefault creates a GroupPolicyObjectCentralAccessRuleGetDefault with default headers values
func NewGroupPolicyObjectCentralAccessRuleGetDefault(code int) *GroupPolicyObjectCentralAccessRuleGetDefault {
	return &GroupPolicyObjectCentralAccessRuleGetDefault{
		_statusCode: code,
	}
}

/*
GroupPolicyObjectCentralAccessRuleGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupPolicyObjectCentralAccessRuleGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group policy object central access rule get default response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group policy object central access rule get default response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group policy object central access rule get default response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group policy object central access rule get default response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessRuleGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group policy object central access rule get default response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessRuleGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group policy object central access rule get default response
func (o *GroupPolicyObjectCentralAccessRuleGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupPolicyObjectCentralAccessRuleGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules/{name}][%d] group_policy_object_central_access_rule_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules/{name}][%d] group_policy_object_central_access_rule_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessRuleGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
