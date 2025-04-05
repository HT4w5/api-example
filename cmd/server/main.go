package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HT4w5/api-example/config"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "pong")
}

func main() {
	// Load config.
	cfg := config.New()
	cfg.Load()

	http.HandleFunc("/ping", pingHandler)
	log.Println("Starting server on " + cfg.HttpAddr)
	if err := http.ListenAndServe(cfg.HttpAddr, nil); err != nil {
		log.Println("Error starting server:", err)
	}

}
