package experiment

import (
	"sync"
	"time"

	"github.com/EthanMendel/Grass2.0/utils"
	"github.com/EthanMendel/Grass2.0/worker"
)

func AvgExp(plant *utils.Plant, numThreads int) time.Duration {
	rtp := len(plant.Dates) / numThreads //rows to process
	start := 0
	var wg sync.WaitGroup
	wg.Add(numThreads)
	startTime := time.Now()
	for i := 0; i < numThreads; i++ {
		go worker.CalculateNSaverage(plant, start, start+rtp, &wg)
		start += rtp
	}
	var wgo sync.WaitGroup //wait group overflow
	if start < len(plant.Dates) {
		wgo.Add(1)
		go worker.CalculateNSaverage(plant, start, len(plant.Dates), &wgo)
	}
	wg.Wait()
	wgo.Wait()
	expTime := time.Since(startTime)
	return expTime

}
