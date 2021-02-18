package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var err error

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	handler := cors.Default().Handler(a.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}

func (a *App) initializeRoutes() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	a.Router.HandleFunc("/healthz", a.healthz).Methods("GET")
}

func (a *App) healthz(w http.ResponseWriter, req *http.Request) {
	log.Info("API health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
