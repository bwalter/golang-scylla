package db

import (
	"fmt"
	"time"

	"bwa.com/hello/model"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type ScyllaQueries struct {
	session *gocqlx.Session
}

func CreateScyllaKeyspace(host string, keyspace string, deleteExisting bool) error {
	cluster := gocql.NewCluster(host)
	cluster.Timeout = 30 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	defer session.Close()

	if deleteExisting {
		if err := session.Query(fmt.Sprintf("DROP KEYSPACE IF EXISTS %s", keyspace)).Exec(); err != nil {
			return err
		}
	}

	cql := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1}", keyspace)
	if err := session.Query(cql).Exec(); err != nil {
		return err
	}

	return nil
}

func StartScyllaSessionAndCreateQueries(host string, keyspace string) (Queries, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	queries, err := newScyllaQueries(session)
	if err != nil {
		return nil, nil
	}

	if err := queries.CreateTablesIfNotExist(); err != nil {
		return nil, err
	}

	return queries, nil
}

func newScyllaQueries(session *gocql.Session) (Queries, error) {
	sessionx, err := gocqlx.WrapSession(session, nil)
	if err != nil {
		return nil, err
	}

	return &ScyllaQueries{
		session: &sessionx,
	}, nil
}

func (queries *ScyllaQueries) CreateTablesIfNotExist() error {
	if err := queries.session.ExecStmt("CREATE TYPE IF NOT EXISTS ev_data (battery_capacity_in_kwh int, soc_in_percent int)"); err != nil {
		return err
	}

	cql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (vin text primary key, engine_type text, ev_data ev_data)", vehicleMetadata.Name)
	if err := queries.session.ExecStmt(cql); err != nil {
		return err
	}

	return nil
}

func (queries *ScyllaQueries) CloseSession() {
	queries.session.Close()
}

func (queries *ScyllaQueries) CreateVehicle(vehicle model.Vehicle) error {
	scylla_vehicle, err := NewScyllaVehicle(vehicle)
	if err != nil {
		return err
	}

	applied, err := queries.session.Query(vehicleTable.InsertBuilder().Unique().ToCql()).BindStruct(scylla_vehicle).ExecCAS()
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
