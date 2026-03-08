package main

import (
	"log"
	"time"

	"github.com/kr-Sanket/hotreload/internal/config"
	"github.com/kr-Sanket/hotreload/internal/debounce"
	"github.com/kr-Sanket/hotreload/internal/watcher"
)

func main() {

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	w, err := watcher.New()
	if err != nil {
		log.Fatal(err)
	}

	err = w.Watch(cfg.Root)
	if err != nil {
		log.Fatal(err)
	}

	db := debounce.New(300 * time.Millisecond)

	go w.Start()

	log.Println("Watcher started")

	for range w.Events {

		db.Trigger(func() {
			log.Println("Change detected → rebuild triggered")
		})

	}
}