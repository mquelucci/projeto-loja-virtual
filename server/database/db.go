package database

import (
	"github.com/mquelucci/projeto-loja-virtual/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	DB, err = gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
	if err != nil {
		panic("Não foi possível conectar ao banco de dados" + err.Error())
	}
	DB.AutoMigrate(models.Produto{}, models.Cliente{}, models.Config{}, models.Admin{})
}
