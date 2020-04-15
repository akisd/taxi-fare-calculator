package tests

import (
	. "../taxiFareCalculator"
	"fmt"
	"math"
	"testing"
	"time"
)

var t1 = time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC)
var t2 = time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)

func TestCalculateSegmentFare(t *testing.T) {
	v := 11.9
	p1 := NewRidePosition(0, 37.966660, 23.728308, t1)
	p2 := NewRidePosition(0, 37.966627, 23.728263, t1.Add(time.Hour))

	f := CalculateSegmentFare(p1, p2)
	f = math.Round(f*100) / 100

	if f != v {
		t.Errorf("Expected: %v\nReturned: %v", v, f)
	}
}

func TestCalculateSegmentIncreasedFare(t *testing.T) {
	v := 1.44
	p1 := NewRidePosition(0, 37.966660, 23.728308, t2)
	p2 := NewRidePosition(0, 37.976627, 23.728263, t2.Add(time.Second*30))

	f := CalculateSegmentFare(p1, p2)
	f = math.Round(f*100) / 100

	if f != v {
		t.Errorf("Expected: %v\nReturned: %v", v, f)
	}
}

func TestCalculateSegmentSingleFare(t *testing.T) {
	v := 0.82
	p1 := NewRidePosition(0, 37.966660, 23.728308, t1)
	p2 := NewRidePosition(0, 37.976627, 23.728263, t1.Add(time.Second*30))

	fmt.Println(p1.GetTime().Hour())
	fmt.Println(t2.Hour())

	f := CalculateSegmentFare(p1, p2)
	f = math.Round(f*100) / 100

	if f != v {
		t.Errorf("Expected: %v\nReturned: %v", v, f)
	}
}

func TestCalculateSegmentFare2(t *testing.T) {
	v := 106.46
	p1 := NewRidePosition(0, 37.966660, 23.728308, t1.Add(time.Hour*11))
	p2 := NewRidePosition(0, 38.986627, 24.728263, t2.Add(time.Second*30))

	f := CalculateSegmentFare(p1, p2)
	f = math.Round(f*100) / 100

	if f != v {
		t.Errorf("Expected: %v\nReturned: %v", v, f)
	}
}
