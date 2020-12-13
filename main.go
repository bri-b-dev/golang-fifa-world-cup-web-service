package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"golang-fifa-world-cup-web-service/data"
	_ "golang-fifa-world-cup-web-service/docs"
	"golang-fifa-world-cup-web-service/handlers"
)

// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support

// @license.name MIT
// @license.url https://github.com/bri-b-dev/golang-fifa-world-cup-web-service/master/LICENSE

// @BasePath /
func main() {
	data.PrintUsage()

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", handlers.RootHandler)
	r.GET("/winners/get", handlers.ListWinners)
	r.POST("/winners/post", handlers.AddNewWinner)
	r.Any("/winners", handlers.WinnersHandler)

	r.Run(":8000")
}
