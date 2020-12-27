package main

import (
	"github.com/eduardompinto/feiqineus/internal"
	"log"
	"net/http"
	"time"
)

func main() {
	db := internal.NewDatabase()
	defer db.Close()
	mux := http.DefaultServeMux
	vmdb := internal.NewVerifiedMessageDB(db)
	replier := internal.NewReplier(vmdb, vmdb)
	internal.NewTelegramIntegration(&replier).Init()
	mux.Handle("/suspicious", internal.NewSuspiciousReceiver(&replier))
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    30 * time.Second,
	}
	log.Printf("Server started!")
	log.Fatal(s.ListenAndServe())
}
