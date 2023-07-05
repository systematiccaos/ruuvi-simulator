package model

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/ruuvi-simulator/pendulum"
)

type Sensor interface {
	GetMeasurements() []Measurement
	Update()
}

type AccelerationSensor struct {
	Measurements []Measurement
	frame        int
	framerate    float64
	dp           pendulum.DoublePendulum
	dp2          pendulum.DoublePendulum
}

func NewAccelerationSensor() AccelerationSensor {
	as := AccelerationSensor{
		frame:     0,
		framerate: 60.0,
		dp:        pendulum.NewDoublePendulum(50, 50, 2, 2.5, 5, 3),
		dp2:       pendulum.NewDoublePendulum(80, 50, 2, 3.5, 5, 8),
	}
	return as
}

func (as *AccelerationSensor) runMeasurementCalculation() {
	as.calcPendulums()
}

func (as *AccelerationSensor) GetMeasurements() []Measurement {
	return as.Measurements
}

func (as *AccelerationSensor) Update() {
	as.runMeasurementCalculation()
}

func (as *AccelerationSensor) calcPendulums() {
	as.dp.MoveObjects(1.0 / as.framerate)
	as.dp2.MoveObjects(1.0 / as.framerate)
	as.dp.UpdatePos()
	as.dp2.UpdatePos()
	if as.frame%600 == 0 {
		start := len(as.dp.P1.Accelerations) - 600
		if start < 0 {
			start = 0
		}
		logrus.Println(as.dp.P1.Accelerations[start:])
	}
	time.Sleep(time.Second / time.Duration(as.framerate))
	as.frame++
}
