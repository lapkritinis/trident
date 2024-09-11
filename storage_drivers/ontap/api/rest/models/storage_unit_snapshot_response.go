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

// StorageUnitSnapshotResponse storage unit snapshot response
//
// swagger:model storage_unit_snapshot_response
type StorageUnitSnapshotResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty"`

	// Number of Records
	NumRecords *int64 `json:"num_records,omitempty"`

	// Space reclaimed when the snapshot is deleted, in bytes.
	// Read Only: true
	ReclaimableSpace *int64 `json:"reclaimable_space,omitempty"`

	// storage unit snapshot response inline records
	StorageUnitSnapshotResponseInlineRecords []*StorageUnitSnapshot `json:"records,omitempty"`
}

// Validate validates this storage unit snapshot response
func (m *StorageUnitSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageUnitSnapshotResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUnitSnapshotResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *StorageUnitSnapshotResponse) validateStorageUnitSnapshotResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageUnitSnapshotResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageUnitSnapshotResponseInlineRecords); i++ {
		if swag.IsZero(m.StorageUnitSnapshotResponseInlineRecords[i]) { // not required
			continue
		}

		if m.StorageUnitSnapshotResponseInlineRecords[i] != nil {
			if err := m.StorageUnitSnapshotResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this storage unit snapshot response based on the context it is used
func (m *StorageUnitSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReclaimableSpace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageUnitSnapshotResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUnitSnapshotResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *StorageUnitSnapshotResponse) contextValidateReclaimableSpace(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reclaimable_space", "body", m.ReclaimableSpace); err != nil {
		return err
	}

	return nil
}

func (m *StorageUnitSnapshotResponse) contextValidateStorageUnitSnapshotResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageUnitSnapshotResponseInlineRecords); i++ {

		if m.StorageUnitSnapshotResponseInlineRecords[i] != nil {
			if err := m.StorageUnitSnapshotResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *StorageUnitSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUnitSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res StorageUnitSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
