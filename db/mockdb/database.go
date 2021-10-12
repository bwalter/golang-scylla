package mockdb

import (
	"bwa.com/hello/db"
)

type Database struct {
	vehicleDAO VehicleDAO
}

func NewDatabase() db.Database {
	return &Database{vehicleDAO: newVehicleDAO()}
}

func (db *Database) CloseSession() {
}

func (db *Database) VehicleDAO() db.VehicleDAO {
	return &db.vehicleDAO
}
