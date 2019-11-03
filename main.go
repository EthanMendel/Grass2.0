package main

import (
	"log"

	"github.com/EthanMendel/Grass2.0/utils"
	"github.com/EthanMendel/Grass2.0/worker"
)

func main() {
	//worker.Hello()
	plant, de := utils.ReadPlant("./Data/SyntheticData1Data.csv", "./Data/SyntheticData1Shading.csv")
	if de != nil {
		log.Fatal(de)
		return
	}
	worker.CalculateNSaverage(plant, 0, len(plant.Dates))
}
