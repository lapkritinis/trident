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
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FpolicyPersistentStoreModifyCollectionReader is a Reader for the FpolicyPersistentStoreModifyCollection structure.
type FpolicyPersistentStoreModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPersistentStoreModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyPersistentStoreModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPersistentStoreModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPersistentStoreModifyCollectionOK creates a FpolicyPersistentStoreModifyCollectionOK with default headers values
func NewFpolicyPersistentStoreModifyCollectionOK() *FpolicyPersistentStoreModifyCollectionOK {
	return &FpolicyPersistentStoreModifyCollectionOK{}
}

/*
FpolicyPersistentStoreModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyPersistentStoreModifyCollectionOK struct {
}

// IsSuccess returns true when this fpolicy persistent store modify collection o k response has a 2xx status code
func (o *FpolicyPersistentStoreModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy persistent store modify collection o k response has a 3xx status code
func (o *FpolicyPersistentStoreModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy persistent store modify collection o k response has a 4xx status code
func (o *FpolicyPersistentStoreModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy persistent store modify collection o k response has a 5xx status code
func (o *FpolicyPersistentStoreModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy persistent store modify collection o k response a status code equal to that given
func (o *FpolicyPersistentStoreModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy persistent store modify collection o k response
func (o *FpolicyPersistentStoreModifyCollectionOK) Code() int {
	return 200
}

func (o *FpolicyPersistentStoreModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicyPersistentStoreModifyCollectionOK", 200)
}

func (o *FpolicyPersistentStoreModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicyPersistentStoreModifyCollectionOK", 200)
}

func (o *FpolicyPersistentStoreModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyPersistentStoreModifyCollectionDefault creates a FpolicyPersistentStoreModifyCollectionDefault with default headers values
func NewFpolicyPersistentStoreModifyCollectionDefault(code int) *FpolicyPersistentStoreModifyCollectionDefault {
	return &FpolicyPersistentStoreModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	FpolicyPersistentStoreModifyCollectionDefault describes a response with status code -1, with default header values.

	| Error Code | Description |

| ---------- | ----------- |
| 9765053    | Volume mentioned does not exist in the SVM |
| 9765056    | Specified Persistent Store name does not exist in the SVM |
| 9765069    | Volume is not a FlexVol volume. Only FlexVol volumes are supported for this operation. |
| 9765070    | Cannot update a Persistent Store volume when it is part of an enabled policy |
| 9765072    | Volume is not of type RW. Only volumes of type RW are supported. |
| 9765074    | Size is a required parameter for the creation of the Persistent Store. |
| 9765077    | The SVM is not associated with any aggregates. |
| 524849     | Attribute cannot be modified for the target volume because of the specified reason. |
| 917533     | The target volume is not eligible for modification due to certain constraints, such as the volume being offline or currently engaged in a volume move operation. |
| 917826     | Volume minimum autosize must be greater than zero. |
| 918647     | The designated size exceeds the maximum available capacity of the intended volume. |
| 917534     | The value specified for the field \\"-size\\" is too small. Update the field \\"-size\\" with the minimum size allowed and retry the operation. |
*/
type FpolicyPersistentStoreModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy persistent store modify collection default response has a 2xx status code
func (o *FpolicyPersistentStoreModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy persistent store modify collection default response has a 3xx status code
func (o *FpolicyPersistentStoreModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy persistent store modify collection default response has a 4xx status code
func (o *FpolicyPersistentStoreModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy persistent store modify collection default response has a 5xx status code
func (o *FpolicyPersistentStoreModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy persistent store modify collection default response a status code equal to that given
func (o *FpolicyPersistentStoreModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy persistent store modify collection default response
func (o *FpolicyPersistentStoreModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPersistentStoreModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicy_persistent_store_modify_collection default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicy_persistent_store_modify_collection default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPersistentStoreModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
FpolicyPersistentStoreModifyCollectionBody fpolicy persistent store modify collection body
swagger:model FpolicyPersistentStoreModifyCollectionBody
*/
type FpolicyPersistentStoreModifyCollectionBody struct {

	// Autosize mode for the volume.<br>grow &dash; Volume automatically grows in response to the amount of space used.<br>grow_shrink &dash; Volume grows or shrinks in response to the amount of space used.<br>off &dash; Autosizing of the volume is disabled.
	// Enum: ["grow","grow_shrink","off"]
	AutosizeMode *string `json:"autosize_mode,omitempty"`

	// fpolicy persistent store response inline records
	FpolicyPersistentStoreResponseInlineRecords []*models.FpolicyPersistentStore `json:"records,omitempty"`

	// The name specified for the FPolicy Persistent Store.
	// Example: ps1
	Name *string `json:"name,omitempty"`

	// The size of the Persistent Store volume, in bytes.
	// Example: 100M
	Size *int64 `json:"size,omitempty"`

	// svm
	Svm *models.FpolicyPersistentStoreInlineSvm `json:"svm,omitempty"`

	// The specified volume to store the events for the FPolicy Persistent Store.
	// Example: psvol
	Volume *string `json:"volume,omitempty"`
}

// Validate validates this fpolicy persistent store modify collection body
func (o *FpolicyPersistentStoreModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAutosizeMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFpolicyPersistentStoreResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fpolicyPersistentStoreModifyCollectionBodyTypeAutosizeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["grow","grow_shrink","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyPersistentStoreModifyCollectionBodyTypeAutosizeModePropEnum = append(fpolicyPersistentStoreModifyCollectionBodyTypeAutosizeModePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// FpolicyPersistentStoreModifyCollectionBody
	// FpolicyPersistentStoreModifyCollectionBody
	// autosize_mode
	// AutosizeMode
	// grow
	// END DEBUGGING
	// FpolicyPersistentStoreModifyCollectionBodyAutosizeModeGrow captures enum value "grow"
	FpolicyPersistentStoreModifyCollectionBodyAutosizeModeGrow string = "grow"

	// BEGIN DEBUGGING
	// FpolicyPersistentStoreModifyCollectionBody
	// FpolicyPersistentStoreModifyCollectionBody
	// autosize_mode
	// AutosizeMode
	// grow_shrink
	// END DEBUGGING
	// FpolicyPersistentStoreModifyCollectionBodyAutosizeModeGrowShrink captures enum value "grow_shrink"
	FpolicyPersistentStoreModifyCollectionBodyAutosizeModeGrowShrink string = "grow_shrink"

	// BEGIN DEBUGGING
	// FpolicyPersistentStoreModifyCollectionBody
	// FpolicyPersistentStoreModifyCollectionBody
	// autosize_mode
	// AutosizeMode
	// off
	// END DEBUGGING
	// FpolicyPersistentStoreModifyCollectionBodyAutosizeModeOff captures enum value "off"
	FpolicyPersistentStoreModifyCollectionBodyAutosizeModeOff string = "off"
)

// prop value enum
func (o *FpolicyPersistentStoreModifyCollectionBody) validateAutosizeModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyPersistentStoreModifyCollectionBodyTypeAutosizeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *FpolicyPersistentStoreModifyCollectionBody) validateAutosizeMode(formats strfmt.Registry) error {
	if swag.IsZero(o.AutosizeMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateAutosizeModeEnum("info"+"."+"autosize_mode", "body", *o.AutosizeMode); err != nil {
		return err
	}

	return nil
}

func (o *FpolicyPersistentStoreModifyCollectionBody) validateFpolicyPersistentStoreResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.FpolicyPersistentStoreResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.FpolicyPersistentStoreResponseInlineRecords); i++ {
		if swag.IsZero(o.FpolicyPersistentStoreResponseInlineRecords[i]) { // not required
			continue
		}

		if o.FpolicyPersistentStoreResponseInlineRecords[i] != nil {
			if err := o.FpolicyPersistentStoreResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FpolicyPersistentStoreModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(o.Svm) { // not required
		return nil
	}

	if o.Svm != nil {
		if err := o.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fpolicy persistent store modify collection body based on the context it is used
func (o *FpolicyPersistentStoreModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFpolicyPersistentStoreResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FpolicyPersistentStoreModifyCollectionBody) contextValidateFpolicyPersistentStoreResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FpolicyPersistentStoreResponseInlineRecords); i++ {

		if o.FpolicyPersistentStoreResponseInlineRecords[i] != nil {
			if err := o.FpolicyPersistentStoreResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FpolicyPersistentStoreModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if o.Svm != nil {
		if err := o.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FpolicyPersistentStoreModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FpolicyPersistentStoreModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res FpolicyPersistentStoreModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FpolicyPersistentStoreInlineSvm fpolicy persistent store inline svm
swagger:model fpolicy_persistent_store_inline_svm
*/
type FpolicyPersistentStoreInlineSvm struct {

	// SVM UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this fpolicy persistent store inline svm
func (o *FpolicyPersistentStoreInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fpolicy persistent store inline svm based on the context it is used
func (o *FpolicyPersistentStoreInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FpolicyPersistentStoreInlineSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"svm"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FpolicyPersistentStoreInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FpolicyPersistentStoreInlineSvm) UnmarshalBinary(b []byte) error {
	var res FpolicyPersistentStoreInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
