package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mycandys-order-analytics/docs"
	"mycandys-order-analytics/internal/database"
	"mycandys-order-analytics/internal/handlers"
	"mycandys-order-analytics/internal/utils"
)

func main() {
	_ = godotenv.Load(".env")

	mongoUrl, _ := utils.GetEnvVar("MONGO_URI")
	println(mongoUrl)

	conn := database.Connect(mongoUrl)
	defer database.Disconnect(conn, context.Background())

	handler := handlers.NewHandler()

	app := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	docs.SwaggerInfo.Title = "MyCandy's Orders Analytics API"
	docs.SwaggerInfo.Description = "This is MyCandy's Orders Analytics  API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.Use(cors.New(config))
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.GET("/health", handler.HealthCheck)
	app.GET("/analytics/last", handler.GetLastCalledEndpoint)
	app.POST("/analytics", handler.AddCallEndpoint)
	app.GET("/analytics/most", handler.GetMostCalledEndpoint)
	app.GET("/analytics/each", handler.GetNumberOfCallForEachEndpoint)

	app.Run(":8000")
}
