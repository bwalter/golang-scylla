package db

import (
	"fmt"
	"strings"

	"bwa.com/hello/model"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type ScyllaQueries struct {
	session *gocqlx.Session
}

func NewQueries(session *gocql.Session) (Queries, error) {
	sessionx, err := gocqlx.WrapSession(session, nil)
	if err != nil {
		return nil, err
	}

	return &ScyllaQueries{
		session: &sessionx,
	}, nil
}

func (queries *ScyllaQueries) CreateTablesIfNotExist() error {
	if err := queries.session.ExecStmt("CREATE KEYSPACE IF NOT EXISTS hello WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1}"); err != nil {
		return err
	}

	if err := queries.session.ExecStmt("CREATE TYPE IF NOT EXISTS ev_data (battery_capacity_in_kwh int, soc_in_percent int)"); err != nil {
		return err
	}

	cql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", vehicleMetadata.Name, strings.Join(vehicleMetadata.Columns, ","))
	if err := queries.session.ExecStmt(cql); err != nil {
		return err
	}

	return nil
}

func (queries *ScyllaQueries) CreateVehicle(vehicle model.Vehicle) error {
	scylla_vehicle, err := NewScyllaVehicle(vehicle)
	if err != nil {
		return err
	}

	q := queries.session.Query(vehicleTable.InsertBuilder().Unique().ToCql()).BindStruct(scylla_vehicle)
	applied, err := q.ExecCASRelease()
	if err != nil {
		return err
	}

	if !applied {
		return gocql.RequestErrAlreadyExists{}
	}

	return nil
}

func (queries *ScyllaQueries) FindVehicle(vin string) (*model.Vehicle, error) {
	q := queries.session.Query(vehicleTable.Get()).BindStruct(&ScyllaVehicle{
		Vin: vin,
	})

	scylla_vehicle := ScyllaVehicle{}
	if err := q.GetRelease(&scylla_vehicle); err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}

		return nil, err
	}

	vehicle, err := scylla_vehicle.ToModelVehicle()
	if err != nil {
		return nil, err
	}

	return &vehicle, err
}
