package gogeo

import (
	"github.com/ClementDelgrange/gogeo/geojson"
)

// A point is define as a simple coordinate.
type Point struct {
	coord Coordinate
}

// Create a new Point with provided coordinate
func NewPoint(c *Coordinate) *Point {
	return &Point{*c}
}

// Returns the GeoJSON type of a point.
func (p Point) GeoJSONType() string {
	return string(geojson.GeometryPoint)

}

// Returns the GeoJSON feature for the point.
func (p Point) GeoJSON() string {
	mapT, _ := geojson.NewPointGeoJSON(&[]float64{p.X(), p.Y()})
	return string(mapT)
}

// Returns the bounding box of the point.
func (p Point) BBox() BBox {
	bbox := BBox{p.X(), p.Y(), p.X(), p.Y()}
	return bbox
}

// Returns the first coordinate of the point.
func (p *Point) X() float64 {
	return p.coord.X()
}

// Return the second coordinate of the point.
func (p *Point) Y() float64 {
	return p.coord.Y()
}

// Determines if a point is equals to an other.
func (p *Point) Equals(other *Point) bool {
	return p.coord.Equals(&other.coord)
}

// Return the distance between to points.
func (p *Point) DistanceTo(other *Point) float64 {
	return p.coord.DistanceTo(&other.coord)
}