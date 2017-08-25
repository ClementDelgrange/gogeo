package algorithms

import (
	"testing"
	"github.com/ClementDelgrange/gogeo"
	"math"
)

func TestConvexHullJarvis(t *testing.T) {

}

func TestConvexHullGraham(t *testing.T) {

}


func TestFarLeftId(t *testing.T) {
	coords := make([]gogeo.Coordinate, 6)
	coords[0] = *gogeo.NewCoordinate(0, 0)
	coords[1] = *gogeo.NewCoordinate(1, 1)
	coords[2] = *gogeo.NewCoordinate(-1, 1)
	coords[3] = *gogeo.NewCoordinate(-1, -1)
	coords[4] = *gogeo.NewCoordinate(0, -5)
	coords[5] = *gogeo.NewCoordinate(0, 2)

	want := 3

	if a := FarLeftId(coords); a != want {
		t.Errorf("FarLeftId(%v) = %v, want %v", coords, a, want)
	}
}


func TestCounterclockwise(t *testing.T) {
	coords := make([]gogeo.Coordinate, 6)
	coords[0] = *gogeo.NewCoordinate(0, 0)
	coords[1] = *gogeo.NewCoordinate(1, 1)
	coords[2] = *gogeo.NewCoordinate(-1, 1)
	coords[3] = *gogeo.NewCoordinate(-1, -1)
	coords[4] = *gogeo.NewCoordinate(0, -5)
	coords[5] = *gogeo.NewCoordinate(0, 2)

	tests := []struct{
		i, j, k int
		want float64
	}{
		{0, 1, 2, 2},
		{0, 5, 1, -2},
		{0, 1, 3, 0},
	}

	for _, test := range tests {
		if a := Counterclockwise(coords, test.i, test.j, test.k); a != test.want {
			t.Errorf("Counterclockwise(%v, %v, %v, %v) = %v, want %v", coords, test.i, test.j, test.k, a, test.want)
		}
	}
}


func TestFarAway(t *testing.T) {
	coords := make([]gogeo.Coordinate, 6)
	coords[0] = *gogeo.NewCoordinate(0, 0)
	coords[1] = *gogeo.NewCoordinate(1, 1)
	coords[2] = *gogeo.NewCoordinate(-1, 1)
	coords[3] = *gogeo.NewCoordinate(-1, -1)
	coords[4] = *gogeo.NewCoordinate(0, -5)
	coords[5] = *gogeo.NewCoordinate(0, 2)

	tests := []struct{
		p, i, j int
		want int
	}{
		{0, 1, 2, 1},
		{0, 1, 4, 4},
		{0, 0, 3, 3},
	}

	for _, test := range tests {
		if a := FarAway(coords, test.p, test.i, test.j); a != test.want {
			t.Errorf("FarAway(%v, %v, %v, %v) = %v, want %v", coords, test.p, test.i, test.j, a, test.want)
		}
	}
}

//
func TestPolarAngle(t *testing.T) {
	coords := make([]gogeo.Coordinate, 6)
	coords[0] = *gogeo.NewCoordinate(0, 0)
	coords[1] = *gogeo.NewCoordinate(1, 1)
	coords[2] = *gogeo.NewCoordinate(-1, 1)
	coords[3] = *gogeo.NewCoordinate(-1, -1)
	coords[4] = *gogeo.NewCoordinate(0, -5)
	coords[5] = *gogeo.NewCoordinate(0, 2)

	tests := []struct{
		p, i int
		want float64
	}{
		{0, 0, 0},
		{0, 4, -math.Pi / 2},
		{0, 1, math.Pi / 4},
		{0, 3, -3 * math.Pi / 4},
	}

	for _, test := range tests {
		if a := PolarAngle(coords, test.p, test.i); a != test.want {
			t.Errorf("PolarAngle(%v, %v, %v) = %v, want %v", coords, test.p, test.i, a, test.want)
		}
	}
}