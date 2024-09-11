// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetroclusterInterconnect Data for a MetroCluster interconnect. REST: /api/cluster/metrocluster/interconnects
//
// swagger:model metrocluster_interconnect
type MetroclusterInterconnect struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Adapter
	// Required: true
	// Read Only: true
	Adapter *string `json:"adapter"`

	// List of objects which contain interface information such as its IP address, netmask and gateway.
	MetroclusterInterconnectInlineInterfaces []*MetroclusterInterconnectInlineInterfacesInlineArrayItem `json:"interfaces,omitempty"`

	// mirror
	Mirror *MetroclusterInterconnectInlineMirror `json:"mirror,omitempty"`

	// Displays the NVRAM mirror multipath policy for the nodes configured in a MetroCluster.
	// Read Only: true
	// Enum: ["no_mp","static_map","dynamic_map","round_robin"]
	MultipathPolicy *string `json:"multipath_policy,omitempty"`

	// node
	Node *MetroclusterInterconnectInlineNode `json:"node,omitempty"`

	// Partner type
	// Required: true
	// Read Only: true
	// Enum: ["aux","dr","ha"]
	PartnerType *string `json:"partner_type"`

	// Adapter status
	// Read Only: true
	// Enum: ["down","up"]
	State *string `json:"state,omitempty"`

	// Adapter type
	// Read Only: true
	// Enum: ["roce","iwarp","unknown"]
	Type *string `json:"type,omitempty"`

	// VLAN ID
	// Read Only: true
	// Maximum: 4095
	// Minimum: 1
	VlanID *int64 `json:"vlan_id,omitempty"`
}

// Validate validates this metrocluster interconnect
func (m *MetroclusterInterconnect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdapter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetroclusterInterconnectInlineInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMirror(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultipathPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartnerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnect) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterInterconnect) validateAdapter(formats strfmt.Registry) error {

	if err := validate.Required("adapter", "body", m.Adapter); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) validateMetroclusterInterconnectInlineInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.MetroclusterInterconnectInlineInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.MetroclusterInterconnectInlineInterfaces); i++ {
		if swag.IsZero(m.MetroclusterInterconnectInlineInterfaces[i]) { // not required
			continue
		}

		if m.MetroclusterInterconnectInlineInterfaces[i] != nil {
			if err := m.MetroclusterInterconnectInlineInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetroclusterInterconnect) validateMirror(formats strfmt.Registry) error {
	if swag.IsZero(m.Mirror) { // not required
		return nil
	}

	if m.Mirror != nil {
		if err := m.Mirror.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mirror")
			}
			return err
		}
	}

	return nil
}

var metroclusterInterconnectTypeMultipathPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_mp","static_map","dynamic_map","round_robin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterInterconnectTypeMultipathPolicyPropEnum = append(metroclusterInterconnectTypeMultipathPolicyPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// multipath_policy
	// MultipathPolicy
	// no_mp
	// END DEBUGGING
	// MetroclusterInterconnectMultipathPolicyNoMp captures enum value "no_mp"
	MetroclusterInterconnectMultipathPolicyNoMp string = "no_mp"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// multipath_policy
	// MultipathPolicy
	// static_map
	// END DEBUGGING
	// MetroclusterInterconnectMultipathPolicyStaticMap captures enum value "static_map"
	MetroclusterInterconnectMultipathPolicyStaticMap string = "static_map"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// multipath_policy
	// MultipathPolicy
	// dynamic_map
	// END DEBUGGING
	// MetroclusterInterconnectMultipathPolicyDynamicMap captures enum value "dynamic_map"
	MetroclusterInterconnectMultipathPolicyDynamicMap string = "dynamic_map"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// multipath_policy
	// MultipathPolicy
	// round_robin
	// END DEBUGGING
	// MetroclusterInterconnectMultipathPolicyRoundRobin captures enum value "round_robin"
	MetroclusterInterconnectMultipathPolicyRoundRobin string = "round_robin"
)

