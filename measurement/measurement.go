package measurement

import (
	"errors"
	"math"

	"github.com/tomchavakis/turf-go/constants"
	"github.com/tomchavakis/turf-go/conversions"
	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	meta "github.com/tomchavakis/turf-go/meta/coordAll"
)

// Distance calculates the distance between two points in kilometers. This uses the Haversine formula
func Distance(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {

	dLat := conversions.DegreesToRadians(lat2 - lat1)
	dLng := conversions.DegreesToRadians(lon2 - lon1)
	lat1R := conversions.DegreesToRadians(lat1)
	lat2R := conversions.DegreesToRadians(lat2)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1R)*math.Cos(lat2R)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := constants.EarthRadius * c
	return d
}

// PointDistance calculates the distance between two points
func PointDistance(p1 geometry.Point, p2 geometry.Point) float64 {
	return Distance(p1.Lng, p1.Lat, p2.Lng, p2.Lat)
}

// Bearing finds the geographic bearing between two given points.
func Bearing(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {
	dLng := conversions.DegreesToRadians(lon2 - lon1)
	lat1R := conversions.DegreesToRadians(lat1)
	lat2R := conversions.DegreesToRadians(lat2)
	y := math.Sin(dLng) * math.Cos(lat2R)
	x := math.Cos(lat1R)*math.Sin(lat2R) - math.Sin(lat1R)*math.Cos(lat2R)*math.Cos(dLng)

	// convert to degrees
	bd := conversions.RadiansToDegrees(math.Atan2(y, x))

	if bd < 0.0 {
		bd += 360.0
	}

	if bd >= 360.0 {
		bd -= 360.0

	}
	return bd

}

// PointBearing finds the geographic bearing between two points.
func PointBearing(p1 geometry.Point, p2 geometry.Point) float64 {
	return Bearing(p1.Lng, p1.Lat, p2.Lng, p2.Lat)
}

// MidPoint finds the point midway between them.
func MidPoint(p1 geometry.Point, p2 geometry.Point) geometry.Point {
	dLon := conversions.DegreesToRadians(p2.Lng - p1.Lng)
	lat1R := conversions.DegreesToRadians(p1.Lat)
	lon1R := conversions.DegreesToRadians(p1.Lng)
	lat2R := conversions.DegreesToRadians(p2.Lat)
	Bx := math.Cos(lat2R) * math.Cos(dLon)
	By := math.Cos(lat2R) * math.Sin(dLon)
	midLat := math.Atan2(math.Sin(lat1R)+math.Sin(lat2R), math.Sqrt((math.Cos(lat1R)+Bx)*(math.Cos(lat1R)+Bx)+By*By))
	midLng := lon1R + math.Atan2(By, math.Cos(lat1R)+Bx)

	return geometry.Point{Lat: conversions.RadiansToDegrees(midLat), Lng: conversions.RadiansToDegrees(midLng)}
}

// Destination returns a destination point according to a reference point, a distance in km and a bearing in degrees from True North.
func Destination(p1 geometry.Point, distance float64, bearing float64) geometry.Point {
	lonR := conversions.DegreesToRadians(p1.Lng)
	latR := conversions.DegreesToRadians(p1.Lat)
	bR := conversions.DegreesToRadians(bearing)
	dLat := math.Asin(math.Sin(latR)*math.Cos(distance/constants.EarthRadius) + math.Cos(latR)*math.Sin(distance/constants.EarthRadius)*math.Cos(bR))
	dLng := lonR + math.Atan2(math.Sin(bR)*math.Sin(distance/constants.EarthRadius)*math.Cos(latR), math.Cos(distance/constants.EarthRadius)-math.Sin(latR)*math.Sin(dLat))

	return geometry.Point{Lat: conversions.RadiansToDegrees(dLat), Lng: conversions.RadiansToDegrees(dLng)}
}

// Length measures the length of a geometry.
func Length(t interface{}) float64 {

	result := 0.0
	switch gtp := t.(type) {
	case []geometry.Point:
		result = lenth(gtp)
	case geometry.LineString:
		result = lenth(gtp.Coordinates)
	case geometry.MultiLineString:
		coords := gtp.Coordinates // []LineString
		for _, c := range coords {
			result += lenth(c.Coordinates)
		}
	case geometry.Polygon:
		for _, c := range gtp.Coordinates {
			result += lenth(c.Coordinates)
		}
	case geometry.MultiPolygon:
		coords := gtp.Coordinates
		for _, coord := range coords {
			for _, pl := range coord.Coordinates {
				result += lenth(pl.Coordinates)
			}
		}
	}
	return result
}

// http://turfjs.org/docs/#linedistance
func lenth(coords []geometry.Point) float64 {
	travelled := 0.0
	prevCoords := coords[0]
	var currentCoords geometry.Point
	for i := 1; i < len(coords); i++ {
		currentCoords = coords[i]
		travelled += PointDistance(prevCoords, currentCoords)
		prevCoords = currentCoords
	}
	return travelled
}

// Area takes a geometry type and returns its area in square meters
func Area(t interface{}) (float64, error) {
	switch gtp := t.(type) {
	case *feature.Feature:
		return calculateArea(gtp.Geometry)
	case *feature.Collection:
		features := gtp.Features
		total := 0.0
		if len(features) > 0 {
			for _, f := range features {
				ar, err := calculateArea(f.Geometry)
				if err != nil {
					return 0, err
				}
				total += ar
			}
		}
		return total, nil
	case *geometry.Geometry:
		return calculateArea(*gtp)
	case *geometry.Polygon:
		return polygonArea(gtp.Coordinates), nil
	case *geometry.MultiPolygon:
		total := 0.0
		for i := 0; i < len(gtp.Coordinates); i++ {
			total += polygonArea(gtp.Coordinates[i].Coordinates)
		}
		return total, nil
	}
	return 0.0, nil
}

func calculateArea(g geometry.Geometry) (float64, error) {
	total := 0.0
	if g.GeoJSONType == geojson.Polygon {

		poly, err := g.ToPolygon()
		if err != nil {
			return 0.0, errors.New("cannot convert geometry to Polygon")
		}
		return polygonArea(poly.Coordinates), nil
	} else if g.GeoJSONType == geojson.MultiPolygon {
		multiPoly, err := g.ToMultiPolygon()
		if err != nil {
			return 0.0, errors.New("cannot convert geometry to MultiPolygon")
		}
		for i := 0; i < len(multiPoly.Coordinates); i++ {
			total += polygonArea(multiPoly.Coordinates[i].Coordinates)
		}

		return total, nil
	} else {
		// area should be 0 for Point, MultiPoint, LineString and MultiLineString
		return total, nil
	}
}

func polygonArea(coords []geometry.LineString) float64 {
	total := 0.0
	if len(coords) > 0 {
		total += math.Abs(ringArea(coords[0].Coordinates))
		for i := 1; i < len(coords); i++ {
			total -= math.Abs(ringArea(coords[i].Coordinates))
		}
	}
	return total
}

// calculate the approximate area of the polygon were it projected onto the earth.
// Note that this area will be positive if ring is oriented clockwise, otherwise
// it will be negative.
//
// Reference:
// Robert. G. Chamberlain and William H. Duquette, "Some Algorithms for Polygons on a Sphere",
// JPL Publication 07-03, Jet Propulsion
// Laboratory, Pasadena, CA, June 2007 https://trs.jpl.nasa.gov/handle/2014/41271
func ringArea(coords []geometry.Point) float64 {
	var p1 geometry.Point
	var p2 geometry.Point
	var p3 geometry.Point
	var lowerIndex int
	var middleIndex int
	var upperIndex int
	total := 0.0
	coordsLength := len(coords)

	if coordsLength > 2 {
		for i := 0; i < coordsLength; i++ {
			if i == coordsLength-2 { // i = N-2
				lowerIndex = coordsLength - 2
				middleIndex = coordsLength - 1
				upperIndex = 0
			} else if i == coordsLength-1 { //i = N-1
				lowerIndex = coordsLength - 1
				middleIndex = 0
				upperIndex = 1
			} else { // i =0 to N-3
				lowerIndex = i
				middleIndex = i + 1
				upperIndex = i + 2
			}
			p1 = coords[lowerIndex]
			p2 = coords[middleIndex]
			p3 = coords[upperIndex]
			total += (conversions.DegreesToRadians(p3.Lng) - conversions.DegreesToRadians(p1.Lng)) * math.Sin(conversions.DegreesToRadians(p2.Lat))
		}
		total = total * constants.EarthRadius * constants.EarthRadius / 2
	}
	return total
}

// BBox takes a set of features, calculates the bbox of all input features, and returns a bounding box.
func BBox(t interface{}) ([]float64, error) {
	return bboxGeom(t, false)
}

// Along Takes a line and returns a point at a specified distance along the line.
func Along(ln geometry.LineString, distance float64) geometry.Point {
	travelled := 0.0
	for i := 0; i < len(ln.Coordinates); i++ {
		if distance >= travelled && i == len(ln.Coordinates)-1 {
			break
		} else if travelled >= distance {
			overshot := distance - travelled
			if overshot == 0 {
				return ln.Coordinates[i]
			}
			direction := PointBearing(ln.Coordinates[i], ln.Coordinates[i-1]) - 180
			return Destination(ln.Coordinates[i], overshot, direction)
		} else {
			travelled += PointDistance(ln.Coordinates[i], ln.Coordinates[i+1])
		}
	}

	return ln.Coordinates[len(ln.Coordinates)-1]
}

func bboxGeom(t interface{}, excludeWrapCoord bool) ([]float64, error) {
	coords, err := meta.CoordAll(t, &excludeWrapCoord)
	if err != nil {
		return nil, errors.New("cannot get coords")
	}

	return bboxCalculator(coords), nil
}

func bboxCalculator(coords []geometry.Point) []float64 {
	var bbox []float64
	bbox = append(bbox, math.Inf(+1))
	bbox = append(bbox, math.Inf(+1))
	bbox = append(bbox, math.Inf(-1))
	bbox = append(bbox, math.Inf(-1))

	for _, p := range coords {
		if bbox[0] > p.Lng {
			bbox[0] = p.Lng
		}
		if bbox[1] > p.Lat {
			bbox[1] = p.Lat
		}
		if bbox[2] < p.Lng {
			bbox[2] = p.Lng
		}
		if bbox[3] < p.Lat {
			bbox[3] = p.Lat
		}
	}
	return bbox
}
