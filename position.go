package taxifarecalculator

import (
	"errors"
	"math"
	"time"
)

type Position struct {
	id        uint64
	lat, lng  float64
	timestamp time.Time
}

const earthRadius = float64(6371)

func NewPosition(id uint64, lat, lng float64, timestamp time.Time) *Position {
	return &Position{id, lat, lng, timestamp}
}

func (p1 *Position) GetId() uint64 {
	return p1.id
}

func (p1 *Position) GetLat() float64 {
	return p1.lat
}

func (p1 *Position) GetLng() float64 {
	return p1.lng
}

func (p1 *Position) GetTime() time.Time {
	return p1.timestamp
}

func (p1 *Position) CalculateElapsedTime(p2 *Position) float64 {
	return math.Abs(p1.timestamp.Sub(p2.timestamp).Hours())
}

func (p1 *Position) CalculateHaversineDistance(p2 *Position) float64 {
	dLat := (p2.lat - p1.lat) * (math.Pi / 180)
	dLng := (p2.lng - p1.lng) * (math.Pi / 180)

	lat1 := p1.lat * (math.Pi / 180)
	lat2 := p2.lat * (math.Pi / 180)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLng/2) * math.Sin(dLng/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRadius
}

func (p1 *Position) CalculateSpeed(p2 *Position) (float64, error) {
	dist := p1.CalculateHaversineDistance(p2)
	dtime := p1.CalculateElapsedTime(p2)

	if dtime == 0 {
		return math.NaN(), errors.New("could not proceed with the speed calculation: elapsed time = 0")
	}

	return dist / dtime, nil
}
