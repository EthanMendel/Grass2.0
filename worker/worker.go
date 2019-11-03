package worker

import (
	"sync"

	"github.com/EthanMendel/Grass2.0/utils"
)

func CalculateNSaverage(plant *utils.Plant, start int, end int, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
	if end > len(plant.Dates) {
		end = len(plant.Dates)
	}
	//fmt.Printf("Processing %d entries\n", end-start)
	for row := start; row < end; row++ {
		total := 0.0
		count := 0.0
		for _, panel := range plant.Panels {
			if panel.Shading == "Good" {
				total += panel.Data[row]
				count++
			}
		}
		plant.Averages[row] = total / count
	}

}

func CalculateDifPer(plant *utils.Plant, start int, end int, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
	if end > len(plant.Dates) {
		end = len(plant.Dates)
	}
	//fmt.Printf("Processing %d entries\n", end-start)
	for _, panel := range plant.Panels {
		for row := start; row < end; row++ {
			if panel.Shading == "Bad" {
				panel.DifPer[row] = (plant.Averages[row] - panel.Data[row]) / plant.Averages[row]
				//fmt.Printf("Difference as a Percent of Panel %s is %d\n", panel.Name, panel.DifPer[row])
			}
		}
	}
}
