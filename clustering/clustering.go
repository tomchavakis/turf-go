package clustering

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/measurement"
	meta "github.com/tomchavakis/turf-go/meta/coordAll"
)

// Different Distance Approaches
// Euclidian distance or Haversine Distance
// TODO: Add Vertical Dimension

// iterations

type Distance string

const (
	Euclidian Distance = "Euclidian"
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

func getClusters(params Parameters) (map[geometry.Point][]geometry.Point, error) {
	ctrIdx, centroids := getCentroids(params)

	tmpCluster := make(map[geometry.Point][]geometry.Point)
	meanCluster := make(map[geometry.Point][]geometry.Point)

	init := false
	for {
		tmpMeanCluster := make(map[geometry.Point][]geometry.Point)
		if !init {
			// create a cluster of points based on random centroids
			for i, p := range params.points {
				if _, isCentroid := ctrIdx[i]; isCentroid {
					tmpCluster[p] = tmpCluster[p]
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

// getDistance returns the distance between the point and the centroids
func getDistance(p geometry.Point, centroids []geometry.Point, dt Distance) (map[int]float64, error) {
	//TODO: Implement Euclidian Distance
	if dt == Euclidian {
		return nil, nil
	} else { // haversine distance implemented
		ds := make(map[int]float64)

		for i, c := range centroids {
			d, err := measurement.Distance(p.Lng, p.Lat, c.Lng, c.Lat, "")
			if err != nil {
				return nil, err
			}
			ds[i] = d
		}

		return ds, nil
	}
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
