package watcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	fsWatcher *fsnotify.Watcher
	Events    chan struct{}
}

func New() (*Watcher, error) {

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		fsWatcher: w,
		Events:    make(chan struct{}),
	}, nil
}

func (w *Watcher) Watch(root string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {

			if shouldIgnore(path) {
				return filepath.SkipDir
			}

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

			if shouldIgnore(event.Name) {
				continue
			}

			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
				w.Events <- struct{}{}
			}

		case err, ok := <-w.fsWatcher.Errors:
			if !ok {
				return
			}

			log.Println("Watcher error:", err)
		}
	}
}

func shouldIgnore(path string) bool {

	ignored := []string{
		".git",
		"bin",
		"node_modules",
	}

	for _, dir := range ignored {
		if strings.Contains(path, dir) {
			return true
		}
	}

	return false
}