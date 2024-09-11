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

// StoragePort storage port
//
// swagger:model storage_port
type StoragePort struct {

	// board name
	// Read Only: true
	BoardName *string `json:"board_name,omitempty"`

	// cable
	Cable *StoragePortInlineCable `json:"cable,omitempty"`

	// description
	// Example: SAS Host Adapter 2a (PMC-Sierra PM8072 rev. C)
	// Read Only: true
	Description *string `json:"description,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// error
	Error *StoragePortInlineError `json:"error,omitempty"`

	// firmware version
	// Example: 03.08.09.00
	// Read Only: true
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	// force
	Force *bool `json:"force,omitempty"`

	// Specifies whether any devices are connected through this port
	// Read Only: true
	InUse *bool `json:"in_use,omitempty"`

	// mac address
	// Read Only: true
	MacAddress *string `json:"mac_address,omitempty"`

	// Operational mode of a non-dedicated Ethernet port
	// Example: storage
	// Enum: ["network","storage"]
	Mode *string `json:"mode,omitempty"`

	// name
	// Example: 2a
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// node
	Node *StoragePortInlineNode `json:"node,omitempty"`

	// part number
	// Example: 111-03801
	// Read Only: true
	PartNumber *string `json:"part_number,omitempty"`

	// Specifies whether all devices connected through this port have a redundant path from another port
	// Read Only: true
	Redundant *bool `json:"redundant,omitempty"`

	// serial number
	// Example: 7A2463CC45B
	// Read Only: true
	SerialNumber *string `json:"serial_number,omitempty"`

	// Operational port speed in Gbps
	// Example: 6
	// Read Only: true
	Speed *float64 `json:"speed,omitempty"`

	// state
	// Example: online
	// Read Only: true
	// Enum: ["online","offline","error"]
	State *string `json:"state,omitempty"`

	// type
	// Example: sas
	// Read Only: true
	// Enum: ["sas","fc","enet"]
	Type *string `json:"type,omitempty"`

	// World Wide Name
	// Example: 50000d1703544b80
	// Read Only: true
	Wwn *string `json:"wwn,omitempty"`

	// World Wide Port Name
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this storage port
func (m *StoragePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePort) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

var storagePortTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["network","storage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagePortTypeModePropEnum = append(storagePortTypeModePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// mode
	// Mode
	// network
	// END DEBUGGING
	// StoragePortModeNetwork captures enum value "network"
	StoragePortModeNetwork string = "network"

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// mode
	// Mode
	// storage
	// END DEBUGGING
	// StoragePortModeStorage captures enum value "storage"
	StoragePortModeStorage string = "storage"
)

// prop value enum
func (m *StoragePort) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storagePortTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StoragePort) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) validateNode(formats strfmt.Registry) error {
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

var storagePortTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagePortTypeStatePropEnum = append(storagePortTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// state
	// State
	// online
	// END DEBUGGING
	// StoragePortStateOnline captures enum value "online"
	StoragePortStateOnline string = "online"

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// state
	// State
	// offline
	// END DEBUGGING
	// StoragePortStateOffline captures enum value "offline"
	StoragePortStateOffline string = "offline"

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// state
	// State
	// error
	// END DEBUGGING
	// StoragePortStateError captures enum value "error"
	StoragePortStateError string = "error"
)

// prop value enum
func (m *StoragePort) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storagePortTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StoragePort) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var storagePortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sas","fc","enet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagePortTypeTypePropEnum = append(storagePortTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// type
	// Type
	// sas
	// END DEBUGGING
	// StoragePortTypeSas captures enum value "sas"
	StoragePortTypeSas string = "sas"

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// type
	// Type
	// fc
	// END DEBUGGING
	// StoragePortTypeFc captures enum value "fc"
	StoragePortTypeFc string = "fc"

	// BEGIN DEBUGGING
	// storage_port
	// StoragePort
	// type
	// Type
	// enet
	// END DEBUGGING
	// StoragePortTypeEnet captures enum value "enet"
	StoragePortTypeEnet string = "enet"
)

// prop value enum
func (m *StoragePort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storagePortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StoragePort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this storage port based on the context it is used
func (m *StoragePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoardName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmwareVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInUse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMacAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedundant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpeed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePort) contextValidateBoardName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "board_name", "body", m.BoardName); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePort) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePort) contextValidateFirmwareVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "firmware_version", "body", m.FirmwareVersion); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateInUse(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "in_use", "body", m.InUse); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateMacAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mac_address", "body", m.MacAddress); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *StoragePort) contextValidatePartNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "part_number", "body", m.PartNumber); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateRedundant(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "redundant", "body", m.Redundant); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "serial_number", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateSpeed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "speed", "body", m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateWwn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwn", "body", m.Wwn); err != nil {
		return err
	}

	return nil
}

func (m *StoragePort) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePort) UnmarshalBinary(b []byte) error {
	var res StoragePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePortInlineCable storage port inline cable
//
// swagger:model storage_port_inline_cable
type StoragePortInlineCable struct {

	// identifier
	// Example: 500a0980000b6c3f-50000d1703544b80
	Identifier *string `json:"identifier,omitempty"`

	// length
	// Example: 2m
	Length *string `json:"length,omitempty"`

	// part number
	// Example: 112-00431+A0
	PartNumber *string `json:"part_number,omitempty"`

	// serial number
	// Example: 616930439
	SerialNumber *string `json:"serial_number,omitempty"`

	// transceiver
	// Example: mini_sas_hd
	// Enum: ["qsfp","qsfp_plus","qsfp28","mini_sas_hd","sfp"]
	Transceiver *string `json:"transceiver,omitempty"`

	// vendor
	// Example: Molex Inc.
	Vendor *string `json:"vendor,omitempty"`
}

// Validate validates this storage port inline cable
func (m *StoragePortInlineCable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransceiver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storagePortInlineCableTypeTransceiverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["qsfp","qsfp_plus","qsfp28","mini_sas_hd","sfp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagePortInlineCableTypeTransceiverPropEnum = append(storagePortInlineCableTypeTransceiverPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// storage_port_inline_cable
	// StoragePortInlineCable
	// transceiver
	// Transceiver
	// qsfp
	// END DEBUGGING
	// StoragePortInlineCableTransceiverQsfp captures enum value "qsfp"
	StoragePortInlineCableTransceiverQsfp string = "qsfp"

	// BEGIN DEBUGGING
	// storage_port_inline_cable
	// StoragePortInlineCable
	// transceiver
	// Transceiver
	// qsfp_plus
	// END DEBUGGING
	// StoragePortInlineCableTransceiverQsfpPlus captures enum value "qsfp_plus"
	StoragePortInlineCableTransceiverQsfpPlus string = "qsfp_plus"

	// BEGIN DEBUGGING
	// storage_port_inline_cable
	// StoragePortInlineCable
	// transceiver
	// Transceiver
	// qsfp28
	// END DEBUGGING
	// StoragePortInlineCableTransceiverQsfp28 captures enum value "qsfp28"
	StoragePortInlineCableTransceiverQsfp28 string = "qsfp28"

	// BEGIN DEBUGGING
	// storage_port_inline_cable
	// StoragePortInlineCable
	// transceiver
	// Transceiver
	// mini_sas_hd
	// END DEBUGGING
	// StoragePortInlineCableTransceiverMiniSasHd captures enum value "mini_sas_hd"
	StoragePortInlineCableTransceiverMiniSasHd string = "mini_sas_hd"

	// BEGIN DEBUGGING
	// storage_port_inline_cable
	// StoragePortInlineCable
	// transceiver
	// Transceiver
	// sfp
	// END DEBUGGING
	// StoragePortInlineCableTransceiverSfp captures enum value "sfp"
	StoragePortInlineCableTransceiverSfp string = "sfp"
)

// prop value enum
func (m *StoragePortInlineCable) validateTransceiverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storagePortInlineCableTypeTransceiverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StoragePortInlineCable) validateTransceiver(formats strfmt.Registry) error {
	if swag.IsZero(m.Transceiver) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransceiverEnum("cable"+"."+"transceiver", "body", *m.Transceiver); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this storage port inline cable based on the context it is used
func (m *StoragePortInlineCable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StoragePortInlineCable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePortInlineCable) UnmarshalBinary(b []byte) error {
	var res StoragePortInlineCable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePortInlineError storage port inline error
//
// swagger:model storage_port_inline_error
type StoragePortInlineError struct {

	// Error corrective action
	CorrectiveAction *string `json:"corrective_action,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`
}

// Validate validates this storage port inline error
func (m *StoragePortInlineError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this storage port inline error based on the context it is used
func (m *StoragePortInlineError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StoragePortInlineError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePortInlineError) UnmarshalBinary(b []byte) error {
	var res StoragePortInlineError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePortInlineNode storage port inline node
//
// swagger:model storage_port_inline_node
type StoragePortInlineNode struct {

	// links
	Links *StoragePortInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this storage port inline node
func (m *StoragePortInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePortInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this storage port inline node based on the context it is used
func (m *StoragePortInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePortInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *StoragePortInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePortInlineNode) UnmarshalBinary(b []byte) error {
	var res StoragePortInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePortInlineNodeInlineLinks storage port inline node inline links
//
// swagger:model storage_port_inline_node_inline__links
type StoragePortInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this storage port inline node inline links
func (m *StoragePortInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePortInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this storage port inline node inline links based on the context it is used
func (m *StoragePortInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePortInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *StoragePortInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePortInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res StoragePortInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
