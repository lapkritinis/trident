// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// ActiveDirectoryPreferredDcDeleteCollectionReader is a Reader for the ActiveDirectoryPreferredDcDeleteCollection structure.
type ActiveDirectoryPreferredDcDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActiveDirectoryPreferredDcDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActiveDirectoryPreferredDcDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActiveDirectoryPreferredDcDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActiveDirectoryPreferredDcDeleteCollectionOK creates a ActiveDirectoryPreferredDcDeleteCollectionOK with default headers values
func NewActiveDirectoryPreferredDcDeleteCollectionOK() *ActiveDirectoryPreferredDcDeleteCollectionOK {
	return &ActiveDirectoryPreferredDcDeleteCollectionOK{}
}

/*
ActiveDirectoryPreferredDcDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type ActiveDirectoryPreferredDcDeleteCollectionOK struct {
}

// IsSuccess returns true when this active directory preferred dc delete collection o k response has a 2xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this active directory preferred dc delete collection o k response has a 3xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this active directory preferred dc delete collection o k response has a 4xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this active directory preferred dc delete collection o k response has a 5xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this active directory preferred dc delete collection o k response a status code equal to that given
func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the active directory preferred dc delete collection o k response
func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) Code() int {
	return 200
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] activeDirectoryPreferredDcDeleteCollectionOK", 200)
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] activeDirectoryPreferredDcDeleteCollectionOK", 200)
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActiveDirectoryPreferredDcDeleteCollectionDefault creates a ActiveDirectoryPreferredDcDeleteCollectionDefault with default headers values
func NewActiveDirectoryPreferredDcDeleteCollectionDefault(code int) *ActiveDirectoryPreferredDcDeleteCollectionDefault {
	return &ActiveDirectoryPreferredDcDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	ActiveDirectoryPreferredDcDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655507 | Failed to remove preferred-dc. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ActiveDirectoryPreferredDcDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this active directory preferred dc delete collection default response has a 2xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this active directory preferred dc delete collection default response has a 3xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this active directory preferred dc delete collection default response has a 4xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this active directory preferred dc delete collection default response has a 5xx status code
func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this active directory preferred dc delete collection default response a status code equal to that given
func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the active directory preferred dc delete collection default response
func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] active_directory_preferred_dc_delete_collection default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] active_directory_preferred_dc_delete_collection default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ActiveDirectoryPreferredDcDeleteCollectionBody active directory preferred dc delete collection body
swagger:model ActiveDirectoryPreferredDcDeleteCollectionBody
*/
type ActiveDirectoryPreferredDcDeleteCollectionBody struct {

	// active directory preferred dc response inline records
	ActiveDirectoryPreferredDcResponseInlineRecords []*models.ActiveDirectoryPreferredDc `json:"records,omitempty"`
}

// Validate validates this active directory preferred dc delete collection body
func (o *ActiveDirectoryPreferredDcDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActiveDirectoryPreferredDcResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionBody) validateActiveDirectoryPreferredDcResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.ActiveDirectoryPreferredDcResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.ActiveDirectoryPreferredDcResponseInlineRecords); i++ {
		if swag.IsZero(o.ActiveDirectoryPreferredDcResponseInlineRecords[i]) { // not required
			continue
		}

		if o.ActiveDirectoryPreferredDcResponseInlineRecords[i] != nil {
			if err := o.ActiveDirectoryPreferredDcResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this active directory preferred dc delete collection body based on the context it is used
func (o *ActiveDirectoryPreferredDcDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateActiveDirectoryPreferredDcResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ActiveDirectoryPreferredDcDeleteCollectionBody) contextValidateActiveDirectoryPreferredDcResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ActiveDirectoryPreferredDcResponseInlineRecords); i++ {

		if o.ActiveDirectoryPreferredDcResponseInlineRecords[i] != nil {
			if err := o.ActiveDirectoryPreferredDcResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *ActiveDirectoryPreferredDcDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ActiveDirectoryPreferredDcDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryPreferredDcDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
