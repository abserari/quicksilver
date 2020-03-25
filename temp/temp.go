package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	c := make(chan struct{})

	go func() {
		log.Println("Start")
		time.Sleep(time.Second)
		close(c)
		log.Println("end")
	}()
	<-c
	fmt.Println("Done")
}
