package integrationTests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"bwa.com/hello/app"
	"bwa.com/hello/db"
	"bwa.com/hello/db/mockdb"
	"bwa.com/hello/helpers"
	"bwa.com/hello/model"
	"github.com/stretchr/testify/require"
)

type Ctx struct {
	db  db.Database
	app app.App
}

var ctx Ctx

func setUp(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Start session
	d := mockdb.NewDatabase()
	ctx.db = d

	// App
	ctx.app = app.NewApp(d)
}

func tearDown(t *testing.T) {
	t.Cleanup(func() {
		ctx.db.CloseSession()
		ctx = Ctx{}
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
	ctx.app.Router.ServeHTTP(rr, req)

	// Check code
	require.Equal(t, 201, rr.Code)

	// Check body
	var responseVehicle model.Vehicle
	err = helpers.DecodeJSONBody(rr.Body, &responseVehicle)
	require.NoError(t, err)
	require.Equal(t, vehicle, responseVehicle)

	// Check vehicle
	vehiclePtr, err := ctx.db.VehicleDAO().FindVehicle("vin1")
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
	err := ctx.db.VehicleDAO().CreateVehicle(vehicle)
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
	ctx.app.Router.ServeHTTP(rr, req)
	return *rr
}
