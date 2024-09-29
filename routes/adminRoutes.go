package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers"
	"github.com/mquelucci/projeto-loja-virtual/middlewares"
)

func AdminRoutes(router *gin.Engine) {
	noAuthAdmin := router.Group("/admin")
	{
		noAuthAdmin.POST("/autenticar", controllers.Autenticar)

	}

	authAdmin := router.Group("/admin").Use(middlewares.Auth())
	{
		authAdmin.GET("/configuracoes", controllers.BuscarConfiguracoes)
		authAdmin.POST("/configuracoes", controllers.AlterarConfiguracoes)
		authAdmin.POST("/logout", controllers.FazerLogout)
	}
}
