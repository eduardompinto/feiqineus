package main

import (
	"github.com/eduardompinto/feiqineus/internal"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.DefaultServeMux
	mux.Handle("/suspicious", &internal.SuspiciousReceiver{})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    30 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
