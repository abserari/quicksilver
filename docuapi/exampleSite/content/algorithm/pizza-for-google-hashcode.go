package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

const filename = "e.in"

func main() {
	var buf []string

	date := time.Now()
	// Open save file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(file))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		buf = append(buf, string(scanner.Bytes()))
	}

	maxSlice := atoi(buf[0])

	fmt.Println(maxSlice)

	for i := 0; i < atoi(buf[1]); i++ {
		if maxSlice-atoi(buf[i+2]) < 0 {
			goto is
		}
		maxSlice -= atoi(buf[i+2])
	}
is:
	point := atoi(buf[0]) - maxSlice
	time := time.Now().Sub(date)
	fmt.Println(point, time)
}

func atoi(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
