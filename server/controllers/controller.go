package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
)

func ExibeHTMLIndex(c *gin.Context) {
	database.DB.Find()
}
