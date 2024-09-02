package utils

import (
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func BuscarConfigs() models.Config {
	var configs models.Config
	database.DB.First(&configs)
	return configs
}
