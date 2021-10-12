package routing

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bwa.com/hello/db"
	"bwa.com/hello/helpers"
	"bwa.com/hello/mock"
	"bwa.com/hello/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var mockVehicleDAO *mock.MockVehicleDAO
var handlers VehicleHandlers

func setUp(t *testing.T) {
	// Database
	ctrl := gomock.NewController(t)
	mockDatabase := mock.NewMockDatabase(ctrl)
	mockVehicleDAO = mock.NewMockVehicleDAO(ctrl)
	mockDatabase.EXPECT().VehicleDAO().DoAndReturn(func() db.VehicleDAO { return mockVehicleDAO }).AnyTimes()

	// Handlers
	handlers = NewVehicleHandlers(mockDatabase)
}

func tearDown() {
}

// POST /vehicle => OK
func TestPostVehicle(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Vehicle JSON
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion", EvData: nil}
	vehicleJSON, err := json.Marshal(vehicle)
	require.NoError(t, err)

	// Expect DB CreateVehicle
	mockVehicleDAO.EXPECT().CreateVehicle(gomock.Eq(vehicle)).Return(nil)

	// Send POST request
	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(vehicleJSON))
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handlers.PostVehicle(rr, req)

	// Check code
	require.Equal(t, 201, rr.Code)

	// Check body
	var responseVehicle model.Vehicle
	err = helpers.DecodeJSONBody(rr.Body, &responseVehicle)
	require.NoError(t, err)
	require.Equal(t, vehicle, responseVehicle)
}

// POST vehicle => InternalError
func TestPostVehicleInternalError(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Vehicle JSON
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion", EvData: nil}
	vehicleJSON, err := json.Marshal(vehicle)
	require.NoError(t, err)

	// Expect DB CreateVehicle
	mockVehicleDAO.EXPECT().CreateVehicle(gomock.Eq(vehicle)).Return(errors.New("test CreateVehicle error"))

	// Send POST request
	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(vehicleJSON))
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handlers.PostVehicle(rr, req)

	// Check code
	require.Equal(t, 500, rr.Code)

	// Check body
	var responseError helpers.ErrorObject
	err = helpers.DecodeJSONBody(rr.Body, &responseError)
	require.NoError(t, err)
	require.Equal(t, helpers.NewErrorObject("Could not create vehicle: test CreateVehicle error"), responseError)
}

// GET vehicle => OK
func TestGetVehicle(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Expect DB FindVehicle
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion", EvData: nil}
	mockVehicleDAO.EXPECT().FindVehicle(gomock.Eq("vin1")).Return(&vehicle, nil)

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle?vin=vin1", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handlers.GetVehicle(rr, req)

	// Check code
	require.Equal(t, 200, rr.Code)

	// Check body
	var responseVehicle model.Vehicle
	err = helpers.DecodeJSONBody(rr.Body, &responseVehicle)
	require.NoError(t, err)
	require.Equal(t, vehicle, responseVehicle)
}

// GET vehicle => InternalError
func TestGetVehicleNoQuery(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handlers.GetVehicle(rr, req)

	// Check code
	require.Equal(t, 500, rr.Code)

	// Check body
	var response map[string]string
	err = helpers.DecodeJSONBody(rr.Body, &response)
	require.NoError(t, err)
	require.Equal(t, map[string]string{"error": "Could not find vehicle (missing 'vin' query parameter)"}, response)
}

// GET vehicle => NotFound
func TestGetVehicleNotFound(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Expect DB FindVehicle
	mockVehicleDAO.EXPECT().FindVehicle(gomock.Eq("wrong")).Return(nil, nil)

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle?vin=wrong", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handlers.GetVehicle(rr, req)

	// Check code
	require.Equal(t, 404, rr.Code)

	// Check body
	var response map[string]string
	err = helpers.DecodeJSONBody(rr.Body, &response)
	require.NoError(t, err)
	require.Equal(t, map[string]string{"vin": "wrong"}, response)
}

// GET vehicle => InternalError
func TestGetVehicleDbError(t *testing.T) {
	setUp(t)
	defer tearDown()

	// Expect DB FindVehicle
	mockVehicleDAO.EXPECT().FindVehicle(gomock.Eq("vin1")).Return(nil, errors.New("test FindVehicle error"))

	// Send GET request
	req, err := http.NewRequest("GET", "/vehicle?vin=vin1", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	handlers.GetVehicle(rr, req)

	// Check code
	require.Equal(t, 500, rr.Code)

	// Check body
	var responseError helpers.ErrorObject
	err = helpers.DecodeJSONBody(rr.Body, &responseError)
	require.NoError(t, err)
	require.Equal(t, helpers.NewErrorObject("Could not find vehicle: test FindVehicle error"), responseError)
}
