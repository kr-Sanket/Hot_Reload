package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kr-Sanket/hotreload/internal/config"
	"github.com/kr-Sanket/hotreload/internal/watcher"
)

func main() {

	cfg, err := config.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	w, err := watcher.New()
	if err != nil {
		log.Fatal(err)
	}

	err = w.Watch(cfg.Root)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Watcher started...")

	w.Start()
}