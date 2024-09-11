// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ClusterNtpServersModifyCollectionReader is a Reader for the ClusterNtpServersModifyCollection structure.
type ClusterNtpServersModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpServersModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNtpServersModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewClusterNtpServersModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpServersModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpServersModifyCollectionOK creates a ClusterNtpServersModifyCollectionOK with default headers values
func NewClusterNtpServersModifyCollectionOK() *ClusterNtpServersModifyCollectionOK {
	return &ClusterNtpServersModifyCollectionOK{}
}

/*
ClusterNtpServersModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNtpServersModifyCollectionOK struct {
	Payload *models.NtpServerJobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers modify collection o k response has a 2xx status code
func (o *ClusterNtpServersModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers modify collection o k response has a 3xx status code
func (o *ClusterNtpServersModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers modify collection o k response has a 4xx status code
func (o *ClusterNtpServersModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers modify collection o k response has a 5xx status code
func (o *ClusterNtpServersModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers modify collection o k response a status code equal to that given
func (o *ClusterNtpServersModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ntp servers modify collection o k response
func (o *ClusterNtpServersModifyCollectionOK) Code() int {
	return 200
}

func (o *ClusterNtpServersModifyCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers][%d] clusterNtpServersModifyCollectionOK %s", 200, payload)
}

func (o *ClusterNtpServersModifyCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers][%d] clusterNtpServersModifyCollectionOK %s", 200, payload)
}

func (o *ClusterNtpServersModifyCollectionOK) GetPayload() *models.NtpServerJobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NtpServerJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersModifyCollectionAccepted creates a ClusterNtpServersModifyCollectionAccepted with default headers values
func NewClusterNtpServersModifyCollectionAccepted() *ClusterNtpServersModifyCollectionAccepted {
	return &ClusterNtpServersModifyCollectionAccepted{}
}

/*
ClusterNtpServersModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterNtpServersModifyCollectionAccepted struct {
	Payload *models.NtpServerJobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers modify collection accepted response has a 2xx status code
func (o *ClusterNtpServersModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers modify collection accepted response has a 3xx status code
func (o *ClusterNtpServersModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers modify collection accepted response has a 4xx status code
func (o *ClusterNtpServersModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers modify collection accepted response has a 5xx status code
func (o *ClusterNtpServersModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers modify collection accepted response a status code equal to that given
func (o *ClusterNtpServersModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cluster ntp servers modify collection accepted response
func (o *ClusterNtpServersModifyCollectionAccepted) Code() int {
	return 202
}

func (o *ClusterNtpServersModifyCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers][%d] clusterNtpServersModifyCollectionAccepted %s", 202, payload)
}

func (o *ClusterNtpServersModifyCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers][%d] clusterNtpServersModifyCollectionAccepted %s", 202, payload)
}

func (o *ClusterNtpServersModifyCollectionAccepted) GetPayload() *models.NtpServerJobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NtpServerJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersModifyCollectionDefault creates a ClusterNtpServersModifyCollectionDefault with default headers values
func NewClusterNtpServersModifyCollectionDefault(code int) *ClusterNtpServersModifyCollectionDefault {
	return &ClusterNtpServersModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	ClusterNtpServersModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2097163 | NTP server address was invalid. |
| 2097164 | NTP server address was invalid. |
| 2097165 | Could not resolve NTP server hostname. |
| 2097166 | NTP server address query returned no valid IP addresses. |
| 2097167 | Failed to connect to NTP server. |
| 2097169 | NTP server provided was not synchronized. |
| 2097174 | NTP server provided had too high of root distance. |
| 2097177 | NTP server provided had an invalid stratum. |
| 2097181 | NTP server address was invalid. |
| 2097182 | NTP server address was invalid. |
| 2097183 | NTP symmetric key authentication cannot be used for a node not in a cluster. |
| 2097185 | NTP key authentication failed for the provided key. |
| 2097188 | An invalid key identifier was provided. Identifiers must be in the range from 1 to 65535. |
| 2097193 | An unknown key was provided. |
| 2097194 | The field \"authentication_enabled\" cannot be false when the field NTP key is given. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterNtpServersModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ntp servers modify collection default response has a 2xx status code
func (o *ClusterNtpServersModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ntp servers modify collection default response has a 3xx status code
func (o *ClusterNtpServersModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ntp servers modify collection default response has a 4xx status code
func (o *ClusterNtpServersModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ntp servers modify collection default response has a 5xx status code
func (o *ClusterNtpServersModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ntp servers modify collection default response a status code equal to that given
func (o *ClusterNtpServersModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ntp servers modify collection default response
func (o *ClusterNtpServersModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpServersModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers][%d] cluster_ntp_servers_modify_collection default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers][%d] cluster_ntp_servers_modify_collection default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpServersModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ClusterNtpServersModifyCollectionBody cluster ntp servers modify collection body
swagger:model ClusterNtpServersModifyCollectionBody
*/
type ClusterNtpServersModifyCollectionBody struct {

	// links
	Links *models.NtpServerInlineLinks `json:"_links,omitempty"`

	// Set NTP symmetric authentication on (true) or off (false).
	// Example: true
	AuthenticationEnabled *bool `json:"authentication_enabled,omitempty"`

	// key
	Key *models.NtpKeyReference `json:"key,omitempty"`

	// ntp server response inline records
	NtpServerResponseInlineRecords []*models.NtpServer `json:"records,omitempty"`

	// NTP server host name, IPv4, or IPv6 address.
	// Example: time.nist.gov
	Server *string `json:"server,omitempty"`

	// NTP protocol version for server. Valid versions are 3, 4, or auto.
	// Example: auto
	// Enum: ["3","4","auto"]
	Version *string `json:"version,omitempty"`
}

