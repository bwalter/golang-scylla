package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bwa.com/hello/helpers"
	"bwa.com/hello/mock"
	"bwa.com/hello/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mock_queries *mock.MockQueries
var app App
var srv *httptest.Server

func setUp(t *testing.T) {
	// Queries
	ctrl := gomock.NewController(t)
	mock_queries = mock.NewMockQueries(ctrl)

	// App
	app = NewApp(mock_queries)

	// Server
	srv = httptest.NewServer(app.Router)
}

func tearDown() {
	srv.Close()
}

// POST /vehicle => OK
func TestPostVehicle(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Vehicle JSON
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion"}
	vehicle_json, err := json.Marshal(vehicle)
	checkError(t, err)

	// Expect DB CreateVehicle
	mock_queries.EXPECT().CreateVehicle(gomock.Eq(vehicle)).Return(nil)

	// Send POST request
	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(vehicle_json))
	checkError(t, err)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	// Check code
	assert.Equal(t, 201, rr.Code)

	// Check body
	var response_vehicle model.Vehicle
	helpers.DecodeJSONBody(rr.Body, &response_vehicle)
	assert.Equal(t, vehicle, response_vehicle)
}

// POST vehicle => InternalError
func TestPostVehicleInternalError(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Vehicle JSON
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion"}
	vehicle_json, err := json.Marshal(vehicle)
	checkError(t, err)

	// Expect DB CreateVehicle
	mock_queries.EXPECT().CreateVehicle(gomock.Eq(vehicle)).Return(errors.New("test CreateVehicle error"))

	// Send POST request
	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(vehicle_json))
	checkError(t, err)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	// Check code
	assert.Equal(t, 500, rr.Code)

	// Check body
	var response_error helpers.ErrorObject
	helpers.DecodeJSONBody(rr.Body, &response_error)
	assert.Equal(t, helpers.NewErrorObject("Could not create vehicle: test CreateVehicle error"), response_error)
}

// GET vehicle => OK
func TestGetVehicle(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Expect DB FindVehicle
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion"}
	mock_queries.EXPECT().FindVehicle(gomock.Eq("vin1")).Return(&vehicle, nil)

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle?vin=vin1", nil)
	checkError(t, err)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	// Check code
	assert.Equal(t, 200, rr.Code)

	// Check body
	var response_vehicle model.Vehicle
	helpers.DecodeJSONBody(rr.Body, &response_vehicle)
	assert.Equal(t, vehicle, response_vehicle)
}

// GET vehicle => NotFound
func TestGetVehicleNotFound(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Expect DB FindVehicle
	mock_queries.EXPECT().FindVehicle(gomock.Eq("wrong")).Return(nil, nil)

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle?vin=wrong", nil)
	checkError(t, err)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	// Check code
	assert.Equal(t, 404, rr.Code)

	// Check body
	var response map[string]string
	helpers.DecodeJSONBody(rr.Body, &response)
	assert.Equal(t, map[string]string{"vin": "wrong"}, response)
}

// GET vehicle => InternalError
func TestGetVehicleInternalError(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Expect DB FindVehicle
	mock_queries.EXPECT().FindVehicle(gomock.Eq("vin1")).Return(nil, errors.New("test FindVehicle error"))

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle?vin=vin1", nil)
	checkError(t, err)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	// Check code
	assert.Equal(t, 500, rr.Code)

	// Check body
	var response_error helpers.ErrorObject
	helpers.DecodeJSONBody(rr.Body, &response_error)
	assert.Equal(t, helpers.NewErrorObject("Could not find vehicle: test FindVehicle error"), response_error)
}

func checkError(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Errorf("Err: %v", err)
}
