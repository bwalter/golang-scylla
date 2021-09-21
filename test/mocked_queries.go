// TODO: generated...

package test

import (
	"bwa.com/hello/model"
)

type MockedQueries struct {
}

func (queries *MockedQueries) CreateTablesIfNotExist() error {
	return nil
}

func (queries *MockedQueries) CreateVehicle(vehicle model.Vehicle) error {
	return nil
}

func (queries *MockedQueries) FindVehicle(vin string) (*model.Vehicle, error) {
	var vehicle model.Vehicle

	return &vehicle, nil
}
