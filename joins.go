package turf

import (
	"github.com/tomchavakis/geojson"
	"github.com/tomchavakis/geojson/geometry"
)

// PointInPolygon takes a Point and a Polygon and determines if the point resides inside the polygon
func PointInPolygon(point geometry.Point, polygon geometry.Polygon) (bool, error) {

	pArr := []geometry.Polygon{}
	pArr = append(pArr, polygon)

	mp, err := geometry.NewMultiPolygon(pArr)
	if err != nil {
		return false, err
	}

	return PointInMultiPolygon(point, *mp), nil
}

// PointInMultiPolygon takes a Point and a MultiPolygon and determines if the point resides inside the polygon
func PointInMultiPolygon(p geometry.Point, mp geometry.MultiPolygon) bool {

	insidePoly := false
	polys := mp.Coordinates

	for i := 0; i < len(polys) && !insidePoly; i++ {
		//check if it is in the outer ring first
		if inRing(p, polys[i].Coordinates[0].Coordinates) {
			inHole := false
			temp := 1
			// check for the point in any of the holes
			for temp < len(polys[i].Coordinates) && !inHole {
				if inRing(p, polys[i].Coordinates[temp].Coordinates) {
					inHole = true
				}
				temp++
			}
			if !inHole {
				insidePoly = true
			}
		}
	}

	return insidePoly
}

// InBBOX returns true if the point is within the Bounding Box
func InBBOX(pt geometry.Point, bbox geojson.BBOX) bool {
	return bbox.West <= pt.Lng &&
		bbox.South <= pt.Lat &&
		bbox.East >= pt.Lng &&
		bbox.North >= pt.Lat
}

func inRing(pt geometry.Point, ring []geometry.Point) bool {

	isInside := false
	j := 0
	for i := 0; i < len(ring); i++ {

		xi := ring[i].Lng
		yi := ring[i].Lat
		xj := ring[j].Lng
		yj := ring[j].Lat

		intersect := (yi > pt.Lat) != (yj > pt.Lat) && (pt.Lng < (xj-xi)*(pt.Lat-yi)/(yj-yi)+xi)

		if intersect {
			isInside = !isInside
		}

		j = i
	}
	return isInside
}
