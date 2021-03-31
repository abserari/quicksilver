package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	result := int64(0)

	var kint = make([]int64, k)

	var slice = make([]int64, 0)
	for i := 0; i < n; i++ {
		var integer int64 = 0
		fmt.Scan(&integer)
		slice = append(slice, integer)
	}

	for i := 0; i < len(slice); i++ {
		for s := 0; s < len(kint); s++ {
			if i+s >= len(slice) {
				break
			}
			kint[s] = slice[i] ^ slice[i+s]
		}
		for _, v := range kint {
			if v > result {
				result = v
			}
		}
	}

	fmt.Println(result)
}
