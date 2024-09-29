package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers"
	"github.com/mquelucci/projeto-loja-virtual/middlewares"
)

func VendaRoutes(router *gin.Engine) {
	authAdminVendas := router.Group("/admin/vendas").Use(middlewares.Auth())
	{
		authAdminVendas.POST("/criar", controllers.CriarVenda)
		authAdminVendas.GET("/buscar/:id", controllers.BuscarVendaPorId)
	}
}
