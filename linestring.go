package gogeo

import (
	"github.com/ClementDelgrange/gogeo/geojson"
)


// A LineString is a figure comprising one or more line segments.
// It is defined by a sequence of coordinates.
type LineString struct {
	coords []Coordinate
}

// Creates a new LineString with the given coordinates
func NewLineString(c []Coordinate) *LineString {
	return &LineString{c}
}

// Returns the GeoJSON type of a polyline.
func (l LineString) GeoJSONType() string {
	return string(geojson.GeometryLineString)
}

// Returns the GeoJSON feature for the LineString.
func (l LineString) GeoJSON() string {
	c := make([][]float64, 0, len(l.coords))

	for _, coord := range l.coords {
		c = append(c, []float64{coord[0], coord[1]})
	}

	mapT, _ := geojson.NewLineStringGeoJSON(&c)
	return string(mapT)
}

// Returns the bounding box of the LineString.
func (l LineString) BBox() BBox {
		xmin := l.coords[0].X()
		ymin := l.coords[0].Y()
		xmax := l.coords[0].X()
		ymax := l.coords[0].Y()

		for _, coord := range l.coords {
			if coord.X() < xmin {
				xmin = coord.X()
			} else if coord.X() > xmax {
				xmax = coord.X()
			}
			if coord.Y() < ymin {
				ymin = coord.Y()
			} else if coord.Y() > ymax {
				ymax = coord.Y()
			}
		}

		return BBox{xmin, ymin, xmax, ymax}
	}

// Returns the points number of the linestring
func Len(l LineString) int {
	return len(l.coords)
}

