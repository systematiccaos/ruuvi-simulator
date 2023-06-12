package model

import "time"

type Measurement struct {
}

type AccelerationMeasurement struct {
	Measurement
	Acc_x           float64   `json:"acc_x"`
	Acc_y           float64   `json:"acc_y"`
	Acc_z           float64   `json:"acc_z"`
	MovementCounter int       `json:"movement_counter"`
	RecordedTime    time.Time `json:"recorded_time"`
}
