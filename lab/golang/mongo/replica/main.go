package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type article struct {
	Title    string `bson:"title,omitempty"`
	Abstract string `bson:"abstract"`
	ReadCnt  int64  `bson:"reads,omitempty"`
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo-rs-n1:27017, mongo-rs-n2:27017, mongo-rs-n3:27017"))

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("simple").Collection("article")

	for i := 0; i < 200; i++ {
		art := article{
			Title:    fmt.Sprintf("T-%d", i),
			Abstract: fmt.Sprintf("Abstract %d", i),
			ReadCnt:  int64(i + 20),
		}

		if _, err := collection.InsertOne(context.Background(), &art); err != nil {
			log.Fatal(err)
		}
	}
}