package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const DateLayout = "1/2/06"

type Panel struct {
	Name    string
	Shading string
	Data    []float64
	DifPer  []float64 //difference percentages
}

type Plant struct {
	Dates    []time.Time
	Averages []float64
	Panels   map[string]*Panel
}

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

func ReadPlant(dataFileName string, shadingFilename string) (*Plant, error) {
	records, err := ReadCSV(dataFileName)
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
			DifPer:  []float64{},
		}
		panel.DifPer = make([]float64, len(plant.Dates), len(plant.Dates))
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
	plant.Averages = make([]float64, len(plant.Dates), len(plant.Dates))
	ReadShading(&plant, shadingFilename)
	return &plant, nil
}

func ReadShading(plant *Plant, fileName string) error {
	records, err := ReadCSV(fileName)
	if err != nil {
		log.Fatalf("Could not read Shading file[%s]: %s", fileName, err.Error())
		return err
	}
	for col := 1; col < len(records[0]); col++ {
		plant.Panels[records[0][col]].Shading = records[1][col]
	}
	return nil
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Millisecond)
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	if s == 0 {
		return fmt.Sprintf("%02d", ms)
	} else {
		return fmt.Sprintf("%02d:%02d", s, ms)
	}
}

func CreateCSV(fileName string, data [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not Create File %s", err.Error())
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatalf("Could not Create File %s", err.Error())
			return err
		}
	}
	return nil
}
