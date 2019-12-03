package main

import (
	"fmt"

	s2 "github.com/golang/geo/s2"
)

func main() {
	p1 := s2.PointFromLatLng(s2.LatLngFromDegrees(38.8804099451, 115.5554008484))
	p2 := s2.PointFromLatLng(s2.LatLngFromDegrees(38.8769270981, 115.5556261539))
	p3 := s2.PointFromLatLng(s2.LatLngFromDegrees(38.8768936887, 115.5601751804))
	p4 := s2.PointFromLatLng(s2.LatLngFromDegrees(38.8803681853, 115.5602824688))

	points := []s2.Point{}
	points = append(points, p1)
	points = append(points, p2)
	points = append(points, p3)
	points = append(points, p4)

	loop := s2.LoopFromPoints(points)

	defaultCoverer := s2.RegionCoverer{MaxLevel: 20, MaxCells: 1000, MinLevel: 1}

	cvr := defaultCoverer.Covering(loop)

	for _, c3 := range cvr {
		fmt.Printf("%x,\n", c3)
	}
}
