package worker

import (
	"strconv"
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
			panel.DifPer[row] = (plant.Averages[row] - panel.Data[row]) / plant.Averages[row]
			//fmt.Printf("Difference as a Percent of Panel %s is %d\n", panel.Name, panel.DifPer[row])
		}
	}
}

func SetupResults(plant *utils.Plant, outFileName string) error {
	results := [][]string{}
	//worker.CalculateNSaverage(plant, 0, len(plant.Dates))
	headers := []string{"Dates"}
	for _, panel := range plant.Panels {
		if panel.Shading == "Bad" {
			headers = append(headers, panel.Name)
		}
	}
	results = append(results, headers)
	for row := 0; row < len(plant.Dates); row++ {
		data := []string{}
		data = append(data, plant.Dates[row].Format(utils.DateLayout))
		for i := 1; i < len(headers); i++ {
			if plant.Panels[headers[i]].Shading == "Bad" {
				data = append(data, strconv.FormatFloat(plant.Panels[headers[i]].DifPer[row], 'f', 6, 64))
			}
		}
		results = append(results, data)
	}
	err := utils.CreateCSV(outFileName, results)
	return err
}
