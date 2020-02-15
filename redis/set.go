package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	set, err := client.SetNX("userkey", "username", 10*time.Second).Result()

	fmt.Println(set, err)
}
