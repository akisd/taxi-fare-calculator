package taxifarecalculator

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func ReadPositions(r *csv.Reader, d []*Position, b chan []*Position, wg *sync.WaitGroup) {
	go func() {
		var p1 *Position
		firstIteration := true

		for {
			line, err := r.Read()

			if err == io.EOF {
				b <- d
				close(b)
				break
			}

			if err != nil {
				close(b)
				log.Fatal(err)
			}

			if firstIteration {
				p1 = ConvertLineToPosition(line[0], line[1], line[2], line[3])
				d = append(d, p1)
				firstIteration = false
				continue
			}

			id, err := strconv.ParseUint(line[0], 10, 64)
			if err != nil {
				log.Println(err)
				continue
			}

			if id != p1.GetId() {
				p1 = ConvertLineToPosition(line[0], line[1], line[2], line[3])
				b <- d
				d = make([]*Position, 0)
				continue
			}

			p2 := ConvertLineToPosition(line[0], line[1], line[2], line[3])

			speed, err := p1.CalculateSpeed(p2)

			if err != nil {
				log.Println(err)
				continue
			}

			if speed > 100 {
				continue
			}

			d = append(d, p2)
			p1 = p2
		}

		wg.Done()
	}()
}

func WriteFareToFile(rideFare chan *Fare, wg *sync.WaitGroup) {
	csvf, err := os.Create("fare.csv")

	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(csvf)

	go func() {
		for f := range rideFare {
			fmt.Println("FARE: ", f.GetId(), f.GetFare())

			err := writer.Write([]string{fmt.Sprintf("%d", f.GetId()), fmt.Sprintf("%f", f.GetFare())})
			writer.Flush()

			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Done()
	}()
}

func ConvertLineToPosition(item1, item2, item3, item4 string) *Position {
	id, err := strconv.ParseUint(item1, 10, 64)

	if err != nil {
		log.Println(err)
		return nil
	}

	x, err := strconv.ParseFloat(item2, 64)

	if err != nil {
		log.Println(err)
		return nil
	}

	y, err := strconv.ParseFloat(item3, 64)

	if err != nil {
		log.Println(err)
		return nil
	}

	t, err := strconv.ParseInt(item4, 10, 32)

	if err != nil {
		log.Println(err)
		return nil
	}

	ut := time.Unix(t, 0)

	return NewPosition(id, x, y, ut)
}
