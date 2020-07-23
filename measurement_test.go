package turf

import (
	"testing"
)

func TestDistance(t *testing.T) {
	d := Distance(-77.03653, 38.89768, -77.05173, 38.8973)
	if d != 1.317556974720262 {
		t.Errorf("error calculating the distance")
	}
}

func TestDistancePoint(t *testing.T) {
	p1 := Point{Lng: -77.03653, Lat: 38.89768}
	p2 := Point{Lng: -77.05173, Lat: 38.8973}
	d := DistancePoint(p1, p2)
	if d != 1.317556974720262 {
		t.Errorf("error calculating the distance")
	}
}
