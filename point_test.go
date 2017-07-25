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
		t.Errorf("Expected to be able to specify 2 as point first coordinate value, but got %f instead", p.coord[0])
	}

	if p.coord[1] != 3.45 {
		t.Errorf("Expected to be able to specify 3.45 as point second coordinate value, but got %f instead", p.coord[1])
	}
}

func TestPoint_X(t *testing.T) {
	p := *NewPoint(NewCoordinate(2, 3.45))

	if p.X() != 2 {
		t.Errorf("Expected to be able to read 2 as X coordinate, but got %f instead", p.X())
	}
}

func TestPoint_Y(t *testing.T) {
	p := *NewPoint(NewCoordinate(2, 3.45))

	if p.Y() != 3.45 {
		t.Errorf("Expected to be able to read 3.45 as Y coordinate, but got %f instead", p.Y())
	}
}

func TestPoint_BBox(t *testing.T) {

}

func TestPoint_Equals(t *testing.T) {
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
		p1 := *NewPoint(NewCoordinate(test.x1, test.y1))
		p2 := *NewPoint(NewCoordinate(test.x2, test.y2))

		if a := p1.Equals(&p2); a != test.want {
			t.Errorf("%v.Equals(%v) = %v, want %v", p1, p2, a, test.want)
		}

		if a := p2.Equals(&p1); a != test.want {
			t.Errorf("%v.Equals(%v) = %v, want %v", p2, p1, a, test.want)
		}
	}
}

func TestPoint_DistanceTo(t *testing.T) {
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
		p1 := *NewPoint(NewCoordinate(test.x1, test.y1))
		p2 := *NewPoint(NewCoordinate(test.x2, test.y2))

		if a := p1.DistanceTo(&p2); a != test.want {
			t.Errorf("%v.DistanceTo(%v) = %v, want %v", p1, p2, a, test.want)
		}

		if a := p2.DistanceTo(&p1); a != test.want {
			t.Errorf("%v.DistanceTo(%v) = %v, want %v", p2, p1, a, test.want)
		}
	}
}

func TestPoint_GeoJSONType(t *testing.T) {
	p := *NewPoint(NewCoordinate(2, 3.45))

	if p.GeoJSONType() != "Point" {
		t.Errorf("Expected 'Point' as GeoJSON type, but got %s instead", p.GeoJSONType())
	}
}

func TestPoint_GeoJSON(t *testing.T) {
	p := *NewPoint(NewCoordinate(2, 3.45))
	want := "{\"type\":\"Feature\",\"geometry\":{\"type\":\"Point\",\"coordinates\":[2,3.45]}}"

	if p.GeoJSON() != want {
		t.Errorf("Expected GeoJSON geometry not obtained, want %s, but got %s", want, p.GeoJSON())
	}
}


