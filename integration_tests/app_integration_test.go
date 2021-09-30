package integrationTests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"bwa.com/hello/app"
	"bwa.com/hello/db"
	"bwa.com/hello/helpers"
	"bwa.com/hello/model"
	"github.com/stretchr/testify/require"
)

var queries db.Queries
var a app.App

func setUp(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Scylla URI
	uri := os.Getenv("SCYLLA_URI")
	if uri == "" {
		uri = "localhost"
	}

	host := uri + ":9042"
	keyspace := "hello_test"

	// Create test keyspace
	err := db.CreateScyllaKeyspace(host, keyspace, true)
	require.NoError(t, err)

	// Start session
	q, err := db.StartScyllaSessionAndCreateQueries(host, keyspace)
	require.NoError(t, err)
	queries = q

	// App
	a = app.NewApp(queries)
}

func tearDown(t *testing.T) {
	t.Cleanup(func() {
		queries.CloseSession()
		queries = nil
	})
}

// POST /vehicle => OK
func TestPostVehicle(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	// Vehicle JSON
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion"}
	vehicleJSON, err := json.Marshal(vehicle)
	require.NoError(t, err)

	// Send POST request
	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(vehicleJSON))
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	// Check code
	require.Equal(t, 201, rr.Code)

	// Check body
	var responseVehicle model.Vehicle
	err = helpers.DecodeJSONBody(rr.Body, &responseVehicle)
	require.NoError(t, err)
	require.Equal(t, vehicle, responseVehicle)

	// Check vehicle
	vehiclePtr, err := queries.FindVehicle("vin1")
	require.NoError(t, err)
	require.NotNil(t, vehiclePtr)
	require.Equal(t, vehicle, *vehiclePtr)
}

// GET /vehicle => OK
func TestGetVehicle(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	// Send GET request => NOT_FOUND
	rr := handleRequest(t, "GET", "/vehicle?vin=vin1")
	require.Equal(t, 404, rr.Code)

	// Insert vehicle
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion"}
	err := queries.CreateVehicle(vehicle)
	require.NoError(t, err)

	// Send GET request => OK
	rr = handleRequest(t, "GET", "/vehicle?vin=vin1")
	require.Equal(t, 200, rr.Code)

	// Check body
	var responseVehicle model.Vehicle
	err = helpers.DecodeJSONBody(rr.Body, &responseVehicle)
	require.NoError(t, err)
	require.Equal(t, vehicle, responseVehicle)
}

func handleRequest(t *testing.T, method string, path string) httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/vehicle?vin=vin1", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return *rr
}
