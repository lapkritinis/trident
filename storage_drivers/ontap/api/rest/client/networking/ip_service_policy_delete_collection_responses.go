// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IPServicePolicyDeleteCollectionReader is a Reader for the IPServicePolicyDeleteCollection structure.
type IPServicePolicyDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPServicePolicyDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPServicePolicyDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPServicePolicyDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPServicePolicyDeleteCollectionOK creates a IPServicePolicyDeleteCollectionOK with default headers values
func NewIPServicePolicyDeleteCollectionOK() *IPServicePolicyDeleteCollectionOK {
	return &IPServicePolicyDeleteCollectionOK{}
}

/*
IPServicePolicyDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type IPServicePolicyDeleteCollectionOK struct {
}

// IsSuccess returns true when this ip service policy delete collection o k response has a 2xx status code
func (o *IPServicePolicyDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip service policy delete collection o k response has a 3xx status code
func (o *IPServicePolicyDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip service policy delete collection o k response has a 4xx status code
func (o *IPServicePolicyDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip service policy delete collection o k response has a 5xx status code
func (o *IPServicePolicyDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip service policy delete collection o k response a status code equal to that given
func (o *IPServicePolicyDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ip service policy delete collection o k response
func (o *IPServicePolicyDeleteCollectionOK) Code() int {
	return 200
}

func (o *IPServicePolicyDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ip/service-policies][%d] ipServicePolicyDeleteCollectionOK", 200)
}

func (o *IPServicePolicyDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /network/ip/service-policies][%d] ipServicePolicyDeleteCollectionOK", 200)
}

func (o *IPServicePolicyDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIPServicePolicyDeleteCollectionDefault creates a IPServicePolicyDeleteCollectionDefault with default headers values
func NewIPServicePolicyDeleteCollectionDefault(code int) *IPServicePolicyDeleteCollectionDefault {
	return &IPServicePolicyDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	IPServicePolicyDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621740 | An unexpected error when trying to determine whether the target Vserver was locked or not on this cluster. |
| 53281927 | Service policies owned by the system cannot be deleted. |
| 53281928 | Service policies assigned to LIFs cannot be deleted. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IPServicePolicyDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ip service policy delete collection default response has a 2xx status code
func (o *IPServicePolicyDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip service policy delete collection default response has a 3xx status code
func (o *IPServicePolicyDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip service policy delete collection default response has a 4xx status code
func (o *IPServicePolicyDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip service policy delete collection default response has a 5xx status code
func (o *IPServicePolicyDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip service policy delete collection default response a status code equal to that given
func (o *IPServicePolicyDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ip service policy delete collection default response
func (o *IPServicePolicyDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *IPServicePolicyDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/service-policies][%d] ip_service_policy_delete_collection default %s", o._statusCode, payload)
}

func (o *IPServicePolicyDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/service-policies][%d] ip_service_policy_delete_collection default %s", o._statusCode, payload)
}

func (o *IPServicePolicyDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPServicePolicyDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IPServicePolicyDeleteCollectionBody IP service policy delete collection body
swagger:model IPServicePolicyDeleteCollectionBody
*/
type IPServicePolicyDeleteCollectionBody struct {

	// ip service policy response inline records
	IPServicePolicyResponseInlineRecords []*models.IPServicePolicy `json:"records,omitempty"`
}

// Validate validates this IP service policy delete collection body
func (o *IPServicePolicyDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPServicePolicyResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IPServicePolicyDeleteCollectionBody) validateIPServicePolicyResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.IPServicePolicyResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.IPServicePolicyResponseInlineRecords); i++ {
		if swag.IsZero(o.IPServicePolicyResponseInlineRecords[i]) { // not required
			continue
		}

		if o.IPServicePolicyResponseInlineRecords[i] != nil {
			if err := o.IPServicePolicyResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this IP service policy delete collection body based on the context it is used
func (o *IPServicePolicyDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPServicePolicyResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IPServicePolicyDeleteCollectionBody) contextValidateIPServicePolicyResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.IPServicePolicyResponseInlineRecords); i++ {

		if o.IPServicePolicyResponseInlineRecords[i] != nil {
			if err := o.IPServicePolicyResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IPServicePolicyDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IPServicePolicyDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res IPServicePolicyDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
