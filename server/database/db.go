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
	DB, err := gorm.Open(sqlite.Open("banco.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(models.Produto{}, models.Cliente{})
}
