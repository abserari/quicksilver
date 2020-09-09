package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	hang := make(chan bool)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	watcher.Add("./")

	go func() {
		for {
			select {
			case e := <-watcher.Events:
				log.Println(e.Op.String(), e.Name)
			case err := <-watcher.Errors:
				log.Println(err)
			}
		}
	}()

	<-hang
}
