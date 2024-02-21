package models

type Location struct {
	DeviceID      int64   `json:"device_id"`
	ClientID      string  `json:"client_id"`
	Lat           float64 `json:"lat"`
	Long          float64 `json:"long"`
	BatteryStatus int     `json:"battery_status"`
	Timestamp     int64   `json:"timestamp"`
}
