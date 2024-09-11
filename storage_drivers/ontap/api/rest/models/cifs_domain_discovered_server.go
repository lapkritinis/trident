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

// CifsDomainDiscoveredServer cifs domain discovered server
//
// swagger:model cifs_domain_discovered_server
type CifsDomainDiscoveredServer struct {

	// Fully Qualified Domain Name.
	//
	// Example: test.com
	Domain *string `json:"domain,omitempty"`

	// node
	Node *CifsDomainDiscoveredServerInlineNode `json:"node,omitempty"`

	// Server Preference
	//
	// Enum: ["unknown","preferred","favored","adequate"]
	Preference *string `json:"preference,omitempty"`

	// Server IP address
	//
	ServerIP *string `json:"server_ip,omitempty"`

	// Server Name
	//
	ServerName *string `json:"server_name,omitempty"`

	// Server Type
	//
	// Enum: ["unknown","kerberos","ms_ldap","ms_dc","ldap"]
	ServerType *string `json:"server_type,omitempty"`

	// Server status
	//
	// Enum: ["ok","unavailable","slow","expired","undetermined","unreachable"]
	State *string `json:"state,omitempty"`
}

// Validate validates this cifs domain discovered server
func (m *CifsDomainDiscoveredServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainDiscoveredServer) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

var cifsDomainDiscoveredServerTypePreferencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","preferred","favored","adequate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsDomainDiscoveredServerTypePreferencePropEnum = append(cifsDomainDiscoveredServerTypePreferencePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// preference
	// Preference
	// unknown
	// END DEBUGGING
	// CifsDomainDiscoveredServerPreferenceUnknown captures enum value "unknown"
	CifsDomainDiscoveredServerPreferenceUnknown string = "unknown"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// preference
	// Preference
	// preferred
	// END DEBUGGING
	// CifsDomainDiscoveredServerPreferencePreferred captures enum value "preferred"
	CifsDomainDiscoveredServerPreferencePreferred string = "preferred"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// preference
	// Preference
	// favored
	// END DEBUGGING
	// CifsDomainDiscoveredServerPreferenceFavored captures enum value "favored"
	CifsDomainDiscoveredServerPreferenceFavored string = "favored"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// preference
	// Preference
	// adequate
	// END DEBUGGING
	// CifsDomainDiscoveredServerPreferenceAdequate captures enum value "adequate"
	CifsDomainDiscoveredServerPreferenceAdequate string = "adequate"
)

// prop value enum
func (m *CifsDomainDiscoveredServer) validatePreferenceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsDomainDiscoveredServerTypePreferencePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsDomainDiscoveredServer) validatePreference(formats strfmt.Registry) error {
	if swag.IsZero(m.Preference) { // not required
		return nil
	}

	// value enum
	if err := m.validatePreferenceEnum("preference", "body", *m.Preference); err != nil {
		return err
	}

	return nil
}

var cifsDomainDiscoveredServerTypeServerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","kerberos","ms_ldap","ms_dc","ldap"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsDomainDiscoveredServerTypeServerTypePropEnum = append(cifsDomainDiscoveredServerTypeServerTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// server_type
	// ServerType
	// unknown
	// END DEBUGGING
	// CifsDomainDiscoveredServerServerTypeUnknown captures enum value "unknown"
	CifsDomainDiscoveredServerServerTypeUnknown string = "unknown"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// server_type
	// ServerType
	// kerberos
	// END DEBUGGING
	// CifsDomainDiscoveredServerServerTypeKerberos captures enum value "kerberos"
	CifsDomainDiscoveredServerServerTypeKerberos string = "kerberos"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// server_type
	// ServerType
	// ms_ldap
	// END DEBUGGING
	// CifsDomainDiscoveredServerServerTypeMsLdap captures enum value "ms_ldap"
	CifsDomainDiscoveredServerServerTypeMsLdap string = "ms_ldap"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// server_type
	// ServerType
	// ms_dc
	// END DEBUGGING
	// CifsDomainDiscoveredServerServerTypeMsDc captures enum value "ms_dc"
	CifsDomainDiscoveredServerServerTypeMsDc string = "ms_dc"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// server_type
	// ServerType
	// ldap
	// END DEBUGGING
	// CifsDomainDiscoveredServerServerTypeLdap captures enum value "ldap"
	CifsDomainDiscoveredServerServerTypeLdap string = "ldap"
)

