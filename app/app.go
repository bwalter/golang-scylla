package app

import (
	"bwa.com/hello/db"
	"bwa.com/hello/handlers"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	handlers handlers.Handlers
	queries  db.Queries
}

// Create app with a router for handling requests
func NewApp(queries db.Queries) App {
	r := mux.NewRouter()
	h := handlers.NewHandlers(queries)

	a := App{
		Router:   r,
		handlers: h,
		queries:  queries,
	}

	// CreateVehicle handler
	// @summary Create a new vehicle.
	// @desc Create a new vehicle.
	// @produce json
	// @success 201 {json} CREATED
	// @failure 500 {json}
	r.HandleFunc("/vehicle", h.PostVehicle).Methods("POST")

	// FindVehicle handler
	// @summary Get vehicle info.
	// @desc Get info of a vehicle with the given VIN.
	// @produce json
	// @success 200 {json} OK
	// @failure 404 {string} The vehicle was not found
	// @failure 500 {json}
	r.HandleFunc("/vehicle", h.GetVehicle).Methods("GET")

	return a
}

func (a *App) CloseSession() {
	a.queries.CloseSession()
}
