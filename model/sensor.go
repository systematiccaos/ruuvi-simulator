package model

type Sensor interface {
	GetMeasurements() []Measurement
	Update()
}
