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

// SoftwareReference software reference
//
// swagger:model software_reference
type SoftwareReference struct {

	// links
	Links *SoftwareReferenceInlineLinks `json:"_links,omitempty"`

	// User triggered action to apply to the install operation
	// Example: pause
	// Enum: ["pause","cancel","resume"]
	Action *string `json:"action,omitempty"`

	// Elapsed time during the upgrade or validation operation
	// Example: 2140
	// Read Only: true
	ElapsedDuration *int64 `json:"elapsed_duration,omitempty"`

	// Overall estimated time for completion of the upgrade or validation operation.
	// Example: 5220
	// Read Only: true
	EstimatedDuration *int64 `json:"estimated_duration,omitempty"`

	// metrocluster
	Metrocluster *SoftwareReferenceInlineMetrocluster `json:"metrocluster,omitempty"`

	// Version being installed on the system.
	// Example: ONTAP_X_1
	// Read Only: true
	PendingVersion *string `json:"pending_version,omitempty"`

	// List of nodes, active versions, and firmware update progressions.
	// Read Only: true
	SoftwareReferenceInlineNodes []*SoftwareNodeReference `json:"nodes,omitempty"`

	// List of failed post-update checks' warnings, errors, and advice.
	// Read Only: true
	SoftwareReferenceInlinePostUpdateChecks []*SoftwareValidationReference `json:"post_update_checks,omitempty"`

	// Display status details.
	// Read Only: true
	SoftwareReferenceInlineStatusDetails []*SoftwareStatusDetailsReference `json:"status_details,omitempty"`

	// Display update progress details.
	// Read Only: true
	SoftwareReferenceInlineUpdateDetails []*SoftwareUpdateDetailsReference `json:"update_details,omitempty"`

	// List of validation warnings, errors, and advice.
	// Read Only: true
	SoftwareReferenceInlineValidationResults []*SoftwareValidationReference `json:"validation_results,omitempty"`

	// Operational state of the upgrade
	// Example: completed
	// Read Only: true
	// Enum: ["in_progress","waiting","paused_by_user","paused_on_error","completed","canceled","failed","pause_pending","cancel_pending"]
	State *string `json:"state,omitempty"`

	// Version of ONTAP installed and currently active on the system. During PATCH, using the 'validate_only' parameter on the request executes pre-checks, but does not perform the full installation.
	// Example: ONTAP_X
	Version *string `json:"version,omitempty"`
}

