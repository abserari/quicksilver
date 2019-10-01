package main

import (
	"fmt"
	"log"
	"reflect"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("./data.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}

		if err = b.Put([]byte("answer"), []byte("42")); err != nil {
			return err
		}

		if err = b.Put([]byte("zero"), []byte("")); err != nil {
			return err
		}

		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("noexists"))
		fmt.Println(reflect.DeepEqual(v, nil)) // false
		fmt.Println(v == nil)                  // true

		v = b.Get([]byte("zero"))
		fmt.Println(reflect.DeepEqual(v, nil)) // false
		fmt.Println(v == nil)                  // true

		c := b.Cursor()
		fmt.Println(c.First())
		k, v := c.Prev()
		fmt.Println(k == nil, v == nil) // true,true

		return nil
	})
}
