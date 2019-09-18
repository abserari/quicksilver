package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("i'm doing for sum: ", sum)
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println("叫醒第一个人工作 17")
	go sum(s[:len(s)/2], c)
	fmt.Println("叫醒下一个人工作 -5")
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
