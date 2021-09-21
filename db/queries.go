package db

import (
	"fmt"

	"bwa.com/hello/model"
	"github.com/gocql/gocql"
)

type Queries interface {
	CreateTablesIfNotExist() error
	CreateVehicle(vehicle model.Vehicle) error
	FindVehicle(vin string) (*model.Vehicle, error)
}

func StartDbSessionAndCreateQueries(host string) (Queries, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = "hello"

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Database!")

	queries, err := NewQueries(session)
	if err != nil {
		return nil, nil
	}
	queries.CreateTablesIfNotExist()

	return queries, nil
}
