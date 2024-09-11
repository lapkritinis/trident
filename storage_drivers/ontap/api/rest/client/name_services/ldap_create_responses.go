// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// LdapCreateReader is a Reader for the LdapCreate structure.
type LdapCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LdapCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLdapCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLdapCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLdapCreateCreated creates a LdapCreateCreated with default headers values
func NewLdapCreateCreated() *LdapCreateCreated {
	return &LdapCreateCreated{}
}

/*
LdapCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LdapCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.LdapServiceResponse
}

// IsSuccess returns true when this ldap create created response has a 2xx status code
func (o *LdapCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap create created response has a 3xx status code
func (o *LdapCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap create created response has a 4xx status code
func (o *LdapCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap create created response has a 5xx status code
func (o *LdapCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap create created response a status code equal to that given
func (o *LdapCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ldap create created response
func (o *LdapCreateCreated) Code() int {
	return 201
}

func (o *LdapCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap][%d] ldapCreateCreated %s", 201, payload)
}

func (o *LdapCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap][%d] ldapCreateCreated %s", 201, payload)
}

func (o *LdapCreateCreated) GetPayload() *models.LdapServiceResponse {
	return o.Payload
}

func (o *LdapCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.LdapServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLdapCreateDefault creates a LdapCreateDefault with default headers values
func NewLdapCreateDefault(code int) *LdapCreateDefault {
	return &LdapCreateDefault{
		_statusCode: code,
	}
}

/*
	LdapCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262186     | LDAP Servers cannot be used with Active Directory domain and/or preferred Active Directory servers |
| 2621488    | Invalid SVM context |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 4915203    | The specified LDAP schema does not exist |
| 262222     | The specified LDAP servers or preferred Active Directory servers contain duplicate server entries |
| 4915229    | DNS resolution failed due to an internal error. Contact technical support if this issue persists |
| 4915231    | DNS resolution failed for one or more of the specified LDAP servers. Verify that a valid DNS server is configured |
| 23724132   | DNS resolution failed for all the specified LDAP servers. Verify that a valid DNS server is configured |
| 4915234    | The specified LDAP server or preferred Active Directory server is not supported because it is one of the following: multicast, loopback, 0.0.0.0, or broadcast |
| 4915248    | LDAP servers cannot be empty or "-". Specified Active Directory domain is invalid because it is empty or "-" or it contains either the special characters or "-" at the start or end of the domain)  |
| 4915251    | STARTTLS and LDAPS cannot be used together |
| 4915257    | The LDAP configuration is invalid. Verify that bind-dn and bind password are correct |
| 4915258    | The LDAP configuration is invalid. Verify that the Active Directory domain or servers are reachable and that the network configuration is correct |
| 4915259    | LDAP configurations with Active Directory domains are not supported on admin SVM. |
| 4915265    | The specified bind password or bind DN is invalid |
| 4915264    | Certificate verification failed. Verify that a valid certificate is installed |
| 13434916   | The SVM is in the process of being created. Wait a few minutes, and then try the command again. |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 LIFs |
| 4915252    | LDAP Referral is not supported with STARTTLS, with session security levels sign, seal or with LDAPS. |
| 4915266    | LDAP site discovery restriction cannot be applied to a mixed version cluster. |
| 656477     | Need default site to be specified to enable site restriction. |
| 4915206    | CIFS server is not configured for the vserver. LDAP client configuration requires CIFS server for binding. |
| 4915261    | Cannot use port "389" when "ldaps_enabled" is "true". |
| 4915255    | Base DN specified in the LDAP client configuration is not available. |
| 4915268    | The bind_as_cifs_server field cannot be set to true when the CIFS server is in workgroup mode. |
| 4915269    | The bind_as_cifs_server field cannot be set to true when the CIFS server is in realm mode. |
*/
type LdapCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ldap create default response has a 2xx status code
func (o *LdapCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ldap create default response has a 3xx status code
func (o *LdapCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ldap create default response has a 4xx status code
func (o *LdapCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ldap create default response has a 5xx status code
func (o *LdapCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ldap create default response a status code equal to that given
func (o *LdapCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ldap create default response
func (o *LdapCreateDefault) Code() int {
	return o._statusCode
}

func (o *LdapCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap][%d] ldap_create default %s", o._statusCode, payload)
}

func (o *LdapCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap][%d] ldap_create default %s", o._statusCode, payload)
}

func (o *LdapCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LdapCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
