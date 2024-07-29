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
	r.GET("/admin", middlewares.Auth(), controllers.ExibeHTMLAdmin)

	noAuth := r.Group("/admin").Use(middlewares.Auth())
	{
		noAuth.GET("/login", controllers.ExibeHTMLAdminLogin)
	}

	auth := r.Group("/admin")
	{
		auth.GET("/", controllers.ExibeHTMLAdmin)
	}

	r.Run()
}
