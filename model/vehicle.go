package model

type Vehicle struct {
	Vin        string `json:"vin" validate:"required"`
	EngineType string `json:"engine_type"`

	// Only for EV engine
	EvData *EvData `json:"ev_data,omitempty"`
}

type EvData struct {
	BatteryCapacityInKwh int `json:"battery_capacity_in_kwh"`
	SocInPercent         int `json:"soc_in_percent"`
}
