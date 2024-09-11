// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// FlexcacheDeleteCollectionReader is a Reader for the FlexcacheDeleteCollection structure.
type FlexcacheDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlexcacheDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlexcacheDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFlexcacheDeleteCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFlexcacheDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFlexcacheDeleteCollectionOK creates a FlexcacheDeleteCollectionOK with default headers values
func NewFlexcacheDeleteCollectionOK() *FlexcacheDeleteCollectionOK {
	return &FlexcacheDeleteCollectionOK{}
}

/*
FlexcacheDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type FlexcacheDeleteCollectionOK struct {
	Payload *models.FlexcacheJobLinkResponse
}

// IsSuccess returns true when this flexcache delete collection o k response has a 2xx status code
func (o *FlexcacheDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flexcache delete collection o k response has a 3xx status code
func (o *FlexcacheDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flexcache delete collection o k response has a 4xx status code
func (o *FlexcacheDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flexcache delete collection o k response has a 5xx status code
func (o *FlexcacheDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flexcache delete collection o k response a status code equal to that given
func (o *FlexcacheDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the flexcache delete collection o k response
func (o *FlexcacheDeleteCollectionOK) Code() int {
	return 200
}

func (o *FlexcacheDeleteCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches][%d] flexcacheDeleteCollectionOK %s", 200, payload)
}

func (o *FlexcacheDeleteCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches][%d] flexcacheDeleteCollectionOK %s", 200, payload)
}

func (o *FlexcacheDeleteCollectionOK) GetPayload() *models.FlexcacheJobLinkResponse {
	return o.Payload
}

func (o *FlexcacheDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlexcacheJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlexcacheDeleteCollectionAccepted creates a FlexcacheDeleteCollectionAccepted with default headers values
func NewFlexcacheDeleteCollectionAccepted() *FlexcacheDeleteCollectionAccepted {
	return &FlexcacheDeleteCollectionAccepted{}
}

/*
FlexcacheDeleteCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FlexcacheDeleteCollectionAccepted struct {
	Payload *models.FlexcacheJobLinkResponse
}

// IsSuccess returns true when this flexcache delete collection accepted response has a 2xx status code
func (o *FlexcacheDeleteCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flexcache delete collection accepted response has a 3xx status code
func (o *FlexcacheDeleteCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flexcache delete collection accepted response has a 4xx status code
func (o *FlexcacheDeleteCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this flexcache delete collection accepted response has a 5xx status code
func (o *FlexcacheDeleteCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this flexcache delete collection accepted response a status code equal to that given
func (o *FlexcacheDeleteCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the flexcache delete collection accepted response
func (o *FlexcacheDeleteCollectionAccepted) Code() int {
	return 202
}

func (o *FlexcacheDeleteCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches][%d] flexcacheDeleteCollectionAccepted %s", 202, payload)
}

func (o *FlexcacheDeleteCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches][%d] flexcacheDeleteCollectionAccepted %s", 202, payload)
}

func (o *FlexcacheDeleteCollectionAccepted) GetPayload() *models.FlexcacheJobLinkResponse {
	return o.Payload
}

func (o *FlexcacheDeleteCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlexcacheJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlexcacheDeleteCollectionDefault creates a FlexcacheDeleteCollectionDefault with default headers values
func NewFlexcacheDeleteCollectionDefault(code int) *FlexcacheDeleteCollectionDefault {
	return &FlexcacheDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	FlexcacheDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 66846879   | The specified volume UUID is not a FlexCache volume |
| 66846731   | Failed to delete the FlexCache volume |
| 524546     | Failed to delete the FlexCache volume because the FlexCache volume is not unmounted |
| 66846980   | Failed to delete the FlexCache volume because it has the writeback property enabled |
*/
type FlexcacheDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this flexcache delete collection default response has a 2xx status code
func (o *FlexcacheDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this flexcache delete collection default response has a 3xx status code
func (o *FlexcacheDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this flexcache delete collection default response has a 4xx status code
func (o *FlexcacheDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this flexcache delete collection default response has a 5xx status code
func (o *FlexcacheDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this flexcache delete collection default response a status code equal to that given
func (o *FlexcacheDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the flexcache delete collection default response
func (o *FlexcacheDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *FlexcacheDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches][%d] flexcache_delete_collection default %s", o._statusCode, payload)
}

func (o *FlexcacheDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/flexcache/flexcaches][%d] flexcache_delete_collection default %s", o._statusCode, payload)
}

func (o *FlexcacheDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FlexcacheDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
FlexcacheDeleteCollectionBody flexcache delete collection body
swagger:model FlexcacheDeleteCollectionBody
*/
type FlexcacheDeleteCollectionBody struct {

	// flexcache response inline records
	FlexcacheResponseInlineRecords []*models.Flexcache `json:"records,omitempty"`
}

// Validate validates this flexcache delete collection body
func (o *FlexcacheDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFlexcacheResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FlexcacheDeleteCollectionBody) validateFlexcacheResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.FlexcacheResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.FlexcacheResponseInlineRecords); i++ {
		if swag.IsZero(o.FlexcacheResponseInlineRecords[i]) { // not required
			continue
		}

		if o.FlexcacheResponseInlineRecords[i] != nil {
			if err := o.FlexcacheResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this flexcache delete collection body based on the context it is used
func (o *FlexcacheDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFlexcacheResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FlexcacheDeleteCollectionBody) contextValidateFlexcacheResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FlexcacheResponseInlineRecords); i++ {

		if o.FlexcacheResponseInlineRecords[i] != nil {
			if err := o.FlexcacheResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *FlexcacheDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FlexcacheDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res FlexcacheDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
