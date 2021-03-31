package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	var result = make([][]int, 0)
	for i := 0; i < n; i++ {
		var count int
		var x int

		fmt.Scan(&count)
		var slices = make([]int, 0)

		for i := 0; i < count; i++ {

			fmt.Scan(&x)

			slices = append(slices, x)
		}

		/// 中间值
		var leftcount = make(map[int]int)
		if count == 1 {
			slices[0] = 0
		} else if count%2 == 0 {
			inter := count / 2
			left := inter - 1
			var index = 0
			for ; inter < count; inter++ {
				if slices[inter] == slices[left] {
					left--
					continue
				} else {
					if slices[inter] > slices[left] {
						leftcount[inter] = slices[left]
						index = inter
					} else {
						leftcount[left] = slices[inter]
						index = left
					}
				}

				left--
			}

			if len(leftcount) == 0 {

			} else if len(leftcount) > 1 {
				for i, v := range slices {
					if v == 0 {
						continue
					} else {
						slices[i] = 0
						break
					}
				}
			} else {
				slices[index] = leftcount[index]
			}
		} else {
			inter := count / 2
			left := inter - 1
			inter++
			var index = 0
			for ; inter < count; inter++ {
				if slices[inter] == slices[left] {
					left--
					continue
				} else {
					if slices[inter] > slices[left] {
						leftcount[inter] = slices[left]
						index = inter
					} else {
						leftcount[left] = slices[inter]
						index = left
					}
				}

				left--
			}

			if len(leftcount) == 0 {
				slices[count/2] = 0
			} else if len(leftcount) > 1 {
				for i, v := range slices {
					if v == 0 {
						continue
					} else {
						slices[i] = 0
						break
					}
				}
			} else {
				slices[index] = leftcount[index]
			}
		}
		result = append(result, slices)
	}

	for _, v := range result {
		var display string
		for _, i := range v {
			display += strconv.Itoa(i)
		}
		fmt.Println(display)
	}
}
