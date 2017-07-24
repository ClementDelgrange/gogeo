package geojson

import (
	"encoding/json"
)

// Geometry types supported by GeoJSON 1.0
type GeometryType string
const (
	GeometryPoint 			GeometryType = "Point"
	GeometryMultiPoint      GeometryType = "MultiPoint"
	GeometryLineString      GeometryType = "LineString"
	GeometryMultiLineString GeometryType = "MultiLineString"
	GeometryPolygon         GeometryType = "Polygon"
	GeometryMultiPolygon    GeometryType = "MultiPolygon"
	GeometryCollection      GeometryType = "GeometryCollection"
)

// GeoJSON feature's structure
type jsonFeature struct {
	Type 				string					`json:"type"`
	Geom 				*jsonGeometry			`json:"geometry"`
	Prop				*jsonProperty			`json:"properties,omitempty"`
	CRS					map[string]interface{}	`json:"crs,omitempty"`
	BBox				[]float64				`json:"bbox,omitempty"`
}

// Returns the GeoJSON feature for the point with the given coordinates.
func NewPointGeoJSON(c *[]float64) ([]byte, error) {
	g := NewPointGeoJSONGeometry(c)
	f := jsonFeature{Type: "Feature", Geom: &g}
	return json.Marshal(f)
}

// Returns the GeoJSON feature for the linestring with the given coordinates.
func NewLineStringGeoJSON(c *[][]float64) ([]byte, error) {
	g := NewLineStringGeoJSONGeometry(c)
	f := jsonFeature{Type: "Feature", Geom: &g}
	return json.Marshal(f)
}

// Returns the GeoJSON feature for the polygon with the given coordinates.
func NewPolygonGeoJSON(c *[][][]float64) ([]byte, error) {
	g := NewPolygonGeoJSONGeometry(c)
	f := jsonFeature{Type: "Feature", Geom: &g}
	return json.Marshal(f)
}

