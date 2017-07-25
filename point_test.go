package gogeo

import (
	"testing"
)


func TestNewPoint(t *testing.T) {
	c := *NewCoordinate(2, 3.45)
	p := *NewPoint(&c)

	//if p == nil {
	//	t.Error("Expected to get a pointer to a new point, but got nil instead.")
	//}

	if p.coord[0] != 2 {
		t.Error("Expected to be able to specify 2 as point first coordinate value, but got %f instead", p.coord[0])
	}

	if p.coord[1] != 3.45 {
		t.Error("Expected to be able to specify 3.45 as point second coordinate value, but got %f instead", p.coord[1])
	}
}

func TestPoint_X(t *testing.T) {
	c := *NewCoordinate(2, 3.45)
	p := *NewPoint(&c)

	if p.X() != 2 {
		t.Error("Expected to be able to read 2 as X coordinate, but got %f instead", p.X())
	}
}

func TestPoint_Y(t *testing.T) {
	c := *NewCoordinate(2, 3.45)
	p := *NewPoint(&c)

	if p.Y() != 3.45 {
		t.Error("Expected to be able to read 3.45 as Y coordinate, but got %f instead", p.Y())
	}
}


