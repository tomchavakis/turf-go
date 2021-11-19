# :hammer: [![GitHub release][release-img]][release] [![Build Status][ci-img]][ci] [![Coverage Status](https://coveralls.io/repos/github/tomchavakis/turf-go/badge.svg?branch=master)](https://coveralls.io/github/tomchavakis/turf-go?branch=master)

# turf-go
A Go language port of [Turfjs](http://turfjs.org/docs/)

## Turf for Go

Turf for Go is a ported library in GoLang ported from the Turf.js library.

# Ported functions

## measurement

- [x] along
- [x] area
- [x] bbox
- [x] bboxPolygon
- [x] bearing
- [x] center
- [ ] centerOfMass
- [x] centroid
- [x] destination
- [x] distance
- [x] envelope
- [x] length
- [x] midpoint
- [ ] pointOnFeature
- [ ] polygonTangents
- [ ] pointToLineDistance
- [ ] rhumbBearing
- [ ] rhumbDestination
- [ ] rhumbDistance
- [ ] square
- [ ] greatCircle

## Coordinate Mutation
- [ ] cleanCoords
- [ ] flip
- [ ] rewind
- [ ] round
- [ ] truncate

## Transformation
- [ ] bboxClip
- [ ] bezierSpline
- [ ] buffer
- [ ] circle
- [ ] clone
- [ ] concave
- [ ] convex
- [ ] difference
- [ ] dissolve
- [ ] intersect
- [ ] lineOffset
- [ ] simplify
- [ ] tesselate
- [ ] transformRotate
- [ ] transformTranslate
- [ ] transformScale
- [ ] union
- [ ] voronoi

## Feature Conversion
- [ ] combine
- [ ] explode
- [ ] flatten
- [ ] lineToPolygon
- [ ] polygonize
- [ ] polygonToLine

## Misc
- [ ] kinks
- [ ] lineArc
- [ ] lineChunk
- [ ] lineIntersect
- [ ] lineOverlap
- [ ] lineSegment
- [ ] lineSlice
- [ ] lineSliceAlong
- [ ] lineSplit
- [ ] mask
- [ ] nearestPointOnLine
- [ ] sector
- [ ] shortestPath
- [ ] unkinkPolygon

## Helper
- [x] featureCollection
- [x] feature
- [x] geometryCollection
- [x] lineString
- [x] multiLineString
- [x] multiPoint
- [x] multiPolygon
- [x] point
- [x] polygon        

## Random
- [ ] randomPosition
- [ ] randomPoint
- [ ] randomLineString
- [ ] randomPolygon

## Data
- [ ] sample

## Joins
- [x] pointsWithinPolygon
- [ ] tag

## Grids
- [ ] hexGrid
- [ ] pointGrid
- [ ] squareGrid
- [ ] triangleGrid

## Classification
- [x] nearestPoint

## Aggregation
- [ ] collect
- [ ] clustersDbscan
- [ ] clustersKmeans

## Meta
- [x] coordAll
- [x] coordEach
- [ ] coordReduce
- [ ] featureEach
- [ ] featureReduce
- [ ] flattenEach
- [ ] flattenReduce
- [ ] getCoord
- [ ] getCoords
- [ ] getGeom
- [ ] getType
- [ ] geomEach
- [ ] geomReduce
- [ ] propEach
- [ ] propReduce
- [ ] segmentEach
- [ ] segmentReduce
- [ ] getCluster
- [ ] clusterEach
- [ ] clusterReduce

## Assertions
- [ ] collectionOf
- [ ] containsNumber
- [ ] geojsonType
- [ ] featureOf 

## Booleans
- [ ] booleanClockwise
- [ ] booleanContains
- [ ] booleanCrosses
- [ ] booleanDisjoint
- [ ] booleanEqual
- [ ] booleanOverlap
- [ ] booleanParallel
- [x] booleanPointInPolygon
- [ ] booleanPointOnLine
- [ ] booleanWithin

## Unit Conversion 
- [x] bearingToAzimuth
- [x]  convertArea
- [x]  convertLength
- [x] degreesToRadians
- [x]  lengthToRadians
- [x]  lengthToDegrees
- [x]  radiansToLength
- [x] radiansToDegrees
- [x]  toMercator
- [x]  toWgs84





## References:

https://github.com/mapbox/mapbox-java

https://github.com/Turfjs/turf
