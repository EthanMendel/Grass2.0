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
	experiment.AvgExp(plant, "./Results/ThreadsForAverageCalculation.csv")
	fmt.Printf("\nProcessed %d Panels", len(plant.Panels))
}
