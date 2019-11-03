package main

import (
	"fmt"
	"log"

	"github.com/EthanMendel/Grass2.0/utils"
)

func main() {
	//worker.Hello()
	data, de := utils.ReadData("./Data/SyntheticData1Data.csv")
	if de != nil {
		log.Fatal(de)
		return
	}
	fmt.Printf("Data: rows[%d], col[%d]", len(data.Data), len(data.Header.Headers))
	//shading, se := utils.ReadCSV("./Data/SyntheticData1Shading.csv")
	//fmt.Println(data, de)
}
