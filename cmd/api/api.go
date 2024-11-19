package main

import (
	"log"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	address string
}

func (app *application) run() error {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    app.config.address,
		Handler: mux,
	}

	log.Printf("Server listening on port %s", app.config.address)

	return srv.ListenAndServe()
}
