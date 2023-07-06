// Code generated by MockGen. DO NOT EDIT.
// Source: uptime_checks.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockUptimeChecksService is a mock of UptimeChecksService interface.
type MockUptimeChecksService struct {
	ctrl     *gomock.Controller
	recorder *MockUptimeChecksServiceMockRecorder
}

// MockUptimeChecksServiceMockRecorder is the mock recorder for MockUptimeChecksService.
type MockUptimeChecksServiceMockRecorder struct {
	mock *MockUptimeChecksService
}

// NewMockUptimeChecksService creates a new mock instance.
func NewMockUptimeChecksService(ctrl *gomock.Controller) *MockUptimeChecksService {
	mock := &MockUptimeChecksService{ctrl: ctrl}
	mock.recorder = &MockUptimeChecksServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUptimeChecksService) EXPECT() *MockUptimeChecksServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUptimeChecksService) Create(arg0 *godo.CreateUptimeCheckRequest) (*do.UptimeCheck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*do.UptimeCheck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUptimeChecksServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUptimeChecksService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockUptimeChecksService) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUptimeChecksServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUptimeChecksService)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockUptimeChecksService) Get(arg0 string) (*do.UptimeCheck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*do.UptimeCheck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUptimeChecksServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUptimeChecksService)(nil).Get), arg0)
}

// GetState mocks base method.
func (m *MockUptimeChecksService) GetState(arg0 string) (*do.UptimeCheckState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", arg0)
	ret0, _ := ret[0].(*do.UptimeCheckState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockUptimeChecksServiceMockRecorder) GetState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockUptimeChecksService)(nil).GetState), arg0)
}

// List mocks base method.
func (m *MockUptimeChecksService) List() ([]do.UptimeCheck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]do.UptimeCheck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUptimeChecksServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUptimeChecksService)(nil).List))
}

// Update mocks base method.
func (m *MockUptimeChecksService) Update(arg0 string, arg1 *godo.UpdateUptimeCheckRequest) (*do.UptimeCheck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*do.UptimeCheck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUptimeChecksServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUptimeChecksService)(nil).Update), arg0, arg1)
}