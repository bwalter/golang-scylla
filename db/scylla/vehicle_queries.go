package scylla

import (
	"errors"

	"bwa.com/hello/model"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type VehicleQueries struct {
	session *gocqlx.Session
}

func newVehicleQueries(sessionx *gocqlx.Session) VehicleQueries {
	return VehicleQueries{
		session: sessionx,
	}
}

func (queries *VehicleQueries) CreateVehicle(vehicle model.Vehicle) error {
	scyllaVehicle, err := NewVehicle(vehicle)
	if err != nil {
		return err
	}

	applied, err := queries.session.Query(vehicleTable.InsertBuilder().Unique().ToCql()).BindStruct(scyllaVehicle).ExecCAS()
	if err != nil {
		return err
	}

	if !applied {
		return gocql.RequestErrAlreadyExists{Table: "Vehicle"}
	}

	return nil
}

func (queries *VehicleQueries) FindVehicle(vin string) (*model.Vehicle, error) {
	var findVehicle Vehicle
	findVehicle.Vin = vin

	q := queries.session.Query(vehicleTable.Get()).BindStruct(findVehicle)

	var scyllaVehicle Vehicle
	if err := q.GetRelease(&scyllaVehicle); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	vehicle, err := scyllaVehicle.ToModelVehicle()
	if err != nil {
		return nil, err
	}

	return &vehicle, err
}
