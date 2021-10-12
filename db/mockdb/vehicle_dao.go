package mockdb

import (
	"bwa.com/hello/model"
)

type VehicleDAO struct {
	vehicles map[string]model.Vehicle
}

func newVehicleDAO() VehicleDAO {
	return VehicleDAO{
		vehicles: make(map[string]model.Vehicle),
	}
}

func (dao *VehicleDAO) CreateVehicle(vehicle model.Vehicle) error {
	dao.vehicles[vehicle.Vin] = vehicle
	return nil
}

func (dao *VehicleDAO) FindVehicle(vin string) (*model.Vehicle, error) {
	vehicle, exists := dao.vehicles[vin]
	if !exists {
		return nil, nil
	}

	return &vehicle, nil
}
