package gogeo

import (
	"github.com/ClementDelgrange/gogeo/geojson"
)

// A polygon is a figure bounded by a closed linestring (more than one line segment).
// It is defined by a sequence of coordinates.
// It may have one or more holes which are also bounded by closed linestring.
// If any ring cross each other, the feature is invalid.
type Polygon struct {
	coords [][]Coordinate
}

// Creates a new Polygon with the given coordinates
func NewPolygon(c [][]Coordinate) *Polygon {
	return &Polygon{c}
}

// Returns the GeoJSON type of a polygon.
func (p Polygon) GeoJSONType() string {
	return string(geojson.GeometryPolygon)
}

// Returns the GeoJSON feature for the polygon.
func (p Polygon) GeoJSON() string {
	c := make([][][]float64, 0, len(p.coords))

	for _, ring := range p.coords {
		cring := make([][]float64, 0, len(ring))
		for _, coord := range ring {
			cring = append(cring, []float64{coord[0], coord[1]})
		}
		c = append(c, cring)
	}

	mapT, _ := geojson.NewPolygonGeoJSON(&c)
	return string(mapT)
}

// Returns the bounding box of the polygon.
func (p Polygon) BBox() BBox {
	xmin := p.coords[0][0].X()
	ymin := p.coords[0][0].Y()
	xmax := p.coords[0][0].X()
	ymax := p.coords[0][0].Y()

	for _, ring := range p.coords {
		for _, coord := range ring {
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
	}

	return BBox{xmin, ymin, xmax, ymax}
}

