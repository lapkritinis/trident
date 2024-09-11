// Code generated by go-swagger; DO NOT EDIT.

package support

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

// ConfigurationBackupFileGetReader is a Reader for the ConfigurationBackupFileGet structure.
type ConfigurationBackupFileGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurationBackupFileGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigurationBackupFileGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigurationBackupFileGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurationBackupFileGetOK creates a ConfigurationBackupFileGetOK with default headers values
func NewConfigurationBackupFileGetOK() *ConfigurationBackupFileGetOK {
	return &ConfigurationBackupFileGetOK{}
}

/*
ConfigurationBackupFileGetOK describes a response with status code 200, with default header values.

OK
*/
type ConfigurationBackupFileGetOK struct {
	Payload *models.ConfigurationBackupFile
}

// IsSuccess returns true when this configuration backup file get o k response has a 2xx status code
func (o *ConfigurationBackupFileGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configuration backup file get o k response has a 3xx status code
func (o *ConfigurationBackupFileGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configuration backup file get o k response has a 4xx status code
func (o *ConfigurationBackupFileGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this configuration backup file get o k response has a 5xx status code
func (o *ConfigurationBackupFileGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this configuration backup file get o k response a status code equal to that given
func (o *ConfigurationBackupFileGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the configuration backup file get o k response
func (o *ConfigurationBackupFileGetOK) Code() int {
	return 200
}

func (o *ConfigurationBackupFileGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups/{node.uuid}/{name}][%d] configurationBackupFileGetOK %s", 200, payload)
}

func (o *ConfigurationBackupFileGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups/{node.uuid}/{name}][%d] configurationBackupFileGetOK %s", 200, payload)
}

func (o *ConfigurationBackupFileGetOK) GetPayload() *models.ConfigurationBackupFile {
	return o.Payload
}

func (o *ConfigurationBackupFileGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationBackupFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurationBackupFileGetDefault creates a ConfigurationBackupFileGetDefault with default headers values
func NewConfigurationBackupFileGetDefault(code int) *ConfigurationBackupFileGetDefault {
	return &ConfigurationBackupFileGetDefault{
		_statusCode: code,
	}
}

/*
	ConfigurationBackupFileGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5963777 | Configuration backup file does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ConfigurationBackupFileGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this configuration backup file get default response has a 2xx status code
func (o *ConfigurationBackupFileGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this configuration backup file get default response has a 3xx status code
func (o *ConfigurationBackupFileGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this configuration backup file get default response has a 4xx status code
func (o *ConfigurationBackupFileGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this configuration backup file get default response has a 5xx status code
func (o *ConfigurationBackupFileGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this configuration backup file get default response a status code equal to that given
func (o *ConfigurationBackupFileGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the configuration backup file get default response
func (o *ConfigurationBackupFileGetDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurationBackupFileGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups/{node.uuid}/{name}][%d] configuration_backup_file_get default %s", o._statusCode, payload)
}

func (o *ConfigurationBackupFileGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups/{node.uuid}/{name}][%d] configuration_backup_file_get default %s", o._statusCode, payload)
}

func (o *ConfigurationBackupFileGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurationBackupFileGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
