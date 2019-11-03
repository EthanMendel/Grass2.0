package worker

import "github.com/EthanMendel/Grass2.0/utils"

var PlantNSavg []float64

func CalculateNSaverage(plant *utils.Plant, start int, end int) {
	if end > len(plant.Dates) {
		end = len(plant.Dates)
	}
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
