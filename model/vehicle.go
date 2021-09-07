package model

type Vehicle struct {
	Vin    string `json:"vin" required:"true"`
	Engine string `json:"engine"`

	// Only for EV engine
	EvData *EvData `json:"evData,omitempty"`
}

type EvData struct {
	BatteryCapacityInKwh int `json:"batteryCapacityInKwh"`
	SocInPercent         int `json:"socInPercent"`
}
