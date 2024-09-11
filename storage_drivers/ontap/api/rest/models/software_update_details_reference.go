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

// SoftwareUpdateDetailsReference software update details reference
//
// swagger:model software_update_details_reference
type SoftwareUpdateDetailsReference struct {

	// Elapsed duration for each update phase
	// Example: 2100
	// Read Only: true
	ElapsedDuration *int64 `json:"elapsed_duration,omitempty"`

	// Estimated duration for each update phase
	// Example: 4620
	// Read Only: true
	EstimatedDuration *int64 `json:"estimated_duration,omitempty"`

	// node
	Node *SoftwareUpdateDetailsReferenceInlineNode `json:"node,omitempty"`

	// Phase details
	// Example: Post-update checks
	// Read Only: true
	Phase *string `json:"phase,omitempty"`

	// State of the update phase
	// Example: failed
	// Read Only: true
	// Enum: ["in_progress","waiting","paused_by_user","paused_on_error","completed","canceled","failed","pause_pending","cancel_pending"]
	State *string `json:"state,omitempty"`
}

// Validate validates this software update details reference
func (m *SoftwareUpdateDetailsReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
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

func (m *SoftwareUpdateDetailsReference) validateNode(formats strfmt.Registry) error {
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

var softwareUpdateDetailsReferenceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in_progress","waiting","paused_by_user","paused_on_error","completed","canceled","failed","pause_pending","cancel_pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareUpdateDetailsReferenceTypeStatePropEnum = append(softwareUpdateDetailsReferenceTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// in_progress
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStateInProgress captures enum value "in_progress"
	SoftwareUpdateDetailsReferenceStateInProgress string = "in_progress"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// waiting
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStateWaiting captures enum value "waiting"
	SoftwareUpdateDetailsReferenceStateWaiting string = "waiting"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// paused_by_user
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStatePausedByUser captures enum value "paused_by_user"
	SoftwareUpdateDetailsReferenceStatePausedByUser string = "paused_by_user"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// paused_on_error
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStatePausedOnError captures enum value "paused_on_error"
	SoftwareUpdateDetailsReferenceStatePausedOnError string = "paused_on_error"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// completed
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStateCompleted captures enum value "completed"
	SoftwareUpdateDetailsReferenceStateCompleted string = "completed"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// canceled
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStateCanceled captures enum value "canceled"
	SoftwareUpdateDetailsReferenceStateCanceled string = "canceled"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// failed
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStateFailed captures enum value "failed"
	SoftwareUpdateDetailsReferenceStateFailed string = "failed"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// pause_pending
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStatePausePending captures enum value "pause_pending"
	SoftwareUpdateDetailsReferenceStatePausePending string = "pause_pending"

	// BEGIN DEBUGGING
	// software_update_details_reference
	// SoftwareUpdateDetailsReference
	// state
	// State
	// cancel_pending
	// END DEBUGGING
	// SoftwareUpdateDetailsReferenceStateCancelPending captures enum value "cancel_pending"
	SoftwareUpdateDetailsReferenceStateCancelPending string = "cancel_pending"
)

// prop value enum
func (m *SoftwareUpdateDetailsReference) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareUpdateDetailsReferenceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareUpdateDetailsReference) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this software update details reference based on the context it is used
func (m *SoftwareUpdateDetailsReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElapsedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEstimatedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareUpdateDetailsReference) contextValidateElapsedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "elapsed_duration", "body", m.ElapsedDuration); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateDetailsReference) contextValidateEstimatedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "estimated_duration", "body", m.EstimatedDuration); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateDetailsReference) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SoftwareUpdateDetailsReference) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateDetailsReference) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareUpdateDetailsReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareUpdateDetailsReference) UnmarshalBinary(b []byte) error {
	var res SoftwareUpdateDetailsReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareUpdateDetailsReferenceInlineNode software update details reference inline node
//
// swagger:model software_update_details_reference_inline_node
type SoftwareUpdateDetailsReferenceInlineNode struct {

	// Name of the node to be retrieved for update details.
	// Example: node1
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this software update details reference inline node
func (m *SoftwareUpdateDetailsReferenceInlineNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software update details reference inline node based on the context it is used
func (m *SoftwareUpdateDetailsReferenceInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareUpdateDetailsReferenceInlineNode) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "node"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareUpdateDetailsReferenceInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareUpdateDetailsReferenceInlineNode) UnmarshalBinary(b []byte) error {
	var res SoftwareUpdateDetailsReferenceInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
