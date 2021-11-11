package main

import (
	"fmt"
	"gifhelper"
	"flag"
)

func main() {
	
	flag.Parse()
	arguments := flag.Args()
	command := arguments[0]
	fmt.Println("Value of command flag is : ",command)

	width := 1.0e23
	galaxies := []Galaxy{}
	var initialUniverse *Universe
	numGens := 100000
	time := 2e14

	if command == "galaxy" { //TO BE RUN FOR 500 
		//g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22) //changed from 500 to 10 
		g0 := InitializeGalaxy(300, 4e21, 7e22, 2e22) //changed from 500 to 10 
		galaxies = []Galaxy{g0}
		initialUniverse = InitializeUniverse(galaxies, width) //returning adress of inital universe
		numGens = 50000
	}

	if command == "jupiter" {  
		g0 := CreateJupiterSystem(4e21, 7e22, 2e22) 
		galaxies = []Galaxy{g0}
		initialUniverse = InitializeUniverse(galaxies, width) //returning adress of inital universe
		//numGens = 100000
	}

	if command == "collision" { //TO BE RUN FOR 500
		g0 := InitializeGalaxy(10, 4e21, 4e22, 3e22) //change to 500 
		g1 := InitializeGalaxy(10, 4e21, 3e22, 4e22)//change to 500 
		galaxies = []Galaxy{g0, g1}
		initialUniverse = InitializeUniverse(galaxies, width) //returning adress of inital universe
		//numGens = 100000
 	}

	theta = 0.5

	timePoints := BarnesHut(initialUniverse, numGens, time, theta)

	fmt.Println("Simulation run. Now drawing images.")
	canvasWidth := 1000
	frequency := 1000
	scalingFactor := 3e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "galaxy")
	fmt.Println("GIF drawn.")

}
