package model

import "time"

type Measurement interface {
	Update()
}

type AccelerationMeasurement struct {
	Acc_x           float64   `json:"acc_x"`
	Acc_y           float64   `json:"acc_y"`
	Acc_z           float64   `json:"acc_z"`
	MovementCounter int       `json:"movement_counter"`
	RecordedTime    time.Time `json:"recorded_time"`
}

func (am *AccelerationMeasurement) Update() {
	return
}

type TemperatureMeasurement struct {
	Temperature      float64   `json:"temperature"`
	TemperatureScale float64   `json:"temperature_scale"`
	RecordedTime     time.Time `json:"recorded_time"`
}

func (tm *TemperatureMeasurement) Update() {
	return
}

type HumidityMeasurement struct {
	Humidity      float64   `json:"humidity"`
	HumidityScale float64   `json:"humidity_scale"`
	RecordedTime  time.Time `json:"recorded_time"`
}

func (tm *HumidityMeasurement) Update() {
	return
}

type VoltageMeasurement struct {
	Voltage      float64   `json:"voltage"`
	VoltageScale float64   `json:"voltage_scale"`
	RecordedTime time.Time `json:"recorded_time"`
}

func (tm *VoltageMeasurement) Update() {
	return
}
