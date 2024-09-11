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

// MetroclusterPartner metrocluster partner
//
// swagger:model metrocluster_partner
type MetroclusterPartner struct {

	// node
	Node *MetroclusterPartnerInlineNode `json:"node,omitempty"`

	// type
	// Enum: ["ha","dr","aux"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this metrocluster partner
func (m *MetroclusterPartner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
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

func (m *MetroclusterPartner) validateNode(formats strfmt.Registry) error {
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

var metroclusterPartnerTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ha","dr","aux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterPartnerTypeTypePropEnum = append(metroclusterPartnerTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_partner
	// MetroclusterPartner
	// type
	// Type
	// ha
	// END DEBUGGING
	// MetroclusterPartnerTypeHa captures enum value "ha"
	MetroclusterPartnerTypeHa string = "ha"

	// BEGIN DEBUGGING
	// metrocluster_partner
	// MetroclusterPartner
	// type
	// Type
	// dr
	// END DEBUGGING
	// MetroclusterPartnerTypeDr captures enum value "dr"
	MetroclusterPartnerTypeDr string = "dr"

	// BEGIN DEBUGGING
	// metrocluster_partner
	// MetroclusterPartner
	// type
	// Type
	// aux
	// END DEBUGGING
	// MetroclusterPartnerTypeAux captures enum value "aux"
	MetroclusterPartnerTypeAux string = "aux"
)

// prop value enum
func (m *MetroclusterPartner) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterPartnerTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterPartner) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster partner based on the context it is used
func (m *MetroclusterPartner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterPartner) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterPartner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterPartner) UnmarshalBinary(b []byte) error {
	var res MetroclusterPartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterPartnerInlineNode metrocluster partner inline node
//
// swagger:model metrocluster_partner_inline_node
type MetroclusterPartnerInlineNode struct {

	// links
	Links *MetroclusterPartnerInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this metrocluster partner inline node
func (m *MetroclusterPartnerInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterPartnerInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster partner inline node based on the context it is used
func (m *MetroclusterPartnerInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterPartnerInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterPartnerInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterPartnerInlineNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterPartnerInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterPartnerInlineNodeInlineLinks metrocluster partner inline node inline links
//
// swagger:model metrocluster_partner_inline_node_inline__links
type MetroclusterPartnerInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster partner inline node inline links
func (m *MetroclusterPartnerInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterPartnerInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster partner inline node inline links based on the context it is used
func (m *MetroclusterPartnerInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterPartnerInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterPartnerInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterPartnerInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterPartnerInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
