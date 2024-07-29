package main

import (
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandleRequests()
}
