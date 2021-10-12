package routing

import (
	"fmt"
	"net/http"

	"bwa.com/hello/db"
	"bwa.com/hello/helpers"
	"bwa.com/hello/model"
)

type VehicleHandlers struct {
	db db.Database
}

func NewVehicleHandlers(d db.Database) VehicleHandlers {
	return VehicleHandlers{
		db: d,
	}

}

// POST (body: vehicle JSON) => (200 body: vehicle JSON) or (500 body: vehicle JSON)
func (h *VehicleHandlers) PostVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle model.Vehicle

	if err := helpers.DecodeJSONBody(r.Body, &vehicle); err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not decode vehicle: %s", err))
		return
	}

	if err := h.db.VehicleDAO().CreateVehicle(vehicle); err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not create vehicle: %s", err))
		return
	}

	helpers.RespondWithJSON(w, 201, vehicle)
}

// GET (query: vin) => (200 body: vehicle JSON) or (404 body: vehicle vin JSON) or (500 body: error JSON)
func (h *VehicleHandlers) GetVehicle(w http.ResponseWriter, r *http.Request) {
	vins := r.URL.Query()["vin"]
	if vins == nil {
		helpers.RespondWithError(w, 500, "Could not find vehicle (missing 'vin' query parameter)")
		return
	}

	vin := vins[0]
	vehicle, err := h.db.VehicleDAO().FindVehicle(vin)
	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("Could not find vehicle: %s", err))
	}

	if vehicle == nil {
		helpers.RespondWithJSON(w, 404, map[string]string{"vin": vin})
		return
	}

	helpers.RespondWithJSON(w, 200, vehicle)
}
