// Code generated by MockGen. DO NOT EDIT.
// Source: segment_grpc.pb.go

// Package segment is a generated GoMock package.
package segment

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockSegmentServerClient is a mock of SegmentServerClient interface.
type MockSegmentServerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentServerClientMockRecorder
}

// MockSegmentServerClientMockRecorder is the mock recorder for MockSegmentServerClient.
type MockSegmentServerClientMockRecorder struct {
	mock *MockSegmentServerClient
}

// NewMockSegmentServerClient creates a new mock instance.
func NewMockSegmentServerClient(ctrl *gomock.Controller) *MockSegmentServerClient {
	mock := &MockSegmentServerClient{ctrl: ctrl}
	mock.recorder = &MockSegmentServerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegmentServerClient) EXPECT() *MockSegmentServerClientMockRecorder {
	return m.recorder
}

// ActivateSegment mocks base method.
func (m *MockSegmentServerClient) ActivateSegment(ctx context.Context, in *ActivateSegmentRequest, opts ...grpc.CallOption) (*ActivateSegmentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ActivateSegment", varargs...)
	ret0, _ := ret[0].(*ActivateSegmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateSegment indicates an expected call of ActivateSegment.
func (mr *MockSegmentServerClientMockRecorder) ActivateSegment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateSegment", reflect.TypeOf((*MockSegmentServerClient)(nil).ActivateSegment), varargs...)
}

// AppendToBlock mocks base method.
func (m *MockSegmentServerClient) AppendToBlock(ctx context.Context, in *AppendToBlockRequest, opts ...grpc.CallOption) (*AppendToBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AppendToBlock", varargs...)
	ret0, _ := ret[0].(*AppendToBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendToBlock indicates an expected call of AppendToBlock.
func (mr *MockSegmentServerClientMockRecorder) AppendToBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendToBlock", reflect.TypeOf((*MockSegmentServerClient)(nil).AppendToBlock), varargs...)
}

// CreateBlock mocks base method.
func (m *MockSegmentServerClient) CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBlock", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBlock indicates an expected call of CreateBlock.
func (mr *MockSegmentServerClientMockRecorder) CreateBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlock", reflect.TypeOf((*MockSegmentServerClient)(nil).CreateBlock), varargs...)
}

// GetBlockInfo mocks base method.
func (m *MockSegmentServerClient) GetBlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...grpc.CallOption) (*GetBlockInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockInfo", varargs...)
	ret0, _ := ret[0].(*GetBlockInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockInfo indicates an expected call of GetBlockInfo.
func (mr *MockSegmentServerClientMockRecorder) GetBlockInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockInfo", reflect.TypeOf((*MockSegmentServerClient)(nil).GetBlockInfo), varargs...)
}

// InactivateSegment mocks base method.
func (m *MockSegmentServerClient) InactivateSegment(ctx context.Context, in *InactivateSegmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InactivateSegment", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InactivateSegment indicates an expected call of InactivateSegment.
func (mr *MockSegmentServerClientMockRecorder) InactivateSegment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InactivateSegment", reflect.TypeOf((*MockSegmentServerClient)(nil).InactivateSegment), varargs...)
}

// LookupOffsetInBlock mocks base method.
func (m *MockSegmentServerClient) LookupOffsetInBlock(ctx context.Context, in *LookupOffsetInBlockRequest, opts ...grpc.CallOption) (*LookupOffsetInBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LookupOffsetInBlock", varargs...)
	ret0, _ := ret[0].(*LookupOffsetInBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupOffsetInBlock indicates an expected call of LookupOffsetInBlock.
func (mr *MockSegmentServerClientMockRecorder) LookupOffsetInBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupOffsetInBlock", reflect.TypeOf((*MockSegmentServerClient)(nil).LookupOffsetInBlock), varargs...)
}

// ReadFromBlock mocks base method.
func (m *MockSegmentServerClient) ReadFromBlock(ctx context.Context, in *ReadFromBlockRequest, opts ...grpc.CallOption) (*ReadFromBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadFromBlock", varargs...)
	ret0, _ := ret[0].(*ReadFromBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFromBlock indicates an expected call of ReadFromBlock.
func (mr *MockSegmentServerClientMockRecorder) ReadFromBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromBlock", reflect.TypeOf((*MockSegmentServerClient)(nil).ReadFromBlock), varargs...)
}

// RemoveBlock mocks base method.
func (m *MockSegmentServerClient) RemoveBlock(ctx context.Context, in *RemoveBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveBlock", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveBlock indicates an expected call of RemoveBlock.
func (mr *MockSegmentServerClientMockRecorder) RemoveBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBlock", reflect.TypeOf((*MockSegmentServerClient)(nil).RemoveBlock), varargs...)
}

// Start mocks base method.
func (m *MockSegmentServerClient) Start(ctx context.Context, in *StartSegmentServerRequest, opts ...grpc.CallOption) (*StartSegmentServerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(*StartSegmentServerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockSegmentServerClientMockRecorder) Start(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSegmentServerClient)(nil).Start), varargs...)
}

// Status mocks base method.
func (m *MockSegmentServerClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockSegmentServerClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockSegmentServerClient)(nil).Status), varargs...)
}

