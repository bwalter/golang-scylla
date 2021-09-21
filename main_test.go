package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bwa.com/hello/helpers"
	"bwa.com/hello/model"
	"bwa.com/hello/test"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Queries
	queries := test.MockedQueries{}

	// App
	app := NewApp(&queries)

	// Server
	srv := httptest.NewServer(app.Router)
	defer srv.Close()

	// Vehicle JSON
	vehicle := model.Vehicle{Vin: "vin1", EngineType: "Combustion"}
	vehicle_json, err := json.Marshal(vehicle)
	checkError(t, err)

	// Send POST request
	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(vehicle_json))
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

func checkError(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Errorf("Err: %v", err)
}
