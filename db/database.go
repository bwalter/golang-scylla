package db

import (
	"bwa.com/hello/model"
)

type Database interface {
	CreateTablesIfNotExist() error
	CloseSession()

	VehicleDAO() VehicleDAO
}

type VehicleDAO interface {
	CreateVehicle(vehicle model.Vehicle) error
	FindVehicle(vin string) (*model.Vehicle, error)
}
