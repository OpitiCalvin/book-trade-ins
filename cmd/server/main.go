package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/OpitiCalvin/novelsTradeIn/pkg/api"
	"github.com/OpitiCalvin/novelsTradeIn/pkg/app"
	"github.com/OpitiCalvin/novelsTradeIn/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}

}

// func run is responsible for setting up db connections, routers etc
func run() error {
	// load env params

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err) // error loading env file
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	host := os.Getenv("HOST")
	DBPort := os.Getenv("DBPORT")

	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, DBPort, dbname)

	// setup a connection to database
	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db)

	// run migrations
	// NOTE that we are passing the connectionString again here. This is so we can easily run migrations
	// against another database, say a test version, for out integration and end-to-end tests
	err = storage.RunMigrations(connectionString)

	if err != nil {
		return err
	}

	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	// create user service
	userService := api.NewUserService(storage)

	// create book service
	bookService := api.NewBookService(storage)

	server := app.NewServer(router, userService, bookService)

	// start server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported databse you want to use
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
