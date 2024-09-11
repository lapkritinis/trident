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

// SQLOnSan Microsoft SQL using SAN.
//
// swagger:model sql_on_san
type SQLOnSan struct {

	// db
	// Required: true
	Db *SQLOnSanInlineDb `json:"db"`

	// The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	IgroupName *string `json:"igroup_name"`

	// log
	// Required: true
	Log *SQLOnSanInlineLog `json:"log"`

	// The name of the host OS running the application.
	// Enum: ["windows","windows_2008","windows_gpt"]
	OsType *string `json:"os_type,omitempty"`

	// protection type
	ProtectionType *SQLOnSanInlineProtectionType `json:"protection_type,omitempty"`

	// The number of server cores for the DB.
	ServerCoresCount *int64 `json:"server_cores_count,omitempty"`

	// The list of initiator groups to create.
	// Max Items: 1
	// Min Items: 0
	SQLOnSanInlineNewIgroups []*SQLOnSanNewIgroups `json:"new_igroups,omitempty"`

	// temp db
	TempDb *SQLOnSanInlineTempDb `json:"temp_db,omitempty"`
}

// Validate validates this sql on san
func (m *SQLOnSan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSQLOnSanInlineNewIgroups(formats); err != nil {
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

func (m *SQLOnSan) validateDb(formats strfmt.Registry) error {

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

func (m *SQLOnSan) validateIgroupName(formats strfmt.Registry) error {

	if err := validate.Required("igroup_name", "body", m.IgroupName); err != nil {
		return err
	}

	if err := validate.MinLength("igroup_name", "body", *m.IgroupName, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("igroup_name", "body", *m.IgroupName, 96); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSan) validateLog(formats strfmt.Registry) error {

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

var sqlOnSanTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["windows","windows_2008","windows_gpt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanTypeOsTypePropEnum = append(sqlOnSanTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_san
	// SQLOnSan
	// os_type
	// OsType
	// windows
	// END DEBUGGING
	// SQLOnSanOsTypeWindows captures enum value "windows"
	SQLOnSanOsTypeWindows string = "windows"

	// BEGIN DEBUGGING
	// sql_on_san
	// SQLOnSan
	// os_type
	// OsType
	// windows_2008
	// END DEBUGGING
	// SQLOnSanOsTypeWindows2008 captures enum value "windows_2008"
	SQLOnSanOsTypeWindows2008 string = "windows_2008"

	// BEGIN DEBUGGING
	// sql_on_san
	// SQLOnSan
	// os_type
	// OsType
	// windows_gpt
	// END DEBUGGING
	// SQLOnSanOsTypeWindowsGpt captures enum value "windows_gpt"
	SQLOnSanOsTypeWindowsGpt string = "windows_gpt"
)

// prop value enum
func (m *SQLOnSan) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSan) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSan) validateProtectionType(formats strfmt.Registry) error {
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

func (m *SQLOnSan) validateSQLOnSanInlineNewIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.SQLOnSanInlineNewIgroups) { // not required
		return nil
	}

	iSQLOnSanInlineNewIgroupsSize := int64(len(m.SQLOnSanInlineNewIgroups))

	if err := validate.MinItems("new_igroups", "body", iSQLOnSanInlineNewIgroupsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("new_igroups", "body", iSQLOnSanInlineNewIgroupsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.SQLOnSanInlineNewIgroups); i++ {
		if swag.IsZero(m.SQLOnSanInlineNewIgroups[i]) { // not required
			continue
		}

		if m.SQLOnSanInlineNewIgroups[i] != nil {
			if err := m.SQLOnSanInlineNewIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLOnSan) validateTempDb(formats strfmt.Registry) error {
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

// ContextValidate validate this sql on san based on the context it is used
func (m *SQLOnSan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLOnSanInlineNewIgroups(ctx, formats); err != nil {
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

func (m *SQLOnSan) contextValidateDb(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SQLOnSan) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SQLOnSan) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SQLOnSan) contextValidateSQLOnSanInlineNewIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SQLOnSanInlineNewIgroups); i++ {

		if m.SQLOnSanInlineNewIgroups[i] != nil {
			if err := m.SQLOnSanInlineNewIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLOnSan) contextValidateTempDb(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SQLOnSan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSan) UnmarshalBinary(b []byte) error {
	var res SQLOnSan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineDb sql on san inline db
//
// swagger:model sql_on_san_inline_db
type SQLOnSanInlineDb struct {

	// The size of the DB. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *SQLOnSanInlineDbInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this sql on san inline db
func (m *SQLOnSanInlineDb) Validate(formats strfmt.Registry) error {
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

func (m *SQLOnSanInlineDb) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("db"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSanInlineDb) validateStorageService(formats strfmt.Registry) error {
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

// ContextValidate validate this sql on san inline db based on the context it is used
func (m *SQLOnSanInlineDb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSanInlineDb) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SQLOnSanInlineDb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineDb) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineDb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineDbInlineStorageService sql on san inline db inline storage service
//
// swagger:model sql_on_san_inline_db_inline_storage_service
type SQLOnSanInlineDbInlineStorageService struct {

	// The storage service of the DB.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this sql on san inline db inline storage service
func (m *SQLOnSanInlineDbInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSanInlineDbInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanInlineDbInlineStorageServiceTypeNamePropEnum = append(sqlOnSanInlineDbInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_san_inline_db_inline_storage_service
	// SQLOnSanInlineDbInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// SQLOnSanInlineDbInlineStorageServiceNameExtreme captures enum value "extreme"
	SQLOnSanInlineDbInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// sql_on_san_inline_db_inline_storage_service
	// SQLOnSanInlineDbInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// SQLOnSanInlineDbInlineStorageServiceNamePerformance captures enum value "performance"
	SQLOnSanInlineDbInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// sql_on_san_inline_db_inline_storage_service
	// SQLOnSanInlineDbInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// SQLOnSanInlineDbInlineStorageServiceNameValue captures enum value "value"
	SQLOnSanInlineDbInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *SQLOnSanInlineDbInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanInlineDbInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanInlineDbInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("db"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on san inline db inline storage service based on context it is used
func (m *SQLOnSanInlineDbInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSanInlineDbInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineDbInlineStorageService) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineDbInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineLog sql on san inline log
//
// swagger:model sql_on_san_inline_log
type SQLOnSanInlineLog struct {

	// The size of the log DB. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *SQLOnSanInlineLogInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this sql on san inline log
func (m *SQLOnSanInlineLog) Validate(formats strfmt.Registry) error {
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

func (m *SQLOnSanInlineLog) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("log"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *SQLOnSanInlineLog) validateStorageService(formats strfmt.Registry) error {
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

// ContextValidate validate this sql on san inline log based on the context it is used
func (m *SQLOnSanInlineLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSanInlineLog) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SQLOnSanInlineLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineLog) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineLogInlineStorageService sql on san inline log inline storage service
//
// swagger:model sql_on_san_inline_log_inline_storage_service
type SQLOnSanInlineLogInlineStorageService struct {

	// The storage service of the log DB.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this sql on san inline log inline storage service
func (m *SQLOnSanInlineLogInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSanInlineLogInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanInlineLogInlineStorageServiceTypeNamePropEnum = append(sqlOnSanInlineLogInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_san_inline_log_inline_storage_service
	// SQLOnSanInlineLogInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// SQLOnSanInlineLogInlineStorageServiceNameExtreme captures enum value "extreme"
	SQLOnSanInlineLogInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// sql_on_san_inline_log_inline_storage_service
	// SQLOnSanInlineLogInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// SQLOnSanInlineLogInlineStorageServiceNamePerformance captures enum value "performance"
	SQLOnSanInlineLogInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// sql_on_san_inline_log_inline_storage_service
	// SQLOnSanInlineLogInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// SQLOnSanInlineLogInlineStorageServiceNameValue captures enum value "value"
	SQLOnSanInlineLogInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *SQLOnSanInlineLogInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanInlineLogInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanInlineLogInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("log"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on san inline log inline storage service based on context it is used
func (m *SQLOnSanInlineLogInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSanInlineLogInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineLogInlineStorageService) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineLogInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineProtectionType sql on san inline protection type
//
// swagger:model sql_on_san_inline_protection_type
type SQLOnSanInlineProtectionType struct {

	// The local RPO of the application.
	// Enum: ["hourly","none"]
	LocalRpo *string `json:"local_rpo,omitempty"`

	// The remote RPO of the application.
	// Enum: ["none","zero"]
	RemoteRpo *string `json:"remote_rpo,omitempty"`
}

// Validate validates this sql on san inline protection type
func (m *SQLOnSanInlineProtectionType) Validate(formats strfmt.Registry) error {
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

var sqlOnSanInlineProtectionTypeTypeLocalRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanInlineProtectionTypeTypeLocalRpoPropEnum = append(sqlOnSanInlineProtectionTypeTypeLocalRpoPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_san_inline_protection_type
	// SQLOnSanInlineProtectionType
	// local_rpo
	// LocalRpo
	// hourly
	// END DEBUGGING
	// SQLOnSanInlineProtectionTypeLocalRpoHourly captures enum value "hourly"
	SQLOnSanInlineProtectionTypeLocalRpoHourly string = "hourly"

	// BEGIN DEBUGGING
	// sql_on_san_inline_protection_type
	// SQLOnSanInlineProtectionType
	// local_rpo
	// LocalRpo
	// none
	// END DEBUGGING
	// SQLOnSanInlineProtectionTypeLocalRpoNone captures enum value "none"
	SQLOnSanInlineProtectionTypeLocalRpoNone string = "none"
)

// prop value enum
func (m *SQLOnSanInlineProtectionType) validateLocalRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanInlineProtectionTypeTypeLocalRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanInlineProtectionType) validateLocalRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocalRpoEnum("protection_type"+"."+"local_rpo", "body", *m.LocalRpo); err != nil {
		return err
	}

	return nil
}

var sqlOnSanInlineProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanInlineProtectionTypeTypeRemoteRpoPropEnum = append(sqlOnSanInlineProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_san_inline_protection_type
	// SQLOnSanInlineProtectionType
	// remote_rpo
	// RemoteRpo
	// none
	// END DEBUGGING
	// SQLOnSanInlineProtectionTypeRemoteRpoNone captures enum value "none"
	SQLOnSanInlineProtectionTypeRemoteRpoNone string = "none"

	// BEGIN DEBUGGING
	// sql_on_san_inline_protection_type
	// SQLOnSanInlineProtectionType
	// remote_rpo
	// RemoteRpo
	// zero
	// END DEBUGGING
	// SQLOnSanInlineProtectionTypeRemoteRpoZero captures enum value "zero"
	SQLOnSanInlineProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *SQLOnSanInlineProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanInlineProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanInlineProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", *m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on san inline protection type based on context it is used
func (m *SQLOnSanInlineProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSanInlineProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineProtectionType) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineTempDb sql on san inline temp db
//
// swagger:model sql_on_san_inline_temp_db
type SQLOnSanInlineTempDb struct {

	// The size of the temp DB. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	Size *int64 `json:"size,omitempty"`

	// storage service
	StorageService *SQLOnSanInlineTempDbInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this sql on san inline temp db
func (m *SQLOnSanInlineTempDb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSanInlineTempDb) validateStorageService(formats strfmt.Registry) error {
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

// ContextValidate validate this sql on san inline temp db based on the context it is used
func (m *SQLOnSanInlineTempDb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSanInlineTempDb) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SQLOnSanInlineTempDb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineTempDb) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineTempDb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SQLOnSanInlineTempDbInlineStorageService sql on san inline temp db inline storage service
//
// swagger:model sql_on_san_inline_temp_db_inline_storage_service
type SQLOnSanInlineTempDbInlineStorageService struct {

	// The storage service of the temp DB.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this sql on san inline temp db inline storage service
func (m *SQLOnSanInlineTempDbInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sqlOnSanInlineTempDbInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanInlineTempDbInlineStorageServiceTypeNamePropEnum = append(sqlOnSanInlineTempDbInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// sql_on_san_inline_temp_db_inline_storage_service
	// SQLOnSanInlineTempDbInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// SQLOnSanInlineTempDbInlineStorageServiceNameExtreme captures enum value "extreme"
	SQLOnSanInlineTempDbInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// sql_on_san_inline_temp_db_inline_storage_service
	// SQLOnSanInlineTempDbInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// SQLOnSanInlineTempDbInlineStorageServiceNamePerformance captures enum value "performance"
	SQLOnSanInlineTempDbInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// sql_on_san_inline_temp_db_inline_storage_service
	// SQLOnSanInlineTempDbInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// SQLOnSanInlineTempDbInlineStorageServiceNameValue captures enum value "value"
	SQLOnSanInlineTempDbInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *SQLOnSanInlineTempDbInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanInlineTempDbInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanInlineTempDbInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("temp_db"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on san inline temp db inline storage service based on context it is used
func (m *SQLOnSanInlineTempDbInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSanInlineTempDbInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanInlineTempDbInlineStorageService) UnmarshalBinary(b []byte) error {
	var res SQLOnSanInlineTempDbInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
