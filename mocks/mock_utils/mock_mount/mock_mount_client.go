// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils/mount (interfaces: Mount)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_utils/mock_mount/mock_mount_client.go github.com/netapp/trident/utils/mount Mount
//

// Package mock_mount is a generated GoMock package.
package mock_mount

import (
	context "context"
	reflect "reflect"

	models "github.com/netapp/trident/utils/models"
	gomock "go.uber.org/mock/gomock"
)

// MockMount is a mock of Mount interface.
type MockMount struct {
	ctrl     *gomock.Controller
	recorder *MockMountMockRecorder
}

// MockMountMockRecorder is the mock recorder for MockMount.
type MockMountMockRecorder struct {
	mock *MockMount
}

// NewMockMount creates a new mock instance.
func NewMockMount(ctrl *gomock.Controller) *MockMount {
	mock := &MockMount{ctrl: ctrl}
	mock.recorder = &MockMountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMount) EXPECT() *MockMountMockRecorder {
	return m.recorder
}

// AttachNFSVolume mocks base method.
func (m *MockMount) AttachNFSVolume(arg0 context.Context, arg1, arg2 string, arg3 *models.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachNFSVolume", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachNFSVolume indicates an expected call of AttachNFSVolume.
func (mr *MockMountMockRecorder) AttachNFSVolume(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachNFSVolume", reflect.TypeOf((*MockMount)(nil).AttachNFSVolume), arg0, arg1, arg2, arg3)
}

// AttachSMBVolume mocks base method.
func (m *MockMount) AttachSMBVolume(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 *models.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachSMBVolume", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachSMBVolume indicates an expected call of AttachSMBVolume.
func (mr *MockMountMockRecorder) AttachSMBVolume(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachSMBVolume", reflect.TypeOf((*MockMount)(nil).AttachSMBVolume), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetHostMountInfo mocks base method.
func (m *MockMount) GetHostMountInfo(arg0 context.Context) ([]models.MountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostMountInfo", arg0)
	ret0, _ := ret[0].([]models.MountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostMountInfo indicates an expected call of GetHostMountInfo.
func (mr *MockMountMockRecorder) GetHostMountInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostMountInfo", reflect.TypeOf((*MockMount)(nil).GetHostMountInfo), arg0)
}

// GetSelfMountInfo mocks base method.
func (m *MockMount) GetSelfMountInfo(arg0 context.Context) ([]models.MountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfMountInfo", arg0)
	ret0, _ := ret[0].([]models.MountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelfMountInfo indicates an expected call of GetSelfMountInfo.
func (mr *MockMountMockRecorder) GetSelfMountInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfMountInfo", reflect.TypeOf((*MockMount)(nil).GetSelfMountInfo), arg0)
}

// IsCompatible mocks base method.
func (m *MockMount) IsCompatible(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCompatible", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsCompatible indicates an expected call of IsCompatible.
func (mr *MockMountMockRecorder) IsCompatible(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCompatible", reflect.TypeOf((*MockMount)(nil).IsCompatible), arg0, arg1)
}

// IsLikelyNotMountPoint mocks base method.
func (m *MockMount) IsLikelyNotMountPoint(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLikelyNotMountPoint", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLikelyNotMountPoint indicates an expected call of IsLikelyNotMountPoint.
func (mr *MockMountMockRecorder) IsLikelyNotMountPoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLikelyNotMountPoint", reflect.TypeOf((*MockMount)(nil).IsLikelyNotMountPoint), arg0, arg1)
}

// IsMounted mocks base method.
func (m *MockMount) IsMounted(arg0 context.Context, arg1, arg2, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMounted", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMounted indicates an expected call of IsMounted.
func (mr *MockMountMockRecorder) IsMounted(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMounted", reflect.TypeOf((*MockMount)(nil).IsMounted), arg0, arg1, arg2, arg3)
}

// IsNFSShareMounted mocks base method.
func (m *MockMount) IsNFSShareMounted(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNFSShareMounted", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNFSShareMounted indicates an expected call of IsNFSShareMounted.
func (mr *MockMountMockRecorder) IsNFSShareMounted(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNFSShareMounted", reflect.TypeOf((*MockMount)(nil).IsNFSShareMounted), arg0, arg1, arg2)
}

// ListProcMountinfo mocks base method.
func (m *MockMount) ListProcMountinfo() ([]models.MountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProcMountinfo")
	ret0, _ := ret[0].([]models.MountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProcMountinfo indicates an expected call of ListProcMountinfo.
func (mr *MockMountMockRecorder) ListProcMountinfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProcMountinfo", reflect.TypeOf((*MockMount)(nil).ListProcMountinfo))
}

// ListProcMounts mocks base method.
func (m *MockMount) ListProcMounts(arg0 string) ([]models.MountPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProcMounts", arg0)
	ret0, _ := ret[0].([]models.MountPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProcMounts indicates an expected call of ListProcMounts.
func (mr *MockMountMockRecorder) ListProcMounts(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProcMounts", reflect.TypeOf((*MockMount)(nil).ListProcMounts), arg0)
}

// MountDevice mocks base method.
func (m *MockMount) MountDevice(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountDevice", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountDevice indicates an expected call of MountDevice.
func (mr *MockMountMockRecorder) MountDevice(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountDevice", reflect.TypeOf((*MockMount)(nil).MountDevice), arg0, arg1, arg2, arg3, arg4)
}

// MountFilesystemForResize mocks base method.
func (m *MockMount) MountFilesystemForResize(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountFilesystemForResize", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MountFilesystemForResize indicates an expected call of MountFilesystemForResize.
func (mr *MockMountMockRecorder) MountFilesystemForResize(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountFilesystemForResize", reflect.TypeOf((*MockMount)(nil).MountFilesystemForResize), arg0, arg1, arg2, arg3)
}

// MountNFSPath mocks base method.
func (m *MockMount) MountNFSPath(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountNFSPath", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountNFSPath indicates an expected call of MountNFSPath.
func (mr *MockMountMockRecorder) MountNFSPath(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountNFSPath", reflect.TypeOf((*MockMount)(nil).MountNFSPath), arg0, arg1, arg2, arg3)
}

// MountSMBPath mocks base method.
func (m *MockMount) MountSMBPath(arg0 context.Context, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountSMBPath", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountSMBPath indicates an expected call of MountSMBPath.
func (mr *MockMountMockRecorder) MountSMBPath(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountSMBPath", reflect.TypeOf((*MockMount)(nil).MountSMBPath), arg0, arg1, arg2, arg3, arg4)
}

// PVMountpointMappings mocks base method.
func (m *MockMount) PVMountpointMappings(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PVMountpointMappings", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PVMountpointMappings indicates an expected call of PVMountpointMappings.
func (mr *MockMountMockRecorder) PVMountpointMappings(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PVMountpointMappings", reflect.TypeOf((*MockMount)(nil).PVMountpointMappings), arg0)
}

// RemountDevice mocks base method.
func (m *MockMount) RemountDevice(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemountDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemountDevice indicates an expected call of RemountDevice.
func (mr *MockMountMockRecorder) RemountDevice(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemountDevice", reflect.TypeOf((*MockMount)(nil).RemountDevice), arg0, arg1, arg2)
}

// RemoveMountPoint mocks base method.
func (m *MockMount) RemoveMountPoint(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMountPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMountPoint indicates an expected call of RemoveMountPoint.
func (mr *MockMountMockRecorder) RemoveMountPoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMountPoint", reflect.TypeOf((*MockMount)(nil).RemoveMountPoint), arg0, arg1)
}

// Umount mocks base method.
func (m *MockMount) Umount(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Umount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Umount indicates an expected call of Umount.
func (mr *MockMountMockRecorder) Umount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Umount", reflect.TypeOf((*MockMount)(nil).Umount), arg0, arg1)
}

// UmountAndRemoveMountPoint mocks base method.
func (m *MockMount) UmountAndRemoveMountPoint(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UmountAndRemoveMountPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UmountAndRemoveMountPoint indicates an expected call of UmountAndRemoveMountPoint.
func (mr *MockMountMockRecorder) UmountAndRemoveMountPoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UmountAndRemoveMountPoint", reflect.TypeOf((*MockMount)(nil).UmountAndRemoveMountPoint), arg0, arg1)
}

// UmountAndRemoveTemporaryMountPoint mocks base method.
func (m *MockMount) UmountAndRemoveTemporaryMountPoint(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UmountAndRemoveTemporaryMountPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UmountAndRemoveTemporaryMountPoint indicates an expected call of UmountAndRemoveTemporaryMountPoint.
func (mr *MockMountMockRecorder) UmountAndRemoveTemporaryMountPoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UmountAndRemoveTemporaryMountPoint", reflect.TypeOf((*MockMount)(nil).UmountAndRemoveTemporaryMountPoint), arg0, arg1)
}

// UmountSMBPath mocks base method.
func (m *MockMount) UmountSMBPath(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UmountSMBPath", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UmountSMBPath indicates an expected call of UmountSMBPath.
func (mr *MockMountMockRecorder) UmountSMBPath(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UmountSMBPath", reflect.TypeOf((*MockMount)(nil).UmountSMBPath), arg0, arg1, arg2)
}

// WindowsBindMount mocks base method.
func (m *MockMount) WindowsBindMount(arg0 context.Context, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WindowsBindMount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WindowsBindMount indicates an expected call of WindowsBindMount.
func (mr *MockMountMockRecorder) WindowsBindMount(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WindowsBindMount", reflect.TypeOf((*MockMount)(nil).WindowsBindMount), arg0, arg1, arg2, arg3)
}
