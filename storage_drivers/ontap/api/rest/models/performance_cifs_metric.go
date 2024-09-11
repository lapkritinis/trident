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

// PerformanceCifsMetric Performance numbers, such as IOPS latency and throughput.
//
// swagger:model performance_cifs_metric
type PerformanceCifsMetric struct {

	// links
	Links *PerformanceCifsMetricInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceCifsMetricInlineIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceCifsMetricInlineLatency `json:"latency,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceCifsMetricInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance cifs metric
func (m *PerformanceCifsMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceCifsMetric) validateLinks(formats strfmt.Registry) error {
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

var performanceCifsMetricTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceCifsMetricTypeDurationPropEnum = append(performanceCifsMetricTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceCifsMetricDurationPT15S captures enum value "PT15S"
	PerformanceCifsMetricDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceCifsMetricDurationPT4M captures enum value "PT4M"
	PerformanceCifsMetricDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceCifsMetricDurationPT30M captures enum value "PT30M"
	PerformanceCifsMetricDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceCifsMetricDurationPT2H captures enum value "PT2H"
	PerformanceCifsMetricDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceCifsMetricDurationP1D captures enum value "P1D"
	PerformanceCifsMetricDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceCifsMetricDurationPT5M captures enum value "PT5M"
	PerformanceCifsMetricDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceCifsMetric) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceCifsMetricTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceCifsMetric) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceCifsMetric) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceCifsMetric) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceCifsMetricTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceCifsMetricTypeStatusPropEnum = append(performanceCifsMetricTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceCifsMetricStatusOk captures enum value "ok"
	PerformanceCifsMetricStatusOk string = "ok"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceCifsMetricStatusError captures enum value "error"
	PerformanceCifsMetricStatusError string = "error"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceCifsMetricStatusPartialNoData captures enum value "partial_no_data"
	PerformanceCifsMetricStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceCifsMetricStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceCifsMetricStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceCifsMetricStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceCifsMetricStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceCifsMetricStatusNegativeDelta captures enum value "negative_delta"
	PerformanceCifsMetricStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// not_found
	// END DEBUGGING
	// PerformanceCifsMetricStatusNotFound captures enum value "not_found"
	PerformanceCifsMetricStatusNotFound string = "not_found"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceCifsMetricStatusBackfilledData captures enum value "backfilled_data"
	PerformanceCifsMetricStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceCifsMetricStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceCifsMetricStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceCifsMetricStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceCifsMetricStatusInconsistentOldData string = "inconsistent_old_data"

	// BEGIN DEBUGGING
	// performance_cifs_metric
	// PerformanceCifsMetric
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceCifsMetricStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceCifsMetricStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceCifsMetric) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceCifsMetricTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceCifsMetric) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceCifsMetric) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceCifsMetric) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance cifs metric based on the context it is used
func (m *PerformanceCifsMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceCifsMetric) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceCifsMetric) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceCifsMetric) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceCifsMetric) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceCifsMetric) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceCifsMetric) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceCifsMetric) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceCifsMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceCifsMetric) UnmarshalBinary(b []byte) error {
	var res PerformanceCifsMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceCifsMetricInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_cifs_metric_inline_iops
type PerformanceCifsMetricInlineIops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance cifs metric inline iops
func (m *PerformanceCifsMetricInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance cifs metric inline iops based on the context it is used
func (m *PerformanceCifsMetricInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceCifsMetricInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceCifsMetricInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_cifs_metric_inline_latency
type PerformanceCifsMetricInlineLatency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance cifs metric inline latency
func (m *PerformanceCifsMetricInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance cifs metric inline latency based on the context it is used
func (m *PerformanceCifsMetricInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceCifsMetricInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceCifsMetricInlineLinks performance cifs metric inline links
//
// swagger:model performance_cifs_metric_inline__links
type PerformanceCifsMetricInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance cifs metric inline links
func (m *PerformanceCifsMetricInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceCifsMetricInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance cifs metric inline links based on the context it is used
func (m *PerformanceCifsMetricInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceCifsMetricInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceCifsMetricInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceCifsMetricInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceCifsMetricInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_cifs_metric_inline_throughput
type PerformanceCifsMetricInlineThroughput struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance cifs metric inline throughput
func (m *PerformanceCifsMetricInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance cifs metric inline throughput based on the context it is used
func (m *PerformanceCifsMetricInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceCifsMetricInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceCifsMetricInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
