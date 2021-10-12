package routing

import (
	"bwa.com/hello/db"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, d db.Database) {
	vh := NewVehicleHandlers(d)

	// CreateVehicle handler
	// @summary Create a new vehicle.
	// @desc Create a new vehicle.
	// @produce json
	// @success 201 {json} CREATED
	// @failure 500 {json}
	r.HandleFunc("/vehicle", vh.PostVehicle).Methods("POST")

	// FindVehicle handler
	// @summary Get vehicle info.
	// @desc Get info of a vehicle with the given VIN.
	// @produce json
	// @success 200 {json} OK
	// @failure 404 {string} The vehicle was not found
	// @failure 500 {json}
	r.HandleFunc("/vehicle", vh.GetVehicle).Methods("GET")
}
