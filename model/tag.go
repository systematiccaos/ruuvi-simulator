package model

import (
	"fmt"
	"math/rand"
)

type Tag struct {
	Sensors []Sensor `json:"sensors"`
	Address string   `json:"address"`
	Name    string   `json:"name"`
}

func NewTag() Tag {
	sfx1 := rand.Intn(255)
	sfx2 := rand.Intn(255)
	tag := Tag{
		Name:    "ruuvi_" + fmt.Sprintf("%x%x", sfx1, sfx2),
		Address: fmt.Sprintf("%x:%x:%x:%x:%x:%x", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255), sfx1, sfx2),
	}
	tag.Sensors = append(tag.Sensors, &AccelerationSensor{})
	return tag
}

func (t *Tag) GetAllSensorsMeasurements() []Measurement {
	measurements := []Measurement{}
	for _, sensor := range t.Sensors {
		measurement := sensor.GetMeasurements()
		measurements = append(measurements, measurement...)
	}
	return measurements
}

func (t *Tag) Update() {
	for _, sensor := range t.Sensors {
		sensor.Update()
	}
}
