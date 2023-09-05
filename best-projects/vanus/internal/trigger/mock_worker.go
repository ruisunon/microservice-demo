// Code generated by MockGen. DO NOT EDIT.
// Source: worker.go

// Package trigger is a generated GoMock package.
package trigger

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "github.com/vanus-labs/vanus/internal/primitive"
	vanus "github.com/vanus-labs/vanus/internal/primitive/vanus"
)

// MockWorker is a mock of Worker interface.
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker.
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance.
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// AddSubscription mocks base method.
func (m *MockWorker) AddSubscription(ctx context.Context, subscription *primitive.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscription", ctx, subscription)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSubscription indicates an expected call of AddSubscription.
func (mr *MockWorkerMockRecorder) AddSubscription(ctx, subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscription", reflect.TypeOf((*MockWorker)(nil).AddSubscription), ctx, subscription)
}

// Init mocks base method.
func (m *MockWorker) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockWorkerMockRecorder) Init(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockWorker)(nil).Init), ctx)
}

// PauseSubscription mocks base method.
func (m *MockWorker) PauseSubscription(ctx context.Context, id vanus.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseSubscription", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// PauseSubscription indicates an expected call of PauseSubscription.
func (mr *MockWorkerMockRecorder) PauseSubscription(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseSubscription", reflect.TypeOf((*MockWorker)(nil).PauseSubscription), ctx, id)
}

// Register mocks base method.
func (m *MockWorker) Register(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockWorkerMockRecorder) Register(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockWorker)(nil).Register), ctx)
}

// RemoveSubscription mocks base method.
func (m *MockWorker) RemoveSubscription(ctx context.Context, id vanus.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSubscription", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSubscription indicates an expected call of RemoveSubscription.
func (mr *MockWorkerMockRecorder) RemoveSubscription(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubscription", reflect.TypeOf((*MockWorker)(nil).RemoveSubscription), ctx, id)
}

// Start mocks base method.
func (m *MockWorker) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockWorkerMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWorker)(nil).Start), ctx)
}

// StartSubscription mocks base method.
func (m *MockWorker) StartSubscription(ctx context.Context, id vanus.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSubscription", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartSubscription indicates an expected call of StartSubscription.
func (mr *MockWorkerMockRecorder) StartSubscription(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSubscription", reflect.TypeOf((*MockWorker)(nil).StartSubscription), ctx, id)
}

// Stop mocks base method.
func (m *MockWorker) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockWorkerMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWorker)(nil).Stop), ctx)
}

// Unregister mocks base method.
func (m *MockWorker) Unregister(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister.
func (mr *MockWorkerMockRecorder) Unregister(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockWorker)(nil).Unregister), ctx)
}
