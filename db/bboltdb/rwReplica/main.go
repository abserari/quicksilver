package main

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	bolt "go.etcd.io/bbolt"
)

func main() {
	hang := make(chan bool, 1)
	// go fsnotify.Watch(os.Path())
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	db, err := bolt.Open("./data.db", 0600, nil)
	go db.Update(func(tx *bolt.Tx) error {
		time.Sleep(time.Second)
		b, err := tx.CreateBucketIfNotExists([]byte("cats"))
		if err != nil {
			return err
		}
		b.Put([]byte("techcats"), []byte("100"))
		return nil
	})

	go db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("cats"))
		if err != nil {
			return err
		}
		b.Put([]byte("techcats"), []byte("101"))
		return nil
	})

	<-hang
	// botldb.Opendatabase()
}
