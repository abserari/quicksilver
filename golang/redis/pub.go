/**
 * @author abser
 * @email [abser@foxmail.com]
 * @create date 2020-02-14 22:37:58
 * @modify date 2020-02-14 22:37:58
 * @desc [description]
 */
package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type foo struct {
	boo int `json:"boo"`
	hei int `json:"hei"`
}

func main() {
	var client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var hello = "hello"
	// hello := foo{boo: 1, hei: 2}
	// hello := 23.23
	go func() {
		for {
			t1 := time.Now()
			client.Publish("room1", hello)
			fmt.Println("pub", time.Now().Sub(t1))

			time.Sleep(time.Second)
		}
	}()

	pubsub := client.Subscribe("room1")
	_, err := pubsub.Receive()
	if err != nil {
		return
	}
	ch := pubsub.Channel()
	for msg := range ch {
		t2 := time.Now()
		fmt.Println(msg.Channel, msg.Payload)
		fmt.Println("sub", time.Now().Sub(t2))
	}

}