// Validate validates this software reference
func (m *SoftwareReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrocluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareReferenceInlineNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareReferenceInlinePostUpdateChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareReferenceInlineStatusDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareReferenceInlineUpdateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareReferenceInlineValidationResults(formats); err != nil {
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

func (m *SoftwareReference) validateLinks(formats strfmt.Registry) error {
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

var softwareReferenceTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pause","cancel","resume"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareReferenceTypeActionPropEnum = append(softwareReferenceTypeActionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// action
	// Action
	// pause
	// END DEBUGGING
	// SoftwareReferenceActionPause captures enum value "pause"
	SoftwareReferenceActionPause string = "pause"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// action
	// Action
	// cancel
	// END DEBUGGING
	// SoftwareReferenceActionCancel captures enum value "cancel"
	SoftwareReferenceActionCancel string = "cancel"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// action
	// Action
	// resume
	// END DEBUGGING
	// SoftwareReferenceActionResume captures enum value "resume"
	SoftwareReferenceActionResume string = "resume"
)

// prop value enum
func (m *SoftwareReference) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareReferenceTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareReference) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareReference) validateMetrocluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrocluster) { // not required
		return nil
	}

	if m.Metrocluster != nil {
		if err := m.Metrocluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrocluster")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareReference) validateSoftwareReferenceInlineNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareReferenceInlineNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareReferenceInlineNodes); i++ {
		if swag.IsZero(m.SoftwareReferenceInlineNodes[i]) { // not required
			continue
		}

		if m.SoftwareReferenceInlineNodes[i] != nil {
			if err := m.SoftwareReferenceInlineNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) validateSoftwareReferenceInlinePostUpdateChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareReferenceInlinePostUpdateChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareReferenceInlinePostUpdateChecks); i++ {
		if swag.IsZero(m.SoftwareReferenceInlinePostUpdateChecks[i]) { // not required
			continue
		}

		if m.SoftwareReferenceInlinePostUpdateChecks[i] != nil {
			if err := m.SoftwareReferenceInlinePostUpdateChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_update_checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) validateSoftwareReferenceInlineStatusDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareReferenceInlineStatusDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareReferenceInlineStatusDetails); i++ {
		if swag.IsZero(m.SoftwareReferenceInlineStatusDetails[i]) { // not required
			continue
		}

		if m.SoftwareReferenceInlineStatusDetails[i] != nil {
			if err := m.SoftwareReferenceInlineStatusDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) validateSoftwareReferenceInlineUpdateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareReferenceInlineUpdateDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareReferenceInlineUpdateDetails); i++ {
		if swag.IsZero(m.SoftwareReferenceInlineUpdateDetails[i]) { // not required
			continue
		}

		if m.SoftwareReferenceInlineUpdateDetails[i] != nil {
			if err := m.SoftwareReferenceInlineUpdateDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) validateSoftwareReferenceInlineValidationResults(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareReferenceInlineValidationResults) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareReferenceInlineValidationResults); i++ {
		if swag.IsZero(m.SoftwareReferenceInlineValidationResults[i]) { // not required
			continue
		}

		if m.SoftwareReferenceInlineValidationResults[i] != nil {
			if err := m.SoftwareReferenceInlineValidationResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validation_results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var softwareReferenceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in_progress","waiting","paused_by_user","paused_on_error","completed","canceled","failed","pause_pending","cancel_pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareReferenceTypeStatePropEnum = append(softwareReferenceTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// in_progress
	// END DEBUGGING
	// SoftwareReferenceStateInProgress captures enum value "in_progress"
	SoftwareReferenceStateInProgress string = "in_progress"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// waiting
	// END DEBUGGING
	// SoftwareReferenceStateWaiting captures enum value "waiting"
	SoftwareReferenceStateWaiting string = "waiting"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// paused_by_user
	// END DEBUGGING
	// SoftwareReferenceStatePausedByUser captures enum value "paused_by_user"
	SoftwareReferenceStatePausedByUser string = "paused_by_user"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// paused_on_error
	// END DEBUGGING
	// SoftwareReferenceStatePausedOnError captures enum value "paused_on_error"
	SoftwareReferenceStatePausedOnError string = "paused_on_error"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// completed
	// END DEBUGGING
	// SoftwareReferenceStateCompleted captures enum value "completed"
	SoftwareReferenceStateCompleted string = "completed"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// canceled
	// END DEBUGGING
	// SoftwareReferenceStateCanceled captures enum value "canceled"
	SoftwareReferenceStateCanceled string = "canceled"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// failed
	// END DEBUGGING
	// SoftwareReferenceStateFailed captures enum value "failed"
	SoftwareReferenceStateFailed string = "failed"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// pause_pending
	// END DEBUGGING
	// SoftwareReferenceStatePausePending captures enum value "pause_pending"
	SoftwareReferenceStatePausePending string = "pause_pending"

	// BEGIN DEBUGGING
	// software_reference
	// SoftwareReference
	// state
	// State
	// cancel_pending
	// END DEBUGGING
	// SoftwareReferenceStateCancelPending captures enum value "cancel_pending"
	SoftwareReferenceStateCancelPending string = "cancel_pending"
)

// prop value enum
func (m *SoftwareReference) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareReferenceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareReference) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this software reference based on the context it is used
func (m *SoftwareReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElapsedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEstimatedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrocluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePendingVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareReferenceInlineNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareReferenceInlinePostUpdateChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareReferenceInlineStatusDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareReferenceInlineUpdateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareReferenceInlineValidationResults(ctx, formats); err != nil {
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

func (m *SoftwareReference) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SoftwareReference) contextValidateElapsedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "elapsed_duration", "body", m.ElapsedDuration); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareReference) contextValidateEstimatedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "estimated_duration", "body", m.EstimatedDuration); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareReference) contextValidateMetrocluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrocluster != nil {
		if err := m.Metrocluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrocluster")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareReference) contextValidatePendingVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "pending_version", "body", m.PendingVersion); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareReference) contextValidateSoftwareReferenceInlineNodes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nodes", "body", []*SoftwareNodeReference(m.SoftwareReferenceInlineNodes)); err != nil {
		return err
	}

	for i := 0; i < len(m.SoftwareReferenceInlineNodes); i++ {

		if m.SoftwareReferenceInlineNodes[i] != nil {
			if err := m.SoftwareReferenceInlineNodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) contextValidateSoftwareReferenceInlinePostUpdateChecks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "post_update_checks", "body", []*SoftwareValidationReference(m.SoftwareReferenceInlinePostUpdateChecks)); err != nil {
		return err
	}

	for i := 0; i < len(m.SoftwareReferenceInlinePostUpdateChecks); i++ {

		if m.SoftwareReferenceInlinePostUpdateChecks[i] != nil {
			if err := m.SoftwareReferenceInlinePostUpdateChecks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_update_checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) contextValidateSoftwareReferenceInlineStatusDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status_details", "body", []*SoftwareStatusDetailsReference(m.SoftwareReferenceInlineStatusDetails)); err != nil {
		return err
	}

	for i := 0; i < len(m.SoftwareReferenceInlineStatusDetails); i++ {

		if m.SoftwareReferenceInlineStatusDetails[i] != nil {
			if err := m.SoftwareReferenceInlineStatusDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) contextValidateSoftwareReferenceInlineUpdateDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "update_details", "body", []*SoftwareUpdateDetailsReference(m.SoftwareReferenceInlineUpdateDetails)); err != nil {
		return err
	}

	for i := 0; i < len(m.SoftwareReferenceInlineUpdateDetails); i++ {

		if m.SoftwareReferenceInlineUpdateDetails[i] != nil {
			if err := m.SoftwareReferenceInlineUpdateDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) contextValidateSoftwareReferenceInlineValidationResults(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "validation_results", "body", []*SoftwareValidationReference(m.SoftwareReferenceInlineValidationResults)); err != nil {
		return err
	}

	for i := 0; i < len(m.SoftwareReferenceInlineValidationResults); i++ {

		if m.SoftwareReferenceInlineValidationResults[i] != nil {
			if err := m.SoftwareReferenceInlineValidationResults[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validation_results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReference) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareReference) UnmarshalBinary(b []byte) error {
	var res SoftwareReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareReferenceInlineLinks software reference inline links
//
// swagger:model software_reference_inline__links
type SoftwareReferenceInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this software reference inline links
func (m *SoftwareReferenceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareReferenceInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this software reference inline links based on the context it is used
func (m *SoftwareReferenceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareReferenceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SoftwareReferenceInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareReferenceInlineLinks) UnmarshalBinary(b []byte) error {
	var res SoftwareReferenceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareReferenceInlineMetrocluster software reference inline metrocluster
//
// swagger:model software_reference_inline_metrocluster
type SoftwareReferenceInlineMetrocluster struct {

	// List of MetroCluster sites, statuses, and active ONTAP versions.
	// Read Only: true
	Clusters []*SoftwareMccReference `json:"clusters,omitempty"`

	// progress details
	ProgressDetails *SoftwareReferenceInlineMetroclusterInlineProgressDetails `json:"progress_details,omitempty"`

	// progress summary
	ProgressSummary *SoftwareReferenceInlineMetroclusterInlineProgressSummary `json:"progress_summary,omitempty"`
}

// Validate validates this software reference inline metrocluster
func (m *SoftwareReferenceInlineMetrocluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgressDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgressSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareReferenceInlineMetrocluster) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrocluster" + "." + "clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReferenceInlineMetrocluster) validateProgressDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ProgressDetails) { // not required
		return nil
	}

	if m.ProgressDetails != nil {
		if err := m.ProgressDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrocluster" + "." + "progress_details")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareReferenceInlineMetrocluster) validateProgressSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.ProgressSummary) { // not required
		return nil
	}

	if m.ProgressSummary != nil {
		if err := m.ProgressSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrocluster" + "." + "progress_summary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this software reference inline metrocluster based on the context it is used
func (m *SoftwareReferenceInlineMetrocluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProgressDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProgressSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareReferenceInlineMetrocluster) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "metrocluster"+"."+"clusters", "body", []*SoftwareMccReference(m.Clusters)); err != nil {
		return err
	}

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrocluster" + "." + "clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareReferenceInlineMetrocluster) contextValidateProgressDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ProgressDetails != nil {
		if err := m.ProgressDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrocluster" + "." + "progress_details")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareReferenceInlineMetrocluster) contextValidateProgressSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.ProgressSummary != nil {
		if err := m.ProgressSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrocluster" + "." + "progress_summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareReferenceInlineMetrocluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareReferenceInlineMetrocluster) UnmarshalBinary(b []byte) error {
	var res SoftwareReferenceInlineMetrocluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareReferenceInlineMetroclusterInlineProgressDetails software reference inline metrocluster inline progress details
//
// swagger:model software_reference_inline_metrocluster_inline_progress_details
type SoftwareReferenceInlineMetroclusterInlineProgressDetails struct {

	// MetroCluster update progress details.
	// Example: Switchover in progress
	Message *string `json:"message,omitempty"`
}

// Validate validates this software reference inline metrocluster inline progress details
func (m *SoftwareReferenceInlineMetroclusterInlineProgressDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software reference inline metrocluster inline progress details based on the context it is used
func (m *SoftwareReferenceInlineMetroclusterInlineProgressDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareReferenceInlineMetroclusterInlineProgressDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareReferenceInlineMetroclusterInlineProgressDetails) UnmarshalBinary(b []byte) error {
	var res SoftwareReferenceInlineMetroclusterInlineProgressDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareReferenceInlineMetroclusterInlineProgressSummary software reference inline metrocluster inline progress summary
//
// swagger:model software_reference_inline_metrocluster_inline_progress_summary
type SoftwareReferenceInlineMetroclusterInlineProgressSummary struct {

	// MetroCluster update progress summary.
	// Example: MetroCluster updated successfully.
	Message *string `json:"message,omitempty"`
}

// Validate validates this software reference inline metrocluster inline progress summary
func (m *SoftwareReferenceInlineMetroclusterInlineProgressSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software reference inline metrocluster inline progress summary based on the context it is used
func (m *SoftwareReferenceInlineMetroclusterInlineProgressSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareReferenceInlineMetroclusterInlineProgressSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareReferenceInlineMetroclusterInlineProgressSummary) UnmarshalBinary(b []byte) error {
	var res SoftwareReferenceInlineMetroclusterInlineProgressSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
