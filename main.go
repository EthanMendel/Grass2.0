package main

import (
	"fmt"
	"log"

	"github.com/EthanMendel/Grass2.0/experiment"

	"github.com/EthanMendel/Grass2.0/utils"
)

func main() {
	//worker.Hello()
	plant, de := utils.ReadPlant("./Data/SyntheticData1Data.csv", "./Data/SyntheticData1Shading.csv")
	if de != nil {
		log.Fatal(de)
		return
	}
	bestTime := experiment.AvgExp(plant, 1)
	bestThreads := 1
	//worker.CalculateNSaverage(plant, 0, len(plant.Dates))
	for i := 2; i < 21; i++ {
		d := experiment.AvgExp(plant, i)
		if d < bestTime {
			bestTime = d
			bestThreads = i
		}
	}
	fmt.Printf("\nProcessed %d Panels", len(plant.Panels))
	fmt.Printf("\nBest time using %d threads is %s \n", bestThreads, utils.FmtDuration(bestTime))
}
