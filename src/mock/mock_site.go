// Code generated by MockGen. DO NOT EDIT.
// Source: site.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	entities "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISiteRepository is a mock of ISiteRepository interface.
type MockISiteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISiteRepositoryMockRecorder
}

// MockISiteRepositoryMockRecorder is the mock recorder for MockISiteRepository.
type MockISiteRepositoryMockRecorder struct {
	mock *MockISiteRepository
}

// NewMockISiteRepository creates a new mock instance.
func NewMockISiteRepository(ctrl *gomock.Controller) *MockISiteRepository {
	mock := &MockISiteRepository{ctrl: ctrl}
	mock.recorder = &MockISiteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISiteRepository) EXPECT() *MockISiteRepositoryMockRecorder {
	return m.recorder
}

// CheckGPC mocks base method.
func (m *MockISiteRepository) CheckGPC(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckGPC", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckGPC indicates an expected call of CheckGPC.
func (mr *MockISiteRepositoryMockRecorder) CheckGPC(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckGPC", reflect.TypeOf((*MockISiteRepository)(nil).CheckGPC), arg0)
}

// Exists mocks base method.
func (m *MockISiteRepository) Exists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockISiteRepositoryMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockISiteRepository)(nil).Exists), arg0)
}

// FindByURL mocks base method.
func (m *MockISiteRepository) FindByURL(arg0 string) (entities.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByURL", arg0)
	ret0, _ := ret[0].(entities.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByURL indicates an expected call of FindByURL.
func (mr *MockISiteRepositoryMockRecorder) FindByURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByURL", reflect.TypeOf((*MockISiteRepository)(nil).FindByURL), arg0)
}

// Store mocks base method.
func (m *MockISiteRepository) Store(arg0 entities.Site) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockISiteRepositoryMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockISiteRepository)(nil).Store), arg0)
}