// Stop mocks base method.
func (m *MockSegmentServerClient) Stop(ctx context.Context, in *StopSegmentServerRequest, opts ...grpc.CallOption) (*StopSegmentServerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stop", varargs...)
	ret0, _ := ret[0].(*StopSegmentServerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockSegmentServerClientMockRecorder) Stop(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSegmentServerClient)(nil).Stop), varargs...)
}

// MockSegmentServerServer is a mock of SegmentServerServer interface.
type MockSegmentServerServer struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentServerServerMockRecorder
}

// MockSegmentServerServerMockRecorder is the mock recorder for MockSegmentServerServer.
type MockSegmentServerServerMockRecorder struct {
	mock *MockSegmentServerServer
}

// NewMockSegmentServerServer creates a new mock instance.
func NewMockSegmentServerServer(ctrl *gomock.Controller) *MockSegmentServerServer {
	mock := &MockSegmentServerServer{ctrl: ctrl}
	mock.recorder = &MockSegmentServerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegmentServerServer) EXPECT() *MockSegmentServerServerMockRecorder {
	return m.recorder
}

// ActivateSegment mocks base method.
func (m *MockSegmentServerServer) ActivateSegment(arg0 context.Context, arg1 *ActivateSegmentRequest) (*ActivateSegmentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateSegment", arg0, arg1)
	ret0, _ := ret[0].(*ActivateSegmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateSegment indicates an expected call of ActivateSegment.
func (mr *MockSegmentServerServerMockRecorder) ActivateSegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateSegment", reflect.TypeOf((*MockSegmentServerServer)(nil).ActivateSegment), arg0, arg1)
}

// AppendToBlock mocks base method.
func (m *MockSegmentServerServer) AppendToBlock(arg0 context.Context, arg1 *AppendToBlockRequest) (*AppendToBlockResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendToBlock", arg0, arg1)
	ret0, _ := ret[0].(*AppendToBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendToBlock indicates an expected call of AppendToBlock.
func (mr *MockSegmentServerServerMockRecorder) AppendToBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendToBlock", reflect.TypeOf((*MockSegmentServerServer)(nil).AppendToBlock), arg0, arg1)
}

// CreateBlock mocks base method.
func (m *MockSegmentServerServer) CreateBlock(arg0 context.Context, arg1 *CreateBlockRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlock", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBlock indicates an expected call of CreateBlock.
func (mr *MockSegmentServerServerMockRecorder) CreateBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlock", reflect.TypeOf((*MockSegmentServerServer)(nil).CreateBlock), arg0, arg1)
}

// GetBlockInfo mocks base method.
func (m *MockSegmentServerServer) GetBlockInfo(arg0 context.Context, arg1 *GetBlockInfoRequest) (*GetBlockInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockInfo", arg0, arg1)
	ret0, _ := ret[0].(*GetBlockInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockInfo indicates an expected call of GetBlockInfo.
func (mr *MockSegmentServerServerMockRecorder) GetBlockInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockInfo", reflect.TypeOf((*MockSegmentServerServer)(nil).GetBlockInfo), arg0, arg1)
}

// InactivateSegment mocks base method.
func (m *MockSegmentServerServer) InactivateSegment(arg0 context.Context, arg1 *InactivateSegmentRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InactivateSegment", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InactivateSegment indicates an expected call of InactivateSegment.
func (mr *MockSegmentServerServerMockRecorder) InactivateSegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InactivateSegment", reflect.TypeOf((*MockSegmentServerServer)(nil).InactivateSegment), arg0, arg1)
}

// LookupOffsetInBlock mocks base method.
func (m *MockSegmentServerServer) LookupOffsetInBlock(arg0 context.Context, arg1 *LookupOffsetInBlockRequest) (*LookupOffsetInBlockResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupOffsetInBlock", arg0, arg1)
	ret0, _ := ret[0].(*LookupOffsetInBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupOffsetInBlock indicates an expected call of LookupOffsetInBlock.
func (mr *MockSegmentServerServerMockRecorder) LookupOffsetInBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupOffsetInBlock", reflect.TypeOf((*MockSegmentServerServer)(nil).LookupOffsetInBlock), arg0, arg1)
}

// ReadFromBlock mocks base method.
func (m *MockSegmentServerServer) ReadFromBlock(arg0 context.Context, arg1 *ReadFromBlockRequest) (*ReadFromBlockResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromBlock", arg0, arg1)
	ret0, _ := ret[0].(*ReadFromBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFromBlock indicates an expected call of ReadFromBlock.
func (mr *MockSegmentServerServerMockRecorder) ReadFromBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromBlock", reflect.TypeOf((*MockSegmentServerServer)(nil).ReadFromBlock), arg0, arg1)
}

// RemoveBlock mocks base method.
func (m *MockSegmentServerServer) RemoveBlock(arg0 context.Context, arg1 *RemoveBlockRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBlock", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveBlock indicates an expected call of RemoveBlock.
func (mr *MockSegmentServerServerMockRecorder) RemoveBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBlock", reflect.TypeOf((*MockSegmentServerServer)(nil).RemoveBlock), arg0, arg1)
}

// Start mocks base method.
func (m *MockSegmentServerServer) Start(arg0 context.Context, arg1 *StartSegmentServerRequest) (*StartSegmentServerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(*StartSegmentServerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockSegmentServerServerMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSegmentServerServer)(nil).Start), arg0, arg1)
}

// Status mocks base method.
func (m *MockSegmentServerServer) Status(arg0 context.Context, arg1 *emptypb.Empty) (*StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockSegmentServerServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockSegmentServerServer)(nil).Status), arg0, arg1)
}

