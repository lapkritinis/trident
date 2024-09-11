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

// StorageUnitEncryptionState storage unit encryption state
//
// swagger:model storage_unit_encryption_state
type StorageUnitEncryptionState struct {

	// Storage unit data encryption state.<br>_unencrypted_ &dash; Unencrypted.<br>_software_encrypted_ &dash; Software encryption enabled.<br>_software_conversion_queued_ &dash; Queued for software conversion.<br>_software_encrypting_ &dash; Software encryption is in progress.<br>_software_rekeying_ &dash; Encryption with a new key is in progress.<br>_software_conversion_paused_ &dash; Software conversion is paused.<br>_software_rekey_paused_ &dash; Encryption with a new key is paused.
	//
	// Read Only: true
	// Enum: ["unencrypted","software_encrypted","software_conversion_queued","software_encrypting","software_rekeying","software_conversion_paused","software_rekey_paused"]
	State *string `json:"state,omitempty"`
}

// Validate validates this storage unit encryption state
func (m *StorageUnitEncryptionState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageUnitEncryptionStateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unencrypted","software_encrypted","software_conversion_queued","software_encrypting","software_rekeying","software_conversion_paused","software_rekey_paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageUnitEncryptionStateTypeStatePropEnum = append(storageUnitEncryptionStateTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// unencrypted
	// END DEBUGGING
	// StorageUnitEncryptionStateStateUnencrypted captures enum value "unencrypted"
	StorageUnitEncryptionStateStateUnencrypted string = "unencrypted"

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// software_encrypted
	// END DEBUGGING
	// StorageUnitEncryptionStateStateSoftwareEncrypted captures enum value "software_encrypted"
	StorageUnitEncryptionStateStateSoftwareEncrypted string = "software_encrypted"

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// software_conversion_queued
	// END DEBUGGING
	// StorageUnitEncryptionStateStateSoftwareConversionQueued captures enum value "software_conversion_queued"
	StorageUnitEncryptionStateStateSoftwareConversionQueued string = "software_conversion_queued"

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// software_encrypting
	// END DEBUGGING
	// StorageUnitEncryptionStateStateSoftwareEncrypting captures enum value "software_encrypting"
	StorageUnitEncryptionStateStateSoftwareEncrypting string = "software_encrypting"

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// software_rekeying
	// END DEBUGGING
	// StorageUnitEncryptionStateStateSoftwareRekeying captures enum value "software_rekeying"
	StorageUnitEncryptionStateStateSoftwareRekeying string = "software_rekeying"

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// software_conversion_paused
	// END DEBUGGING
	// StorageUnitEncryptionStateStateSoftwareConversionPaused captures enum value "software_conversion_paused"
	StorageUnitEncryptionStateStateSoftwareConversionPaused string = "software_conversion_paused"

	// BEGIN DEBUGGING
	// storage_unit_encryption_state
	// StorageUnitEncryptionState
	// state
	// State
	// software_rekey_paused
	// END DEBUGGING
	// StorageUnitEncryptionStateStateSoftwareRekeyPaused captures enum value "software_rekey_paused"
	StorageUnitEncryptionStateStateSoftwareRekeyPaused string = "software_rekey_paused"
)

// prop value enum
func (m *StorageUnitEncryptionState) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storageUnitEncryptionStateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StorageUnitEncryptionState) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this storage unit encryption state based on the context it is used
func (m *StorageUnitEncryptionState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUnitEncryptionState) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageUnitEncryptionState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUnitEncryptionState) UnmarshalBinary(b []byte) error {
	var res StorageUnitEncryptionState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
