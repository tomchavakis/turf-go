package clustering

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/tomchavakis/geojson/geometry"
	"github.com/tomchavakis/turf-go/measurement"
	meta "github.com/tomchavakis/turf-go/meta/coordAll"
)

type Distance string

const (
	Euclidean Distance = "Euclidean"
	Haversine Distance = "Haversine"
)

// Parameters for the KMean Clustering
type Parameters struct {
	k            int              // number of clusters
	points       []geometry.Point // pointSet
	distanceType Distance
}

// KMeans initialisation
// http://ilpubs.stanford.edu:8090/778/1/2006-13.pdf
func KMeans(params Parameters) (map[geometry.Point][]geometry.Point, error) {

	if params.k < 2 {
		return nil, fmt.Errorf("at least 2 centroids required")
	}

	if params.k > len(params.points) {
		return nil, fmt.Errorf("clusters can't be more than the length of the points")
	}

	return getClusters(params)

}

func initialisation(tmpCluster map[geometry.Point][]geometry.Point, centroids []geometry.Point, ctrIdx map[int]bool, params Parameters) (map[geometry.Point][]geometry.Point, error) {
	// create a cluster of points based on random centroids
	for i, p := range params.points {
		if _, isCentroid := ctrIdx[i]; isCentroid {
			//tmpCluster[p] = tmpCluster[p]
			continue
		}

		// get the distance from all the centroids
		ptCtrsDistances, err := getDistance(p, centroids, params.distanceType)
		if err != nil {
			return nil, err
		}

		// get the minimum distance index for this point
		minDistanceIdx := minDistanceIdx(ptCtrsDistances)
		nearestCentroid := centroids[minDistanceIdx]
		tmpCluster[nearestCentroid] = append(tmpCluster[nearestCentroid], p)
	}
	return tmpCluster, nil
}

func getClusters(params Parameters) (map[geometry.Point][]geometry.Point, error) {
	ctrIdx, centroids := getCentroids(params)

	tmpCluster := make(map[geometry.Point][]geometry.Point)
	meanCluster := make(map[geometry.Point][]geometry.Point)

	init := false
	for {
		tmpMeanCluster := make(map[geometry.Point][]geometry.Point)

		if !init {
			tmpCl, err := initialisation(tmpCluster, centroids, ctrIdx, params)
			if err != nil {
				return nil, err
			}
			tmpCluster = tmpCl
		}

		// calculate the mass mean of each cluster
		clusterMeanPoints := []geometry.Point{}
		for i, c := range tmpCluster {
			// median included
			if !init {
				c = append(c, i)
				init = true
			}
			meanClusterPoint, err := meanClusterPoint(c)
			if err != nil {
				return nil, err
			}
			clusterMeanPoints = append(clusterMeanPoints, *meanClusterPoint)
		}

		for _, p := range params.points {

			// get the distance from all the cluster mean
			meanCtrsDistances, err := getDistance(p, clusterMeanPoints, params.distanceType)
			if err != nil {
				return nil, err
			}

			// get the minimum distance index for this point
			minDistanceIdx := minDistanceIdx(meanCtrsDistances)
			nearestMean := clusterMeanPoints[minDistanceIdx]
			tmpMeanCluster[nearestMean] = append(tmpMeanCluster[nearestMean], p)
		}
		tmpCluster = tmpMeanCluster
		// exit point
		if isEqual(meanCluster, tmpMeanCluster) {
			break
		}

		meanCluster = tmpMeanCluster
	}

	//getDistance(d)
	return meanCluster, nil
}

// TODO: Find the centroid

// https://sites.google.com/site/yangdingqi/home/foursquare-dataset
// https://desktop.arcgis.com/en/arcmap/latest/tools/spatial-statistics-toolbox/h-how-mean-center-spatial-statistics-works.html
// https://postgis.net/docs/ST_Centroid.html
// https://cloud.google.com/bigquery/docs/reference/standard-sql/geography_functions#st_centroid
// https://stackoverflow.com/questions/30299267/geometric-median-of-multidimensional-points
// https://www.pnas.org/content/pnas/97/4/1423.full.pdf
func meanClusterPoint(cluster []geometry.Point) (*geometry.Point, error) {
	var pointSet = []geometry.Point{}
	excludeWrapCoord := true
	for _, v := range cluster {
		coords, err := meta.CoordAll(&v, &excludeWrapCoord)
		if err != nil {
			return nil, err
		}
		pointSet = append(pointSet, coords...)
	}

	coordsLength := len(pointSet)
	if coordsLength < 1 {
		return nil, errors.New("no coordinates found")
	}

	xSum := 0.0
	ySum := 0.0

	for i := 0; i < coordsLength; i++ {
		xSum += pointSet[i].Lng
		ySum += pointSet[i].Lat
	}

	finalCenterLongtitude := xSum / float64(coordsLength)
	finalCenterLatitude := ySum / float64(coordsLength)

	return &geometry.Point{
		Lat: finalCenterLatitude,
		Lng: finalCenterLongtitude,
	}, nil

}

func getCentroids(params Parameters) (map[int]bool, []geometry.Point) {
	var centroids []geometry.Point
	ctrIdx := make(map[int]bool)

	idx := getRandoms(len(params.points), params.k)
	for _, v := range idx {
		centroids = append(centroids, params.points[v])
		ctrIdx[v] = true
	}

	return ctrIdx, centroids
}

// l length, k number of clusters
func getRandoms(l int, k int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(l)[:k]
}

func euclideanDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	dist := math.Sqrt(math.Pow(lat2-lat1, 2) + math.Pow(lon2-lon1, 2))
	return dist
}

// getDistance returns the distance between the point and the centroids
func getDistance(p geometry.Point, centroids []geometry.Point, dt Distance) (map[int]float64, error) {
	ds := make(map[int]float64)

	if dt == Euclidean {
		for i, c := range centroids {
			d := euclideanDistance(p.Lat, p.Lng, c.Lat, c.Lng)
			ds[i] = d
		}
	} else { // haversine distance implemented
		for i, c := range centroids {
			d, err := measurement.Distance(p.Lng, p.Lat, c.Lng, c.Lat, "")
			if err != nil {
				return nil, err
			}
			ds[i] = d
		}
	}
	return ds, nil
}

func memoizeCluster(key geometry.Point, cluster []geometry.Point) map[string]bool {
	mr := make(map[string]bool)
	for _, v := range cluster {
		s := memoizationSignature(key, v)
		mr[s] = true
	}
	return mr
}

func memoizationSignature(key geometry.Point, p geometry.Point) string {
	return fmt.Sprintf("%f_%f_%f", key, p.Lat, p.Lng)
}

func mergeMaps(maps ...map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func isEqual(clusterA map[geometry.Point][]geometry.Point, clusterB map[geometry.Point][]geometry.Point) bool {
	if len(clusterA) != len(clusterB) {
		return false
	}

	memo := make(map[string]bool)
	for i, v1 := range clusterA {
		tmp := memoizeCluster(i, v1)
		memo = mergeMaps(memo, tmp)
	}

	for j, arrB := range clusterB {
		for _, v := range arrB {
			if !memo[memoizationSignature(j, v)] {
				return false
			}
		}
	}

	return true
}

func minDistanceIdx(ptCtrsDistances map[int]float64) int {
	minDistanceIdx := 0
	minD := ptCtrsDistances[minDistanceIdx]
	for i, d := range ptCtrsDistances {
		if d < minD {
			minD = d
			minDistanceIdx = i
		}
	}
	return minDistanceIdx
}
