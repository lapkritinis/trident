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

// FpolicyEventModifyCollectionReader is a Reader for the FpolicyEventModifyCollection structure.
type FpolicyEventModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEventModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEventModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEventModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEventModifyCollectionOK creates a FpolicyEventModifyCollectionOK with default headers values
func NewFpolicyEventModifyCollectionOK() *FpolicyEventModifyCollectionOK {
	return &FpolicyEventModifyCollectionOK{}
}

/*
FpolicyEventModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEventModifyCollectionOK struct {
}

// IsSuccess returns true when this fpolicy event modify collection o k response has a 2xx status code
func (o *FpolicyEventModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy event modify collection o k response has a 3xx status code
func (o *FpolicyEventModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy event modify collection o k response has a 4xx status code
func (o *FpolicyEventModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy event modify collection o k response has a 5xx status code
func (o *FpolicyEventModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy event modify collection o k response a status code equal to that given
func (o *FpolicyEventModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy event modify collection o k response
func (o *FpolicyEventModifyCollectionOK) Code() int {
	return 200
}

func (o *FpolicyEventModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/events][%d] fpolicyEventModifyCollectionOK", 200)
}

func (o *FpolicyEventModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/events][%d] fpolicyEventModifyCollectionOK", 200)
}

func (o *FpolicyEventModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyEventModifyCollectionDefault creates a FpolicyEventModifyCollectionDefault with default headers values
func NewFpolicyEventModifyCollectionDefault(code int) *FpolicyEventModifyCollectionDefault {
	return &FpolicyEventModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	FpolicyEventModifyCollectionDefault describes a response with status code -1, with default header values.

	| Error Code | Description |

| ---------- | ----------- |
| 9764873    | The event is a cluster event |
| 9764929    | The file operation is not supported by the protocol |
| 9764955    | The filter is not supported by the protocol |
| 9764930    | The filter is not supported by any of the file operations |
| 9764946    | The protocol is specified without file operation or a file operation and filter pair |
| 9765048    | The monitor fileop failure option is set without protocol and file operations |
*/
type FpolicyEventModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy event modify collection default response has a 2xx status code
func (o *FpolicyEventModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy event modify collection default response has a 3xx status code
func (o *FpolicyEventModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy event modify collection default response has a 4xx status code
func (o *FpolicyEventModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy event modify collection default response has a 5xx status code
func (o *FpolicyEventModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy event modify collection default response a status code equal to that given
func (o *FpolicyEventModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy event modify collection default response
func (o *FpolicyEventModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEventModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/events][%d] fpolicy_event_modify_collection default %s", o._statusCode, payload)
}

func (o *FpolicyEventModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/events][%d] fpolicy_event_modify_collection default %s", o._statusCode, payload)
}

func (o *FpolicyEventModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEventModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
FpolicyEventModifyCollectionBody fpolicy event modify collection body
swagger:model FpolicyEventModifyCollectionBody
*/
type FpolicyEventModifyCollectionBody struct {

	// file operations
	FileOperations *models.FpolicyEventInlineFileOperations `json:"file_operations,omitempty"`

	// filters
	Filters *models.FpolicyEventInlineFilters `json:"filters,omitempty"`

	// fpolicy event response inline records
	FpolicyEventResponseInlineRecords []*models.FpolicyEvent `json:"records,omitempty"`

	// Specifies whether failed file operations monitoring is required.
	MonitorFileopFailure *bool `json:"monitor_fileop_failure,omitempty"`

	// Specifies the name of the FPolicy event.
	// Example: event_cifs
	Name *string `json:"name,omitempty"`

	// Protocol for which event is created. If you specify protocol, then you
	// must also specify a valid value for the file operation parameters.
	//   The value of this parameter must be one of the following:
	//     * cifs  - for the CIFS protocol.
	//     * nfsv3 - for the NFSv3 protocol.
	//     * nfsv4 - for the NFSv4 protocol.
	//
	// Enum: ["cifs","nfsv3","nfsv4"]
	Protocol *string `json:"protocol,omitempty"`

	// svm
	Svm *models.FpolicyEventInlineSvm `json:"svm,omitempty"`

	// Specifies whether volume operation monitoring is required.
	VolumeMonitoring *bool `json:"volume_monitoring,omitempty"`
}

// Validate validates this fpolicy event modify collection body
func (o *FpolicyEventModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFileOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFpolicyEventResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProtocol(formats); err != nil {
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

func (o *FpolicyEventModifyCollectionBody) validateFileOperations(formats strfmt.Registry) error {
	if swag.IsZero(o.FileOperations) { // not required
		return nil
	}

	if o.FileOperations != nil {
		if err := o.FileOperations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "file_operations")
			}
			return err
		}
	}

	return nil
}

func (o *FpolicyEventModifyCollectionBody) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(o.Filters) { // not required
		return nil
	}

	if o.Filters != nil {
		if err := o.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "filters")
			}
			return err
		}
	}

	return nil
}

