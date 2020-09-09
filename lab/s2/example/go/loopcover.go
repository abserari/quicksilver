package main

import (
	"fmt"
	s2 "github.com/golang/geo/s2"
)
func main() {
	ll1 := s2.LatLngFromDegrees(31.803269, 113.421145)
	ll2 := s2.LatLngFromDegrees(31.461846, 113.695803)
	ll3 := s2.LatLngFromDegrees(31.250756, 113.756228)
	ll4 := s2.LatLngFromDegrees(30.902604, 113.997927)
	ll5 := s2.LatLngFromDegrees(30.817726, 114.464846)
	ll6 := s2.LatLngFromDegrees(30.850743, 114.76697)
	ll7 := s2.LatLngFromDegrees(30.713884, 114.997683)
	ll8 := s2.LatLngFromDegrees(30.430111, 115.42615)
	ll9 := s2.LatLngFromDegrees(30.088491, 115.640384)
	ll10 := s2.LatLngFromDegrees(29.907713, 115.656863)
	ll11 := s2.LatLngFromDegrees(29.783833, 115.135012)
	ll12 := s2.LatLngFromDegrees(29.712295, 114.728518)
	ll13 := s2.LatLngFromDegrees(29.55473, 114.24512)
	ll14 := s2.LatLngFromDegrees(29.530835, 113.717776)
	ll15 := s2.LatLngFromDegrees(29.55473, 113.3772)
	ll16 := s2.LatLngFromDegrees(29.678892, 112.998172)
	ll17 := s2.LatLngFromDegrees(29.941039, 112.349978)
	ll18 := s2.LatLngFromDegrees(30.040949, 112.025882)
	ll19 := s2.LatLngFromDegrees(31.803269, 113.421145)

	point1 := s2.PointFromLatLng(ll1)
	point2 := s2.PointFromLatLng(ll2)
	point3 := s2.PointFromLatLng(ll3)
	point4 := s2.PointFromLatLng(ll4)
	point5 := s2.PointFromLatLng(ll5)
	point6 := s2.PointFromLatLng(ll6)
	point7 := s2.PointFromLatLng(ll7)
	point8 := s2.PointFromLatLng(ll8)
	point9 := s2.PointFromLatLng(ll9)
	point10 := s2.PointFromLatLng(ll10)
	point11 := s2.PointFromLatLng(ll11)
	point12 := s2.PointFromLatLng(ll12)
	point13 := s2.PointFromLatLng(ll13)
	point14 := s2.PointFromLatLng(ll14)
	point15 := s2.PointFromLatLng(ll15)
	point16 := s2.PointFromLatLng(ll16)
	point17 := s2.PointFromLatLng(ll17)
	point18 := s2.PointFromLatLng(ll18)
	point19 := s2.PointFromLatLng(ll19)

	points := []s2.Point{}
	points = append(points, point19)
	points = append(points, point18)
	points = append(points, point17)
	points = append(points, point16)
	points = append(points, point15)
	points = append(points, point14)
	points = append(points, point13)
	points = append(points, point12)
	points = append(points, point11)
	points = append(points, point10)
	points = append(points, point9)
	points = append(points, point8)
	points = append(points, point7)
	points = append(points, point6)
	points = append(points, point5)
	points = append(points, point4)
	points = append(points, point3)
	points = append(points, point2)
	points = append(points, point1)

	loop := s2.LoopFromPoints(points)

	fmt.Println("----  loop search (gets too much) -----")
	// fmt.Printf("Some loop status items: empty:%t   full:%t \n", loop.IsEmpty(), loop.IsFull())

	// ref: https://github.com/golang/geo/issues/14#issuecomment-257064823
	defaultCoverer := &s2.RegionCoverer{MaxLevel: 20, MaxCells: 1000, MinLevel: 1}
	// rg := s2.Region(loop.CapBound())
	// cvr := defaultCoverer.Covering(rg)
	cvr := defaultCoverer.Covering(loop)

	// fmt.Println(poly.CapBound())
	for _, c3 := range cvr {
		fmt.Printf("%d,\n", c3)
	}
}
