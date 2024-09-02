package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers"
	docs "github.com/mquelucci/projeto-loja-virtual/server/docs"
	"github.com/mquelucci/projeto-loja-virtual/server/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.MaxMultipartMemory = 8 << 20

	store := cookie.NewStore([]byte("lojavirtual"))
	r.Use(sessions.Sessions("lojavirtual", store))

	r.LoadHTMLGlob("../client/templates/**/*")
	r.Static("/assets", "../client/assets")
	r.GET("/", controllers.ExibeHTMLIndex)

	noAuth := r.Group("/admin")
	{
		noAuth.GET("/login", controllers.ExibeHTMLAdminLogin)
		noAuth.POST("/autorizar", controllers.Autenticar)
	}

	auth := r.Group("/admin").Use(middlewares.Auth())
	{
		auth.POST("/configuracoes", controllers.AlterarConfiguracoes)
		auth.POST("/produtos/criar", controllers.CriarProduto)
		auth.POST("/produtos/editar", controllers.EditarProduto)
		auth.POST("/logout", controllers.FazerLogout)
		auth.DELETE("/produtos/removeImagem", controllers.RemoverImagemProduto)
		auth.DELETE("/produtos/delete", controllers.DeletarProduto)

		auth.GET("/produtos/criar", controllers.ExibeHTMLAdminCadastrarProduto)
		auth.GET("/produtos/editar", controllers.ExibeHTMLAdminEditarProduto)
		auth.GET("/produtos", controllers.ExibeHTMLAdminProdutos)

		auth.GET("/", controllers.ExibeHTMLAdmin)

	}

	r.NoRoute(controllers.ExibeHTML404)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
