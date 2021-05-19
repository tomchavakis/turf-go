package meta

import (
	"errors"

	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// CoordEach iterate over coordinates in any Geojson object
func CoordEach(geojson interface{}, callbackFn func(geometry.Point) geometry.Point, excludeWrapCoord *bool) ([]geometry.Point, error) {
	if geojson == nil {
		return nil, errors.New("geojson is empty")
	}
	switch gtp := geojson.(type) {
	case *geometry.Point:
		return callbackAllPoint(*gtp, callbackFn), nil
		// case *geometry.MultiPoint:
		// 	return coordAllMultiPoint(*gtp), nil
		// case *geometry.LineString:
		// 	return coordAllLineString(*gtp), nil
		// case *geometry.Polygon:
		// 	if excludeWrapCoord == nil {
		// 		return nil, errors.New("exclude wrap coord can't be null")
		// 	}
		// 	return coordAllPolygon(*gtp, *excludeWrapCoord), nil
		// case *geometry.MultiLineString:
		// 	return coordAllMultiLineString(*gtp), nil
		// case *geometry.MultiPolygon:
		// 	if excludeWrapCoord == nil {
		// 		return nil, errors.New("exclude wrap coord can't be null")
		// 	}
		// 	return coordAllMultiPolygon(*gtp, *excludeWrapCoord), nil
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

func callbackAllPoint(p geometry.Point, callbackFn func(geometry.Point) geometry.Point) []geometry.Point {
	var coords []geometry.Point
	np := callbackFn(p)
	coords = append(coords, np)

	return coords
}
