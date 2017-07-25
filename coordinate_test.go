package gogeo

import (
	"testing"
	"math"
)


var epsilon float64 = 1e-10

func float64AlmostEqual(x, y float64) bool {
	return math.Abs(x - y) <= epsilon
}


func TestNewCoordinate(t *testing.T) {
	c := *NewCoordinate(2, -3.45)

	//if p == nil {
	//	t.Error("Expected to get a pointer to a new coordinate, but got nil instead.")
	//}

	if c[0] != 2 {
		t.Errorf("Expected to be able to specify 2 as first coordinate value, but got %f instead", c[0])
	}

	if c[1] != -3.45 {
		t.Errorf("Expected to be able to specify 3.45 as second coordinate value, but got %f instead", c[1])
	}
}

func TestCoordinate_Equals(t *testing.T) {
	tests := []struct {
		x1, y1 float64
		x2, y2 float64
		want bool
	}{
		{1, 0, 1, 0, true},
		{1.25, 2.37, 1.25, 2.37, true},
		{1, 0, 0, 0, false},
	}

	for _, test := range tests {
		c1 := *NewCoordinate(test.x1, test.y1)
		c2 := *NewCoordinate(test.x2, test.y2)

		if a := c1.Equals(&c2); a != test.want {
			t.Errorf("%v.Equals(%v) = %v, want %v", c1, c2, a, test.want)
		}

		if a := c2.Equals(&c1); a != test.want {
			t.Errorf("%v.Equals(%v) = %v, want %v", c2, c1, a, test.want)
		}
	}
}

func TestCoordinate_DistanceTo(t *testing.T) {
	tests := []struct {
		x1, y1 float64
		x2, y2 float64
		want float64
	}{
		{1, 0, 1, 0, 0},
		{1, 0, 0, 0, 1},
		{1.25, 2.37, 0, 0, 2.679440240050149},
	}

	for _, test := range tests {
		c1 := *NewCoordinate(test.x1, test.y1)
		c2 := *NewCoordinate(test.x2, test.y2)

		if a := c1.DistanceTo(&c2); a != test.want {
			t.Errorf("%v.DistanceTo(%v) = %v, want %v", c1, c2, a, test.want)
		}

		if a := c2.DistanceTo(&c1); a != test.want {
			t.Errorf("%v.DistanceTo(%v) = %v, want %v", c2, c1, a, test.want)
		}
	}
}

func TestCoordinate_ShortestDistanceToLine(t *testing.T) {
	tests := []struct {
		x, y float64
		x1, y1 float64
		x2, y2 float64
		want float64
	}{
		{0, 2, 0, 0, 0, 0, 0}, // the two points defining the line are equals
		{0, 0, 0, 0, 1, 1, 0}, // point is equals to one of the point defining the line
		{1, 2, 0, 0, 1, 0, 2}, // horizontal line
		{0, 1.2, 1.5, 1, 1.5, 2, 1.5}, // vertical line
		{0, 5, 1.5, 1, 1.5, 2, 1.5},  // vertical line, projection doesn't crossing segment
		{0, 2, 1, -1, 3, 1, math.Sqrt(8)}, // general line
		{-2, 0, 1, -1, 3, 1, math.Sqrt(8)}, // general line, projection doesn't corssing segment
	}

	for _, test := range tests {
		c := *NewCoordinate(test.x, test.y)
		c1 := *NewCoordinate(test.x1, test.y1)
		c2 := *NewCoordinate(test.x2, test.y2)


		if a, _ := c.ShortestDistanceToLine(&c1, &c2); !float64AlmostEqual(a, test.want) {
			t.Errorf("%v.ShortestDistanceToLine(%v, %v) = %v, want %v", c, c1, c2, a, test.want)
		}
	}
}

func TestCoordinate_ShortestDistanceToSegment(t *testing.T) {
	tests := []struct {
		x, y float64
		x1, y1 float64
		x2, y2 float64
		want float64
	}{
		{0, 2, 0, 0, 0, 0, 0}, // the two points defining the line are equals
		{0, 0, 0, 0, 1, 1, 0}, // point is equals to one of the point defining the line
		{1, 2, 0, 0, 1, 0, 2}, // horizontal line
		{0, 1.2, 1.5, 1, 1.5, 2, 1.5}, // vertical line
		{0, 5, 1.5, 1, 1.5, 2, math.Sqrt(11.25)}, // vertical line, projection doesn't crossing segment
		{0, 2, 1, -1, 3, 1, math.Sqrt(8)}, // general line
		{-2, 0, 1, -1, 3, 1, math.Sqrt(10)}, // general line, projection doesn't corssing segment
	}

	for _, test := range tests {
		c := *NewCoordinate(test.x, test.y)
		c1 := *NewCoordinate(test.x1, test.y1)
		c2 := *NewCoordinate(test.x2, test.y2)

		if a, _ := c.ShortestDistanceToSegment(&c1, &c2); !float64AlmostEqual(a, test.want) {
			t.Errorf("%v.ShortestDistanceToSegment(%v, %v) = %v, want %v", c, c1, c2, a, test.want)
		}
	}
}

func TestCoordinate_OrthogonalDistanceToSegment(t *testing.T) {
	tests := []struct {
		x, y float64
		x1, y1 float64
		x2, y2 float64
		want float64
	}{
		{1, 2, 0, 0, 1, 0, 2},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 5, 1.5, 1, 1.5, 2, 1.5},
		{0, 1.2, 1.5, 1, 1.5, 2, 1.5},
		{0, 2, 1, 0, 1, 0, math.Sqrt(5)},
		{0, 2, 1, -1, 3, 1, math.Sqrt(8)},
		{-2, 0, 1, -1, 3, 1, math.Sqrt(8)},
	}

	for _, test := range tests {
		c := *NewCoordinate(test.x, test.y)
		c1 := *NewCoordinate(test.x1, test.y1)
		c2 := *NewCoordinate(test.x2, test.y2)

		if a, _ := c.OrthogonalDistanceToSegment(&c1, &c2); !float64AlmostEqual(a, test.want) {
			t.Errorf("%v.OrthogonalDistanceToSegment(%v, %v) = %v, want %v", c, c1, c2, a, test.want)
		}
	}
}

