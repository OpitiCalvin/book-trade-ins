package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	a := App{}

	a.Initialize()

	defer a.DB.Close()
	a.Migrations()

	log.Info("Server running on port 8080...")
	a.Run(":8080")
}