// prop value enum
func (m *CifsDomainDiscoveredServer) validateServerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsDomainDiscoveredServerTypeServerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsDomainDiscoveredServer) validateServerType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerTypeEnum("server_type", "body", *m.ServerType); err != nil {
		return err
	}

	return nil
}

var cifsDomainDiscoveredServerTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","unavailable","slow","expired","undetermined","unreachable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsDomainDiscoveredServerTypeStatePropEnum = append(cifsDomainDiscoveredServerTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// state
	// State
	// ok
	// END DEBUGGING
	// CifsDomainDiscoveredServerStateOk captures enum value "ok"
	CifsDomainDiscoveredServerStateOk string = "ok"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// state
	// State
	// unavailable
	// END DEBUGGING
	// CifsDomainDiscoveredServerStateUnavailable captures enum value "unavailable"
	CifsDomainDiscoveredServerStateUnavailable string = "unavailable"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// state
	// State
	// slow
	// END DEBUGGING
	// CifsDomainDiscoveredServerStateSlow captures enum value "slow"
	CifsDomainDiscoveredServerStateSlow string = "slow"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// state
	// State
	// expired
	// END DEBUGGING
	// CifsDomainDiscoveredServerStateExpired captures enum value "expired"
	CifsDomainDiscoveredServerStateExpired string = "expired"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// state
	// State
	// undetermined
	// END DEBUGGING
	// CifsDomainDiscoveredServerStateUndetermined captures enum value "undetermined"
	CifsDomainDiscoveredServerStateUndetermined string = "undetermined"

	// BEGIN DEBUGGING
	// cifs_domain_discovered_server
	// CifsDomainDiscoveredServer
	// state
	// State
	// unreachable
	// END DEBUGGING
	// CifsDomainDiscoveredServerStateUnreachable captures enum value "unreachable"
	CifsDomainDiscoveredServerStateUnreachable string = "unreachable"
)

// prop value enum
func (m *CifsDomainDiscoveredServer) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsDomainDiscoveredServerTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsDomainDiscoveredServer) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cifs domain discovered server based on the context it is used
func (m *CifsDomainDiscoveredServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainDiscoveredServer) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainDiscoveredServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainDiscoveredServer) UnmarshalBinary(b []byte) error {
	var res CifsDomainDiscoveredServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsDomainDiscoveredServerInlineNode cifs domain discovered server inline node
//
// swagger:model cifs_domain_discovered_server_inline_node
type CifsDomainDiscoveredServerInlineNode struct {

	// links
	Links *CifsDomainDiscoveredServerInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cifs domain discovered server inline node
func (m *CifsDomainDiscoveredServerInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainDiscoveredServerInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs domain discovered server inline node based on the context it is used
func (m *CifsDomainDiscoveredServerInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainDiscoveredServerInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainDiscoveredServerInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainDiscoveredServerInlineNode) UnmarshalBinary(b []byte) error {
	var res CifsDomainDiscoveredServerInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsDomainDiscoveredServerInlineNodeInlineLinks cifs domain discovered server inline node inline links
//
// swagger:model cifs_domain_discovered_server_inline_node_inline__links
type CifsDomainDiscoveredServerInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs domain discovered server inline node inline links
func (m *CifsDomainDiscoveredServerInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainDiscoveredServerInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs domain discovered server inline node inline links based on the context it is used
func (m *CifsDomainDiscoveredServerInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainDiscoveredServerInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainDiscoveredServerInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainDiscoveredServerInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsDomainDiscoveredServerInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
