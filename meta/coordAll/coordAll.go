package meta

import (
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// CoordAll get all coordinates from any GeoJSON object.
func CoordAll(t interface{}, excludeWrapCoord *bool) ([]geometry.Point, error) {
	switch gtp := t.(type) {
	case *geometry.Point:
		return coordAllPoint(*gtp), nil
	case *geometry.MultiPoint:
		return coordAllMultiPoint(*gtp), nil
	case *geometry.LineString:
		return coordAllLineString(*gtp), nil
	case *geometry.Polygon:
		if excludeWrapCoord == nil {
			return nil, errors.New("exclude wrap coord can't be null")
		}
		return coordAllPolygon(*gtp, *excludeWrapCoord), nil
	case *geometry.MultiLineString:
		return coordAllMultiLineString(*gtp), nil
	case *geometry.MultiPolygon:
		if excludeWrapCoord == nil {
			return nil, errors.New("exclude wrap coord can't be null")
		}
		return coordAllMultiPolygon(*gtp, *excludeWrapCoord), nil
	case *feature.Feature:
		return coordAllFeature(*gtp, *excludeWrapCoord)
	case *feature.Collection:
		if excludeWrapCoord == nil {
			return nil, errors.New("exclude wrap coord can't be null")
		}
		return coordAllFeatureCollection(*gtp, *excludeWrapCoord)
	case *geometry.Collection:
		pts := []geometry.Point{}
		for _, gmt := range gtp.Geometries {
			snl, _ := coordsAllFromSingleGeometry(pts, gmt, *excludeWrapCoord)
			pts = append(pts, snl...)
		}
		return pts, nil
	}

	return nil, nil
}

func coordAllPoint(p geometry.Point) []geometry.Point {
	var coords []geometry.Point
	coords = append(coords, p)
	return coords
}

func coordAllMultiPoint(m geometry.MultiPoint) []geometry.Point {
	return appendCoordsToMultiPoint([]geometry.Point{}, m)
}

func appendCoordsToMultiPoint(coords []geometry.Point, m geometry.MultiPoint) []geometry.Point {
	coords = append(coords, m.Coordinates...)
	return coords
}

func coordAllLineString(m geometry.LineString) []geometry.Point {
	return appendCoordsToLineString([]geometry.Point{}, m)
}

func appendCoordsToLineString(coords []geometry.Point, l geometry.LineString) []geometry.Point {
	coords = append(coords, l.Coordinates...)
	return coords
}

func coordAllPolygon(p geometry.Polygon, excludeWrapCoord bool) []geometry.Point {
	return appendCoordsToPolygon([]geometry.Point{}, p, excludeWrapCoord)
}

func appendCoordsToPolygon(coords []geometry.Point, p geometry.Polygon, excludeWrapCoord bool) []geometry.Point {
	wrapShrink := 0
	if excludeWrapCoord {
		wrapShrink = 1
	}
	for i := 0; i < len(p.Coordinates); i++ {
		for j := 0; j < len(p.Coordinates[i].Coordinates)-wrapShrink; j++ {
			coords = append(coords, p.Coordinates[i].Coordinates[j])
		}
	}
	return coords
}

func coordAllMultiLineString(m geometry.MultiLineString) []geometry.Point {
	return appendCoordToMultiLineString([]geometry.Point{}, m)
}

func appendCoordToMultiLineString(coords []geometry.Point, m geometry.MultiLineString) []geometry.Point {
	for i := 0; i < len(m.Coordinates); i++ {
		coords = append(coords, m.Coordinates[i].Coordinates...)
	}
	return coords
}

func coordAllMultiPolygon(mp geometry.MultiPolygon, excludeWrapCoord bool) []geometry.Point {

	return appendCoordToMultiPolygon([]geometry.Point{}, mp, excludeWrapCoord)
}

func appendCoordToMultiPolygon(coords []geometry.Point, mp geometry.MultiPolygon, excludeWrapCoord bool) []geometry.Point {
	wrapShrink := 0
	if excludeWrapCoord {
		wrapShrink = 1
	}

	for i := 0; i < len(mp.Coordinates); i++ {
		for j := 0; j < len(mp.Coordinates[i].Coordinates); j++ {
			for k := 0; k < len(mp.Coordinates[i].Coordinates[j].Coordinates)-wrapShrink; k++ {
				coords = append(coords, mp.Coordinates[i].Coordinates[j].Coordinates[k])
			}
		}
	}
	return coords
}

func coordAllFeature(f feature.Feature, excludeWrapCoord bool) ([]geometry.Point, error) {
	return appendCoordToFeature([]geometry.Point{}, f, excludeWrapCoord)
}

func coordAllFeatureCollection(c feature.Collection, excludeWrapCoord bool) ([]geometry.Point, error) {
	var finalCoordsList []geometry.Point
	for _, f := range c.Features {
		finalCoordsList, _ = appendCoordToFeature(finalCoordsList, f, excludeWrapCoord)
	}
	return finalCoordsList, nil
}

func appendCoordToFeature(pointList []geometry.Point, f feature.Feature, excludeWrapCoord bool) ([]geometry.Point, error) {

	coords, err := coordsAllFromSingleGeometry(pointList, f.Geometry, excludeWrapCoord)
	if err != nil {
		return nil, err
	}
	return coords, nil
}

func coordsAllFromSingleGeometry(pointList []geometry.Point, g geometry.Geometry, excludeWrapCoord bool) ([]geometry.Point, error) {

	if g.GeoJSONType == geojson.Point {
		p, err := g.ToPoint()
		if err != nil {
			return nil, err
		}
		pointList = append(pointList, *p)
	}

	if g.GeoJSONType == geojson.MultiPoint {
		mp, err := g.ToMultiPoint()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordsToMultiPoint(pointList, *mp)
	}

	if g.GeoJSONType == geojson.LineString {
		ln, err := g.ToLineString()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordsToLineString(pointList, *ln)
	}

	if g.GeoJSONType == geojson.MiltiLineString {
		mln, err := g.ToMultiLineString()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordToMultiLineString(pointList, *mln)
	}

	if g.GeoJSONType == geojson.Polygon {
		poly, err := g.ToPolygon()
		if err != nil {
			return nil, err
		}
		return appendCoordsToPolygon(pointList, *poly, excludeWrapCoord), nil

	}

	if g.GeoJSONType == geojson.MultiPolygon {
		multiPoly, err := g.ToMultiPolygon()
		if err != nil {
			return nil, err
		}
		return appendCoordToMultiPolygon(pointList, *multiPoly, excludeWrapCoord), nil
	}

	return pointList, nil
}

// GetCoord unwrap a coordinate from a  Feature with a Point geometry.
func GetCoord(obj feature.Feature) (*geometry.Point, error) {
	if obj.Geometry.GeoJSONType == geojson.Point {
		return obj.Geometry.ToPoint()
	}
	return nil, errors.New("invalid feature")
}
