package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event: ", event)
			case e, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", e)
			}
		}
	}()

	// watch current folder
	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
