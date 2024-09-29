package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/router"
	"github.com/mquelucci/projeto-loja-virtual/seeders"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}
	database.ConectaBanco()
	seeders.RunSeeders(database.DB)
	router.HandleRequests()
}
