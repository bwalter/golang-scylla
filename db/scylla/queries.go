package scylla

import (
	"fmt"
	"time"

	"bwa.com/hello/db"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Queries struct {
	session        *gocqlx.Session
	vehicleQueries VehicleQueries
}

func CreateKeyspace(host string, keyspace string, deleteExisting bool) error {
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

func StartSessionAndCreateQueries(host string, keyspace string) (db.Queries, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	queries, err := newQueries(session)
	if err != nil {
		return nil, nil
	}

	if err := queries.CreateTablesIfNotExist(); err != nil {
		return nil, err
	}

	return queries, nil
}

func newQueries(session *gocql.Session) (db.Queries, error) {
	sessionx, err := gocqlx.WrapSession(session, nil)
	if err != nil {
		return nil, err
	}

	return &Queries{
		session:        &sessionx,
		vehicleQueries: newVehicleQueries(&sessionx),
	}, nil
}

func (queries *Queries) CreateTablesIfNotExist() error {
	if err := queries.session.ExecStmt("CREATE TYPE IF NOT EXISTS ev_data (battery_capacity_in_kwh int, soc_in_percent int)"); err != nil {
		return err
	}

	cql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (vin text primary key, engine_type text, ev_data ev_data)", vehicleMetadata.Name)
	if err := queries.session.ExecStmt(cql); err != nil {
		return err
	}

	return nil
}

func (queries *Queries) CloseSession() {
	queries.session.Close()
}

func (queries *Queries) VehicleQueries() db.VehicleQueries {
	return &queries.vehicleQueries
}
