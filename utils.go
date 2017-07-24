package gogeo

import (
	"math/rand"
)

func RandomCoord(bbox BBox) Coordinate {
	x := rand.Float64() * (bbox.Xmax - bbox.Xmin) + bbox.Xmin
	y := rand.Float64() * (bbox.Ymax - bbox.Ymin) + bbox.Ymin
	return Coordinate{x, y}
}

func RandomCoords(n int, bbox BBox) []Coordinate {
	var coords []Coordinate
	for i := 0; i < n; i++ {
		coords = append(coords, RandomCoord(bbox))
	}
	return coords
}

func RandomPoint(bbox BBox) Point {
	x := rand.Float64() * (bbox.Xmax - bbox.Xmin) + bbox.Xmin
	y := rand.Float64() * (bbox.Ymax - bbox.Ymin) + bbox.Ymin
	return Point{coord: Coordinate{x, y}}
}

func RandomPoints(n int, bbox BBox) []Point {
	var points []Point
	for i := 0; i < n; i++ {
		points = append(points, RandomPoint(bbox))
	}
	return points
}

// GetSlope returns the slope of the straight line.
func GetSlope(c1, c2 *Coordinate) float64 {
	return (c1.Y() - c2.Y()) / (c1.X() - c2.X())
}

// GetYIntercept returns the ordonate where the line crosses the Y axis.
func GetYIntercept(c *Coordinate, m float64) float64 {
	return c.Y() - m * c.X()
}

