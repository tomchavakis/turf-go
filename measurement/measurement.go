package measurement

import (
	"errors"
	"math"

	"github.com/tomchavakis/turf-go/constants"
	"github.com/tomchavakis/turf-go/conversions"
	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/internal/common"
	"github.com/tomchavakis/turf-go/invariant"
	meta "github.com/tomchavakis/turf-go/meta/coordAll"
)

// Distance calculates the distance between two points in kilometers. This uses the Haversine formula
func Distance(lon1 float64, lat1 float64, lon2 float64, lat2 float64, units string) (float64, error) {

	dLat := conversions.DegreesToRadians(lat2 - lat1)
	dLng := conversions.DegreesToRadians(lon2 - lon1)
	lat1R := conversions.DegreesToRadians(lat1)
	lat2R := conversions.DegreesToRadians(lat2)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1R)*math.Cos(lat2R)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	// d := constants.EarthRadius * c

	return conversions.RadiansToLength(c, units)
}

// PointDistance calculates the distance between two points
func PointDistance(p1 geometry.Point, p2 geometry.Point, units string) (float64, error) {
	return Distance(p1.Lng, p1.Lat, p2.Lng, p2.Lat, units)
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
func Destination(p1 geometry.Point, distance float64, bearing float64, units string) (*geometry.Point, error) {
	lonR := conversions.DegreesToRadians(p1.Lng)
	latR := conversions.DegreesToRadians(p1.Lat)
	bR := conversions.DegreesToRadians(bearing)
	radians, err := conversions.LengthToRadians(distance, units)
	if err != nil {
		return nil, err
	}
	dLat := math.Asin(math.Sin(latR)*math.Cos(radians) + math.Cos(latR)*math.Sin(radians)*math.Cos(bR))
	dLng := lonR + math.Atan2(math.Sin(bR)*math.Sin(radians)*math.Cos(latR), math.Cos(radians)-math.Sin(latR)*math.Sin(dLat))

	return &geometry.Point{Lat: conversions.RadiansToDegrees(dLat), Lng: conversions.RadiansToDegrees(dLng)}, nil
}

// Length measures the length of a geometry.
func Length(t interface{}, units string) (float64, error) {

	result := 0.0
	var err error
	var l float64
	switch gtp := t.(type) {
	case []geometry.Point:
		l, err = lenth(gtp, units)
		result = l
	case geometry.LineString:
		l, err = lenth(gtp.Coordinates, units)
		result = l
	case geometry.MultiLineString:
		coords := gtp.Coordinates // []LineString
		for _, c := range coords {
			l, err = lenth(c.Coordinates, units)
			if err != nil {
				break
			}
			result += l
		}
	case geometry.Polygon:
		for _, c := range gtp.Coordinates {
			l, err = lenth(c.Coordinates, units)
			if err != nil {
				break
			}
			result += l
		}
	case geometry.MultiPolygon:
		coords := gtp.Coordinates
		for _, coord := range coords {
			for _, pl := range coord.Coordinates {
				l, err = lenth(pl.Coordinates, units)
				if err != nil {
					break
				}
				result += l
			}
		}
	}
	return result, err
}

// http://turfjs.org/docs/#linedistance
func lenth(coords []geometry.Point, units string) (float64, error) {
	travelled := 0.0
	prevCoords := coords[0]
	var currentCoords geometry.Point
	for i := 1; i < len(coords); i++ {
		currentCoords = coords[i]
		pd, err := PointDistance(prevCoords, currentCoords, units)
		if err != nil {
			return 0.0, err
		}
		travelled += pd
		prevCoords = currentCoords
	}
	return travelled, nil
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
		total = total * constants.EarthRadius * constants.EarthRadius / 2.0
	}
	return total
}

// BBox takes a set of features, calculates the bbox of all input features, and returns a bounding box.
func BBox(t interface{}) ([]float64, error) {
	return bboxGeom(t, false)
}

