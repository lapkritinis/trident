// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

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

// SnapmirrorPolicyModifyReader is a Reader for the SnapmirrorPolicyModify structure.
type SnapmirrorPolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorPolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapmirrorPolicyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapmirrorPolicyModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorPolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorPolicyModifyOK creates a SnapmirrorPolicyModifyOK with default headers values
func NewSnapmirrorPolicyModifyOK() *SnapmirrorPolicyModifyOK {
	return &SnapmirrorPolicyModifyOK{}
}

/*
SnapmirrorPolicyModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnapmirrorPolicyModifyOK struct {
	Payload *models.SnapmirrorPolicyJobLinkResponse
}

// IsSuccess returns true when this snapmirror policy modify o k response has a 2xx status code
func (o *SnapmirrorPolicyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror policy modify o k response has a 3xx status code
func (o *SnapmirrorPolicyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror policy modify o k response has a 4xx status code
func (o *SnapmirrorPolicyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror policy modify o k response has a 5xx status code
func (o *SnapmirrorPolicyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror policy modify o k response a status code equal to that given
func (o *SnapmirrorPolicyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapmirror policy modify o k response
func (o *SnapmirrorPolicyModifyOK) Code() int {
	return 200
}

func (o *SnapmirrorPolicyModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/policies/{uuid}][%d] snapmirrorPolicyModifyOK %s", 200, payload)
}

func (o *SnapmirrorPolicyModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/policies/{uuid}][%d] snapmirrorPolicyModifyOK %s", 200, payload)
}

func (o *SnapmirrorPolicyModifyOK) GetPayload() *models.SnapmirrorPolicyJobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapmirrorPolicyJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorPolicyModifyAccepted creates a SnapmirrorPolicyModifyAccepted with default headers values
func NewSnapmirrorPolicyModifyAccepted() *SnapmirrorPolicyModifyAccepted {
	return &SnapmirrorPolicyModifyAccepted{}
}

/*
SnapmirrorPolicyModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorPolicyModifyAccepted struct {
	Payload *models.SnapmirrorPolicyJobLinkResponse
}

// IsSuccess returns true when this snapmirror policy modify accepted response has a 2xx status code
func (o *SnapmirrorPolicyModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror policy modify accepted response has a 3xx status code
func (o *SnapmirrorPolicyModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror policy modify accepted response has a 4xx status code
func (o *SnapmirrorPolicyModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror policy modify accepted response has a 5xx status code
func (o *SnapmirrorPolicyModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror policy modify accepted response a status code equal to that given
func (o *SnapmirrorPolicyModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapmirror policy modify accepted response
func (o *SnapmirrorPolicyModifyAccepted) Code() int {
	return 202
}

func (o *SnapmirrorPolicyModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/policies/{uuid}][%d] snapmirrorPolicyModifyAccepted %s", 202, payload)
}

func (o *SnapmirrorPolicyModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/policies/{uuid}][%d] snapmirrorPolicyModifyAccepted %s", 202, payload)
}

func (o *SnapmirrorPolicyModifyAccepted) GetPayload() *models.SnapmirrorPolicyJobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapmirrorPolicyJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorPolicyModifyDefault creates a SnapmirrorPolicyModifyDefault with default headers values
func NewSnapmirrorPolicyModifyDefault(code int) *SnapmirrorPolicyModifyDefault {
	return &SnapmirrorPolicyModifyDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorPolicyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 6619714     | Schedule specified is an interval schedule. SnapMirror does not support interval schedules. |
| 13303842    | SnapMirror policy is not supported. |
| 13303843    | Conflicting values between SnapMirror policy and SnapMirror relationships for either 'transfer_schedule, throttle or identity_preservation' properties |
| 13303850    | Invalid input parameter |
| 13303887    | Failed to create SnapMirror policy. Reason: Maximum number of allowed retention rules reached |
| 13304050    | Retention cannot be empty for a SnapMirror policy with 'create_snapshot_on_source' set to false. |
| 13304092    | Input value of the retention period property is invalid. For relationships with FlexVol volume or FlexGroup volume destinations, the duration must be in ISO 6801 format or can be infinite. For relationships with object store destinations, only duration values with Y, M or D and supported and must be in the specified range. |

| 6621045     | The property retention.warn is not supported for a policy when the retention.preserve property is false. |
| 13304129    | The property retention.warn value must be less than the property retention.count value for a rule in a policy. |
| 13304126    | Enter a value for "count" in the range indicated in the error message. |
| 13304130    | The total retention.count value for all rules in a policy cannot exceed the value indicated in the error message." |
| 13304131    | Modifying property "retention.creation_schedule" for the policy is not supported because the policy is associated with a relationship with a FlexGroup volume endpoint. |
| 13304011    | Failed to create or modify the specified Snapmirror policy. Reason: The property retention.count cannot have a value greater than 1. |
| 6621091     | The specified Snapmirror policy cannot have a rule with a preserve value of true. |
| 6621088     | The specified Snapmirror policy must have at least one rule without a schedule. |
*/
type SnapmirrorPolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapmirror policy modify default response has a 2xx status code
func (o *SnapmirrorPolicyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror policy modify default response has a 3xx status code
func (o *SnapmirrorPolicyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror policy modify default response has a 4xx status code
func (o *SnapmirrorPolicyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror policy modify default response has a 5xx status code
func (o *SnapmirrorPolicyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror policy modify default response a status code equal to that given
func (o *SnapmirrorPolicyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapmirror policy modify default response
func (o *SnapmirrorPolicyModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorPolicyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/policies/{uuid}][%d] snapmirror_policy_modify default %s", o._statusCode, payload)
}

func (o *SnapmirrorPolicyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/policies/{uuid}][%d] snapmirror_policy_modify default %s", o._statusCode, payload)
}

func (o *SnapmirrorPolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
