package main

import (
	"os"
	//"github.com/Arafatk/glot"
	"github.com/sirupsen/logrus"
	chart "github.com/wcharczuk/go-chart"
)

func ScatterPlot(x, y []float64) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	filename := "output.png"
	f, err := os.Create(filename)
	if err != nil {
		logrus.Error("Error: Cannot create output file")
		return
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		logrus.Error("Graph cannot be rendered")
		return
	}
}

func main() {
	x := []float64{1.0, 2.0, 3.0}
	y := []float64{2.0, 5.0, 7.0}
	ScatterPlot(x, y)

}