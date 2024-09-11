// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CifsDomainPreferredDc cifs domain preferred dc
//
// swagger:model cifs_domain_preferred_dc
type CifsDomainPreferredDc struct {

	// Fully Qualified Domain Name.
	//
	// Example: test.com
	Fqdn *string `json:"fqdn,omitempty"`

	// IP address of the preferred domain controller (DC). The address can be either an IPv4 or an IPv6 address.
	//
	// Example: 4.4.4.4
	ServerIP *string `json:"server_ip,omitempty"`

	// status
	Status *CifsDomainPreferredDcInlineStatus `json:"status,omitempty"`

	// svm
	Svm *CifsDomainPreferredDcInlineSvm `json:"svm,omitempty"`
}

// Validate validates this cifs domain preferred dc
func (m *CifsDomainPreferredDc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDc) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *CifsDomainPreferredDc) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs domain preferred dc based on the context it is used
func (m *CifsDomainPreferredDc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDc) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *CifsDomainPreferredDc) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDc) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsDomainPreferredDcInlineStatus Status of CIFS preferred domain controller.
//
// swagger:model cifs_domain_preferred_dc_inline_status
type CifsDomainPreferredDcInlineStatus struct {

	// Provides a detailed description of the state if the state is 'down' or
	// the response time of the DNS server if the state is 'up'.
	//
	// Example: Response time (msec): 111
	// Read Only: true
	Details *string `json:"details,omitempty"`

	// Indicates whether or not the domain controller is reachable.
	// Example: true
	// Read Only: true
	Reachable *bool `json:"reachable,omitempty"`
}

// Validate validates this cifs domain preferred dc inline status
func (m *CifsDomainPreferredDcInlineStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cifs domain preferred dc inline status based on the context it is used
func (m *CifsDomainPreferredDcInlineStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDcInlineStatus) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status"+"."+"details", "body", m.Details); err != nil {
		return err
	}

	return nil
}

func (m *CifsDomainPreferredDcInlineStatus) contextValidateReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status"+"."+"reachable", "body", m.Reachable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineStatus) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDcInlineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsDomainPreferredDcInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model cifs_domain_preferred_dc_inline_svm
type CifsDomainPreferredDcInlineSvm struct {

	// links
	Links *CifsDomainPreferredDcInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cifs domain preferred dc inline svm
func (m *CifsDomainPreferredDcInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDcInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs domain preferred dc inline svm based on the context it is used
func (m *CifsDomainPreferredDcInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDcInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineSvm) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDcInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsDomainPreferredDcInlineSvmInlineLinks cifs domain preferred dc inline svm inline links
//
// swagger:model cifs_domain_preferred_dc_inline_svm_inline__links
type CifsDomainPreferredDcInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs domain preferred dc inline svm inline links
func (m *CifsDomainPreferredDcInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDcInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs domain preferred dc inline svm inline links based on the context it is used
func (m *CifsDomainPreferredDcInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDcInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDcInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
