package scylla

import (
	"errors"

	"bwa.com/hello/model"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type VehicleDAO struct {
	session *gocqlx.Session
}

func newVehicleDAO(sessionx *gocqlx.Session) VehicleDAO {
	return VehicleDAO{
		session: sessionx,
	}
}

func (dao *VehicleDAO) CreateVehicle(vehicle model.Vehicle) error {
	scyllaVehicle, err := NewVehicle(vehicle)
	if err != nil {
		return err
	}

	applied, err := dao.session.Query(vehicleTable.InsertBuilder().Unique().ToCql()).BindStruct(scyllaVehicle).ExecCAS()
	if err != nil {
		return err
	}

	if !applied {
		return gocql.RequestErrAlreadyExists{Table: "Vehicle"}
	}

	return nil
}

func (dao *VehicleDAO) FindVehicle(vin string) (*model.Vehicle, error) {
	var findVehicle Vehicle
	findVehicle.Vin = vin

	q := dao.session.Query(vehicleTable.Get()).BindStruct(findVehicle)

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