func (o *FpolicyEventModifyCollectionBody) validateFpolicyEventResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.FpolicyEventResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.FpolicyEventResponseInlineRecords); i++ {
		if swag.IsZero(o.FpolicyEventResponseInlineRecords[i]) { // not required
			continue
		}

		if o.FpolicyEventResponseInlineRecords[i] != nil {
			if err := o.FpolicyEventResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var fpolicyEventModifyCollectionBodyTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cifs","nfsv3","nfsv4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEventModifyCollectionBodyTypeProtocolPropEnum = append(fpolicyEventModifyCollectionBodyTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// FpolicyEventModifyCollectionBody
	// FpolicyEventModifyCollectionBody
	// protocol
	// Protocol
	// cifs
	// END DEBUGGING
	// FpolicyEventModifyCollectionBodyProtocolCifs captures enum value "cifs"
	FpolicyEventModifyCollectionBodyProtocolCifs string = "cifs"

	// BEGIN DEBUGGING
	// FpolicyEventModifyCollectionBody
	// FpolicyEventModifyCollectionBody
	// protocol
	// Protocol
	// nfsv3
	// END DEBUGGING
	// FpolicyEventModifyCollectionBodyProtocolNfsv3 captures enum value "nfsv3"
	FpolicyEventModifyCollectionBodyProtocolNfsv3 string = "nfsv3"

	// BEGIN DEBUGGING
	// FpolicyEventModifyCollectionBody
	// FpolicyEventModifyCollectionBody
	// protocol
	// Protocol
	// nfsv4
	// END DEBUGGING
	// FpolicyEventModifyCollectionBodyProtocolNfsv4 captures enum value "nfsv4"
	FpolicyEventModifyCollectionBodyProtocolNfsv4 string = "nfsv4"
)

// prop value enum
func (o *FpolicyEventModifyCollectionBody) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEventModifyCollectionBodyTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *FpolicyEventModifyCollectionBody) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(o.Protocol) { // not required
		return nil
	}

	// value enum
	if err := o.validateProtocolEnum("info"+"."+"protocol", "body", *o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *FpolicyEventModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this fpolicy event modify collection body based on the context it is used
func (o *FpolicyEventModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFileOperations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFpolicyEventResponseInlineRecords(ctx, formats); err != nil {
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

func (o *FpolicyEventModifyCollectionBody) contextValidateFileOperations(ctx context.Context, formats strfmt.Registry) error {

	if o.FileOperations != nil {
		if err := o.FileOperations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "file_operations")
			}
			return err
		}
	}

	return nil
}

func (o *FpolicyEventModifyCollectionBody) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if o.Filters != nil {
		if err := o.Filters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "filters")
			}
			return err
		}
	}

	return nil
}

