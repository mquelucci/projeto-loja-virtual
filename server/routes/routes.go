package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers"
	"github.com/mquelucci/projeto-loja-virtual/server/middlewares"
)

func HandleRequests() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	store := memstore.NewStore([]byte("apikey"))
	r.Use(sessions.Sessions("mysession", store))

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
		auth.GET("/produtos", controllers.ExibeHTMLAdminProdutos)
		auth.GET("/produtos/new", controllers.ExibeHTMLAdminCadastrarProduto)
		auth.POST("/produtos/new", controllers.CriarProduto)
		auth.POST("/produtos/edit", controllers.EditarProduto)
		auth.GET("/produtos/delete", controllers.DeletarProduto)
		auth.DELETE("/produtos/delete", controllers.DeletarProduto)
		auth.POST("/logout", controllers.FazerLogout)
	}

	r.NoRoute(controllers.ExibeHTML404)
	r.Run()
}
