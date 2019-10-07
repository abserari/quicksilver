package main

import (
	"encoding/binary"
	"log"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

var insertNum uint64 = 200000

func main() {
	hang := make(chan bool)
	db, err := bolt.Open("./data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		times := time.Now()
		for {
			db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte("cats"))
				if err != nil {
					return err
				}

				num, err := b.NextSequence()
				log.Println(num)
				byteid := make([]byte, 8)
				binary.BigEndian.PutUint64(byteid, num)

				b.Put(byteid, byteid)

				if num == insertNum {
					log.Println(time.Now().Sub(times))
					os.Exit(0)
				}
				return nil
			})
		}

	}()

	<-hang
	// go db.Update(func(tx *bolt.Tx) error {
	// 	b, err := tx.CreateBucketIfNotExists([]byte("cats"))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	b.Put([]byte("techcats"), []byte("101"))
	// 	return nil
	// })

	// botldb.Opendatabase()
}
