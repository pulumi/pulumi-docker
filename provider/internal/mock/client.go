// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -package mock -source client.go -destination mock/client.go
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	pb "github.com/docker/buildx/controller/pb"
	types "github.com/docker/cli/cli/manifest/types"
	types0 "github.com/docker/docker/api/types"
	client "github.com/moby/buildkit/client"
	properties "github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockClient) Auth(ctx context.Context, creds properties.RegistryAuth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", ctx, creds)
	ret0, _ := ret[0].(error)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MockClientMockRecorder) Auth(ctx, creds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockClient)(nil).Auth), ctx, creds)
}

// Build mocks base method.
func (m *MockClient) Build(ctx context.Context, opts pb.BuildOptions) (*client.SolveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, opts)
	ret0, _ := ret[0].(*client.SolveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockClientMockRecorder) Build(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockClient)(nil).Build), ctx, opts)
}

// BuildKitEnabled mocks base method.
func (m *MockClient) BuildKitEnabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildKitEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildKitEnabled indicates an expected call of BuildKitEnabled.
func (mr *MockClientMockRecorder) BuildKitEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildKitEnabled", reflect.TypeOf((*MockClient)(nil).BuildKitEnabled))
}

// Delete mocks base method.
func (m *MockClient) Delete(ctx context.Context, id string) ([]types0.ImageDeleteResponseItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].([]types0.ImageDeleteResponseItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), ctx, id)
}

// Inspect mocks base method.
func (m *MockClient) Inspect(ctx context.Context, id string) ([]types.ImageManifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", ctx, id)
	ret0, _ := ret[0].([]types.ImageManifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect.
func (mr *MockClientMockRecorder) Inspect(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockClient)(nil).Inspect), ctx, id)
}
