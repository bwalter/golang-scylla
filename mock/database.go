// Code generated by MockGen. DO NOT EDIT.
// Source: db/database.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	db "bwa.com/hello/db"
	model "bwa.com/hello/model"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CloseSession mocks base method.
func (m *MockDatabase) CloseSession() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseSession")
}

// CloseSession indicates an expected call of CloseSession.
func (mr *MockDatabaseMockRecorder) CloseSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSession", reflect.TypeOf((*MockDatabase)(nil).CloseSession))
}

// CreateTablesIfNotExist mocks base method.
func (m *MockDatabase) CreateTablesIfNotExist() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTablesIfNotExist")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTablesIfNotExist indicates an expected call of CreateTablesIfNotExist.
func (mr *MockDatabaseMockRecorder) CreateTablesIfNotExist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTablesIfNotExist", reflect.TypeOf((*MockDatabase)(nil).CreateTablesIfNotExist))
}

// VehicleDAO mocks base method.
func (m *MockDatabase) VehicleDAO() db.VehicleDAO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VehicleDAO")
	ret0, _ := ret[0].(db.VehicleDAO)
	return ret0
}

// VehicleDAO indicates an expected call of VehicleDAO.
func (mr *MockDatabaseMockRecorder) VehicleDAO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VehicleDAO", reflect.TypeOf((*MockDatabase)(nil).VehicleDAO))
}

// MockVehicleDAO is a mock of VehicleDAO interface.
type MockVehicleDAO struct {
	ctrl     *gomock.Controller
	recorder *MockVehicleDAOMockRecorder
}

// MockVehicleDAOMockRecorder is the mock recorder for MockVehicleDAO.
type MockVehicleDAOMockRecorder struct {
	mock *MockVehicleDAO
}

// NewMockVehicleDAO creates a new mock instance.
func NewMockVehicleDAO(ctrl *gomock.Controller) *MockVehicleDAO {
	mock := &MockVehicleDAO{ctrl: ctrl}
	mock.recorder = &MockVehicleDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVehicleDAO) EXPECT() *MockVehicleDAOMockRecorder {
	return m.recorder
}

// CreateVehicle mocks base method.
func (m *MockVehicleDAO) CreateVehicle(vehicle model.Vehicle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVehicle", vehicle)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVehicle indicates an expected call of CreateVehicle.
func (mr *MockVehicleDAOMockRecorder) CreateVehicle(vehicle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVehicle", reflect.TypeOf((*MockVehicleDAO)(nil).CreateVehicle), vehicle)
}

// FindVehicle mocks base method.
func (m *MockVehicleDAO) FindVehicle(vin string) (*model.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVehicle", vin)
	ret0, _ := ret[0].(*model.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVehicle indicates an expected call of FindVehicle.
func (mr *MockVehicleDAOMockRecorder) FindVehicle(vin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVehicle", reflect.TypeOf((*MockVehicleDAO)(nil).FindVehicle), vin)
}