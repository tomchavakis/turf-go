package meta

import (
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// CoordEach iterate over coordinates in any Geojson object and apply the callbackFn
// geojson can be a FeatureCollection | Feature | Geometry
// callbackFn is a method that takes a point and returns a point
// excludeWrapCoord whether or not to include the final coordinate of LinearRings that wraps the ring in its iteration.
func CoordEach(geojson interface{}, callbackFn func(geometry.Point) geometry.Point, excludeWrapCoord *bool) ([]geometry.Point, error) {
	if geojson == nil {
		return nil, errors.New("geojson is empty")
	}
	switch gtp := geojson.(type) {
	case nil:
		break
	case *geometry.Point:
		return callbackEachPoint(gtp, callbackFn), nil
	case *geometry.MultiPoint:
		return coordEachMultiPoint(gtp, callbackFn), nil
	case *geometry.LineString:
		return coordEachLineString(gtp, callbackFn), nil
	case *geometry.Polygon:
		if excludeWrapCoord == nil {
			return coordEachPolygon(gtp, false, callbackFn), nil
		}
		return coordEachPolygon(gtp, *excludeWrapCoord, callbackFn), nil
	case *geometry.MultiLineString:
		return coordEachMultiLineString(gtp, callbackFn), nil
	case *geometry.MultiPolygon:
		if excludeWrapCoord == nil {
			return coordEachMultiPolygon(gtp, false, callbackFn), nil
		}
		return coordEachMultiPolygon(gtp, *excludeWrapCoord, callbackFn), nil
	case *feature.Feature:
		if excludeWrapCoord == nil {
			return coordEachFeature(gtp, false, callbackFn)
		}
		return coordEachFeature(gtp, *excludeWrapCoord, callbackFn)
	case *feature.Collection:
		if excludeWrapCoord == nil {
			return coordEachFeatureCollection(gtp, false, callbackFn)
		}
		return coordEachFeatureCollection(gtp, *excludeWrapCoord, callbackFn)
	case *geometry.Collection:
		if excludeWrapCoord == nil {
			return coordEachGeometryCollection(gtp, false, callbackFn)
		}
		return coordEachGeometryCollection(gtp, *excludeWrapCoord, callbackFn)
	}

	return nil, nil
}

func callbackEachPoint(p *geometry.Point, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	var coords []geometry.Point
	np := callbackFn(*p)
	// Conversion assignment
	p.Lat = np.Lat
	p.Lng = np.Lng
	coords = append(coords, np)
	return coords
}

func coordEachMultiPoint(m *geometry.MultiPoint, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordsToMultiPoint([]geometry.Point{}, m, callbackFn)
}

func appendCoordsToMultiPoint(coords []geometry.Point, m *geometry.MultiPoint, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	for _, v := range m.Coordinates {
		np := callbackFn(v)
		coords = append(coords, np)
	}
	m.Coordinates = coords
	return coords
}

func coordEachLineString(m *geometry.LineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordsToLineString([]geometry.Point{}, m, callbackFn)
}

func appendCoordsToLineString(coords []geometry.Point, l *geometry.LineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	for _, v := range l.Coordinates {
		np := callbackFn(v)
		coords = append(coords, np)
	}

	l.Coordinates = coords
	return coords
}

func coordEachMultiLineString(m *geometry.MultiLineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordToMultiLineString([]geometry.Point{}, m, callbackFn)
}

func appendCoordToMultiLineString(coords []geometry.Point, m *geometry.MultiLineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	for i := 0; i < len(m.Coordinates); i++ {
		for j := 0; j < len(m.Coordinates[i].Coordinates); j++ {
			np := callbackFn(m.Coordinates[i].Coordinates[j])
			m.Coordinates[i].Coordinates[j] = np
			coords = append(coords, np)
		}

	}
	return coords
}

func coordEachPolygon(p *geometry.Polygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordsToPolygon([]geometry.Point{}, p, excludeWrapCoord, callbackFn)
}

func appendCoordsToPolygon(coords []geometry.Point, p *geometry.Polygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	wrapShrink := 0
	if excludeWrapCoord {
		wrapShrink = 1
	}
	for i := 0; i < len(p.Coordinates); i++ {
		for j := 0; j < len(p.Coordinates[i].Coordinates)-wrapShrink; j++ {
			np := callbackFn(p.Coordinates[i].Coordinates[j])
			p.Coordinates[i].Coordinates[j] = np
			coords = append(coords, np)
		}
	}
	return coords
}

func coordEachMultiPolygon(mp *geometry.MultiPolygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {

	return appendCoordToMultiPolygon([]geometry.Point{}, mp, excludeWrapCoord, callbackFn)
}

func appendCoordToMultiPolygon(coords []geometry.Point, mp *geometry.MultiPolygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	wrapShrink := 0
	if excludeWrapCoord {
		wrapShrink = 1
	}

	for i := 0; i < len(mp.Coordinates); i++ {
		for j := 0; j < len(mp.Coordinates[i].Coordinates); j++ {
			for k := 0; k < len(mp.Coordinates[i].Coordinates[j].Coordinates)-wrapShrink; k++ {
				np := callbackFn(mp.Coordinates[i].Coordinates[j].Coordinates[k])
				mp.Coordinates[i].Coordinates[j].Coordinates[k] = np
				coords = append(coords, np)
			}
		}
	}
	return coords
}

func coordEachFeature(f *feature.Feature, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) ([]geometry.Point, error) {
	return appendCoordToFeature([]geometry.Point{}, f, excludeWrapCoord, callbackFn)
}

func appendCoordToFeature(pointList []geometry.Point, f *feature.Feature, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) ([]geometry.Point, error) {

	coords, err := coordsEachFromSingleGeometry(pointList, &f.Geometry, excludeWrapCoord, callbackFn)
	if err != nil {
		return nil, err
	}
	return coords, nil
}

func coordsEachFromSingleGeometry(pointList []geometry.Point, g *geometry.Geometry, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) ([]geometry.Point, error) {

	if g.GeoJSONType == geojson.Point {
		p, err := g.ToPoint()
		if err != nil {
			return nil, err
		}
		np := callbackFn(*p)
		pointList = append(pointList, np)
		g.Coordinates = np
	}

	if g.GeoJSONType == geojson.MultiPoint {
		mp, err := g.ToMultiPoint()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordsToMultiPoint(pointList, mp, callbackFn)

		g.Coordinates = mp.Coordinates
	}

	if g.GeoJSONType == geojson.LineString {
		ln, err := g.ToLineString()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordsToLineString(pointList, ln, callbackFn)
		g.Coordinates = ln.Coordinates
	}

	if g.GeoJSONType == geojson.MultiLineString {
		mln, err := g.ToMultiLineString()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordToMultiLineString(pointList, mln, callbackFn)
		g.Coordinates = mln.Coordinates
	}

	if g.GeoJSONType == geojson.Polygon {
		poly, err := g.ToPolygon()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordsToPolygon(pointList, poly, excludeWrapCoord, callbackFn)
		g.Coordinates = poly.Coordinates
	}

	if g.GeoJSONType == geojson.MultiPolygon {
		multiPoly, err := g.ToMultiPolygon()
		if err != nil {
			return nil, err
		}
		pointList = appendCoordToMultiPolygon(pointList, multiPoly, excludeWrapCoord, callbackFn)
		g.Coordinates = multiPoly.Coordinates
	}

	return pointList, nil
}

func coordEachFeatureCollection(c *feature.Collection, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) ([]geometry.Point, error) {
	var finalCoordsList []geometry.Point

	for i := 0; i < len(c.Features); i++ {
		var tempCoordsList []geometry.Point
		tempCoordsList, _ = appendCoordToFeature(tempCoordsList, &c.Features[i], excludeWrapCoord, callbackFn)
		finalCoordsList = append(finalCoordsList, tempCoordsList...)
	}

	return finalCoordsList, nil
}

func coordEachGeometryCollection(g *geometry.Collection, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) ([]geometry.Point, error) {
	var finalCoordsList []geometry.Point

	for i := 0; i < len(g.Geometries); i++ {
		var tempCoordsList []geometry.Point
		tempCoordsList, _ = coordsEachFromSingleGeometry(tempCoordsList, &g.Geometries[i], excludeWrapCoord, callbackFn)
		finalCoordsList = append(finalCoordsList, tempCoordsList...)
	}

	return finalCoordsList, nil
}
