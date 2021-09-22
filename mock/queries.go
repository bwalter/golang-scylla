// Code generated by MockGen. DO NOT EDIT.
// Source: db/queries.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "bwa.com/hello/model"
	gomock "github.com/golang/mock/gomock"
)

// MockQueries is a mock of Queries interface.
type MockQueries struct {
	ctrl     *gomock.Controller
	recorder *MockQueriesMockRecorder
}

// MockQueriesMockRecorder is the mock recorder for MockQueries.
type MockQueriesMockRecorder struct {
	mock *MockQueries
}

// NewMockQueries creates a new mock instance.
func NewMockQueries(ctrl *gomock.Controller) *MockQueries {
	mock := &MockQueries{ctrl: ctrl}
	mock.recorder = &MockQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueries) EXPECT() *MockQueriesMockRecorder {
	return m.recorder
}

// CreateTablesIfNotExist mocks base method.
func (m *MockQueries) CreateTablesIfNotExist() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTablesIfNotExist")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTablesIfNotExist indicates an expected call of CreateTablesIfNotExist.
func (mr *MockQueriesMockRecorder) CreateTablesIfNotExist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTablesIfNotExist", reflect.TypeOf((*MockQueries)(nil).CreateTablesIfNotExist))
}

// CreateVehicle mocks base method.
func (m *MockQueries) CreateVehicle(vehicle model.Vehicle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVehicle", vehicle)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVehicle indicates an expected call of CreateVehicle.
func (mr *MockQueriesMockRecorder) CreateVehicle(vehicle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVehicle", reflect.TypeOf((*MockQueries)(nil).CreateVehicle), vehicle)
}

// FindVehicle mocks base method.
func (m *MockQueries) FindVehicle(vin string) (*model.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVehicle", vin)
	ret0, _ := ret[0].(*model.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVehicle indicates an expected call of FindVehicle.
func (mr *MockQueriesMockRecorder) FindVehicle(vin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVehicle", reflect.TypeOf((*MockQueries)(nil).FindVehicle), vin)
}