package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()

	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// A quadratic function x^2
	quad := plotter.NewFunction(func(x float64) float64 { return x * x })
	quad.Color = color.RGBA{B: 255, A: 255}

	// An exponential function 2^x
	exp := plotter.NewFunction(func(x float64) float64 { return math.Pow(2, x) })
	exp.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	exp.Width = vg.Points(2)
	exp.Color = color.RGBA{G: 255, A: 255}

	// The sine function, shifted and scaled
	// to be nicely visible on the plot.
	sin := plotter.NewFunction(func(x float64) float64 { return 10*math.Sin(x) + 50 })
	sin.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
	sin.Width = vg.Points(4)
	sin.Color = color.RGBA{R: 255, A: 255}
	
	// The sine function, shifted and scaled
	// to be nicely visible on the plot.
	sin2 := plotter.NewFunction(func(x float64) float64 { return 10*math.Sin(x) + 80 })
	sin2.Dashes = []vg.Length{vg.Points(4), vg.Points(8)}
	sin2.Width = vg.Points(4)
	sin2.Color = color.RGBA{G: 90,R: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(quad, exp, sin,sin2)
	p.Legend.Add("x^2", quad)
	p.Legend.Add("2^x", exp)
	p.Legend.Add("10*sin(x)+50", sin)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = 0
	p.X.Max = 10
	p.Y.Min = 0
	p.Y.Max = 100


	x := []float64{0, 0.25 * math.Pi, 0.75 * math.Pi}
	weights := []float64{1, 2, 2.5}
	cmean := stat.CircularMean(x, weights)

	fmt.Printf("The circular mean is %.5f.\n", cmean)


	x2 := []float64{0, 0.25 * math.Pi, 99}
	cdf := stat.CDF(10, 1, x2, weights)
	
	
	fmt.Printf("The cdf  %1.1f.\n", cdf)

	dist2 := distuv.Normal{
		Mu:    20,
		Sigma: 50,
	}.CDF(2)
	fmt.Printf("d %1.1f ± %0.1v\n", dist2, dist2)
	// Create a normal distribution
	dist := distuv.Normal{
		Mu:    100,
		Sigma: 5,
	}

	data := make([]float64, 1e5)

	// Draw some random values from the standard normal distribution
	for i := range data {
		data[i] = dist.Rand()
	}

	mean, std := stat.MeanStdDev(data, nil)
	meanErr := stat.StdErr(std, float64(len(data)))

	fmt.Printf("mean= %1.1f ± %0.1v\n", mean, meanErr)





	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "functions.png"); err != nil {
		panic(err)
	}
}