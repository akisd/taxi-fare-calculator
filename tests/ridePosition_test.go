package tests

import (
	. "../taxiFareCalculator"
	"math"
	"testing"
	"time"
)

var p1 = NewRidePosition(0, 37.966660, 23.728308, time.Now())
var p2 = NewRidePosition(0, 37.966627, 23.728263, time.Now().Add(time.Hour))

func TestCalculateElapsedTime(t *testing.T) {
	v := 1.0
	tm := math.Floor(p1.CalculateElapsedTime(p2))

	if tm != v {
		t.Errorf("Expected: %v\nReturned: %v", v, tm)
	}
}

func TestCalculateSpeed(t *testing.T) {
	v := 0.0054
	s, _ := p1.CalculateSpeed(p2)
	s = math.Round(s*10000) / 10000

	if s != v {
		t.Errorf("Expected: %v\nReturned: %v", v, s)
	}
}

func TestCalculateHaversineDistance(t *testing.T) {
	v := 0.0054
	d := p1.CalculateHaversineDistance(p2)
	d = math.Round(d*10000) / 10000

	if d != v {
		t.Errorf("Expected: %v\nReturned: %v", v, d)
	}
}
