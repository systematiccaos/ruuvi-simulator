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
	Measurements []Measurement            `json:"measurements,omitempty"`
	Tp           string                   `json:"type"`
	frame        int                      `json:"-"`
	framerate    float64                  `json:"-"`
	dp           *pendulum.DoublePendulum `json:"-"`
	dp2          *pendulum.DoublePendulum `json:"-"`
}

func NewAccelerationSensor() AccelerationSensor {
	dp1 := pendulum.NewDoublePendulum(50, 50, 2, 2.5, 5, 3)
	dp2 := pendulum.NewDoublePendulum(80, 50, 2, 3.5, 5, 8)
	as := AccelerationSensor{
		frame:     0,
		framerate: 60.0,
		dp:        &dp1,
		dp2:       &dp2,
		Tp:        "AccelerationSensor",
	}
	return as
}

func (as *AccelerationSensor) GetMeasurements() []Measurement {
	cleanedMeasurements := []Measurement{}
	for _, m := range as.Measurements {
		if m != nil {
			cleanedMeasurements = append(cleanedMeasurements, m)
		}
	}
	return cleanedMeasurements
}

func (as *AccelerationSensor) Update() {
	as.calcPendulums()
	as.Measurements = append(as.Measurements, &AccelerationMeasurement{
		Acc_x:           as.dp.P1.Accelerations[len(as.dp.P1.Accelerations)-1].X,
		Acc_y:           as.dp.P1.Accelerations[len(as.dp.P1.Accelerations)-1].Y,
		Acc_z:           9.81,
		MovementCounter: len(as.dp.P1.Accelerations),
		RecordedTime:    time.Now(),
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
