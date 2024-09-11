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

// GroupPolicyObjectCentralAccessRuleCollectionGetReader is a Reader for the GroupPolicyObjectCentralAccessRuleCollectionGet structure.
type GroupPolicyObjectCentralAccessRuleCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPolicyObjectCentralAccessRuleCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPolicyObjectCentralAccessRuleCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPolicyObjectCentralAccessRuleCollectionGetOK creates a GroupPolicyObjectCentralAccessRuleCollectionGetOK with default headers values
func NewGroupPolicyObjectCentralAccessRuleCollectionGetOK() *GroupPolicyObjectCentralAccessRuleCollectionGetOK {
	return &GroupPolicyObjectCentralAccessRuleCollectionGetOK{}
}

/*
GroupPolicyObjectCentralAccessRuleCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupPolicyObjectCentralAccessRuleCollectionGetOK struct {
	Payload *models.GroupPolicyObjectCentralAccessRuleResponse
}

// IsSuccess returns true when this group policy object central access rule collection get o k response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group policy object central access rule collection get o k response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group policy object central access rule collection get o k response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group policy object central access rule collection get o k response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group policy object central access rule collection get o k response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group policy object central access rule collection get o k response
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) Code() int {
	return 200
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules][%d] groupPolicyObjectCentralAccessRuleCollectionGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules][%d] groupPolicyObjectCentralAccessRuleCollectionGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) GetPayload() *models.GroupPolicyObjectCentralAccessRuleResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPolicyObjectCentralAccessRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPolicyObjectCentralAccessRuleCollectionGetDefault creates a GroupPolicyObjectCentralAccessRuleCollectionGetDefault with default headers values
func NewGroupPolicyObjectCentralAccessRuleCollectionGetDefault(code int) *GroupPolicyObjectCentralAccessRuleCollectionGetDefault {
	return &GroupPolicyObjectCentralAccessRuleCollectionGetDefault{
		_statusCode: code,
	}
}

/*
GroupPolicyObjectCentralAccessRuleCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupPolicyObjectCentralAccessRuleCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group policy object central access rule collection get default response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group policy object central access rule collection get default response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group policy object central access rule collection get default response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group policy object central access rule collection get default response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group policy object central access rule collection get default response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group policy object central access rule collection get default response
func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules][%d] group_policy_object_central_access_rule_collection_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-rules][%d] group_policy_object_central_access_rule_collection_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessRuleCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
