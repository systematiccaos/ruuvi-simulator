package main

import (
	"image/color"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/ajstarks/giocanvas"
	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/going-forward/util"
	"github.com/systematiccaos/ruuvi-simulator/pendulum"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vggio"
)

func main() {
	util.SetupLogs()
	w := app.NewWindow()
	run(w)
}

func run(w *app.Window) error {
	var ops op.Ops
	frame := 1
	framerate := 60.0
	dp := pendulum.NewDoublePendulum(50, 50, 2, 2.5, 5, 3)
	dp2 := pendulum.NewDoublePendulum(80, 50, 2, 3.5, 5, 8)
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			// p := plot.New()
			// p.Title.Text = "Pendulum"
			// p.X.Label.Text = "X"
			// p.Y.Label.Text = "Y"

			gc := giocanvas.Canvas{
				Height:  1000.0,
				Width:   1000.0,
				Context: gtx,
			}
			drawPendulum(dp, gc)
			drawPendulum(dp2, gc)
			// logrus.Printf("X: %f, Y: %f", dp.P1.Position.X, dp.P1.Position.Y)

			// pts := plotter.XYs{
			// 	{X: dp.P1.Position.X, Y: dp.P1.Position.Y},
			// 	{X: dp.P2.Position.X, Y: dp.P2.Position.Y},
			// }
			// scatter, err := plotter.NewScatter(pts)
			// if err != nil {
			// 	logrus.Fatalln(err)
			// }
			// p.Add(scatter)
			cnv := vggio.New(gtx, 20*vg.Centimeter, 20*vg.Centimeter, vggio.UseDPI(96))
			// p.Draw(draw.New(cnv))

			e.Frame(cnv.Paint())
			w.Invalidate()
			dp.MoveObjects(1.0 / framerate)
			dp2.MoveObjects(1.0 / framerate)
			// logrus.Printf("frametime: %d", frame/60)
			dp.UpdatePos()
			dp2.UpdatePos()
			frame++
			time.Sleep(time.Second / time.Duration(framerate))
			if frame%600 == 0 {
				logrus.Println(dp.P1.Accelerations)
			}
		}
	}
}

func drawPendulum(dp pendulum.DoublePendulum, gc giocanvas.Canvas) {
	gc.Line(float32(dp.P1.Start.X), float32(dp.P1.Start.Y), float32(dp.P1.Position.X), float32(dp.P1.Position.Y), 1.0, color.NRGBA{100, 100, 100, 255})
	gc.Line(float32(dp.P2.Start.X), float32(dp.P2.Start.Y), float32(dp.P2.Position.X), float32(dp.P2.Position.Y), 1.0, color.NRGBA{100, 100, 100, 255})
	gc.Circle(float32(dp.P1.Start.X), float32(dp.P1.Start.Y), 1, color.NRGBA{255, 255, 0, 255})

	gc.Circle(float32(dp.P1.Position.X), float32(dp.P1.Position.Y), 1, color.NRGBA{255, 0, 0, 255})
	gc.Circle(float32(dp.P2.Position.X), float32(dp.P2.Position.Y), 1, color.NRGBA{0, 0, 255, 255})
}
