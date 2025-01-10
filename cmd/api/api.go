package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	// Add a new field to the application struct
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)
	return mux

}
func (app *application) run(mux *http.ServeMux) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting server on %s", app.config.addr)
	return server.ListenAndServe()
}
