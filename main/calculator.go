package main

import (
	. "../../taxi-fare-calculator"
	"encoding/csv"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	batch := make(chan []*Position)

	log.Print("Reading rides from csv file...")

	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/testdata/positions.csv")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	var data []*Position

	log.Print("Calculating fare...")

	wg.Add(3)

	ReadPositions(reader, data, batch, &wg)
	rideFare := CalculateFare(batch, &wg)
	WriteFareToFile(rideFare, &wg)

	wg.Wait()

	log.Print("Calculation finished.")
}
