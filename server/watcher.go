package server

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/arexon/fsnotify"
	"github.com/ink0rr/rockide/internal/protocol"
)

func watchFiles() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	if err := watcher.Add(filepath.Join(project.BP, "...")); err != nil {
		return errors.Join(err, errors.New("failed to watch BP path"))
	}
	if err := watcher.Add(filepath.Join(project.RP, "...")); err != nil {
		return errors.Join(err, errors.New("failed to watch RP path"))
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				uri := protocol.URIFromPath(event.Name)
				if event.Op.Has(fsnotify.Remove | fsnotify.Rename) {
					onDelete(uri)
					continue
				}
				if stat, err := os.Stat(event.Name); err != nil || stat.IsDir() {
					continue
				}
				switch {
				case event.Op.Has(fsnotify.Create):
					onCreate(uri)
				case event.Op.Has(fsnotify.Write):
					onChange(uri)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()
	return nil
}

func onCreate(uri protocol.DocumentURI) {
	store := findStore(uri)
	if store == nil {
		return
	}
	log.Printf("create: %s", uri)
	if err := store.Parse(uri); err != nil {
		log.Println(err)
	}
}

func onChange(uri protocol.DocumentURI) {
	store := findStore(uri)
	if store == nil {
		return
	}
	log.Printf("change: %s", uri)
	store.Delete(uri)
	if err := store.Parse(uri); err != nil {
		log.Println(err)
	}
}

func onDelete(uri protocol.DocumentURI) {
	store := findStore(uri)
	if store == nil {
		return
	}
	log.Printf("delete: %s", uri)
	store.Delete(uri)
}
