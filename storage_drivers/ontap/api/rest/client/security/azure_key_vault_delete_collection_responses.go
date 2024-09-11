// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AzureKeyVaultDeleteCollectionReader is a Reader for the AzureKeyVaultDeleteCollection structure.
type AzureKeyVaultDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureKeyVaultDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAzureKeyVaultDeleteCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultDeleteCollectionOK creates a AzureKeyVaultDeleteCollectionOK with default headers values
func NewAzureKeyVaultDeleteCollectionOK() *AzureKeyVaultDeleteCollectionOK {
	return &AzureKeyVaultDeleteCollectionOK{}
}

/*
AzureKeyVaultDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AzureKeyVaultDeleteCollectionOK struct {
}

// IsSuccess returns true when this azure key vault delete collection o k response has a 2xx status code
func (o *AzureKeyVaultDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault delete collection o k response has a 3xx status code
func (o *AzureKeyVaultDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault delete collection o k response has a 4xx status code
func (o *AzureKeyVaultDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault delete collection o k response has a 5xx status code
func (o *AzureKeyVaultDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault delete collection o k response a status code equal to that given
func (o *AzureKeyVaultDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the azure key vault delete collection o k response
func (o *AzureKeyVaultDeleteCollectionOK) Code() int {
	return 200
}

func (o *AzureKeyVaultDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /security/azure-key-vaults][%d] azureKeyVaultDeleteCollectionOK", 200)
}

func (o *AzureKeyVaultDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /security/azure-key-vaults][%d] azureKeyVaultDeleteCollectionOK", 200)
}

func (o *AzureKeyVaultDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAzureKeyVaultDeleteCollectionAccepted creates a AzureKeyVaultDeleteCollectionAccepted with default headers values
func NewAzureKeyVaultDeleteCollectionAccepted() *AzureKeyVaultDeleteCollectionAccepted {
	return &AzureKeyVaultDeleteCollectionAccepted{}
}

/*
AzureKeyVaultDeleteCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AzureKeyVaultDeleteCollectionAccepted struct {
	Payload *models.AzureKeyVaultJobLinkResponse
}

// IsSuccess returns true when this azure key vault delete collection accepted response has a 2xx status code
func (o *AzureKeyVaultDeleteCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault delete collection accepted response has a 3xx status code
func (o *AzureKeyVaultDeleteCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault delete collection accepted response has a 4xx status code
func (o *AzureKeyVaultDeleteCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault delete collection accepted response has a 5xx status code
func (o *AzureKeyVaultDeleteCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault delete collection accepted response a status code equal to that given
func (o *AzureKeyVaultDeleteCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the azure key vault delete collection accepted response
func (o *AzureKeyVaultDeleteCollectionAccepted) Code() int {
	return 202
}

func (o *AzureKeyVaultDeleteCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/azure-key-vaults][%d] azureKeyVaultDeleteCollectionAccepted %s", 202, payload)
}

func (o *AzureKeyVaultDeleteCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/azure-key-vaults][%d] azureKeyVaultDeleteCollectionAccepted %s", 202, payload)
}

func (o *AzureKeyVaultDeleteCollectionAccepted) GetPayload() *models.AzureKeyVaultJobLinkResponse {
	return o.Payload
}

func (o *AzureKeyVaultDeleteCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureKeyVaultJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultDeleteCollectionDefault creates a AzureKeyVaultDeleteCollectionDefault with default headers values
func NewAzureKeyVaultDeleteCollectionDefault(code int) *AzureKeyVaultDeleteCollectionDefault {
	return &AzureKeyVaultDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536242 | One or more self-encrypting drives are assigned an authentication key. |
| 65536243 | Cannot determine authentication key presence on one or more self-encrypting drives. |
| 65536817 | Internal error. Failed to determine if key manager is safe to disable. |
| 65536827 | Internal error. Failed to determine if the given SVM has any encrypted volumes. |
| 65536834 | Internal error. Failed to get existing key-server details for the given SVM. |
| 65536867 | Volume encryption keys (VEK) for one or more encrypted volumes are stored on the key manager configured for the given SVM. |
| 65536883 | Internal error. Volume encryption key is missing for a volume. |
| 65536884 | Internal error. Volume encryption key is invalid for a volume. |
| 65536924 | Cannot remove key manager that still contains one or more NSE authentication keys. |
| 65537120 | Azure Key Vault is not configured for the given SVM. |
| 196608080 | One or more nodes in the cluster have the root volume encrypted using NVE (NetApp Volume Encryption). |
| 196608301 | Internal error. Failed to get encryption type. |
| 196608305 | NAE aggregates found in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AzureKeyVaultDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this azure key vault delete collection default response has a 2xx status code
func (o *AzureKeyVaultDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault delete collection default response has a 3xx status code
func (o *AzureKeyVaultDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault delete collection default response has a 4xx status code
func (o *AzureKeyVaultDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault delete collection default response has a 5xx status code
func (o *AzureKeyVaultDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault delete collection default response a status code equal to that given
func (o *AzureKeyVaultDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the azure key vault delete collection default response
func (o *AzureKeyVaultDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *AzureKeyVaultDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/azure-key-vaults][%d] azure_key_vault_delete_collection default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/azure-key-vaults][%d] azure_key_vault_delete_collection default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AzureKeyVaultDeleteCollectionBody azure key vault delete collection body
swagger:model AzureKeyVaultDeleteCollectionBody
*/
type AzureKeyVaultDeleteCollectionBody struct {

	// azure key vault response inline records
	AzureKeyVaultResponseInlineRecords []*models.AzureKeyVault `json:"records,omitempty"`
}

// Validate validates this azure key vault delete collection body
func (o *AzureKeyVaultDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAzureKeyVaultResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AzureKeyVaultDeleteCollectionBody) validateAzureKeyVaultResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureKeyVaultResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.AzureKeyVaultResponseInlineRecords); i++ {
		if swag.IsZero(o.AzureKeyVaultResponseInlineRecords[i]) { // not required
			continue
		}

		if o.AzureKeyVaultResponseInlineRecords[i] != nil {
			if err := o.AzureKeyVaultResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this azure key vault delete collection body based on the context it is used
func (o *AzureKeyVaultDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAzureKeyVaultResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AzureKeyVaultDeleteCollectionBody) contextValidateAzureKeyVaultResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AzureKeyVaultResponseInlineRecords); i++ {

		if o.AzureKeyVaultResponseInlineRecords[i] != nil {
			if err := o.AzureKeyVaultResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *AzureKeyVaultDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AzureKeyVaultDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
