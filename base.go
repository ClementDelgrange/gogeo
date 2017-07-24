// Package gogeo contains geometry primitives and basic geometrical operations
package gogeo

import (
	"github.com/ClementDelgrange/gogeo/geojson"
)


// A geometry is an interface that represents methods shared by all of the geometries.
type Geometry interface {
	GeoJSONType() string
	GeoJSON() string
	BBox() BBox
}

// A Collection is a collection of heterogenous geometries.
type Collection struct {
	geoms []Geometry
}

// Creates a new Collection with provided geometries
func NewCollection(g ...Geometry) *Collection {
	coll := Collection{}
	coll.geoms = append(coll.geoms, g...)
	return &coll
}

// Returns the GeoJSON type of a geometry collection
func (c Collection) GeoJSONType() string {
	return string(geojson.GeometryCollection)
}

// Returns the global bounding box for of all geometries of a collection.
func (c Collection) BBox() BBox {
	bbox := c.geoms[0].BBox()
	for i := 1; i < len(c.geoms); i++ {
		bbox = bbox.Union(c.geoms[i].BBox())
	}
	return bbox
}
