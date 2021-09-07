package main

import (
	"bwa.com/hello/db"
	"bwa.com/hello/model"

	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"net/http"
)

type App struct {
	Router  *mux.Router
	queries db.Queries
}

// Create app with a router for handling requests
func NewApp(queries db.Queries) App {
	r := mux.NewRouter()

	a := App{
		Router:  r,
		queries: queries,
	}

	r.HandleFunc("/vehicle", a.createVehicle).Methods("POST")
	r.HandleFunc("/vehicle", a.findVehicle).Methods("GET")

	return a
}

// POST (body: vehicle JSON) => (200 body: vehicle JSON) or (500 body: vehicle JSON)
func (a *App) createVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle model.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not decode vehicle: %s", err))
		return
	}

	a.queries.CreateVehicle(vehicle)

	respondWithJSON(w, 200, vehicle)
}

// GET (query: vin) => (200 body: vehicle JSON) or (404 body: vehicle vin JSON) or (500 body: error JSON)
func (a *App) findVehicle(w http.ResponseWriter, r *http.Request) {
	vins := r.URL.Query()["vin"]
	if vins == nil {
		respondWithError(w, 500, fmt.Sprintf("Could not find vehicle (missing 'vin' query parameter)"))
		return
	}

	vin := vins[0]
	vehicle, err := a.queries.FindVehicle(vin)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not find vehicle: %s", err))
	}

	if vehicle == nil {
		respondWithJSON(w, 404, vin)
		return
	}

	respondWithJSON(w, 200, vehicle)
}

// => (<code> body: error JSON)
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// => (<code> body: payload JSON)
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
