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

// AutosupportIssues autosupport issues
//
// swagger:model autosupport_issues
type AutosupportIssues struct {

	// The name of the component where the issue occurred.
	//
	// Example: mail_server
	// Enum: ["https_put_destination","https_post_destination","mail_server","ondemand_server"]
	Component *string `json:"component,omitempty"`

	// corrective action
	CorrectiveAction *AutosupportConnectivityCorrectiveAction `json:"corrective_action,omitempty"`

	// The HTTPS/SMTP/AOD AutoSupport Destination.
	// Example: mailhost1.example.com
	// Read Only: true
	Destination *string `json:"destination,omitempty"`

	// issue
	Issue *AutosupportConnectivityIssue `json:"issue,omitempty"`

	// node
	Node *AutosupportIssuesInlineNode `json:"node,omitempty"`
}

// Validate validates this autosupport issues
func (m *AutosupportIssues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCorrectiveAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
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

var autosupportIssuesTypeComponentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["https_put_destination","https_post_destination","mail_server","ondemand_server"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autosupportIssuesTypeComponentPropEnum = append(autosupportIssuesTypeComponentPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// autosupport_issues
	// AutosupportIssues
	// component
	// Component
	// https_put_destination
	// END DEBUGGING
	// AutosupportIssuesComponentHTTPSPutDestination captures enum value "https_put_destination"
	AutosupportIssuesComponentHTTPSPutDestination string = "https_put_destination"

	// BEGIN DEBUGGING
	// autosupport_issues
	// AutosupportIssues
	// component
	// Component
	// https_post_destination
	// END DEBUGGING
	// AutosupportIssuesComponentHTTPSPostDestination captures enum value "https_post_destination"
	AutosupportIssuesComponentHTTPSPostDestination string = "https_post_destination"

	// BEGIN DEBUGGING
	// autosupport_issues
	// AutosupportIssues
	// component
	// Component
	// mail_server
	// END DEBUGGING
	// AutosupportIssuesComponentMailServer captures enum value "mail_server"
	AutosupportIssuesComponentMailServer string = "mail_server"

	// BEGIN DEBUGGING
	// autosupport_issues
	// AutosupportIssues
	// component
	// Component
	// ondemand_server
	// END DEBUGGING
	// AutosupportIssuesComponentOndemandServer captures enum value "ondemand_server"
	AutosupportIssuesComponentOndemandServer string = "ondemand_server"
)

// prop value enum
func (m *AutosupportIssues) validateComponentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autosupportIssuesTypeComponentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutosupportIssues) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentEnum("component", "body", *m.Component); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportIssues) validateCorrectiveAction(formats strfmt.Registry) error {
	if swag.IsZero(m.CorrectiveAction) { // not required
		return nil
	}

	if m.CorrectiveAction != nil {
		if err := m.CorrectiveAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("corrective_action")
			}
			return err
		}
	}

	return nil
}

func (m *AutosupportIssues) validateIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

func (m *AutosupportIssues) validateNode(formats strfmt.Registry) error {
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

// ContextValidate validate this autosupport issues based on the context it is used
func (m *AutosupportIssues) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCorrectiveAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssue(ctx, formats); err != nil {
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

func (m *AutosupportIssues) contextValidateCorrectiveAction(ctx context.Context, formats strfmt.Registry) error {

	if m.CorrectiveAction != nil {
		if err := m.CorrectiveAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("corrective_action")
			}
			return err
		}
	}

	return nil
}

func (m *AutosupportIssues) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportIssues) contextValidateIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.Issue != nil {
		if err := m.Issue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

func (m *AutosupportIssues) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AutosupportIssues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportIssues) UnmarshalBinary(b []byte) error {
	var res AutosupportIssues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutosupportIssuesInlineNode autosupport issues inline node
//
// swagger:model autosupport_issues_inline_node
type AutosupportIssuesInlineNode struct {

	// links
	Links *AutosupportIssuesInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this autosupport issues inline node
func (m *AutosupportIssuesInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportIssuesInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this autosupport issues inline node based on the context it is used
func (m *AutosupportIssuesInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportIssuesInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AutosupportIssuesInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportIssuesInlineNode) UnmarshalBinary(b []byte) error {
	var res AutosupportIssuesInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutosupportIssuesInlineNodeInlineLinks autosupport issues inline node inline links
//
// swagger:model autosupport_issues_inline_node_inline__links
type AutosupportIssuesInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this autosupport issues inline node inline links
func (m *AutosupportIssuesInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportIssuesInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this autosupport issues inline node inline links based on the context it is used
func (m *AutosupportIssuesInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportIssuesInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AutosupportIssuesInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportIssuesInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res AutosupportIssuesInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
