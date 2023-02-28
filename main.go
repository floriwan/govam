package main

import (
	"log"

	"github.com/floriwan/govam/handler"
	"github.com/floriwan/govam/models/controllers"
	"github.com/floriwan/govam/pkg/initializer"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("starting ...")

	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	if err := initializer.ConnectDB(config); err != nil {
		log.Fatal(err)
	}
	initializer.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8081")
}

func initRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// https://github.com/gin-gonic/gin/blob/v1.9.0/docs/doc.md#dont-trust-all-proxies
	router.SetTrustedProxies([]string{"localhost"})

	// hompage
	router.LoadHTMLGlob("templates/**/*.tmpl")
	router.GET("/", handler.Homepage)
	router.GET("/user", handler.User)
	router.GET("/registration", handler.Registration)

	router.Static("/css", "templates/css")

	// api
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/register/form", controllers.RegisterUserForm)
	}

	return router
}

// func initRouter() *gin.Engine {
// 	router := gin.Default()
// 	api := router.Group("/api")
// 	{
// 		api.POST("/token", controllers.GenerateToken)
// 		api.POST("/user/register", controllers.RegisterUser)
// 		secured := api.Group("/secured").Use(middleware.Auth())
// 		{
// 			secured.GET("/ping", controllers.Ping)
// 		}
// 	}
// 	return router
// }
