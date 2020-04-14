package main

import (
	. "../../taxiFareCalculator"
	"encoding/csv"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	batch := make(chan []*RidePosition)

	log.Print("Reading rides from csv file...")

	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/taxiFareCalculator/data/paths.csv")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	var data []*RidePosition

	log.Print("Calculating fare...")

	wg.Add(3)

	ReadRidePositions(reader, data, batch, &wg)
	rideFare := CalculateRideFare(batch, &wg)
	WriteFareToFile(rideFare, &wg)

	wg.Wait()

	log.Print("Calculation finished.")
}
