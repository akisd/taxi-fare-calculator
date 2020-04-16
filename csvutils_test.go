package taxifarecalculator

import (
	"fmt"
	"testing"
	"time"
)

func TestConvertLineToPosition(t *testing.T) {
	p := ConvertLineToPosition("1", "37.966660", "23.728308", "1405594957")

	fmt.Println(time.Unix(1405594957, 0))

	if p.GetId() != 1 && p.GetLat() != 37.966660 && p.GetLng() != 23.728308 && p.GetTime().Unix() == 1405594957 {
		t.Error("Wrong values were returned from function")
	}

	if p == nil {
		t.Error("Error while returning RidePosition")
	}
}

func TestFailedConvertLineToPosition(t *testing.T) {
	p := ConvertLineToPosition("", "37.966660", "23.728308", "1405594957")

	if p != nil {
		t.Error("Expected nil RidePosition\n Returned: ", p)
	}
}
