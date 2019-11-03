package experiment

import (
	"strconv"
	"sync"
	"time"

	"github.com/EthanMendel/Grass2.0/utils"
	"github.com/EthanMendel/Grass2.0/worker"
)

const MaxThreads = 20

func FindAvg(plant *utils.Plant, numThreads int) time.Duration {
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

func AvgExp(plant *utils.Plant, outFileName string) error {
	results := [][]string{{"Threads", "Duration"}}
	//worker.CalculateNSaverage(plant, 0, len(plant.Dates))
	for i := 1; i < MaxThreads+1; i++ {
		d := FindAvg(plant, i)
		h := []string{strconv.Itoa(i), strconv.Itoa(int(d))}
		results = append(results, h)
	}
	err := utils.CreateCSV(outFileName, results)
	return err
}

func FindDifPer(plant *utils.Plant, numThreads int) time.Duration {
	rtp := len(plant.Dates) / numThreads //rows to process
	start := 0
	var wg sync.WaitGroup
	wg.Add(numThreads)
	startTime := time.Now()
	for i := 0; i < numThreads; i++ {
		go worker.CalculateDifPer(plant, start, start+rtp, &wg)
		start += rtp
	}
	var wgo sync.WaitGroup //wait group overflow
	if start < len(plant.Dates) {
		wgo.Add(1)
		go worker.CalculateDifPer(plant, start, len(plant.Dates), &wgo)
	}
	wg.Wait()
	wgo.Wait()
	expTime := time.Since(startTime)
	return expTime

}

func DifPerExp(plant *utils.Plant, outFileName string) error {
	results := [][]string{{"Threads", "Duration"}}
	//worker.CalculateNSaverage(plant, 0, len(plant.Dates))
	for i := 1; i < MaxThreads+1; i++ {
		d := FindDifPer(plant, i)
		h := []string{strconv.Itoa(i), strconv.Itoa(int(d))}
		results = append(results, h)
	}
	err := utils.CreateCSV(outFileName, results)
	return err
}
