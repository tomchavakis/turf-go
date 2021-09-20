package meta

import (
	"errors"

	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// CoordEach iterate over coordinates in any Geojson object
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
		return callbackEachPoint(*gtp, callbackFn), nil
	case *geometry.MultiPoint:
		return coordEachMultiPoint(*gtp, callbackFn), nil
	case *geometry.LineString:
		return callbackLineString(*gtp, callbackFn), nil
	case *geometry.Polygon:
		if excludeWrapCoord == nil {
			return nil, errors.New("exclude wrap coord can't be null")
		}
		return coordEachPolygon(*gtp, *excludeWrapCoord, callbackFn), nil
	case *geometry.MultiLineString:
		return coordEachMultiLineString(*gtp, callbackFn), nil
	case *geometry.MultiPolygon:
		if excludeWrapCoord == nil {
			return nil, errors.New("exclude wrap coord can't be null")
		}
		return coordEachMultiPolygon(*gtp, *excludeWrapCoord, callbackFn), nil
		// case *feature.Feature:
		// 	return coordAllFeature(*gtp, *excludeWrapCoord)
		// case *feature.Collection:
		// 	if excludeWrapCoord == nil {
		// 		return nil, errors.New("exclude wrap coord can't be null")
		// 	}
		// 	return coordAllFeatureCollection(*gtp, *excludeWrapCoord)
		// case *geometry.Collection:
		// 	pts := []geometry.Point{}
		// 	for _, gmt := range gtp.Geometries {
		// 		snl, _ := coordsAllFromSingleGeometry(pts, gmt, *excludeWrapCoord)
		// 		pts = append(pts, snl...)
		// 	}
		// 	return pts, nil
		// }

	}
	return nil, nil
}

func callbackEachPoint(p geometry.Point, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	var coords []geometry.Point
	np := callbackFn(p)
	coords = append(coords, np)

	return coords
}

func coordEachMultiPoint(m geometry.MultiPoint, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordsToMultiPoint([]geometry.Point{}, m, callbackFn)
}

func appendCoordsToMultiPoint(coords []geometry.Point, m geometry.MultiPoint, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	for _, v := range m.Coordinates {
		np := callbackFn(v)
		coords = append(coords, np)
	}
	return coords
}

func callbackLineString(ln geometry.LineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	var coords []geometry.Point
	for _, v := range ln.Coordinates {
		np := callbackFn(v)
		coords = append(coords, np)
	}
	return coords
}

func coordEachMultiLineString(m geometry.MultiLineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordToMultiLineString([]geometry.Point{}, m, callbackFn)
}

func appendCoordToMultiLineString(coords []geometry.Point, m geometry.MultiLineString, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	for i := 0; i < len(m.Coordinates); i++ {
		for j := 0; j < len(m.Coordinates[i].Coordinates); j++ {
			np := callbackFn(m.Coordinates[i].Coordinates[j])
			coords = append(coords, np)
		}

	}
	return coords
}

func coordEachPolygon(p geometry.Polygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	return appendCoordsToPolygon([]geometry.Point{}, p, excludeWrapCoord, callbackFn)
}

func appendCoordsToPolygon(coords []geometry.Point, p geometry.Polygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	wrapShrink := 0
	if excludeWrapCoord {
		wrapShrink = 1
	}
	for i := 0; i < len(p.Coordinates); i++ {
		for j := 0; j < len(p.Coordinates[i].Coordinates)-wrapShrink; j++ {
			np := callbackFn(p.Coordinates[i].Coordinates[j])
			coords = append(coords, np)
		}
	}
	return coords
}

func coordEachMultiPolygon(mp geometry.MultiPolygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {

	return appendCoordToMultiPolygon([]geometry.Point{}, mp, excludeWrapCoord, callbackFn)
}

func appendCoordToMultiPolygon(coords []geometry.Point, mp geometry.MultiPolygon, excludeWrapCoord bool, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	wrapShrink := 0
	if excludeWrapCoord {
		wrapShrink = 1
	}

	for i := 0; i < len(mp.Coordinates); i++ {
		for j := 0; j < len(mp.Coordinates[i].Coordinates); j++ {
			for k := 0; k < len(mp.Coordinates[i].Coordinates[j].Coordinates)-wrapShrink; k++ {
				np := callbackFn(mp.Coordinates[i].Coordinates[j].Coordinates[k])
				coords = append(coords, np)
			}
		}
	}
	return coords
}
