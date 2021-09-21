package main

import (
	"bwa.com/hello/db"
	"bwa.com/hello/helpers"
	"bwa.com/hello/model"

	"github.com/gorilla/mux"

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

	if err := helpers.DecodeBodyToJSON(r.Body, &vehicle); err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not decode vehicle: %s", err))
		return
	}

	if err := a.queries.CreateVehicle(vehicle); err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not create vehicle: %s", err))
		return
	}

	helpers.RespondWithJSON(w, 201, vehicle)
}

// GET (query: vin) => (200 body: vehicle JSON) or (404 body: vehicle vin JSON) or (500 body: error JSON)
func (a *App) findVehicle(w http.ResponseWriter, r *http.Request) {
	vins := r.URL.Query()["vin"]
	if vins == nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not find vehicle (missing 'vin' query parameter)"))
		return
	}

	vin := vins[0]
	vehicle, err := a.queries.FindVehicle(vin)
	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not find vehicle: %s", err))
	}

	if vehicle == nil {
		helpers.RespondWithJSON(w, 404, map[string]string{"vin": vin})
		return
	}

	helpers.RespondWithJSON(w, 200, vehicle)
}
