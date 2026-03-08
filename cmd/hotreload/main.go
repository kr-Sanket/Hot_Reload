package main

import (
	"log"
	"time"

	"github.com/kr-Sanket/hotreload/internal/builder"
	"github.com/kr-Sanket/hotreload/internal/config"
	"github.com/kr-Sanket/hotreload/internal/debounce"
	"github.com/kr-Sanket/hotreload/internal/loghub"
	"github.com/kr-Sanket/hotreload/internal/process"
	"github.com/kr-Sanket/hotreload/internal/watcher"
)

func main() {

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Start log hub + dashboard
	hub := loghub.New()
	loghub.StartServer(hub)

	msg := "HotReload log dashboard running at http://localhost:8090"
	log.Println(msg)
	hub.Publish(msg)

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

	msg = "Watcher started"
	log.Println(msg)
	hub.Publish(msg)

	// Initial build (required by assignment)
	err = b.Build()
	if err != nil {

		msg = "Initial build failed"
		log.Println(msg)
		hub.Publish(msg)

	} else {

		msg = "Initial build successful"
		log.Println(msg)
		hub.Publish(msg)

		err = p.Start()
		if err != nil {
			log.Println("Failed to start server:", err)
		}
	}

	for range w.Events {

		db.Trigger(func() {

			msg := "File change detected"
			log.Println(msg)
			hub.Publish(msg)

			err := b.Build()
			if err != nil {

				msg := "Build failed — server will not restart"
				log.Println(msg)
				hub.Publish(msg)

				return
			}

			msg = "Build successful"
			log.Println(msg)
			hub.Publish(msg)

			err = p.Restart()
			if err != nil {

				msg := "Failed to restart server"
				log.Println(msg)
				hub.Publish(msg)

				return
			}

			msg = "Server restarted successfully"
			log.Println(msg)
			hub.Publish(msg)

		})

	}
}