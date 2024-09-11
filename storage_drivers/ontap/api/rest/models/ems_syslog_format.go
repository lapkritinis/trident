// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsSyslogFormat ems syslog format
//
// swagger:model ems_syslog_format
type EmsSyslogFormat struct {

	// Syslog Hostname Format Override. The supported hostname formats are no_override (hostname format based on the syslog.format.message property i.e. fqdn if syslog.format.message is rfc_5424, hostname_only if syslog.format.message is legacy_netapp), fqdn (Fully Qualified Domain Name) and hostname_only.
	//
	// Enum: ["no_override","fqdn","hostname_only"]
	HostnameOverride *string `json:"hostname_override,omitempty"`

	// Syslog Message Format. The supported message formats are legacy_netapp (format: &lt;PRIVAL&gt;TIMESTAMP [HOSTNAME:Event-name:Event-severity]: MSG) and rfc_5424 (format: &lt;PRIVAL&gt;VERSION TIMESTAMP HOSTNAME Event-source - Event-name - MSG).
	//
	// Enum: ["legacy_netapp","rfc_5424"]
	Message *string `json:"message,omitempty"`

	// Syslog Timestamp Format Override. The supported timestamp formats are no_override (timestamp format based on the syslog.format.message property i.e. rfc_3164 if syslog.format.message is legacy_netapp, iso_8601_local_time if syslog.format.message is rfc_5424), rfc_3164 (format: Mmm dd hh:mm:ss), iso_8601_local_time (format: YYYY-MM-DDThh:mm:ss+/-hh:mm) and iso_8601_utc (format: YYYY-MM-DDThh:mm:ssZ).
	//
	// Enum: ["no_override","rfc_3164","iso_8601_local_time","iso_8601_utc"]
	TimestampOverride *string `json:"timestamp_override,omitempty"`
}

// Validate validates this ems syslog format
func (m *EmsSyslogFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostnameOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestampOverride(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var emsSyslogFormatTypeHostnameOverridePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_override","fqdn","hostname_only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsSyslogFormatTypeHostnameOverridePropEnum = append(emsSyslogFormatTypeHostnameOverridePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// hostname_override
	// HostnameOverride
	// no_override
	// END DEBUGGING
	// EmsSyslogFormatHostnameOverrideNoOverride captures enum value "no_override"
	EmsSyslogFormatHostnameOverrideNoOverride string = "no_override"

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// hostname_override
	// HostnameOverride
	// fqdn
	// END DEBUGGING
	// EmsSyslogFormatHostnameOverrideFqdn captures enum value "fqdn"
	EmsSyslogFormatHostnameOverrideFqdn string = "fqdn"

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// hostname_override
	// HostnameOverride
	// hostname_only
	// END DEBUGGING
	// EmsSyslogFormatHostnameOverrideHostnameOnly captures enum value "hostname_only"
	EmsSyslogFormatHostnameOverrideHostnameOnly string = "hostname_only"
)

// prop value enum
func (m *EmsSyslogFormat) validateHostnameOverrideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsSyslogFormatTypeHostnameOverridePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsSyslogFormat) validateHostnameOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.HostnameOverride) { // not required
		return nil
	}

	// value enum
	if err := m.validateHostnameOverrideEnum("hostname_override", "body", *m.HostnameOverride); err != nil {
		return err
	}

	return nil
}

var emsSyslogFormatTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["legacy_netapp","rfc_5424"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsSyslogFormatTypeMessagePropEnum = append(emsSyslogFormatTypeMessagePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// message
	// Message
	// legacy_netapp
	// END DEBUGGING
	// EmsSyslogFormatMessageLegacyNetapp captures enum value "legacy_netapp"
	EmsSyslogFormatMessageLegacyNetapp string = "legacy_netapp"

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// message
	// Message
	// rfc_5424
	// END DEBUGGING
	// EmsSyslogFormatMessageRfc5424 captures enum value "rfc_5424"
	EmsSyslogFormatMessageRfc5424 string = "rfc_5424"
)

// prop value enum
func (m *EmsSyslogFormat) validateMessageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsSyslogFormatTypeMessagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsSyslogFormat) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageEnum("message", "body", *m.Message); err != nil {
		return err
	}

	return nil
}

var emsSyslogFormatTypeTimestampOverridePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_override","rfc_3164","iso_8601_local_time","iso_8601_utc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsSyslogFormatTypeTimestampOverridePropEnum = append(emsSyslogFormatTypeTimestampOverridePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// timestamp_override
	// TimestampOverride
	// no_override
	// END DEBUGGING
	// EmsSyslogFormatTimestampOverrideNoOverride captures enum value "no_override"
	EmsSyslogFormatTimestampOverrideNoOverride string = "no_override"

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// timestamp_override
	// TimestampOverride
	// rfc_3164
	// END DEBUGGING
	// EmsSyslogFormatTimestampOverrideRfc3164 captures enum value "rfc_3164"
	EmsSyslogFormatTimestampOverrideRfc3164 string = "rfc_3164"

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// timestamp_override
	// TimestampOverride
	// iso_8601_local_time
	// END DEBUGGING
	// EmsSyslogFormatTimestampOverrideIso8601LocalTime captures enum value "iso_8601_local_time"
	EmsSyslogFormatTimestampOverrideIso8601LocalTime string = "iso_8601_local_time"

	// BEGIN DEBUGGING
	// ems_syslog_format
	// EmsSyslogFormat
	// timestamp_override
	// TimestampOverride
	// iso_8601_utc
	// END DEBUGGING
	// EmsSyslogFormatTimestampOverrideIso8601Utc captures enum value "iso_8601_utc"
	EmsSyslogFormatTimestampOverrideIso8601Utc string = "iso_8601_utc"
)

// prop value enum
func (m *EmsSyslogFormat) validateTimestampOverrideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsSyslogFormatTypeTimestampOverridePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsSyslogFormat) validateTimestampOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.TimestampOverride) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimestampOverrideEnum("timestamp_override", "body", *m.TimestampOverride); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ems syslog format based on context it is used
func (m *EmsSyslogFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmsSyslogFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsSyslogFormat) UnmarshalBinary(b []byte) error {
	var res EmsSyslogFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
