// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BroadcastDomain Set of ports that will receive a broadcast Ethernet packet from any of them
//
// swagger:model broadcast_domain
type BroadcastDomain struct {

	// links
	Links *BroadcastDomainInlineLinks `json:"_links,omitempty"`

	// Ports that belong to the broadcast domain
	// Read Only: true
	BroadcastDomainInlinePorts []*BroadcastDomainInlinePortsInlineArrayItem `json:"ports,omitempty"`

	// ipspace
	Ipspace *BroadcastDomainInlineIpspace `json:"ipspace,omitempty"`

	// Maximum transmission unit, largest packet size on this network
	// Example: 1500
	// Minimum: 68
	Mtu *int64 `json:"mtu,omitempty"`

	// Name of the broadcast domain, scoped to its IPspace
	// Example: bd1
	Name *string `json:"name,omitempty"`

	// Broadcast domain UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this broadcast domain
func (m *BroadcastDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBroadcastDomainInlinePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpspace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomain) validateLinks(formats strfmt.Registry) error {
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

func (m *BroadcastDomain) validateBroadcastDomainInlinePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.BroadcastDomainInlinePorts) { // not required
		return nil
	}

	for i := 0; i < len(m.BroadcastDomainInlinePorts); i++ {
		if swag.IsZero(m.BroadcastDomainInlinePorts[i]) { // not required
			continue
		}

		if m.BroadcastDomainInlinePorts[i] != nil {
			if err := m.BroadcastDomainInlinePorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BroadcastDomain) validateIpspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipspace) { // not required
		return nil
	}

	if m.Ipspace != nil {
		if err := m.Ipspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace")
			}
			return err
		}
	}

	return nil
}

func (m *BroadcastDomain) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", *m.Mtu, 68, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this broadcast domain based on the context it is used
func (m *BroadcastDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBroadcastDomainInlinePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomain) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BroadcastDomain) contextValidateBroadcastDomainInlinePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ports", "body", []*BroadcastDomainInlinePortsInlineArrayItem(m.BroadcastDomainInlinePorts)); err != nil {
		return err
	}

	for i := 0; i < len(m.BroadcastDomainInlinePorts); i++ {

		if m.BroadcastDomainInlinePorts[i] != nil {
			if err := m.BroadcastDomainInlinePorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BroadcastDomain) contextValidateIpspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipspace != nil {
		if err := m.Ipspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace")
			}
			return err
		}
	}

	return nil
}

func (m *BroadcastDomain) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomain) UnmarshalBinary(b []byte) error {
	var res BroadcastDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainInlineIpspace Applies to both SVM and cluster-scoped objects. Either the UUID or name is supplied on input.
//
// swagger:model broadcast_domain_inline_ipspace
type BroadcastDomainInlineIpspace struct {

	// links
	Links *BroadcastDomainInlineIpspaceInlineLinks `json:"_links,omitempty"`

	// IPspace name
	// Example: Default
	Name *string `json:"name,omitempty"`

	// IPspace UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this broadcast domain inline ipspace
func (m *BroadcastDomainInlineIpspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlineIpspace) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ipspace based on the context it is used
func (m *BroadcastDomainInlineIpspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlineIpspace) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainInlineIpspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainInlineIpspace) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlineIpspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainInlineIpspaceInlineLinks broadcast domain inline ipspace inline links
//
// swagger:model broadcast_domain_inline_ipspace_inline__links
type BroadcastDomainInlineIpspaceInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain inline ipspace inline links
func (m *BroadcastDomainInlineIpspaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlineIpspaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ipspace inline links based on the context it is used
func (m *BroadcastDomainInlineIpspaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlineIpspaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainInlineIpspaceInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainInlineIpspaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlineIpspaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainInlineLinks broadcast domain inline links
//
// swagger:model broadcast_domain_inline__links
type BroadcastDomainInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain inline links
func (m *BroadcastDomainInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline links based on the context it is used
func (m *BroadcastDomainInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainInlineLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainInlinePortsInlineArrayItem Port UUID along with readable names
//
// swagger:model broadcast_domain_inline_ports_inline_array_item
type BroadcastDomainInlinePortsInlineArrayItem struct {

	// links
	Links *BroadcastDomainInlinePortsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// name
	// Example: e1b
	Name *string `json:"name,omitempty"`

	// node
	Node *BroadcastDomainInlinePortsInlineArrayItemInlineNode `json:"node,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this broadcast domain inline ports inline array item
func (m *BroadcastDomainInlinePortsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlinePortsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *BroadcastDomainInlinePortsInlineArrayItem) validateNode(formats strfmt.Registry) error {
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

// ContextValidate validate this broadcast domain inline ports inline array item based on the context it is used
func (m *BroadcastDomainInlinePortsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlinePortsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BroadcastDomainInlinePortsInlineArrayItem) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BroadcastDomainInlinePortsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainInlinePortsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlinePortsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainInlinePortsInlineArrayItemInlineLinks broadcast domain inline ports inline array item inline links
//
// swagger:model broadcast_domain_inline_ports_inline_array_item_inline__links
type BroadcastDomainInlinePortsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain inline ports inline array item inline links
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ports inline array item inline links based on the context it is used
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlinePortsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainInlinePortsInlineArrayItemInlineNode broadcast domain inline ports inline array item inline node
//
// swagger:model broadcast_domain_inline_ports_inline_array_item_inline_node
type BroadcastDomainInlinePortsInlineArrayItemInlineNode struct {

	// Name of node on which the port is located.
	// Example: node1
	Name *string `json:"name,omitempty"`
}

// Validate validates this broadcast domain inline ports inline array item inline node
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broadcast domain inline ports inline array item inline node based on context it is used
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainInlinePortsInlineArrayItemInlineNode) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlinePortsInlineArrayItemInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
