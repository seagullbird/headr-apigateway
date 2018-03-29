// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/seagullbird/headr-repoctl/service (interfaces: Service)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DeleteSite mocks base method
func (m *MockService) DeleteSite(arg0 context.Context, arg1 uint) error {
	ret := m.ctrl.Call(m, "DeleteSite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSite indicates an expected call of DeleteSite
func (mr *MockServiceMockRecorder) DeleteSite(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSite", reflect.TypeOf((*MockService)(nil).DeleteSite), arg0, arg1)
}

// NewSite mocks base method
func (m *MockService) NewSite(arg0 context.Context, arg1 uint, arg2 string) error {
	ret := m.ctrl.Call(m, "NewSite", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewSite indicates an expected call of NewSite
func (mr *MockServiceMockRecorder) NewSite(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSite", reflect.TypeOf((*MockService)(nil).NewSite), arg0, arg1, arg2)
}

// ReadConfig mocks base method
func (m *MockService) ReadConfig(arg0 context.Context, arg1 uint) (string, error) {
	ret := m.ctrl.Call(m, "ReadConfig", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadConfig indicates an expected call of ReadConfig
func (mr *MockServiceMockRecorder) ReadConfig(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConfig", reflect.TypeOf((*MockService)(nil).ReadConfig), arg0, arg1)
}

// ReadPost mocks base method
func (m *MockService) ReadPost(arg0 context.Context, arg1 uint, arg2 string) (string, error) {
	ret := m.ctrl.Call(m, "ReadPost", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPost indicates an expected call of ReadPost
func (mr *MockServiceMockRecorder) ReadPost(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPost", reflect.TypeOf((*MockService)(nil).ReadPost), arg0, arg1, arg2)
}

// RemovePost mocks base method
func (m *MockService) RemovePost(arg0 context.Context, arg1 uint, arg2 string) error {
	ret := m.ctrl.Call(m, "RemovePost", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePost indicates an expected call of RemovePost
func (mr *MockServiceMockRecorder) RemovePost(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePost", reflect.TypeOf((*MockService)(nil).RemovePost), arg0, arg1, arg2)
}

// WriteConfig mocks base method
func (m *MockService) WriteConfig(arg0 context.Context, arg1 uint, arg2 string) error {
	ret := m.ctrl.Call(m, "WriteConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteConfig indicates an expected call of WriteConfig
func (mr *MockServiceMockRecorder) WriteConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteConfig", reflect.TypeOf((*MockService)(nil).WriteConfig), arg0, arg1, arg2)
}

// WritePost mocks base method
func (m *MockService) WritePost(arg0 context.Context, arg1 uint, arg2, arg3 string) error {
	ret := m.ctrl.Call(m, "WritePost", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePost indicates an expected call of WritePost
func (mr *MockServiceMockRecorder) WritePost(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePost", reflect.TypeOf((*MockService)(nil).WritePost), arg0, arg1, arg2, arg3)
}
