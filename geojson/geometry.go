package geojson

import (
	"encoding/json"
	"fmt"
)

// GeoJSON geometry structure
type jsonGeometry struct {
	Type 				GeometryType		`json:"type"`
	Coordinates 		json.RawMessage		`json:"coordinates"`
}

// Returns the GeoJSON geometry for the point with the given coordinates.
func NewPointGeoJSONGeometry(c *[]float64) jsonGeometry {
	raw, err := json.Marshal(c)
	if err != nil {
		fmt.Println("ERROR %s", err.Error())
	}
	return jsonGeometry{Type: GeometryPoint, Coordinates: json.RawMessage(raw)}

}

// Returns the GeoJSON geometry for the linestring with the given coordinates.
func NewLineStringGeoJSONGeometry(c *[][]float64) jsonGeometry {
	raw, err := json.Marshal(c)
	if err != nil {
		fmt.Println("ERROR %s", err.Error())
	}
	return jsonGeometry{Type: GeometryLineString, Coordinates: raw}
}

// Returns the GeoJSON geometry for the polygon with the given coordinates.
func NewPolygonGeoJSONGeometry(c *[][][]float64) jsonGeometry {
	raw, err := json.Marshal(c)
	if err != nil {
		fmt.Println("ERROR %s", err.Error())
	}
	return jsonGeometry{Type: GeometryPolygon, Coordinates: raw}
}