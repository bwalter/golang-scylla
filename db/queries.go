package db

import (
	"bwa.com/hello/model"
)

type Queries interface {
	CreateTablesIfNotExist() error
	CloseSession()
	CreateVehicle(vehicle model.Vehicle) error
	FindVehicle(vin string) (*model.Vehicle, error)
}
