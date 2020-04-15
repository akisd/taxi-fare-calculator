package taxiFareCalculator

import (
	"errors"
	"math"
	"time"
)

type RidePosition struct {
	id        uint64
	lat, lng  float64
	timestamp time.Time
}

const earthRadius = float64(6371)

func NewRidePosition(id uint64, lat, lng float64, timestamp time.Time) *RidePosition {
	return &RidePosition{id, lat, lng, timestamp}
}

func (p1 *RidePosition) GetId() uint64 {
	return p1.id
}

func (p1 *RidePosition) GetLat() float64 {
	return p1.lat
}

func (p1 *RidePosition) GetLng() float64 {
	return p1.lng
}

func (p1 *RidePosition) GetTime() time.Time {
	return p1.timestamp
}

func (p1 *RidePosition) CalculateElapsedTime(p2 *RidePosition) float64 {
	return math.Abs(p1.timestamp.Sub(p2.timestamp).Hours())
}

func (p1 *RidePosition) CalculateHaversineDistance(p2 *RidePosition) float64 {
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

func (p1 *RidePosition) CalculateSpeed(p2 *RidePosition) (float64, error) {
	dist := p1.CalculateHaversineDistance(p2)
	dtime := p1.CalculateElapsedTime(p2)

	if dtime == 0 {
		return math.NaN(), errors.New("could not proceed with the speed calculation: elapsed time = 0")
	}

	return dist / dtime, nil
}
