// Code generated by MockGen. DO NOT EDIT.
// Source: offset.go

// Package offset is a generated GoMock package.
package offset

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	info "github.com/vanus-labs/vanus/internal/primitive/info"
	vanus "github.com/vanus-labs/vanus/internal/primitive/vanus"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// GetOffset mocks base method.
func (m *MockManager) GetOffset(ctx context.Context, subscriptionID vanus.ID) (info.ListOffsetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffset", ctx, subscriptionID)
	ret0, _ := ret[0].(info.ListOffsetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffset indicates an expected call of GetOffset.
func (mr *MockManagerMockRecorder) GetOffset(ctx, subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffset", reflect.TypeOf((*MockManager)(nil).GetOffset), ctx, subscriptionID)
}

// Offset mocks base method.
func (m *MockManager) Offset(ctx context.Context, subscriptionID vanus.ID, offsets info.ListOffsetInfo, commit bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", ctx, subscriptionID, offsets, commit)
	ret0, _ := ret[0].(error)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *MockManagerMockRecorder) Offset(ctx, subscriptionID, offsets, commit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockManager)(nil).Offset), ctx, subscriptionID, offsets, commit)
}

// RemoveRegisterSubscription mocks base method.
func (m *MockManager) RemoveRegisterSubscription(ctx context.Context, id vanus.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRegisterSubscription", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRegisterSubscription indicates an expected call of RemoveRegisterSubscription.
func (mr *MockManagerMockRecorder) RemoveRegisterSubscription(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRegisterSubscription", reflect.TypeOf((*MockManager)(nil).RemoveRegisterSubscription), ctx, id)
}

// Start mocks base method.
func (m *MockManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop))
}
