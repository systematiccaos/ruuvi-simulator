package model

type Sensor struct {
	Measurements []Measurement
}

type AccelerationSensor struct {
	Sensor
	Measurements []AccelerationMeasurement
}
