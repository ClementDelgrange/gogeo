package gogeo

import (
	"math"
	"fmt"
)


// A Coordinate represents an 2-dimensional coordinate.
// It is distinct from Point which is a Geometry and have a lot of other properties like spatial reference system.
type Coordinate [2]float64

// NewCoordinate create a new Coordinate with provided values
func NewCoordinate(x, y float64) *Coordinate {
	return &Coordinate{x, y}
}

// Equals determines if two coordinates are equals or not.
func (c Coordinate) Equals(other *Coordinate) bool {
	if c[0] != other[0] || c[1] != other[1] {
			return false
	}
	return true
}

// X returns the x coordinate. It assumed to be the first of the list.
func (c Coordinate) X() float64 {
	return c[0]
}

// Y returns the y coordinate. It assumed to be the second of the list.
func (c Coordinate) Y() float64 {
	return c[1]
}

// DistanceTo returns distance to an other coordinate.
func (c Coordinate) DistanceTo(other *Coordinate) float64 {
	d := math.Sqrt(math.Pow((c.X() - other.X()), 2) + math.Pow((c.Y() - other.Y()), 2))
	return d
}

// ShortestDistanceToLine returns the shortest distance from the coordinate to the line (c1-c2).
func (c Coordinate) ShortestDistanceToLine(c1 *Coordinate, c2 *Coordinate) (float64, error) {
	if c1.Equals(c2) {
		// c1 and c2 are equals : no line
		err := fmt.Errorf("The two points %v, %v defining the line are identical", c1, c2)
		return 0, err
	}

	if c1.X() == c2.X() {
		// vertical line
		return math.Abs(c1.X() - c.X()), nil
	}
	if c1.Y() == c2.Y() {
		// horizontal line
		return math.Abs(c1.Y() - c.Y()), nil
	}

	// general case : equation of the straight line c1-c2 (y = m.x + b)
	m := GetSlope(c1, c2)
	b := GetYIntercept(c1, m)
	// orthogonal projection
	return math.Abs((m*c.X() - c.Y() + b) / math.Sqrt(1 + m*m)), nil
}

// OrthogonalDistanceToSegment returns, if it exists, the orthogonal distance from the coordinate to the segment [c1-c2].
func (c Coordinate) OrthogonalDistanceToSegment(c1 *Coordinate, c2 *Coordinate) (float64, error) {
	dmin, err := c.ShortestDistanceToLine(c1, c2)
	if err != nil {
		return 0, err
	}

	// need to check if the shortest distance to the line is also shortest distance to the segment
	d1 := c.DistanceTo(c1)
	d2 := c.DistanceTo(c2)
	if d1 < dmin || d2 < dmin {
		err := fmt.Errorf("OrthogonalDistanceToSegment: orthogonal projection of %v isn't on the segment [%v-%v]", c, c1, c2)
		return math.Inf(1), err
	}
	return dmin, nil
}

// ShortestDistanceToSegment returns the shortest distance from the coordinate to the segment [c1-c2].
func (c Coordinate) ShortestDistanceToSegment(c1 *Coordinate, c2 *Coordinate) (float64, error) {
	dmin, err := c.ShortestDistanceToLine(c1, c2)
	if err != nil {
		return 0, err
	}

	// need to check if the shortest distance to the line is also shortest distance to the segment
	d1 := c.DistanceTo(c1)
	d2 := c.DistanceTo(c2)
	if d1 < dmin {
		return d1, nil
	}
	if d2 < dmin {
		return d2, nil
	}
	return dmin, nil
}

//
func AngularCoordinate(c Coordinate) float64 {
	return math.Atan2(c[1], c[0])
}