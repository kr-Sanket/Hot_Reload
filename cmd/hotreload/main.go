package main

import (
	"log"
	"time"

	"github.com/kr-Sanket/hotreload/internal/builder"
	"github.com/kr-Sanket/hotreload/internal/config"
	"github.com/kr-Sanket/hotreload/internal/debounce"
	"github.com/kr-Sanket/hotreload/internal/process"
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

	b := builder.New(cfg.Build)
	p := process.New(cfg.Exec)

	db := debounce.New(300 * time.Millisecond)

	go w.Start()

	log.Println("Watcher started")

	err = b.Build()
	if err != nil {
		log.Println("Initial build failed")
	} else {
		p.Start()
	}

	for range w.Events {

		db.Trigger(func() {

			log.Println("File change detected")

			err := b.Build()
			if err != nil {
				log.Println("Build failed — server will not restart")
				return
			}

			err = p.Restart()
			if err != nil {
				log.Println("Failed to restart server:", err)
			}

		})

	}
}