package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	xys,b,a, err := readData("Neutrality.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	_ = xys
	flag := 1
	err = plotData("Neutrality.png", xys,b,a,flag) // Regression Line 
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}


	xys,b,a, err = readData("Parity.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	_ = xys

	flag = 0
	err = plotData("Parity.png", xys,b,a,flag)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}

}

type xy struct{ x []float64 ; y[]float64 }


func readData(path string) (plotter.XYs,float64, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil,0,0, err
	}
	defer f.Close()
	var listofx []float64
	var listofy []float64
	
	data := xy{ 

        x: listofx, 

        y: listofy, 

    } 


	var xys plotter.XYs
	s := bufio.NewScanner(f)
	//var i int
	//i = 0

	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
        data.y =  append(data.y,y)
        data.x =  append(data.x,x)
	}
	if err := s.Err(); err != nil {
		return nil,0,0, fmt.Errorf("could not scan: %v", err)
	}

    b, a := stat.LinearRegression(data.x, data.y, nil, false) 
    fmt.Println(b,a)

	return xys, b, a, nil
}

func plotData(path string, xys plotter.XYs,b,a float64,flag int) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	if flag == 1{
		// y = ax + b ; have a,b
		// create a linear regression line using a and b values obtained
		l, err := plotter.NewLine(plotter.XYs{
			{0.3, 0.3*a +b}, {0.9, 0.9*a + b},
		})
		if err != nil {
			return fmt.Errorf("could not create line: %v", err)
		}
		p.Add(l)
	}
	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}