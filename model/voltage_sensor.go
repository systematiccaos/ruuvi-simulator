package model

import (
	"math/rand"
	"time"
)

type VoltageSensor struct {
	Measurements []Measurement `json:"measurements,omitempty"`
	Tp           string        `json:"type"`
	Voltage      float64       `json:"-"`
	VoltageScale string        `json:"-"`
}

func NewVoltageSensor() VoltageSensor {
	as := VoltageSensor{
		Voltage:      3 * rand.Float64(),
		VoltageScale: "celsius",
		Tp:           "VoltageSensor",
	}
	return as
}

func (vs *VoltageSensor) GetMeasurements() []Measurement {
	cleanedMeasurements := []Measurement{}
	for _, m := range vs.Measurements {
		if m != nil {
			cleanedMeasurements = append(cleanedMeasurements, m)
		}
	}
	return cleanedMeasurements
}

func (vs *VoltageSensor) Update() {
	plusminus := 1.0
	if rand.Float64() > 0.95 {
		plusminus = -1.0
	}
	vs.Voltage = vs.Voltage + rand.Float64()*.001 + plusminus
	vs.Measurements = append(vs.Measurements, &VoltageMeasurement{
		Voltage:      vs.Voltage,
		RecordedTime: time.Now(),
	})
}
