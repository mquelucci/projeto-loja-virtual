package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/routes"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}
	database.ConectaBanco()
	routes.HandleRequests()
}
