package main

import (
	"fmt"
	"log"

	"github.com/EthanMendel/Grass2.0/experiment"
	"github.com/EthanMendel/Grass2.0/utils"
	"github.com/EthanMendel/Grass2.0/worker"
)

func main() {
	//worker.Hello()
	plant, de := utils.ReadPlant("./Data/SyntheticData2Data.csv", "./Data/SyntheticData2Shading.csv")
	if de != nil {
		log.Fatal(de)
		return
	}
	//experiment.FindAvg(plant, 5)
	//experiment.FindDifPer(plant, 5)
	experiment.AvgExp(plant, "./Results/ThreadsForAverageCalculation.csv")
	experiment.DifPerExp(plant, "./Results/ThreadsForDifPerCalculation.csv")
	worker.SetupResults(plant, "./Results/DifferenceAsPercent.csv")

	fmt.Printf("\nProcessed %d Panels", len(plant.Panels))
}