// Validate validates this cluster ntp servers modify collection body
func (o *ClusterNtpServersModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNtpServerResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(o.Key) { // not required
		return nil
	}

	if o.Key != nil {
		if err := o.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "key")
			}
			return err
		}
	}

	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) validateNtpServerResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.NtpServerResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.NtpServerResponseInlineRecords); i++ {
		if swag.IsZero(o.NtpServerResponseInlineRecords[i]) { // not required
			continue
		}

		if o.NtpServerResponseInlineRecords[i] != nil {
			if err := o.NtpServerResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var clusterNtpServersModifyCollectionBodyTypeVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["3","4","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterNtpServersModifyCollectionBodyTypeVersionPropEnum = append(clusterNtpServersModifyCollectionBodyTypeVersionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ClusterNtpServersModifyCollectionBody
	// ClusterNtpServersModifyCollectionBody
	// version
	// Version
	// 3
	// END DEBUGGING
	// ClusterNtpServersModifyCollectionBodyVersionNr3 captures enum value "3"
	ClusterNtpServersModifyCollectionBodyVersionNr3 string = "3"

	// BEGIN DEBUGGING
	// ClusterNtpServersModifyCollectionBody
	// ClusterNtpServersModifyCollectionBody
	// version
	// Version
	// 4
	// END DEBUGGING
	// ClusterNtpServersModifyCollectionBodyVersionNr4 captures enum value "4"
	ClusterNtpServersModifyCollectionBodyVersionNr4 string = "4"

	// BEGIN DEBUGGING
	// ClusterNtpServersModifyCollectionBody
	// ClusterNtpServersModifyCollectionBody
	// version
	// Version
	// auto
	// END DEBUGGING
	// ClusterNtpServersModifyCollectionBodyVersionAuto captures enum value "auto"
	ClusterNtpServersModifyCollectionBodyVersionAuto string = "auto"
)

// prop value enum
func (o *ClusterNtpServersModifyCollectionBody) validateVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterNtpServersModifyCollectionBodyTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(o.Version) { // not required
		return nil
	}

	// value enum
	if err := o.validateVersionEnum("info"+"."+"version", "body", *o.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster ntp servers modify collection body based on the context it is used
func (o *ClusterNtpServersModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNtpServerResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if o.Key != nil {
		if err := o.Key.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "key")
			}
			return err
		}
	}

	return nil
}

func (o *ClusterNtpServersModifyCollectionBody) contextValidateNtpServerResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.NtpServerResponseInlineRecords); i++ {

		if o.NtpServerResponseInlineRecords[i] != nil {
			if err := o.NtpServerResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *ClusterNtpServersModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClusterNtpServersModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res ClusterNtpServersModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NtpServerInlineLinks ntp server inline links
swagger:model ntp_server_inline__links
*/
type NtpServerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this ntp server inline links
func (o *NtpServerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NtpServerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ntp server inline links based on the context it is used
func (o *NtpServerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NtpServerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NtpServerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NtpServerInlineLinks) UnmarshalBinary(b []byte) error {
	var res NtpServerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
