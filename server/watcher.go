package server

import (
	"log"

	"github.com/ink0rr/rockide/internal/protocol"
)

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
