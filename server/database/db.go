package database

import (
	"github.com/mquelucci/projeto-loja-virtual/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	DB, err = gorm.Open(postgres.Open("user=postgres.hmvdnlleavjmxqdzztzp password=q1w2Q!W@z1x2Z!X@ host=aws-0-sa-east-1.pooler.supabase.com port=5432 dbname=postgres"), &gorm.Config{})
	if err != nil {
		panic("Não foi possível conectar ao banco de dados" + err.Error())
	}
	DB.AutoMigrate(models.Produto{}, models.Cliente{}, models.Config{}, models.Admin{})
}