// Stop mocks base method.
func (m *MockSegmentServerServer) Stop(arg0 context.Context, arg1 *StopSegmentServerRequest) (*StopSegmentServerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(*StopSegmentServerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockSegmentServerServerMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSegmentServerServer)(nil).Stop), arg0, arg1)
}

// MockUnsafeSegmentServerServer is a mock of UnsafeSegmentServerServer interface.
type MockUnsafeSegmentServerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSegmentServerServerMockRecorder
}

// MockUnsafeSegmentServerServerMockRecorder is the mock recorder for MockUnsafeSegmentServerServer.
type MockUnsafeSegmentServerServerMockRecorder struct {
	mock *MockUnsafeSegmentServerServer
}

// NewMockUnsafeSegmentServerServer creates a new mock instance.
func NewMockUnsafeSegmentServerServer(ctrl *gomock.Controller) *MockUnsafeSegmentServerServer {
	mock := &MockUnsafeSegmentServerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSegmentServerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSegmentServerServer) EXPECT() *MockUnsafeSegmentServerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSegmentServerServer mocks base method.
func (m *MockUnsafeSegmentServerServer) mustEmbedUnimplementedSegmentServerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSegmentServerServer")
}

// mustEmbedUnimplementedSegmentServerServer indicates an expected call of mustEmbedUnimplementedSegmentServerServer.
func (mr *MockUnsafeSegmentServerServerMockRecorder) mustEmbedUnimplementedSegmentServerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSegmentServerServer", reflect.TypeOf((*MockUnsafeSegmentServerServer)(nil).mustEmbedUnimplementedSegmentServerServer))
}