package main

import (
	"math/rand"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"

	"github.com/JoshuaKolden/interp"
)

func main() {
	rand.Seed(int64(0))

	stepsplot, err := plot.New()
	if err != nil {
		panic(err)
	}
	easeout, err := plot.New()
	if err != nil {
		panic(err)
	}
	easein, err := plot.New()
	if err != nil {
		panic(err)
	}
	smoothease, err := plot.New()
	if err != nil {
		panic(err)
	}

	stepsplot.Title.Text = "Step Functions"
	stepsplot.X.Label.Text = "t"
	stepsplot.Y.Label.Text = "r"

	easeout.Title.Text = "EaseOut Functions"
	easeout.X.Label.Text = "t"
	easeout.Y.Label.Text = "r"

	easein.Title.Text = "EaseIn Functions"
	easein.X.Label.Text = "t"
	easein.Y.Label.Text = "r"

	smoothease.Title.Text = "Smoothstep vs Easein/out"
	smoothease.X.Label.Text = "t"
	smoothease.Y.Label.Text = "r"

	err = plotutil.AddLinePoints(stepsplot,
		"Smoothstep", smoothstepPoints(20),
		"Linearstep", linearstepPoints(20),
		"Step", stepPoints(20))
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLinePoints(easeout,
		"Linearstep", linearstepPoints(20),
		"Easeoutstep 2", easeoutPoints(20, 2),
		"Easeoutstep 3", easeoutPoints(20, 3),
		"Easeoutstep 4", easeoutPoints(20, 4),
		"Easeoutstep 5", easeoutPoints(20, 5))
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLinePoints(easein,
		"Linearstep", linearstepPoints(20),
		"Easeoutstep 2", easeinPoints(20, 2),
		"Easeoutstep 3", easeinPoints(20, 3),
		"Easeoutstep 4", easeinPoints(20, 4),
		"Easeoutstep 5", easeinPoints(20, 5))
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLinePoints(smoothease,
		"Linearstep", linearstepPoints(20),
		"Smoothstep", smoothstepPoints(20),
		"Easeinstep 2", easeinPoints(20, 2),
		"Easeoutstep 2", easeoutPoints(20, 2))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := stepsplot.Save(4*vg.Inch, 4*vg.Inch, "steps.png"); err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := easeout.Save(4*vg.Inch, 4*vg.Inch, "easeout.png"); err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := easein.Save(4*vg.Inch, 4*vg.Inch, "easein.png"); err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := smoothease.Save(4*vg.Inch, 4*vg.Inch, "smoothease.png"); err != nil {
		panic(err)
	}
}

func easeinPoints(n int, exp float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		// range from 0.0 to 1.0
		t := float64(i) / float64(n-1)
		pts[i].X = t
		pts[i].Y = interp.Easeinstep(t, exp)
	}
	return pts
}

func easeoutPoints(n int, exp float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		// range from 0.0 to 1.0
		t := float64(i) / float64(n-1)
		pts[i].X = t
		pts[i].Y = interp.Easeoutstep(t, exp)
	}
	return pts
}

func stepPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		// range from 0.0 to 1.0
		t := float64(i) / float64(n-1)
		pts[i].X = t
		pts[i].Y = interp.Step(t, 0.5)
	}
	return pts
}

func linearstepPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		// range from 0.0 to 1.0
		t := float64(i) / float64(n-1)
		pts[i].X = t
		pts[i].Y = interp.Linearstep(t)
	}
	return pts
}

func smoothstepPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		// range from 0.0 to 1.0
		t := float64(i) / float64(n-1)
		pts[i].X = t
		pts[i].Y = interp.Smoothstep(t)
	}
	return pts
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
