package utils

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

const DateLayout = "1/2/06"

func ReadCSV(fileName string) ([][]string, error) {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return records, nil
}

type Panel struct {
	Name    string
	Shading string
	Data    []float64
}

type Plant struct {
	Dates  []time.Time
	Panels map[string]*Panel
}

func ReadPlant(fileName string) (*Plant, error) {
	records, err := ReadCSV(fileName)
	if err != nil {
		return nil, err
	}
	plant := Plant{
		Dates:  []time.Time{},
		Panels: map[string]*Panel{},
	}
	for row := 1; row < len(records); row++ {
		dt, te := time.Parse(DateLayout, records[row][0])
		if te != nil {
			log.Fatalf("Could not parse[%s] as date, row %d", records[row][0], row)
			return nil, te
		}
		//fmt.Println(dt.Format(DateLayout))
		plant.Dates = append(plant.Dates, dt)
	}
	for col := 1; col < len(records[0]); col++ {
		panel := Panel{
			Name:    records[0][col],
			Shading: "TBD",
			Data:    []float64{},
		}
		for row := 1; row < len(records); row++ {
			f, fe := strconv.ParseFloat(records[row][col], 64)
			if fe != nil {
				log.Fatalf("Could not parse[%s], row:%d, col:%d", records[row][col], row, col)
				return nil, fe
			}
			panel.Data = append(panel.Data, f)
		}
		plant.Panels[panel.Name] = &panel
	}
	return &plant, nil
}