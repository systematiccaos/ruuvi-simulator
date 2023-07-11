package model

import (
	"math/rand"
	"time"
)

type TemperatureSensor struct {
	Measurements     []Measurement `json:"measurements,omitempty"`
	Tp               string        `json:"type"`
	Temperature      float64       `json:"-"`
	TemperatureScale string        `json:"-"`
}

func NewTemperatureSensor() TemperatureSensor {
	as := TemperatureSensor{
		Temperature:      50 * rand.Float64(),
		TemperatureScale: "celsius",
		Tp:               "TemperatureSensor",
	}
	return as
}

func (as *TemperatureSensor) GetMeasurements() []Measurement {
	cleanedMeasurements := []Measurement{}
	for _, m := range as.Measurements {
		if m != nil {
			cleanedMeasurements = append(cleanedMeasurements, m)
		}
	}
	return cleanedMeasurements
}

func (ts *TemperatureSensor) Update() {
	plusminus := 1.0
	if rand.Float64() > 0.95 {
		plusminus = -1.0
	}
	ts.Temperature = ts.Temperature + rand.Float64()*.1 + plusminus
	ts.Measurements = append(ts.Measurements, &TemperatureMeasurement{
		Temperature:  ts.Temperature,
		RecordedTime: time.Now(),
	})
}
