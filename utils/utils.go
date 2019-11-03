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

type DataHeader struct {
	Headers []string
}

type DataRow struct {
	Date time.Time
	Data []float64
}

type Data struct {
	Header DataHeader
	Data   []DataRow
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

func ReadData(fileName string) (*Data, error) {
	records, err := ReadCSV(fileName)
	if err != nil {
		return nil, err
	}
	Header := DataHeader{}
	for i := 0; i < len(records[0]); i++ {
		Header.Headers = append(Header.Headers, records[0][i])
	}
	Data := Data{
		Header: Header,
		Data:   []DataRow{},
	}
	for i := 1; i < len(records); i++ {
		dt, te := time.Parse(DateLayout, records[i][0])
		if te != nil {
			log.Fatalf("Could not parse[%s] as date, row %d", records[i][0], i)
			return nil, te
		}
		//fmt.Println(dt.Format(DateLayout))
		dr := DataRow{
			Date: dt,
			Data: []float64{},
		}
		for j := 1; j < len(records[i]); j++ {
			f, fe := strconv.ParseFloat(records[i][j], 64)
			if fe != nil {
				log.Fatalf("Could not parse[%s], row:%d, col:%d", records[i][j], i, j)
				return nil, fe
			}
			dr.Data = append(dr.Data, f)
		}
		Data.Data = append(Data.Data, dr)
	}
	return &Data, nil
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
