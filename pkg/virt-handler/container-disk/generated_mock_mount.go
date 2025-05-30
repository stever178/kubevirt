// Code generated by MockGen. DO NOT EDIT.
// Source: mount.go
//
// Generated by this command:
//
//	mockgen -source mount.go -package=container_disk -destination=generated_mock_mount.go
//

// Package container_disk is a generated GoMock package.
package container_disk

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
	v1 "kubevirt.io/api/core/v1"
)

// MockMounter is a mock of Mounter interface.
type MockMounter struct {
	ctrl     *gomock.Controller
	recorder *MockMounterMockRecorder
	isgomock struct{}
}

// MockMounterMockRecorder is the mock recorder for MockMounter.
type MockMounterMockRecorder struct {
	mock *MockMounter
}

// NewMockMounter creates a new mock instance.
func NewMockMounter(ctrl *gomock.Controller) *MockMounter {
	mock := &MockMounter{ctrl: ctrl}
	mock.recorder = &MockMounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMounter) EXPECT() *MockMounterMockRecorder {
	return m.recorder
}

// ComputeChecksums mocks base method.
func (m *MockMounter) ComputeChecksums(vmi *v1.VirtualMachineInstance) (*DiskChecksums, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeChecksums", vmi)
	ret0, _ := ret[0].(*DiskChecksums)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeChecksums indicates an expected call of ComputeChecksums.
func (mr *MockMounterMockRecorder) ComputeChecksums(vmi any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeChecksums", reflect.TypeOf((*MockMounter)(nil).ComputeChecksums), vmi)
}

// ContainerDisksReady mocks base method.
func (m *MockMounter) ContainerDisksReady(vmi *v1.VirtualMachineInstance, notInitializedSince time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerDisksReady", vmi, notInitializedSince)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerDisksReady indicates an expected call of ContainerDisksReady.
func (mr *MockMounterMockRecorder) ContainerDisksReady(vmi, notInitializedSince any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerDisksReady", reflect.TypeOf((*MockMounter)(nil).ContainerDisksReady), vmi, notInitializedSince)
}

// MountAndVerify mocks base method.
func (m *MockMounter) MountAndVerify(vmi *v1.VirtualMachineInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountAndVerify", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountAndVerify indicates an expected call of MountAndVerify.
func (mr *MockMounterMockRecorder) MountAndVerify(vmi any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountAndVerify", reflect.TypeOf((*MockMounter)(nil).MountAndVerify), vmi)
}

// Unmount mocks base method.
func (m *MockMounter) Unmount(vmi *v1.VirtualMachineInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount.
func (mr *MockMounterMockRecorder) Unmount(vmi any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockMounter)(nil).Unmount), vmi)
}
