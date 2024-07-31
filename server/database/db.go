package database

import (
	"os"

	"github.com/mquelucci/projeto-loja-virtual/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	tipoBanco := os.Getenv("TYPE")

	switch tipoBanco {
	case "sqlite3":
		nomeBanco := os.Getenv("FILEDATABASE")
		DB, err = gorm.Open(sqlite.Open(nomeBanco), &gorm.Config{})
		if err != nil {
			panic("Não foi possível conectar ao banco de dados" + err.Error())
		}
		DB.AutoMigrate(models.Produto{}, models.Cliente{}, models.Config{}, models.Admin{})
	case "postgres":
		stringDeConexao := os.Getenv("CONNECTIONSTRING")
		DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
		if err != nil {
			panic("Não foi possível conectar ao banco de dados" + err.Error())
		}
		DB.AutoMigrate(models.Produto{}, models.Cliente{}, models.Config{}, models.Admin{})
	default:
		panic("Nenhum tipo de banco de dados informado no arquivo config.env")
	}
}