func (o *FpolicyEventModifyCollectionBody) contextValidateFpolicyEventResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FpolicyEventResponseInlineRecords); i++ {

		if o.FpolicyEventResponseInlineRecords[i] != nil {
			if err := o.FpolicyEventResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FpolicyEventModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (o *FpolicyEventModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FpolicyEventModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res FpolicyEventModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FpolicyEventInlineFileOperations Specifies the file operations for the FPolicy event. You must specify a valid protocol in the protocol parameter.
// The event will check the operations specified from all client requests using the protocol.
//
swagger:model fpolicy_event_inline_file_operations
*/
type FpolicyEventInlineFileOperations struct {

	// Access operations
	Access *bool `json:"access,omitempty"`

	// File close operations
	Close *bool `json:"close,omitempty"`

	// File create operations
	Create *bool `json:"create,omitempty"`

	// Directory create operations
	CreateDir *bool `json:"create_dir,omitempty"`

	// File delete operations
	Delete *bool `json:"delete,omitempty"`

	// Directory delete operations
	DeleteDir *bool `json:"delete_dir,omitempty"`

	// Get attribute operations
	Getattr *bool `json:"getattr,omitempty"`

	// Link operations
	Link *bool `json:"link,omitempty"`

	// Lookup operations
	Lookup *bool `json:"lookup,omitempty"`

	// File open operations
	Open *bool `json:"open,omitempty"`

	// File read operations
	Read *bool `json:"read,omitempty"`

	// File rename operations
	Rename *bool `json:"rename,omitempty"`

	// Directory rename operations
	RenameDir *bool `json:"rename_dir,omitempty"`

	// Set attribute operations
	Setattr *bool `json:"setattr,omitempty"`

	// Symbolic link operations
	Symlink *bool `json:"symlink,omitempty"`

	// File write operations
	Write *bool `json:"write,omitempty"`
}

// Validate validates this fpolicy event inline file operations
func (o *FpolicyEventInlineFileOperations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy event inline file operations based on context it is used
func (o *FpolicyEventInlineFileOperations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FpolicyEventInlineFileOperations) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FpolicyEventInlineFileOperations) UnmarshalBinary(b []byte) error {
	var res FpolicyEventInlineFileOperations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FpolicyEventInlineFilters Specifies the list of filters for a given file operation for the specified protocol.
// When you specify the filters, you must specify the valid protocols and a valid file operations.
//
swagger:model fpolicy_event_inline_filters
*/
type FpolicyEventInlineFilters struct {

	// Filter the client request for close with modification.
	CloseWithModification *bool `json:"close_with_modification,omitempty"`

	// Filter the client request for close with read.
	CloseWithRead *bool `json:"close_with_read,omitempty"`

	// Filter the client request for close without modification.
	CloseWithoutModification *bool `json:"close_without_modification,omitempty"`

	// Filter the client requests for directory operations. When this filter is specified directory operations are not monitored.
	ExcludeDirectory *bool `json:"exclude_directory,omitempty"`

	// Filter the client requests for the first-read.
	FirstRead *bool `json:"first_read,omitempty"`

	// Filter the client requests for the first-write.
	FirstWrite *bool `json:"first_write,omitempty"`

	// Filter the client request for alternate data stream.
	MonitorAds *bool `json:"monitor_ads,omitempty"`

	// Filter the client request for offline bit set. FPolicy server receives notification only when offline files are accessed.
	OfflineBit *bool `json:"offline_bit,omitempty"`

	// Filter the client request for open with delete intent.
	OpenWithDeleteIntent *bool `json:"open_with_delete_intent,omitempty"`

	// Filter the client request for open with write intent.
	OpenWithWriteIntent *bool `json:"open_with_write_intent,omitempty"`

	// Filter the client setattr requests for changing the access time of a file or directory.
	SetattrWithAccessTimeChange *bool `json:"setattr_with_access_time_change,omitempty"`

	// Filter the client setattr requests for changing the allocation size of a file.
	SetattrWithAllocationSizeChange *bool `json:"setattr_with_allocation_size_change,omitempty"`

	// Filter the client setattr requests for changing the creation time of a file or directory.
	SetattrWithCreationTimeChange *bool `json:"setattr_with_creation_time_change,omitempty"`

	// Filter the client setattr requests for changing dacl on a file or directory.
	SetattrWithDaclChange *bool `json:"setattr_with_dacl_change,omitempty"`

	// Filter the client setattr requests for changing group of a file or directory.
	SetattrWithGroupChange *bool `json:"setattr_with_group_change,omitempty"`

	// Filter the client setattr requests for changing the mode bits on a file or directory.
	SetattrWithModeChange *bool `json:"setattr_with_mode_change,omitempty"`

	// Filter the client setattr requests for changing the modification time of a file or directory.
	SetattrWithModifyTimeChange *bool `json:"setattr_with_modify_time_change,omitempty"`

	// Filter the client setattr requests for changing owner of a file or directory.
	SetattrWithOwnerChange *bool `json:"setattr_with_owner_change,omitempty"`

	// Filter the client setattr requests for changing sacl on a file or directory.
	SetattrWithSaclChange *bool `json:"setattr_with_sacl_change,omitempty"`

	// Filter the client setattr requests for changing the size of a file.
	SetattrWithSizeChange *bool `json:"setattr_with_size_change,omitempty"`

	// Filter the client request for write with size change.
	WriteWithSizeChange *bool `json:"write_with_size_change,omitempty"`
}

// Validate validates this fpolicy event inline filters
func (o *FpolicyEventInlineFilters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy event inline filters based on context it is used
func (o *FpolicyEventInlineFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FpolicyEventInlineFilters) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FpolicyEventInlineFilters) UnmarshalBinary(b []byte) error {
	var res FpolicyEventInlineFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FpolicyEventInlineSvm fpolicy event inline svm
swagger:model fpolicy_event_inline_svm
*/
type FpolicyEventInlineSvm struct {

	// SVM UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this fpolicy event inline svm
func (o *FpolicyEventInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fpolicy event inline svm based on the context it is used
func (o *FpolicyEventInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FpolicyEventInlineSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"svm"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FpolicyEventInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FpolicyEventInlineSvm) UnmarshalBinary(b []byte) error {
	var res FpolicyEventInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
