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

// IPInterfaceSvm Interface parameters.  Name and home_node are optional.
//
// swagger:model ip_interface_svm
type IPInterfaceSvm struct {

	// links
	Links *IPInterfaceSvmInlineLinks `json:"_links,omitempty"`

	// ip
	IP *IPInterfaceSvmInlineIP `json:"ip,omitempty"`

	// The services associated with the interface.
	// Read Only: true
	IPInterfaceSvmInlineServices []*IPService `json:"services,omitempty"`

	// location
	Location *IPInterfaceSvmInlineLocation `json:"location,omitempty"`

	// The name of the interface (optional).
	// Example: lif1
	Name *string `json:"name,omitempty"`

	// service policy
	ServicePolicy *IPServicePolicySvmEnum `json:"service_policy,omitempty"`

	// Allocates an interface address from a subnet.
	Subnet *IPSubnetReference `json:"subnet,omitempty"`

	// The UUID that uniquely identifies the interface.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ip interface svm
func (m *IPInterfaceSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPInterfaceSvmInlineServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvm) validateLinks(formats strfmt.Registry) error {
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

func (m *IPInterfaceSvm) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) validateIPInterfaceSvmInlineServices(formats strfmt.Registry) error {
	if swag.IsZero(m.IPInterfaceSvmInlineServices) { // not required
		return nil
	}

	for i := 0; i < len(m.IPInterfaceSvmInlineServices); i++ {
		if swag.IsZero(m.IPInterfaceSvmInlineServices[i]) { // not required
			continue
		}

		if m.IPInterfaceSvmInlineServices[i] != nil {
			if err := m.IPInterfaceSvmInlineServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IPInterfaceSvm) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) validateServicePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ServicePolicy) { // not required
		return nil
	}

	if m.ServicePolicy != nil {
		if err := m.ServicePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_policy")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm based on the context it is used
func (m *IPInterfaceSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPInterfaceSvmInlineServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServicePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnet(ctx, formats); err != nil {
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

func (m *IPInterfaceSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IPInterfaceSvm) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) contextValidateIPInterfaceSvmInlineServices(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "services", "body", []*IPService(m.IPInterfaceSvmInlineServices)); err != nil {
		return err
	}

	for i := 0; i < len(m.IPInterfaceSvmInlineServices); i++ {

		if m.IPInterfaceSvmInlineServices[i] != nil {
			if err := m.IPInterfaceSvmInlineServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IPInterfaceSvm) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) contextValidateServicePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ServicePolicy != nil {
		if err := m.ServicePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_policy")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) contextValidateSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.Subnet != nil {
		if err := m.Subnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvm) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineIP IP information
//
// swagger:model ip_interface_svm_inline_ip
type IPInterfaceSvmInlineIP struct {

	// address
	Address *IPAddressReadcreate `json:"address,omitempty"`

	// netmask
	Netmask *IPNetmaskCreateonly `json:"netmask,omitempty"`
}

// Validate validates this ip interface svm inline ip
func (m *IPInterfaceSvmInlineIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineIP) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvmInlineIP) validateNetmask(formats strfmt.Registry) error {
	if swag.IsZero(m.Netmask) { // not required
		return nil
	}

	if m.Netmask != nil {
		if err := m.Netmask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "netmask")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm inline ip based on the context it is used
func (m *IPInterfaceSvmInlineIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetmask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineIP) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvmInlineIP) contextValidateNetmask(ctx context.Context, formats strfmt.Registry) error {

	if m.Netmask != nil {
		if err := m.Netmask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "netmask")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvmInlineIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineIP) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineLinks ip interface svm inline links
//
// swagger:model ip_interface_svm_inline__links
type IPInterfaceSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ip interface svm inline links
func (m *IPInterfaceSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ip interface svm inline links based on the context it is used
func (m *IPInterfaceSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IPInterfaceSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineLocation Home_node is optional.
//
// swagger:model ip_interface_svm_inline_location
type IPInterfaceSvmInlineLocation struct {

	// broadcast domain
	BroadcastDomain *IPInterfaceSvmInlineLocationInlineBroadcastDomain `json:"broadcast_domain,omitempty"`

	// home node
	HomeNode *IPInterfaceSvmInlineLocationInlineHomeNode `json:"home_node,omitempty"`

	// home port
	HomePort *PortSvm `json:"home_port,omitempty"`
}

// Validate validates this ip interface svm inline location
func (m *IPInterfaceSvmInlineLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBroadcastDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomeNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocation) validateBroadcastDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.BroadcastDomain) { // not required
		return nil
	}

	if m.BroadcastDomain != nil {
		if err := m.BroadcastDomain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "broadcast_domain")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvmInlineLocation) validateHomeNode(formats strfmt.Registry) error {
	if swag.IsZero(m.HomeNode) { // not required
		return nil
	}

	if m.HomeNode != nil {
		if err := m.HomeNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_node")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvmInlineLocation) validateHomePort(formats strfmt.Registry) error {
	if swag.IsZero(m.HomePort) { // not required
		return nil
	}

	if m.HomePort != nil {
		if err := m.HomePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_port")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm inline location based on the context it is used
func (m *IPInterfaceSvmInlineLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBroadcastDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHomeNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHomePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocation) contextValidateBroadcastDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.BroadcastDomain != nil {
		if err := m.BroadcastDomain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "broadcast_domain")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvmInlineLocation) contextValidateHomeNode(ctx context.Context, formats strfmt.Registry) error {

	if m.HomeNode != nil {
		if err := m.HomeNode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_node")
			}
			return err
		}
	}

	return nil
}

func (m *IPInterfaceSvmInlineLocation) contextValidateHomePort(ctx context.Context, formats strfmt.Registry) error {

	if m.HomePort != nil {
		if err := m.HomePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_port")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocation) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineLocationInlineBroadcastDomain Broadcast domain UUID along with a readable name.
//
// swagger:model ip_interface_svm_inline_location_inline_broadcast_domain
type IPInterfaceSvmInlineLocationInlineBroadcastDomain struct {

	// links
	Links *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks `json:"_links,omitempty"`

	// Name of the broadcast domain, scoped to its IPspace
	// Example: bd1
	Name *string `json:"name,omitempty"`

	// Broadcast domain UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ip interface svm inline location inline broadcast domain
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomain) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "broadcast_domain" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm inline location inline broadcast domain based on the context it is used
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomain) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "broadcast_domain" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomain) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineLocationInlineBroadcastDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks ip interface svm inline location inline broadcast domain inline links
//
// swagger:model ip_interface_svm_inline_location_inline_broadcast_domain_inline__links
type IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ip interface svm inline location inline broadcast domain inline links
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "broadcast_domain" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm inline location inline broadcast domain inline links based on the context it is used
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "broadcast_domain" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineLocationInlineBroadcastDomainInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineLocationInlineHomeNode ip interface svm inline location inline home node
//
// swagger:model ip_interface_svm_inline_location_inline_home_node
type IPInterfaceSvmInlineLocationInlineHomeNode struct {

	// links
	Links *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ip interface svm inline location inline home node
func (m *IPInterfaceSvmInlineLocationInlineHomeNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineHomeNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm inline location inline home node based on the context it is used
func (m *IPInterfaceSvmInlineLocationInlineHomeNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineHomeNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineHomeNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineHomeNode) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineLocationInlineHomeNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks ip interface svm inline location inline home node inline links
//
// swagger:model ip_interface_svm_inline_location_inline_home_node_inline__links
type IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ip interface svm inline location inline home node inline links
func (m *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip interface svm inline location inline home node inline links based on the context it is used
func (m *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "home_node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res IPInterfaceSvmInlineLocationInlineHomeNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
