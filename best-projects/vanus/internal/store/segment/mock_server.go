// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package segment is a generated GoMock package.
package segment

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "github.com/vanus-labs/vanus/internal/primitive"
	vanus "github.com/vanus-labs/vanus/internal/primitive/vanus"
	cloudevents "github.com/vanus-labs/vanus/proto/pkg/cloudevents"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// ActivateSegment mocks base method.
func (m *MockServer) ActivateSegment(ctx context.Context, logID, segID vanus.ID, replicas map[vanus.ID]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateSegment", ctx, logID, segID, replicas)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivateSegment indicates an expected call of ActivateSegment.
func (mr *MockServerMockRecorder) ActivateSegment(ctx, logID, segID, replicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateSegment", reflect.TypeOf((*MockServer)(nil).ActivateSegment), ctx, logID, segID, replicas)
}

// AppendToBlock mocks base method.
func (m *MockServer) AppendToBlock(ctx context.Context, id vanus.ID, events []*cloudevents.CloudEvent) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendToBlock", ctx, id, events)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendToBlock indicates an expected call of AppendToBlock.
func (mr *MockServerMockRecorder) AppendToBlock(ctx, id, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendToBlock", reflect.TypeOf((*MockServer)(nil).AppendToBlock), ctx, id, events)
}

// CreateBlock mocks base method.
func (m *MockServer) CreateBlock(ctx context.Context, id vanus.ID, size int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlock", ctx, id, size)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBlock indicates an expected call of CreateBlock.
func (mr *MockServerMockRecorder) CreateBlock(ctx, id, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlock", reflect.TypeOf((*MockServer)(nil).CreateBlock), ctx, id, size)
}

// InactivateSegment mocks base method.
func (m *MockServer) InactivateSegment(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InactivateSegment", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InactivateSegment indicates an expected call of InactivateSegment.
func (mr *MockServerMockRecorder) InactivateSegment(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InactivateSegment", reflect.TypeOf((*MockServer)(nil).InactivateSegment), ctx)
}

// Initialize mocks base method.
func (m *MockServer) Initialize(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockServerMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockServer)(nil).Initialize), arg0)
}

// LookupOffsetInBlock mocks base method.
func (m *MockServer) LookupOffsetInBlock(ctx context.Context, id vanus.ID, stime int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupOffsetInBlock", ctx, id, stime)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupOffsetInBlock indicates an expected call of LookupOffsetInBlock.
func (mr *MockServerMockRecorder) LookupOffsetInBlock(ctx, id, stime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupOffsetInBlock", reflect.TypeOf((*MockServer)(nil).LookupOffsetInBlock), ctx, id, stime)
}

// ReadFromBlock mocks base method.
func (m *MockServer) ReadFromBlock(ctx context.Context, id vanus.ID, seq int64, num int, pollingTimeout uint32) ([]*cloudevents.CloudEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromBlock", ctx, id, seq, num, pollingTimeout)
	ret0, _ := ret[0].([]*cloudevents.CloudEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFromBlock indicates an expected call of ReadFromBlock.
func (mr *MockServerMockRecorder) ReadFromBlock(ctx, id, seq, num, pollingTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromBlock", reflect.TypeOf((*MockServer)(nil).ReadFromBlock), ctx, id, seq, num, pollingTimeout)
}

// RegisterToController mocks base method.
func (m *MockServer) RegisterToController(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterToController", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterToController indicates an expected call of RegisterToController.
func (mr *MockServerMockRecorder) RegisterToController(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterToController", reflect.TypeOf((*MockServer)(nil).RegisterToController), ctx)
}

// RemoveBlock mocks base method.
func (m *MockServer) RemoveBlock(ctx context.Context, id vanus.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBlock", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBlock indicates an expected call of RemoveBlock.
func (mr *MockServerMockRecorder) RemoveBlock(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBlock", reflect.TypeOf((*MockServer)(nil).RemoveBlock), ctx, id)
}

// Serve mocks base method.
func (m *MockServer) Serve(lis net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve", lis)
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockServerMockRecorder) Serve(lis interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockServer)(nil).Serve), lis)
}

// Start mocks base method.
func (m *MockServer) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockServerMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockServer)(nil).Start), ctx)
}

// Status mocks base method.
func (m *MockServer) Status() primitive.ServerState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(primitive.ServerState)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockServerMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockServer)(nil).Status))
}

// Stop mocks base method.
func (m *MockServer) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockServerMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockServer)(nil).Stop), ctx)
}