func bboxGeom(t interface{}, excludeWrapCoord bool) ([]float64, error) {
	coords, err := meta.CoordAll(t, &excludeWrapCoord)
	if err != nil {
		return nil, errors.New("cannot get coords")
	}

	return bboxCalculator(coords), nil
}

// Along Takes a line and returns a point at a specified distance along the line.
func Along(ln geometry.LineString, distance float64, units string) (*geometry.Point, error) {
	travelled := 0.0
	for i := 0; i < len(ln.Coordinates); i++ {
		if distance >= travelled && i == len(ln.Coordinates)-1 {
			break
		} else if travelled >= distance {
			overshot := distance - travelled
			if overshot == 0 {
				return &ln.Coordinates[i], nil
			}
			direction := PointBearing(ln.Coordinates[i], ln.Coordinates[i-1]) - 180

			d, err := Destination(ln.Coordinates[i], overshot, direction, units)
			if err != nil {
				return nil, err
			}
			return d, nil
		} else {
			pd, err := PointDistance(ln.Coordinates[i], ln.Coordinates[i+1], units)
			if err != nil {
				return nil, err
			}
			travelled += pd
		}
	}

	return &ln.Coordinates[len(ln.Coordinates)-1], nil
}

// BBoxPolygon takes a BoundingBox and returns an equivalent polygon.
func BBoxPolygon(bbox geojson.BBOX, id string) (*feature.Feature, error) {

	var cds [][][]float64
	coords := [][]float64{
		{
			bbox.South,
			bbox.West,
		},
		{
			bbox.South,
			bbox.East,
		},
		{
			bbox.North,
			bbox.East,
		},
		{
			bbox.North,
			bbox.West,
		},
		{
			bbox.South,
			bbox.West,
		},
	}
	cds = append(cds, coords)
	bbbox, err := BBox(bbox)
	if err != nil {
		return nil, err
	}
	geom := geometry.Geometry{
		GeoJSONType: geojson.Polygon,
		Coordinates: cds,
	}

	f, err := feature.New(geom, bbbox, nil, id)
	if err != nil {
		return nil, err
	}

	return f, nil
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

// CenterFeature takes a Feature and returns the absolute center of the Feature. Return a Feature with a Point geometry type.
func CenterFeature(f feature.Feature, properties map[string]interface{}, id string) (*feature.Feature, error) {
	fs := []feature.Feature{}
	fs = append(fs, f)
	fc, err := feature.NewFeatureCollection(fs)
	if err != nil {
		return nil, err
	}
	return CenterFeatureCollection(*fc, properties, id)
}

// CenterFeatureCollection takes a FeatureCollection and returns the absolute center of the Feature(s) in the FeatureCollection.
func CenterFeatureCollection(fc feature.Collection, properties map[string]interface{}, id string) (*feature.Feature, error) {
	ext, err := BBox(&fc)
	if err != nil {
		return nil, err
	}

	finalCenterLongtitude := (ext[0] + ext[2]) / 2
	finalCenterLatitude := (ext[1] + ext[3]) / 2

	coords := []float64{finalCenterLongtitude, finalCenterLatitude}
	g := geometry.Geometry{
		GeoJSONType: geojson.Point,
		Coordinates: coords,
	}
	f, err := feature.New(g, ext, properties, id)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Envelope takes a FeatureCollection and returns a rectangular Polygon than encompasses all vertices.
func Envelope(fc feature.Collection) (*feature.Feature, error) {
	excludeWrapCoord := false
	coords, err := meta.CoordAll(&fc, &excludeWrapCoord)
	if err != nil {
		return nil, errors.New("cannot get coords")
	}

	return calcEnvelopeCoords(coords)
}

func calcEnvelopeCoords(coords []geometry.Point) (*feature.Feature, error) {
	if len(coords) == 0 {
		return nil, errors.New("Empty coordinates")
	}
	northStar := coords[0]
	southStar := coords[0]
	westStar := coords[0]
	eastStar := coords[0]

	for _, p := range coords {

		if northStar.Lat < p.Lat {
			northStar = p
		}

		if southStar.Lat > p.Lat {
			southStar = p
		}

		if westStar.Lng > p.Lng {
			westStar = p
		}

		if eastStar.Lng < p.Lng {
			eastStar = p
		}
	}

	var cds [][][]float64
	polygonCoords := [][]float64{
		{
			northStar.Lat,
			westStar.Lng,
		},
		{
			southStar.Lat,
			westStar.Lng,
		},
		{
			southStar.Lat,
			eastStar.Lng,
		},
		{
			northStar.Lat,
			eastStar.Lng,
		},
		{
			northStar.Lat,
			westStar.Lng,
		},
	}

	stars := []geometry.Point{
		northStar,
		southStar,
		eastStar,
		westStar,
	}

	cds = append(cds, polygonCoords)
	bbox := bboxCalculator(stars)
	bbbox, err := BBox(bbox)
	if err != nil {
		return nil, err
	}
	geom := geometry.Geometry{
		GeoJSONType: geojson.Polygon,
		Coordinates: cds,
	}

	f, err := feature.New(geom, bbbox, nil, "")
	if err != nil {
		return nil, err
	}

	return f, nil
}

// CentroidFeature takes a Feature and returns the centroid of the Feature. Return a Feature with a Point geometry type.
func CentroidFeature(f feature.Feature, properties map[string]interface{}, id string) (*feature.Feature, error) {
	fs := []feature.Feature{}
	fs = append(fs, f)
	fc, err := feature.NewFeatureCollection(fs)
	if err != nil {
		return nil, err
	}
	return CentroidFeatureCollection(*fc, properties, id)
}

// CentroidFeatureCollection takes a FeatureCollection and returns the centroid of the Feature(s) in the FeatureCollection.
func CentroidFeatureCollection(fc feature.Collection, properties map[string]interface{}, id string) (*feature.Feature, error) {
	excludeWrapCoord := true
	coords, err := meta.CoordAll(&fc, &excludeWrapCoord)

	if err != nil {
		return nil, errors.New("cannot get coords")
	}

	coordsLength := len(coords)
	if coordsLength < 1 {
		return nil, errors.New("no coordinates found")
	}

	xSum := 0.0
	ySum := 0.0

	for i := 0; i < coordsLength; i++ {
		xSum += coords[i].Lng
		ySum += coords[i].Lat
	}

	finalCenterLongtitude := xSum / float64(coordsLength)
	finalCenterLatitude := ySum / float64(coordsLength)

	coordinates := []float64{finalCenterLongtitude, finalCenterLatitude}
	g := geometry.Geometry{
		GeoJSONType: geojson.Point,
		Coordinates: coordinates,
	}
	f, err := feature.New(g, bboxCalculator(coords), properties, id)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// RhumbBearing takes two points and finds the bearing angle between them along a Rhumb line
// final option calculates the final bearing if true
// returns a bearing from north in decimal degrees, between -180 and 180 degrees (positive clockwise)
// i.e the angle measured in degrees start the north line (0 degrees)
// https://en.wikipedia.org/wiki/Rhumb_line
// In navigation, a rhumb line or loxodrome is an arc crossing all meridians of longitude at the same angle, that is, a path with constant
// bearing as measured relative to true north.
func RhumbBearing(start geometry.Point, end geometry.Point, final bool) (*float64, error) {
	var bear360 float64
	e, err := invariant.GetCoord(&end)
	if err != nil {
		return nil, err
	}
	s, err := invariant.GetCoord(&start)
	if err != nil {
		return nil, err
	}

	if final {
		bear360 = calculateRhumbBearing(e, s)
	} else {
		bear360 = calculateRhumbBearing(s, e)
	}
	var bear180 float64
	if bear360 > 180.0 {
		bear180 = -(360.0 - bear360)
	} else {
		bear180 = bear360
	}
	return &bear180, nil
}

// returns the bearing from 'this' point to destination point along a rhumb line.
// adapted from geodesy https://github.com/chrisveness/geodesy/blob/master/latlon-spherical.js
func calculateRhumbBearing(from []float64, to []float64) float64 {
	φ := conversions.DegreesToRadians(from[1])
	φ2 := conversions.DegreesToRadians(to[1])
	Δλ := conversions.DegreesToRadians(to[0] - from[0])
	// if Δλ is over 180° take shorter rhumb line across the anti-meridian:
	if Δλ > math.Pi {
		Δλ -= 2 * math.Pi
	}
	if Δλ < -math.Pi {
		Δλ += 2 * math.Pi
	}

	Δψ := math.Log(math.Tan(φ2/2+math.Pi/4) / math.Tan(φ/2+math.Pi/4))
	θ := math.Atan2(Δλ, Δψ)
	tmp := conversions.RadiansToDegrees(θ) + 360.0
	return math.Mod(tmp, 360)
}

// RhumbDestination returns the destination having travelled the given distance along a Rhumb line from the origin Point with the (varant) given bearing.
// If you maintain a constant bearing along a rhumb line, you will gradually spiral towards one of the poles. ref. http://www.movable-type.co.uk/scripts/latlong.html#rhumblines
func RhumbDestination(origin geometry.Point, distance float64, bearing float64, units string, properties map[string]interface{}) (*feature.Feature, error) {
	wasNegativeDistance := distance < 0
	distanceInMeters, err := conversions.ConvertLength(math.Abs(distance), units, "meters")
	if err != nil {
		return nil, err
	}

	if wasNegativeDistance {
		distanceInMeters = -math.Abs(distanceInMeters)
	}
	coords, err := invariant.GetCoord(&origin)
	if err != nil {
		return nil, err
	}

	destination := calculateRhumbDestination(coords, distanceInMeters, bearing, nil)

	// compensate the crossing of the 180th meridian (https://macwright.org/2016/09/26/the-180th-meridian.html)
	// solution from https://github.com/mapbox/mapbox-gl-js/issues/3250#issuecomment-294887678

	if destination[0]-coords[0] > 180.0 {
		destination[0] += -360.0
	} else {
		if coords[0]-destination[0] > 180.0 {
			destination[0] += 360
		} else {
			destination[0] = 0
		}
	}

	result := feature.Feature{
		Type:       "Feature",
		Properties: properties,
		Bbox:       []float64{},
		Geometry: geometry.Geometry{
			GeoJSONType: "Point",
			Coordinates: destination,
		},
	}

	return &result, nil
}

// Adapted from Geodesy: http://www.movable-type.co.uk/scripts/latlong.html#rhumblines
func calculateRhumbDestination(origin []float64, distance float64, bearing float64, radius *float64) []float64 {
	if radius == nil {
		radius = common.Float64Ptr(constants.EarthRadius)
	}

	// angular distance in radians
	δ := distance / *radius
	// to radians, but without normalize to π
	λ1 := (origin[0] * math.Pi) / 180.0
	φ1 := conversions.DegreesToRadians(origin[1])
	θ := conversions.DegreesToRadians(bearing)

	Δφ := δ * math.Cos(θ)
	φ2 := φ1 + Δφ

	if math.Abs(φ2) > math.Pi/2 {
		if φ2 > 0 {
			φ2 = math.Pi - φ2
		} else {
			φ2 = -math.Pi - φ2
		}
	}

	Δψ := math.Log(math.Tan((φ2/2)+(math.Pi/4)) / math.Tan((φ1/2)+(math.Pi/4)))
	q := 0.0
	if math.Abs(Δψ) > 10e-12 {
		q = Δφ / Δψ
	} else {
		q = math.Cos(φ1)
	}

	Δλ := (δ * math.Sin(θ)) / q
	λ2 := λ1 + Δλ

	// normalise to -180...+180
	return []float64{(math.Mod(((λ2*180.0)/math.Pi+540), 360) - 180.0), (φ2 * 180.0) / math.Pi}
}
