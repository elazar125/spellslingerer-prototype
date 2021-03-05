package main

import (
	"api/migrate"
	"api/router"

	"log"
	"os"
)

func main() {
	if err := migrate.SetupAndMigrateDatabase(); err != nil {
		log.Fatalln(err)
	}

	r := router.SetupRouter()
	err := r.Run(os.Getenv("API_PORT"))
	// TODO: Set up ssl cert
	// err := r.RunTLS(os.Getenv("API_PORT"), os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"))

	if err != nil {
		log.Fatalln(err)
	}
}
