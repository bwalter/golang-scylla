package main

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

type error_object struct {
	Error string
}

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
	assert.Equal(t, rr.Code, 201)

	// Check body
	var response_vehicle model.Vehicle
	helpers.DecodeBodyToJSON(rr.Body, &response_vehicle)
	assert.Equal(t, response_vehicle, vehicle)
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
	assert.Equal(t, rr.Code, 500)

	// Check body
	var response_error error_object
	helpers.DecodeBodyToJSON(rr.Body, &response_error)
	assert.Equal(t, response_error, error_object{Error: "Could not create vehicle: test CreateVehicle error"})
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
	assert.Equal(t, rr.Code, 200)

	// Check body
	var response_vehicle model.Vehicle
	helpers.DecodeBodyToJSON(rr.Body, &response_vehicle)
	assert.Equal(t, response_vehicle, vehicle)
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
	assert.Equal(t, rr.Code, 404)

	// Check body
	var response map[string]string
	helpers.DecodeBodyToJSON(rr.Body, &response)
	assert.Equal(t, response, map[string]string{"vin": "wrong"})
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
	assert.Equal(t, rr.Code, 500)

	// Check body
	var response_error error_object
	helpers.DecodeBodyToJSON(rr.Body, &response_error)
	assert.Equal(t, response_error, error_object{Error: "Could not find vehicle: test FindVehicle error"})
}

func checkError(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Errorf("Err: %v", err)
}
