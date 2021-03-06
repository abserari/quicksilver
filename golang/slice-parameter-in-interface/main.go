package main

import "log"

type point struct {
	x int
	y int
}

type ruler interface {
	distance(p *point, origin ...point) ([]int, error)
}

type myruler struct{}

// // print !ok
// func (*myruler) distance(p *point) ([]int, error) {
// 	v := make([]int, 1)
// 	v = append(v, p.x+p.y)
// 	return v, nil
// }

// print : ok
func (*myruler) distance(p *point, origin ...point) ([]int, error) {
	v := make([]int, 1)
	v = append(v, p.x+p.y)
	return v, nil
}

func main() {
	var valid interface{}
	valid = new(myruler)
	_, ok := valid.(ruler)
	if ok {
		log.Println("ok")
	} else {
		log.Println("not")
	}

	p := &point{x: 1, y: 3}

	valid.(ruler).distance(p)
}
