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

// NodeMetricsResponse node metrics response
//
// swagger:model node_metrics_response
type NodeMetricsResponse struct {

	// links
	Links *NodeMetricsResponseInlineLinks `json:"_links,omitempty"`

	// node metrics response inline records
	NodeMetricsResponseInlineRecords []*NodeMetricsResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this node metrics response
func (m *NodeMetricsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeMetricsResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *NodeMetricsResponse) validateNodeMetricsResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeMetricsResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeMetricsResponseInlineRecords); i++ {
		if swag.IsZero(m.NodeMetricsResponseInlineRecords[i]) { // not required
			continue
		}

		if m.NodeMetricsResponseInlineRecords[i] != nil {
			if err := m.NodeMetricsResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this node metrics response based on the context it is used
func (m *NodeMetricsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeMetricsResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NodeMetricsResponse) contextValidateNodeMetricsResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeMetricsResponseInlineRecords); i++ {

		if m.NodeMetricsResponseInlineRecords[i] != nil {
			if err := m.NodeMetricsResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeMetricsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeMetricsResponse) UnmarshalBinary(b []byte) error {
	var res NodeMetricsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NodeMetricsResponseInlineLinks node metrics response inline links
//
// swagger:model node_metrics_response_inline__links
type NodeMetricsResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this node metrics response inline links
func (m *NodeMetricsResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *NodeMetricsResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this node metrics response inline links based on the context it is used
func (m *NodeMetricsResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *NodeMetricsResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NodeMetricsResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeMetricsResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res NodeMetricsResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NodeMetricsResponseInlineRecordsInlineArrayItem CPU performance for the nodes.
//
// swagger:model node_metrics_response_inline_records_inline_array_item
type NodeMetricsResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Enum: ["PT15S","PT5M","PT30M","PT2H","P1D"]
	Duration *string `json:"duration,omitempty"`

	// Average CPU Utilization for the node
	// Example: 13
	ProcessorUtilization *int64 `json:"processor_utilization,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: ["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]
	Status *string `json:"status,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this node metrics response inline records inline array item
func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

var nodeMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT5M","PT30M","PT2H","P1D"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(nodeMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	NodeMetricsResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	NodeMetricsResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"
)

// prop value enum
func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

var nodeMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(nodeMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// ok
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// error
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// node_metrics_response_inline_records_inline_array_item
	// NodeMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// NodeMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	NodeMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node metrics response inline records inline array item based on the context it is used
func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeMetricsResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res NodeMetricsResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks node metrics response inline records inline array item inline links
//
// swagger:model node_metrics_response_inline_records_inline_array_item_inline__links
type NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this node metrics response inline records inline array item inline links
func (m *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this node metrics response inline records inline array item inline links based on the context it is used
func (m *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res NodeMetricsResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
