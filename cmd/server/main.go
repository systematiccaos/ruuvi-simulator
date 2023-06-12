package main

import (
	"github.com/systematiccaos/going-forward/util"
	"github.com/systematiccaos/ruuvi-simulator/pendulum"
)

func main() {
	util.SetupLogs()

	dp := &pendulum.DoublePendulum{
		G:    980,
		Damp: 1,
	}
	frame := 0.0
	for {
		dp.MoveObjects(frame)
		dp.UpdatePos()

		frame++
	}
}