// prop value enum
func (m *MetroclusterInterconnect) validateMultipathPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterInterconnectTypeMultipathPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterInterconnect) validateMultipathPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MultipathPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateMultipathPolicyEnum("multipath_policy", "body", *m.MultipathPolicy); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) validateNode(formats strfmt.Registry) error {
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

var metroclusterInterconnectTypePartnerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aux","dr","ha"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterInterconnectTypePartnerTypePropEnum = append(metroclusterInterconnectTypePartnerTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// partner_type
	// PartnerType
	// aux
	// END DEBUGGING
	// MetroclusterInterconnectPartnerTypeAux captures enum value "aux"
	MetroclusterInterconnectPartnerTypeAux string = "aux"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// partner_type
	// PartnerType
	// dr
	// END DEBUGGING
	// MetroclusterInterconnectPartnerTypeDr captures enum value "dr"
	MetroclusterInterconnectPartnerTypeDr string = "dr"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// partner_type
	// PartnerType
	// ha
	// END DEBUGGING
	// MetroclusterInterconnectPartnerTypeHa captures enum value "ha"
	MetroclusterInterconnectPartnerTypeHa string = "ha"
)

// prop value enum
func (m *MetroclusterInterconnect) validatePartnerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterInterconnectTypePartnerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterInterconnect) validatePartnerType(formats strfmt.Registry) error {

	if err := validate.Required("partner_type", "body", m.PartnerType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePartnerTypeEnum("partner_type", "body", *m.PartnerType); err != nil {
		return err
	}

	return nil
}

var metroclusterInterconnectTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["down","up"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterInterconnectTypeStatePropEnum = append(metroclusterInterconnectTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// state
	// State
	// down
	// END DEBUGGING
	// MetroclusterInterconnectStateDown captures enum value "down"
	MetroclusterInterconnectStateDown string = "down"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// state
	// State
	// up
	// END DEBUGGING
	// MetroclusterInterconnectStateUp captures enum value "up"
	MetroclusterInterconnectStateUp string = "up"
)

// prop value enum
func (m *MetroclusterInterconnect) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterInterconnectTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterInterconnect) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var metroclusterInterconnectTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roce","iwarp","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterInterconnectTypeTypePropEnum = append(metroclusterInterconnectTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// type
	// Type
	// roce
	// END DEBUGGING
	// MetroclusterInterconnectTypeRoce captures enum value "roce"
	MetroclusterInterconnectTypeRoce string = "roce"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// type
	// Type
	// iwarp
	// END DEBUGGING
	// MetroclusterInterconnectTypeIwarp captures enum value "iwarp"
	MetroclusterInterconnectTypeIwarp string = "iwarp"

	// BEGIN DEBUGGING
	// metrocluster_interconnect
	// MetroclusterInterconnect
	// type
	// Type
	// unknown
	// END DEBUGGING
	// MetroclusterInterconnectTypeUnknown captures enum value "unknown"
	MetroclusterInterconnectTypeUnknown string = "unknown"
)

// prop value enum
func (m *MetroclusterInterconnect) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterInterconnectTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterInterconnect) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) validateVlanID(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanID) { // not required
		return nil
	}

	if err := validate.MinimumInt("vlan_id", "body", *m.VlanID, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vlan_id", "body", *m.VlanID, 4095, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster interconnect based on the context it is used
func (m *MetroclusterInterconnect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdapter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetroclusterInterconnectInlineInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMirror(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMultipathPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartnerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnect) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateAdapter(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "adapter", "body", m.Adapter); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateMetroclusterInterconnectInlineInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetroclusterInterconnectInlineInterfaces); i++ {

		if m.MetroclusterInterconnectInlineInterfaces[i] != nil {
			if err := m.MetroclusterInterconnectInlineInterfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateMirror(ctx context.Context, formats strfmt.Registry) error {

	if m.Mirror != nil {
		if err := m.Mirror.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mirror")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateMultipathPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "multipath_policy", "body", m.MultipathPolicy); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetroclusterInterconnect) contextValidatePartnerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "partner_type", "body", m.PartnerType); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterInterconnect) contextValidateVlanID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterInterconnect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterInterconnect) UnmarshalBinary(b []byte) error {
	var res MetroclusterInterconnect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterInterconnectInlineInterfacesInlineArrayItem Object to setup an interface along with its default router.
//
// swagger:model metrocluster_interconnect_inline_interfaces_inline_array_item
type MetroclusterInterconnectInlineInterfacesInlineArrayItem struct {

	// IPv4 or IPv6 address
	// Example: 10.10.10.7
	Address *string `json:"address,omitempty"`

	// The IPv4 or IPv6 address of the default router.
	// Example: 10.1.1.1
	Gateway *string `json:"gateway,omitempty"`

	// netmask
	Netmask *IPNetmask `json:"netmask,omitempty"`
}

// Validate validates this metrocluster interconnect inline interfaces inline array item
func (m *MetroclusterInterconnectInlineInterfacesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineInterfacesInlineArrayItem) validateNetmask(formats strfmt.Registry) error {
	if swag.IsZero(m.Netmask) { // not required
		return nil
	}

	if m.Netmask != nil {
		if err := m.Netmask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netmask")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster interconnect inline interfaces inline array item based on the context it is used
func (m *MetroclusterInterconnectInlineInterfacesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetmask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineInterfacesInlineArrayItem) contextValidateNetmask(ctx context.Context, formats strfmt.Registry) error {

	if m.Netmask != nil {
		if err := m.Netmask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netmask")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterInterconnectInlineInterfacesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterInterconnectInlineInterfacesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res MetroclusterInterconnectInlineInterfacesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterInterconnectInlineMirror metrocluster interconnect inline mirror
//
// swagger:model metrocluster_interconnect_inline_mirror
type MetroclusterInterconnectInlineMirror struct {

	// Specifies the administrative state of the NVRAM mirror between partner nodes.
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies the operational state of the NVRAM mirror between partner nodes.
	// Read Only: true
	// Enum: ["online","offline","unknown"]
	State *string `json:"state,omitempty"`
}

// Validate validates this metrocluster interconnect inline mirror
func (m *MetroclusterInterconnectInlineMirror) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metroclusterInterconnectInlineMirrorTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterInterconnectInlineMirrorTypeStatePropEnum = append(metroclusterInterconnectInlineMirrorTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_interconnect_inline_mirror
	// MetroclusterInterconnectInlineMirror
	// state
	// State
	// online
	// END DEBUGGING
	// MetroclusterInterconnectInlineMirrorStateOnline captures enum value "online"
	MetroclusterInterconnectInlineMirrorStateOnline string = "online"

	// BEGIN DEBUGGING
	// metrocluster_interconnect_inline_mirror
	// MetroclusterInterconnectInlineMirror
	// state
	// State
	// offline
	// END DEBUGGING
	// MetroclusterInterconnectInlineMirrorStateOffline captures enum value "offline"
	MetroclusterInterconnectInlineMirrorStateOffline string = "offline"

	// BEGIN DEBUGGING
	// metrocluster_interconnect_inline_mirror
	// MetroclusterInterconnectInlineMirror
	// state
	// State
	// unknown
	// END DEBUGGING
	// MetroclusterInterconnectInlineMirrorStateUnknown captures enum value "unknown"
	MetroclusterInterconnectInlineMirrorStateUnknown string = "unknown"
)

// prop value enum
func (m *MetroclusterInterconnectInlineMirror) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterInterconnectInlineMirrorTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterInterconnectInlineMirror) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("mirror"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster interconnect inline mirror based on the context it is used
func (m *MetroclusterInterconnectInlineMirror) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineMirror) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mirror"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterInterconnectInlineMirror) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterInterconnectInlineMirror) UnmarshalBinary(b []byte) error {
	var res MetroclusterInterconnectInlineMirror
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterInterconnectInlineNode metrocluster interconnect inline node
//
// swagger:model metrocluster_interconnect_inline_node
type MetroclusterInterconnectInlineNode struct {

	// links
	Links *MetroclusterInterconnectInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this metrocluster interconnect inline node
func (m *MetroclusterInterconnectInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster interconnect inline node based on the context it is used
func (m *MetroclusterInterconnectInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterInterconnectInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterInterconnectInlineNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterInterconnectInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterInterconnectInlineNodeInlineLinks metrocluster interconnect inline node inline links
//
// swagger:model metrocluster_interconnect_inline_node_inline__links
type MetroclusterInterconnectInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster interconnect inline node inline links
func (m *MetroclusterInterconnectInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster interconnect inline node inline links based on the context it is used
func (m *MetroclusterInterconnectInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterInterconnectInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterInterconnectInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterInterconnectInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterInterconnectInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
