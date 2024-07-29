package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers"
	"github.com/mquelucci/projeto-loja-virtual/server/middlewares"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("client/templates/**/*")
	r.Static("/assets", "./client/assets")
	r.GET("/", controllers.ExibeHTMLIndex)

	noAuth := r.Group("/admin")
	{
		noAuth.GET("/login", controllers.ExibeHTMLAdminLogin)
		noAuth.POST("/login", controllers.FazerLogin)
	}

	auth := r.Group("/admin").Use(middlewares.Auth())
	{
		auth.GET("/", controllers.ExibeHTMLAdmin)
		auth.POST("/logout", controllers.FazerLogout)
	}

	r.Run()
}
