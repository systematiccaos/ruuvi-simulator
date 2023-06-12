package main

import (
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/going-forward/util"
	"github.com/systematiccaos/ruuvi-simulator/pendulum"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vggio"
)

func main() {
	util.SetupLogs()
	w := app.NewWindow()
	run(w)
}

func run(w *app.Window) error {
	var ops op.Ops
	frame := 1.0
	dp := pendulum.NewDoublePendulum()
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			dp.MoveObjects(frame)
			dp.UpdatePos()
			gtx := layout.NewContext(&ops, e)

			p := plot.New()
			p.Title.Text = "Pendulum"
			p.X.Label.Text = "X"
			p.Y.Label.Text = "Y"
			logrus.Printf("X: %f, Y: %f", dp.P1.Position.X, dp.P1.Position.Y)

			pts := plotter.XYs{
				{X: dp.P1.Position.X, Y: dp.P1.Position.Y},
				{X: dp.P2.Position.X, Y: dp.P2.Position.Y},
			}
			scatter, err := plotter.NewScatter(pts)
			if err != nil {
				logrus.Fatalln(err)
			}
			p.Add(scatter)
			cnv := vggio.New(gtx, 20*vg.Centimeter, 20*vg.Centimeter, vggio.UseDPI(96))
			p.Draw(draw.New(cnv))

			e.Frame(cnv.Paint())
			time.Sleep(time.Second / 4)
			w.Invalidate()
			frame++
		}
	}
}
