package model

import (
	"time"

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
	dp           *pendulum.DoublePendulum
	dp2          *pendulum.DoublePendulum
}

func NewAccelerationSensor() AccelerationSensor {
	dp1 := pendulum.NewDoublePendulum(50, 50, 2, 2.5, 5, 3)
	dp2 := pendulum.NewDoublePendulum(80, 50, 2, 3.5, 5, 8)
	as := AccelerationSensor{
		frame:     0,
		framerate: 60.0,
		dp:        &dp1,
		dp2:       &dp2,
	}
	return as
}

func (as *AccelerationSensor) GetMeasurements() []Measurement {
	return as.Measurements
}

func (as *AccelerationSensor) Update() {
	as.calcPendulums()
	as.Measurements = append(as.Measurements, AccelerationMeasurement{
		Acc_x:           as.dp.P1.Accelerations[len(as.dp.P1.Accelerations)-1].X,
		Acc_y:           as.dp.P1.Accelerations[len(as.dp.P1.Accelerations)-1].Y,
		Acc_z:           9.81,
		MovementCounter: len(as.dp.P1.Accelerations),
	})
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
		// logrus.Println(as.dp.P1.Accelerations)
	}
	time.Sleep(time.Second / time.Duration(as.framerate))
	as.frame++
}
