package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers/responses"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, responses.Error{Erro: "URL não encontrada!"})
}
