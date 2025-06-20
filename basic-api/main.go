package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/keploy-api-fellowship/basic-api/controllers"
	"github.com/utkarshkrsingh/keploy-api-fellowship/basic-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/watchlist", controllers.GetList)
	router.POST("/watchlist", controllers.InsertAnime)
	router.PATCH("/watchlist/:name", controllers.Update)
	router.DELETE("/watchlist/:name", controllers.Delete)

	router.Run(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
