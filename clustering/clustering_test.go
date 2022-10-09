package clustering

import (
	"testing"

	"github.com/tomchavakis/geojson/geometry"
	"github.com/tomchavakis/turf-go/assert"
)

// SELECT ST_AsText(ST_Centroid('MULTIPOINT ( 10.0 20.0 , 10.0 25.0, 11.0 18.0, 10.0 18.0 )'));
// SELECT ST_AsText(ST_Centroid('MULTIPOINT ( 10.0 60.0,  11.0 50.0 )'));
func TestKMeans(t *testing.T) {
	params := Parameters{
		k:            2,
		points:       []geometry.Point{{Lat: 20.0, Lng: 10.0}, {Lat: 25.0, Lng: 10.0}, {Lat: 60.0, Lng: 10.0}, {Lat: 18.0, Lng: 11.0}, {Lat: 18.0, Lng: 10.0}, {Lat: 50.0, Lng: 11.0}},
		distanceType: Haversine,
	}

	res, err := KMeans(params)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	clusters := make(map[geometry.Point][]geometry.Point)
	p1 := geometry.Point{
		Lat: 20.25,
		Lng: 10.25,
	}
	p1Cluster := []geometry.Point{}
	p1Cluster = append(p1Cluster, *geometry.NewPoint(20.0, 10.0))
	p1Cluster = append(p1Cluster, *geometry.NewPoint(25.0, 10.0))
	p1Cluster = append(p1Cluster, *geometry.NewPoint(18.0, 11.0))
	p1Cluster = append(p1Cluster, *geometry.NewPoint(18.0, 10.0))

	p2 := geometry.Point{
		Lat: 55.0,
		Lng: 10.5,
	}
	p2Cluster := []geometry.Point{}
	p2Cluster = append(p2Cluster, *geometry.NewPoint(60.0, 10.0))
	p2Cluster = append(p2Cluster, *geometry.NewPoint(50.0, 11.0))

	clusters[p1] = p1Cluster
	clusters[p2] = p2Cluster

	assert.Equal(t, len(res), 2)
	assert.Equal(t, res, clusters)
}
