// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SnapshotPolicySchedule The snapshot policy schedule object is associated with a snapshot policy and it defines the interval at which snapshots are created and deleted.
//
// swagger:model snapshot_policy_schedule
type SnapshotPolicySchedule struct {

	// links
	Links *SnapshotPolicyScheduleInlineLinks `json:"_links,omitempty"`

	// The number of snapshots to maintain for this schedule.
	Count *int64 `json:"count,omitempty"`

	// The prefix to use while creating snapshots at regular intervals.
	Prefix *string `json:"prefix,omitempty"`

	// The retention period of snapshots for this schedule.
	RetentionPeriod *string `json:"retention_period,omitempty"`

	// schedule
	Schedule *SnapshotPolicyScheduleInlineSchedule `json:"schedule,omitempty"`

	// Label for SnapMirror operations
	SnapmirrorLabel *string `json:"snapmirror_label,omitempty"`

	// snapshot policy
	SnapshotPolicy *SnapshotPolicyScheduleInlineSnapshotPolicy `json:"snapshot_policy,omitempty"`
}

// Validate validates this snapshot policy schedule
func (m *SnapshotPolicySchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicySchedule) validateLinks(formats strfmt.Registry) error {
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

func (m *SnapshotPolicySchedule) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotPolicySchedule) validateSnapshotPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotPolicy) { // not required
		return nil
	}

	if m.SnapshotPolicy != nil {
		if err := m.SnapshotPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy schedule based on the context it is used
func (m *SnapshotPolicySchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicySchedule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotPolicySchedule) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotPolicySchedule) contextValidateSnapshotPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotPolicy != nil {
		if err := m.SnapshotPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicySchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicySchedule) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicySchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyScheduleInlineLinks snapshot policy schedule inline links
//
// swagger:model snapshot_policy_schedule_inline__links
type SnapshotPolicyScheduleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapshot policy schedule inline links
func (m *SnapshotPolicyScheduleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snapshot policy schedule inline links based on the context it is used
func (m *SnapshotPolicyScheduleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapshotPolicyScheduleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyScheduleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyScheduleInlineSchedule snapshot policy schedule inline schedule
//
// swagger:model snapshot_policy_schedule_inline_schedule
type SnapshotPolicyScheduleInlineSchedule struct {

	// links
	Links *SnapshotPolicyScheduleInlineScheduleInlineLinks `json:"_links,omitempty"`

	// Job schedule name
	// Example: weekly
	Name *string `json:"name,omitempty"`

	// Job schedule UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapshot policy schedule inline schedule
func (m *SnapshotPolicyScheduleInlineSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineSchedule) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy schedule inline schedule based on the context it is used
func (m *SnapshotPolicyScheduleInlineSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineSchedule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineSchedule) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyScheduleInlineSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyScheduleInlineScheduleInlineLinks snapshot policy schedule inline schedule inline links
//
// swagger:model snapshot_policy_schedule_inline_schedule_inline__links
type SnapshotPolicyScheduleInlineScheduleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapshot policy schedule inline schedule inline links
func (m *SnapshotPolicyScheduleInlineScheduleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineScheduleInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy schedule inline schedule inline links based on the context it is used
func (m *SnapshotPolicyScheduleInlineScheduleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineScheduleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineScheduleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineScheduleInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyScheduleInlineScheduleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyScheduleInlineSnapshotPolicy This is a reference to the snapshot policy.
//
// swagger:model snapshot_policy_schedule_inline_snapshot_policy
type SnapshotPolicyScheduleInlineSnapshotPolicy struct {

	// links
	Links *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks `json:"_links,omitempty"`

	// name
	// Example: default
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapshot policy schedule inline snapshot policy
func (m *SnapshotPolicyScheduleInlineSnapshotPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineSnapshotPolicy) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy schedule inline snapshot policy based on the context it is used
func (m *SnapshotPolicyScheduleInlineSnapshotPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineSnapshotPolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineSnapshotPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineSnapshotPolicy) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyScheduleInlineSnapshotPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks snapshot policy schedule inline snapshot policy inline links
//
// swagger:model snapshot_policy_schedule_inline_snapshot_policy_inline__links
type SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapshot policy schedule inline snapshot policy inline links
func (m *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_policy" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy schedule inline snapshot policy inline links based on the context it is used
func (m *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot_policy" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyScheduleInlineSnapshotPolicyInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
