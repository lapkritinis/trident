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

// SQLOnSmb Microsoft SQL using SMB.
//
// swagger:model sql_on_smb
type SQLOnSmb struct {

	// access
	// Required: true
	Access *SQLOnSmbInlineAccess `json:"access"`

	// db
	// Required: true
	Db *SQLOnSmbInlineDb `json:"db"`

	// log
	// Required: true
	Log *SQLOnSmbInlineLog `json:"log"`

	// protection type
	ProtectionType *SQLOnSmbInlineProtectionType `json:"protection_type,omitempty"`

	// The number of server cores for the DB.
	ServerCoresCount *int64 `json:"server_cores_count,omitempty"`

	// temp db
	TempDb *SQLOnSmbInlineTempDb `json:"temp_db,omitempty"`
}

// Validate validates this sql on smb
func (m *SQLOnSmb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTempDb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmb) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	if m.Access != nil {
		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) validateDb(formats strfmt.Registry) error {

	if err := validate.Required("db", "body", m.Db); err != nil {
		return err
	}

	if m.Db != nil {
		if err := m.Db.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	if m.Log != nil {
		if err := m.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) validateTempDb(formats strfmt.Registry) error {
	if swag.IsZero(m.TempDb) { // not required
		return nil
	}

	if m.TempDb != nil {
		if err := m.TempDb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp_db")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sql on smb based on the context it is used
func (m *SQLOnSmb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTempDb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmb) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.Access != nil {
		if err := m.Access.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) contextValidateDb(ctx context.Context, formats strfmt.Registry) error {

	if m.Db != nil {
		if err := m.Db.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	if m.Log != nil {
		if err := m.Log.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {
		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *SQLOnSmb) contextValidateTempDb(ctx context.Context, formats strfmt.Registry) error {

	if m.TempDb != nil {
		if err := m.TempDb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp_db")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmb) UnmarshalBinary(b []byte) error {
	var res SQLOnSmb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineAccess sql on smb inline access
//
// swagger:model sql_on_smb_inline_access
type SQLOnSmbInlineAccess struct {

	// SQL installer admin user name.
	// Max Length: 256
	// Min Length: 1
	Installer *string `json:"installer,omitempty"`

	// SQL service account user name.
	// Required: true
	// Max Length: 256
	// Min Length: 1
	ServiceAccount *string `json:"service_account"`
}

// Validate validates this sql on smb inline access
func (m *SQLOnSmbInlineAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstaller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineAccess) validateInstaller(formats strfmt.Registry) error {
	if swag.IsZero(m.Installer) { // not required
		return nil
	}

	if err := validate.MinLength("access"+"."+"installer", "body", *m.Installer, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("access"+"."+"installer", "body", *m.Installer, 256); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSmbInlineAccess) validateServiceAccount(formats strfmt.Registry) error {

	if err := validate.Required("access"+"."+"service_account", "body", m.ServiceAccount); err != nil {
		return err
	}

	if err := validate.MinLength("access"+"."+"service_account", "body", *m.ServiceAccount, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("access"+"."+"service_account", "body", *m.ServiceAccount, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on smb inline access based on context it is used
func (m *SQLOnSmbInlineAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineAccess) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineDb sql on smb inline db
//
// swagger:model sql_on_smb_inline_db
type SQLOnSmbInlineDb struct {

	// The size of the DB. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *SQLOnSmbInlineDbInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this sql on smb inline db
func (m *SQLOnSmbInlineDb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineDb) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("db"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSmbInlineDb) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sql on smb inline db based on the context it is used
func (m *SQLOnSmbInlineDb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineDb) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {
		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineDb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineDb) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineDb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineDbInlineStorageService sql on smb inline db inline storage service
//
// swagger:model sql_on_smb_inline_db_inline_storage_service
type SQLOnSmbInlineDbInlineStorageService struct {

	// The storage service of the DB.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this sql on smb inline db inline storage service
func (m *SQLOnSmbInlineDbInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSmbInlineDbInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSmbInlineDbInlineStorageServiceTypeNamePropEnum = append(sqlOnSmbInlineDbInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_smb_inline_db_inline_storage_service
	// SQLOnSmbInlineDbInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// SQLOnSmbInlineDbInlineStorageServiceNameExtreme captures enum value "extreme"
	SQLOnSmbInlineDbInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_db_inline_storage_service
	// SQLOnSmbInlineDbInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// SQLOnSmbInlineDbInlineStorageServiceNamePerformance captures enum value "performance"
	SQLOnSmbInlineDbInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_db_inline_storage_service
	// SQLOnSmbInlineDbInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// SQLOnSmbInlineDbInlineStorageServiceNameValue captures enum value "value"
	SQLOnSmbInlineDbInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *SQLOnSmbInlineDbInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSmbInlineDbInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSmbInlineDbInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("db"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on smb inline db inline storage service based on context it is used
func (m *SQLOnSmbInlineDbInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineDbInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineDbInlineStorageService) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineDbInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineLog sql on smb inline log
//
// swagger:model sql_on_smb_inline_log
type SQLOnSmbInlineLog struct {

	// The size of the log DB. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *SQLOnSmbInlineLogInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this sql on smb inline log
func (m *SQLOnSmbInlineLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineLog) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("log"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSmbInlineLog) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sql on smb inline log based on the context it is used
func (m *SQLOnSmbInlineLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineLog) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {
		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineLog) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineLogInlineStorageService sql on smb inline log inline storage service
//
// swagger:model sql_on_smb_inline_log_inline_storage_service
type SQLOnSmbInlineLogInlineStorageService struct {

	// The storage service of the log DB.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this sql on smb inline log inline storage service
func (m *SQLOnSmbInlineLogInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSmbInlineLogInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSmbInlineLogInlineStorageServiceTypeNamePropEnum = append(sqlOnSmbInlineLogInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_smb_inline_log_inline_storage_service
	// SQLOnSmbInlineLogInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// SQLOnSmbInlineLogInlineStorageServiceNameExtreme captures enum value "extreme"
	SQLOnSmbInlineLogInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_log_inline_storage_service
	// SQLOnSmbInlineLogInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// SQLOnSmbInlineLogInlineStorageServiceNamePerformance captures enum value "performance"
	SQLOnSmbInlineLogInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_log_inline_storage_service
	// SQLOnSmbInlineLogInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// SQLOnSmbInlineLogInlineStorageServiceNameValue captures enum value "value"
	SQLOnSmbInlineLogInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *SQLOnSmbInlineLogInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSmbInlineLogInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSmbInlineLogInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("log"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on smb inline log inline storage service based on context it is used
func (m *SQLOnSmbInlineLogInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineLogInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineLogInlineStorageService) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineLogInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineProtectionType sql on smb inline protection type
//
// swagger:model sql_on_smb_inline_protection_type
type SQLOnSmbInlineProtectionType struct {

	// The local RPO of the application.
	// Enum: ["hourly","none"]
	LocalRpo *string `json:"local_rpo,omitempty"`

	// The remote RPO of the application.
	// Enum: ["none","zero"]
	RemoteRpo *string `json:"remote_rpo,omitempty"`
}

// Validate validates this sql on smb inline protection type
func (m *SQLOnSmbInlineProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalRpo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSmbInlineProtectionTypeTypeLocalRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSmbInlineProtectionTypeTypeLocalRpoPropEnum = append(sqlOnSmbInlineProtectionTypeTypeLocalRpoPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_smb_inline_protection_type
	// SQLOnSmbInlineProtectionType
	// local_rpo
	// LocalRpo
	// hourly
	// END DEBUGGING
	// SQLOnSmbInlineProtectionTypeLocalRpoHourly captures enum value "hourly"
	SQLOnSmbInlineProtectionTypeLocalRpoHourly string = "hourly"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_protection_type
	// SQLOnSmbInlineProtectionType
	// local_rpo
	// LocalRpo
	// none
	// END DEBUGGING
	// SQLOnSmbInlineProtectionTypeLocalRpoNone captures enum value "none"
	SQLOnSmbInlineProtectionTypeLocalRpoNone string = "none"
)

// prop value enum
func (m *SQLOnSmbInlineProtectionType) validateLocalRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSmbInlineProtectionTypeTypeLocalRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSmbInlineProtectionType) validateLocalRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocalRpoEnum("protection_type"+"."+"local_rpo", "body", *m.LocalRpo); err != nil {
		return err
	}

	return nil
}

var sqlOnSmbInlineProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSmbInlineProtectionTypeTypeRemoteRpoPropEnum = append(sqlOnSmbInlineProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_smb_inline_protection_type
	// SQLOnSmbInlineProtectionType
	// remote_rpo
	// RemoteRpo
	// none
	// END DEBUGGING
	// SQLOnSmbInlineProtectionTypeRemoteRpoNone captures enum value "none"
	SQLOnSmbInlineProtectionTypeRemoteRpoNone string = "none"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_protection_type
	// SQLOnSmbInlineProtectionType
	// remote_rpo
	// RemoteRpo
	// zero
	// END DEBUGGING
	// SQLOnSmbInlineProtectionTypeRemoteRpoZero captures enum value "zero"
	SQLOnSmbInlineProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *SQLOnSmbInlineProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSmbInlineProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSmbInlineProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", *m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on smb inline protection type based on context it is used
func (m *SQLOnSmbInlineProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineProtectionType) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineTempDb sql on smb inline temp db
//
// swagger:model sql_on_smb_inline_temp_db
type SQLOnSmbInlineTempDb struct {

	// The size of the temp DB. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	Size *int64 `json:"size,omitempty"`

	// storage service
	StorageService *SQLOnSmbInlineTempDbInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this sql on smb inline temp db
func (m *SQLOnSmbInlineTempDb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineTempDb) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp_db" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sql on smb inline temp db based on the context it is used
func (m *SQLOnSmbInlineTempDb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSmbInlineTempDb) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {
		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp_db" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineTempDb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineTempDb) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineTempDb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSmbInlineTempDbInlineStorageService sql on smb inline temp db inline storage service
//
// swagger:model sql_on_smb_inline_temp_db_inline_storage_service
type SQLOnSmbInlineTempDbInlineStorageService struct {

	// The storage service of the temp DB.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this sql on smb inline temp db inline storage service
func (m *SQLOnSmbInlineTempDbInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSmbInlineTempDbInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSmbInlineTempDbInlineStorageServiceTypeNamePropEnum = append(sqlOnSmbInlineTempDbInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_smb_inline_temp_db_inline_storage_service
	// SQLOnSmbInlineTempDbInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// SQLOnSmbInlineTempDbInlineStorageServiceNameExtreme captures enum value "extreme"
	SQLOnSmbInlineTempDbInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_temp_db_inline_storage_service
	// SQLOnSmbInlineTempDbInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// SQLOnSmbInlineTempDbInlineStorageServiceNamePerformance captures enum value "performance"
	SQLOnSmbInlineTempDbInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// sql_on_smb_inline_temp_db_inline_storage_service
	// SQLOnSmbInlineTempDbInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// SQLOnSmbInlineTempDbInlineStorageServiceNameValue captures enum value "value"
	SQLOnSmbInlineTempDbInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *SQLOnSmbInlineTempDbInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSmbInlineTempDbInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSmbInlineTempDbInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("temp_db"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on smb inline temp db inline storage service based on context it is used
func (m *SQLOnSmbInlineTempDbInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSmbInlineTempDbInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSmbInlineTempDbInlineStorageService) UnmarshalBinary(b []byte) error {
	var res SQLOnSmbInlineTempDbInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
