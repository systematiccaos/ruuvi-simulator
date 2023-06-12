package pendulum

import (
	"math"
	"math/rand"

	"github.com/sirupsen/logrus"
)

type Pendulum struct {
	Radius    float64
	StringLen float64
	Angle     float64
	Start     Vector2
	Position  Vector2
	MoveBob   bool
	Mass      float64
	Velocity  float64
}

type DoublePendulum struct {
	G            float64
	Damp         float64
	P1           Pendulum
	P2           Pendulum
	FrameCounter int
}

// Algorithm ported from: https://github.com/v1nam/DoublePendulum/blob/main/src/pendulum.cpp
func (dp *DoublePendulum) MoveObjects(frametime float64) {
	p1acc := -1 * dp.G * (2.0*dp.P1.Mass + dp.P2.Mass) * math.Sin(dp.P1.Angle)
	p1acc = p1acc - dp.G*dp.P2.Mass*math.Sin(dp.P1.Angle-2.0*dp.P2.Angle)
	logrus.Println(dp.P2.Velocity)
	p1acc = p1acc - 2.0*dp.P2.Mass*dp.P2.Velocity*dp.P2.Velocity*dp.P2.StringLen*math.Sin(dp.P1.Angle-dp.P2.Angle)
	p1acc = p1acc - dp.P2.Mass*dp.P1.Velocity*dp.P1.Velocity*dp.P1.StringLen*math.Sin(2.0*(dp.P1.Angle-dp.P2.Angle))
	p1acc = p1acc / (dp.P1.StringLen * (2.0*dp.P1.Mass + dp.P2.Mass - dp.P2.Mass*math.Cos(2.0*(dp.P1.Angle-dp.P2.Angle))))

	p2acc := (dp.P1.Mass + dp.P2.Mass) * dp.P1.StringLen * dp.P1.Velocity * dp.P1.Velocity
	p2acc = p2acc + dp.G*(dp.P1.Mass+dp.P2.Mass)*math.Cos(dp.P1.Angle)
	p2acc = p2acc + dp.P2.Velocity*dp.P2.Velocity*dp.P2.StringLen*dp.P2.Mass*math.Cos(dp.P1.Angle-dp.P2.Angle)
	p2acc = p2acc * 2.0 * math.Sin(dp.P1.Angle-dp.P2.Angle)
	p2acc = p2acc / (dp.P2.StringLen * (2.0*dp.P1.Mass + dp.P2.Mass - dp.P2.Mass*math.Cos(2.0*(dp.P1.Angle-dp.P2.Angle))))

	dp.P1.Velocity += p1acc * frametime
	dp.P2.Velocity += p2acc * frametime
	dp.P1.Velocity = dp.P1.Velocity * dp.Damp
	dp.P2.Velocity = dp.P2.Velocity * dp.Damp
	dp.P1.Angle += dp.P1.Velocity * frametime
	dp.P2.Angle += dp.P2.Velocity * frametime
}

func (dp *DoublePendulum) UpdatePos() {
	dp.P1.Position.X = dp.P1.Start.X + math.Sin(dp.P1.Angle)*dp.P1.StringLen
	dp.P1.Position.Y = dp.P1.Start.Y + math.Cos(dp.P1.Angle)*dp.P1.StringLen
	dp.P2.Start.X = dp.P1.Position.X
	dp.P2.Start.Y = dp.P1.Position.Y

	dp.P2.Position.X = dp.P2.Start.X + math.Sin(dp.P2.Angle)*dp.P2.StringLen
	dp.P2.Position.Y = dp.P2.Start.Y + math.Cos(dp.P2.Angle)*dp.P2.StringLen
}

func NewDoublePendulum() DoublePendulum {
	startx1 := 100 / rand.Float64()
	starty1 := 100 / rand.Float64()
	dp := DoublePendulum{
		G:    980.0,
		Damp: 0.001,
		P1: Pendulum{
			Start: Vector2{
				X: startx1,
				Y: starty1,
			},
			StringLen: 140,
			Angle:     2 / math.Pi,
			Velocity:  0.1,
			Mass:      1.0,
		},
		P2: Pendulum{
			Start: Vector2{
				X: startx1 + 10/rand.Float64(),
				Y: starty1 + 10/rand.Float64(),
			},
			StringLen: 140,
			Angle:     3 / math.Pi,
			Velocity:  0.1,
			Mass:      1.0,
		},
	}
	dp.P1.Position = Vector2{X: 100.0, Y: dp.P1.Start.Y + 200.0}
	dp.P2.Position = Vector2{X: 100.0, Y: dp.P2.Start.Y + 150.0}
	return dp
}
