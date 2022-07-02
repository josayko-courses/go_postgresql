package main

import (
	"net/http"
)

func (app *Application) serve() error {
	srv := http.Server{
		Handler: app.handlers(),
		Addr:    ":4005",
	}
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (app *Application) handlers() http.Handler {
	mux := http.NewServeMux()
	return mux
}
