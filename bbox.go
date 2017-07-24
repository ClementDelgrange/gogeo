package gogeo

import "math"

// BBox is a bounding box of any geometry.
type BBox struct {
	Xmin, Ymin, Xmax, Ymax float64
}

// Creates a new BBox with provided floats
func NewBBox(c *[4]float64) *BBox {
	return &BBox{c[0], c[1], c[2], c[3]}
}



// Returns the input bounding box expanding to include this passed in.
func (b BBox) Union(other BBox) BBox {
	Xmin := math.Min(b.Xmin, other.Xmin)
	Xmax := math.Max(b.Xmax, other.Xmax)
	Ymin := math.Min(b.Ymin, other.Ymin)
	Ymax := math.Max(b.Ymax, other.Ymax)
	bbox := BBox{Xmin, Ymin, Xmax, Ymax}
	return bbox
}

// Determines if the point is within the bounding box.
// Points on the boundary are considered inside.
func (b *BBox) Contains(p *Point) bool {
	if p.X() <= b.Xmin || p.X() >= b.Xmax {
		return false
	}
	if p.Y() <= b.Ymin || p.Y() >= b.Ymax {
		return false
	}
	return true
}

// Converts the bounding box into a Polygon.
func (b BBox) ToPolygon() Polygon {
	coords := make([][]Coordinate, 1)
	ring := make([]Coordinate, 5)
	ring[0] = Coordinate{b.Xmin, b.Ymin}
	ring[1] = Coordinate{b.Xmin, b.Ymax}
	ring[2] = Coordinate{b.Xmax, b.Ymax}
	ring[3] = Coordinate{b.Xmax, b.Ymin}
	ring[4] = Coordinate{b.Xmin, b.Ymin}
	coords[0] = ring
	return Polygon{coords}
}