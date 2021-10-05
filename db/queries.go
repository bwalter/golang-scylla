package db

import (
	"bwa.com/hello/model"
)

type Queries interface {
	CreateTablesIfNotExist() error
	CloseSession()

	VehicleQueries() VehicleQueries
}

type VehicleQueries interface {
	CreateVehicle(vehicle model.Vehicle) error
	FindVehicle(vin string) (*model.Vehicle, error)
}
