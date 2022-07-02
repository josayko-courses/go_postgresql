package main

import (
	"encoding/json"
	"log"
	"main/models"
	"net/http"
)

func (app *Application) Serve() error {
	srv := http.Server{
		Handler: app.Handlers(),
		Addr:    ":4005",
	}
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (app *Application) Handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HomePage)
	return mux
}

func (app *Application) HomePage(w http.ResponseWriter, req *http.Request) {
	f := models.Filter{
		Page:     1,
		PageSize: 20,
	}
	users, meta, err := app.Models.Users.GetAll(f)

	if err != nil {
		log.Fatalln(err)
	}

	res := struct {
		Users []models.User
		Meta  models.Metadata
	}{
		Users: users,
		Meta:  meta,
	}
	js, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(js)
}
