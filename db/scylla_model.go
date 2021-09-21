package db

import (
	"bwa.com/hello/model"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

var vehicleMetadata = table.Metadata{
	Name:    "vehicles",
	Columns: []string{"vin", "engine_type", "ev_data"},
	PartKey: []string{"vin"},
}

var vehicleTable = table.New(vehicleMetadata)

type ScyllaVehicle struct {
	Vin        string
	EngineType string
	EvData     EvDataUDT // does not work as optional value (pointer)
}

type EvDataUDT struct {
	gocqlx.UDT

	BatteryCapacityInKwh int
	SocInPercent         int
}

func NewScyllaVehicle(vehicle model.Vehicle) (ScyllaVehicle, error) {
	var ev_data_udt EvDataUDT
	if vehicle.EvData != nil {
		ev_data_udt = EvDataUDT{
			BatteryCapacityInKwh: vehicle.EvData.BatteryCapacityInKwh,
			SocInPercent:         vehicle.EvData.SocInPercent,
		}
	}

	return ScyllaVehicle{
		Vin:        vehicle.Vin,
		EngineType: vehicle.EngineType,
		EvData:     ev_data_udt,
	}, nil
}

func (sv *ScyllaVehicle) ToModelVehicle() (model.Vehicle, error) {
	var ev_data_ptr *model.EvData
	if sv.EvData.BatteryCapacityInKwh > 0 {
		ev_data_ptr = &model.EvData{
			BatteryCapacityInKwh: sv.EvData.BatteryCapacityInKwh,
			SocInPercent:         sv.EvData.SocInPercent,
		}
	}

	return model.Vehicle{
		Vin:        sv.Vin,
		EngineType: sv.EngineType,
		EvData:     ev_data_ptr,
	}, nil
}
