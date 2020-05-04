// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	secure "github.com/sysdiglabs/harbor-scanner-sysdig-secure/pkg/secure"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddImage mocks base method
func (m *MockClient) AddImage(image string, force bool) (secure.ScanResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImage", image, force)
	ret0, _ := ret[0].(secure.ScanResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddImage indicates an expected call of AddImage
func (mr *MockClientMockRecorder) AddImage(image, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImage", reflect.TypeOf((*MockClient)(nil).AddImage), image, force)
}

// GetImage mocks base method
func (m *MockClient) GetImage(shaDigest string) (secure.ScanResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", shaDigest)
	ret0, _ := ret[0].(secure.ScanResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImage indicates an expected call of GetImage
func (mr *MockClientMockRecorder) GetImage(shaDigest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockClient)(nil).GetImage), shaDigest)
}

// GetVulnerabilities mocks base method
func (m *MockClient) GetVulnerabilities(shaDigest string) (secure.VulnerabilityReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVulnerabilities", shaDigest)
	ret0, _ := ret[0].(secure.VulnerabilityReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVulnerabilities indicates an expected call of GetVulnerabilities
func (mr *MockClientMockRecorder) GetVulnerabilities(shaDigest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVulnerabilities", reflect.TypeOf((*MockClient)(nil).GetVulnerabilities), shaDigest)
}

// AddRegistry mocks base method
func (m *MockClient) AddRegistry(registry, user, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRegistry", registry, user, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRegistry indicates an expected call of AddRegistry
func (mr *MockClientMockRecorder) AddRegistry(registry, user, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRegistry", reflect.TypeOf((*MockClient)(nil).AddRegistry), registry, user, password)
}

// UpdateRegistry mocks base method
func (m *MockClient) UpdateRegistry(registry, user, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegistry", registry, user, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRegistry indicates an expected call of UpdateRegistry
func (mr *MockClientMockRecorder) UpdateRegistry(registry, user, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegistry", reflect.TypeOf((*MockClient)(nil).UpdateRegistry), registry, user, password)
}

// DeleteRegistry mocks base method
func (m *MockClient) DeleteRegistry(registry string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRegistry", registry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRegistry indicates an expected call of DeleteRegistry
func (mr *MockClientMockRecorder) DeleteRegistry(registry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistry", reflect.TypeOf((*MockClient)(nil).DeleteRegistry), registry)
}
