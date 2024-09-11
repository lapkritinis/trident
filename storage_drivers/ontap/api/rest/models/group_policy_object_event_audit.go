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

// GroupPolicyObjectEventAudit group policy object event audit
//
// swagger:model group_policy_object_event_audit
type GroupPolicyObjectEventAudit struct {

	// Type of logon event to be audited.
	// Example: failure
	// Enum: ["none","success","failure","both"]
	LogonType *string `json:"logon_type,omitempty"`

	// Type of object access to be audited.
	// Example: failure
	// Enum: ["none","success","failure","both"]
	ObjectAccessType *string `json:"object_access_type,omitempty"`
}

// Validate validates this group policy object event audit
func (m *GroupPolicyObjectEventAudit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogonType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectAccessType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var groupPolicyObjectEventAuditTypeLogonTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","success","failure","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectEventAuditTypeLogonTypePropEnum = append(groupPolicyObjectEventAuditTypeLogonTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// logon_type
	// LogonType
	// none
	// END DEBUGGING
	// GroupPolicyObjectEventAuditLogonTypeNone captures enum value "none"
	GroupPolicyObjectEventAuditLogonTypeNone string = "none"

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// logon_type
	// LogonType
	// success
	// END DEBUGGING
	// GroupPolicyObjectEventAuditLogonTypeSuccess captures enum value "success"
	GroupPolicyObjectEventAuditLogonTypeSuccess string = "success"

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// logon_type
	// LogonType
	// failure
	// END DEBUGGING
	// GroupPolicyObjectEventAuditLogonTypeFailure captures enum value "failure"
	GroupPolicyObjectEventAuditLogonTypeFailure string = "failure"

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// logon_type
	// LogonType
	// both
	// END DEBUGGING
	// GroupPolicyObjectEventAuditLogonTypeBoth captures enum value "both"
	GroupPolicyObjectEventAuditLogonTypeBoth string = "both"
)

// prop value enum
func (m *GroupPolicyObjectEventAudit) validateLogonTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectEventAuditTypeLogonTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObjectEventAudit) validateLogonType(formats strfmt.Registry) error {
	if swag.IsZero(m.LogonType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogonTypeEnum("logon_type", "body", *m.LogonType); err != nil {
		return err
	}

	return nil
}

var groupPolicyObjectEventAuditTypeObjectAccessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","success","failure","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectEventAuditTypeObjectAccessTypePropEnum = append(groupPolicyObjectEventAuditTypeObjectAccessTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// object_access_type
	// ObjectAccessType
	// none
	// END DEBUGGING
	// GroupPolicyObjectEventAuditObjectAccessTypeNone captures enum value "none"
	GroupPolicyObjectEventAuditObjectAccessTypeNone string = "none"

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// object_access_type
	// ObjectAccessType
	// success
	// END DEBUGGING
	// GroupPolicyObjectEventAuditObjectAccessTypeSuccess captures enum value "success"
	GroupPolicyObjectEventAuditObjectAccessTypeSuccess string = "success"

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// object_access_type
	// ObjectAccessType
	// failure
	// END DEBUGGING
	// GroupPolicyObjectEventAuditObjectAccessTypeFailure captures enum value "failure"
	GroupPolicyObjectEventAuditObjectAccessTypeFailure string = "failure"

	// BEGIN DEBUGGING
	// group_policy_object_event_audit
	// GroupPolicyObjectEventAudit
	// object_access_type
	// ObjectAccessType
	// both
	// END DEBUGGING
	// GroupPolicyObjectEventAuditObjectAccessTypeBoth captures enum value "both"
	GroupPolicyObjectEventAuditObjectAccessTypeBoth string = "both"
)

// prop value enum
func (m *GroupPolicyObjectEventAudit) validateObjectAccessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectEventAuditTypeObjectAccessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObjectEventAudit) validateObjectAccessType(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectAccessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectAccessTypeEnum("object_access_type", "body", *m.ObjectAccessType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this group policy object event audit based on context it is used
func (m *GroupPolicyObjectEventAudit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupPolicyObjectEventAudit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectEventAudit) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectEventAudit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
