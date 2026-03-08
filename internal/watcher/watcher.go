package watcher

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	fsWatcher *fsnotify.Watcher
}

func New() (*Watcher, error) {

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		fsWatcher: w,
	}, nil
}

func (w *Watcher) Watch(root string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			err = w.fsWatcher.Add(path)
			if err != nil {
				return err
			}
			log.Println("Watching:", path)
		}

		return nil
	})

	return err
}

func (w *Watcher) Start() {

	for {
		select {

		case event, ok := <-w.fsWatcher.Events:
			if !ok {
				return
			}

			log.Println("File event:", event)

		case err, ok := <-w.fsWatcher.Errors:
			if !ok {
				return
			}

			log.Println("Watcher error:", err)
		}
	}
}