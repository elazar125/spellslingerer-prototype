package main

import (
	"api/db"
	"api/models"
	"api/router"

	"log"
	"os"
)

func migrateDatabase() error {
	if err := db.GlobalDB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatalln(err)
	}

	if err := migrateDatabase(); err != nil {
		log.Fatalln(err)
	}

	r := router.SetupRouter()
	err := r.Run(os.Getenv("API_PORT"))
	// err := r.RunTLS(os.Getenv("API_PORT"), os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"))

	if err != nil {
		log.Fatalln(err)
	}
}
