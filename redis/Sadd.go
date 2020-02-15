package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	// set int64
	set, err := client.SAdd("userkey", "username").Result()

	fmt.Println(set, err)
}
