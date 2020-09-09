package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]string)
	for i := 0; i < 1e7; i++ {
		m[string(i)] = fmt.Sprintf("haa")
	}

	start := time.Now()
	for _, _ = range m {

	}
	fmt.Println(time.Now().Sub(start))
}
