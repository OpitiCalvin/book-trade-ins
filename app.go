package main

import (
	"io"
	"net/http"

	datamodels "github.com/OpitiCalvin/novelsTradeIn/datamodels"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var err error

// App structure with pointer to mux router and pointer to gorm database as fields
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize method initializes the mux router, database connection and API routes
func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	// connect to sqlite database
	a.DB, err = gorm.Open("sqlite3", "./novels_trade_in.db")
	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	}

	log.Info("Connected successfully to database")

	a.initializeRoutes()
}

// Run executes the application to start listening and serving API content
func (a *App) Run(addr string) {
	handler := cors.Default().Handler(a.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}

// Migrations handle migrations to the database, modifying table fields where necessary
func (a *App) Migrations() {
	a.DB.AutoMigrate(&datamodels.Book{})
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
