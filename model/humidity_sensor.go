package model

import (
	"math/rand"
	"time"
)

type HumiditySensor struct {
	Measurements  []Measurement `json:"measurements,omitempty"`
	Tp            string        `json:"type"`
	Humidity      float64       `json:"-"`
	HumidityScale string        `json:"-"`
}

func NewHumiditySensor() HumiditySensor {
	hs := HumiditySensor{
		Humidity:      rand.Float64(),
		HumidityScale: "celsius",
		Tp:            "HumiditySensor",
	}
	return hs
}

func (hs *HumiditySensor) GetMeasurements() []Measurement {
	cleanedMeasurements := []Measurement{}
	for _, m := range hs.Measurements {
		if m != nil {
			cleanedMeasurements = append(cleanedMeasurements, m)
		}
	}
	return cleanedMeasurements
}

func (hs *HumiditySensor) Update() {
	plusminus := 1.0
	if rand.Float64() > 0.95 {
		plusminus = -1.0
	}
	hs.Humidity = hs.Humidity + rand.Float64()*0.01 + plusminus
	hs.Measurements = append(hs.Measurements, &HumidityMeasurement{
		Humidity:     hs.Humidity,
		RecordedTime: time.Now(),
	})
}
