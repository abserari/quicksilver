package main

import (
	"fmt"
)

func main() {
	var n, c1, c2 int
	fmt.Scan(&n, &c1, &c2)

	result := 0
	mp := 0
	if c1 > c2 {
		mp = c2
	} else {
		mp = c1
	}
	var count = 0

	var win string
	fmt.Scan(&win)
	for i := 0; i < n; i++ {

		if win[i] == 70 {
			count++
		}
		if count == 3 {
			result += mp
			count = 0
		}
	}

	fmt.Println(result)

}
