package scylla

import (
	"fmt"
	"time"

	"bwa.com/hello/db"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Database struct {
	session    *gocqlx.Session
	vehicleDAO VehicleDAO
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

func StartSessionAndCreateDatabase(host string, keyspace string) (db.Database, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	database, err := newDatabase(session)
	if err != nil {
		return nil, nil
	}

	if err := database.CreateTablesIfNotExist(); err != nil {
		return nil, err
	}

	return database, nil
}

func newDatabase(session *gocql.Session) (db.Database, error) {
	sessionx, err := gocqlx.WrapSession(session, nil)
	if err != nil {
		return nil, err
	}

	return &Database{
		session:    &sessionx,
		vehicleDAO: newVehicleDAO(&sessionx),
	}, nil
}

func (db *Database) CreateTablesIfNotExist() error {
	if err := db.session.ExecStmt("CREATE TYPE IF NOT EXISTS ev_data (battery_capacity_in_kwh int, soc_in_percent int)"); err != nil {
		return err
	}

	cql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (vin text primary key, engine_type text, ev_data ev_data)", vehicleMetadata.Name)
	if err := db.session.ExecStmt(cql); err != nil {
		return err
	}

	return nil
}

func (db *Database) CloseSession() {
	db.session.Close()
}

func (db *Database) VehicleDAO() db.VehicleDAO {
	return &db.vehicleDAO
}
