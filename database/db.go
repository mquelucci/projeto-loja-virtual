package database

import (
	"os"
	"strconv"

	"github.com/mquelucci/projeto-loja-virtual/models"
	"gorm.io/driver/mysql"
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
	case "postgres":
		stringDeConexao := os.Getenv("CONNECTIONSTRING")
		DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
		if err != nil {
			panic("Não foi possível conectar ao banco de dados" + err.Error())
		}
	case "mysql":
		stringDeConexao := os.Getenv("CONNECTIONSTRING")
		DB, err = gorm.Open(mysql.Open(stringDeConexao), &gorm.Config{})
		if err != nil {
			panic("Não foi possível conectar ao banco de dados" + err.Error())
		}

	default:
		panic("Nenhum tipo de banco de dados informado no arquivo config.env")
	}
	sqlDB, _ := DB.DB()

	openConns, err := strconv.Atoi(os.Getenv("OPENCONNS"))
	if err != nil {
		panic("Erro ao tentar ler a variável de ambiente OPENCONNS - " + err.Error())
	}
	idleConns, err := strconv.Atoi(os.Getenv("IDLECONNS"))
	if err != nil {
		panic("Erro ao tentar ler a variável de ambiente IDLECONNS - " + err.Error())
	}

	sqlDB.SetMaxIdleConns(idleConns)
	sqlDB.SetMaxOpenConns(openConns)
	// Mantém a estrutura do banco de dados sempre atualizadas
	DB.AutoMigrate(models.Migration{}, models.Produto{}, models.Cliente{}, models.Config{}, models.Admin{}, models.Venda{}, models.ItensVenda{})
}
