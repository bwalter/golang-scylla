package db

import (
	"bwa.com/hello/model"
	"github.com/gocql/gocql"
)

type ScyllaQueries struct {
	session gocql.Session
}

func NewQueries(session *gocql.Session) Queries {
	return &ScyllaQueries{
		session: *session,
	}
}

func (queries *ScyllaQueries) CreateTablesIfNotExist() error {
	if err := queries.session.Query("CREATE KEYSPACE IF NOT EXISTS hello WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1}").Exec(); err != nil {
		return err
	}

	if err := queries.session.Query("CREATE TYPE IF NOT EXISTS ev_data (battery_capacity_in_kwh int, soc_in_percent int)").Exec(); err != nil {
		return err
	}

	if err := queries.session.Query("CREATE TABLE IF NOT EXISTS vehicles (vin text primary key, engine_type text, ev_data ev_data)").Exec(); err != nil {
		return err
	}

	return nil
}

func (queries *ScyllaQueries) CreateVehicle(vehicle model.Vehicle) error {
	if err := queries.session.Query("INSERT INTO vehicles (vin, engine_type, ev_data) VALUES (?, ?, ?) IF NOT EXISTS", &vehicle.Vin, &vehicle.Engine, &vehicle.EvData).Exec(); err != nil {
		return err
	}

	return nil
}

func (queries *ScyllaQueries) FindVehicle(vin string) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	if err := queries.session.Query("SELECT * FROM vehicles WHERE vin = ?", vin).Scan(&vehicle.Vin, &vehicle.Engine, &vehicle.EvData); err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &vehicle, nil
}
