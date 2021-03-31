package main

import (
	"fmt"
)

type position struct {
	first int
	last  int
	flag  bool
}

func NewPostion(first, last int) *position {
	return &position{
		first: first,
		last:  last,
		flag:  true,
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var maps = make(map[int]*position)
	for i := 0; i < n; i++ {
		var integer = 0
		fmt.Scan(&integer)
		if maps[integer] != nil {
			maps[integer].last = i + 1
		} else {
			maps[integer] = NewPostion(i+1, i+1)
		}
	}

	result := make([]int, 2*m)

	for i := 0; i < m; i++ {
		var integer = 0
		fmt.Scan(&integer)
		if maps[integer] != nil {
			result[2*i] = maps[integer].first
			result[2*i+1] = maps[integer].last
		} else {
			result[2*i] = 0
			result[2*i+1] = 0
		}
	}

	for i := 0; i < len(result); i = i + 2 {
		if result[i] == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(result[i], result[i+1])
		}
	}

}
