package models

type SensorData struct {
	Timestamp int64   `json:"timestamp"`
	RSSI	   int     `json:"rssi"`
	Variance  float64 `json:"variance"`
}

type PresenceState struct {
	State string `json:"state"`
	Variance float64 `json:"variance"`
	Timestamp int64 `json:"timestamp"`
}