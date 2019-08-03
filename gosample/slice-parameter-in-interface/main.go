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

// func (*myruler) distance(p *point) ([]int, error) {
// 	v := make([]int, 1)
// 	v = append(v, p.x+p.y)
// 	return v, nil
// }
func (*myruler) distance(p *point, origin ...point) ([]int, error) {
	v := make([]int, 1)
	v = append(v, p.x+p.y)
	return v, nil
}

func main() {
	var valid interface{}
	valid = new(myruler)
	if _, ok := valid.(ruler); ok {
		log.Println("ok")
	} else {
		log.Println("not")
	}

}
