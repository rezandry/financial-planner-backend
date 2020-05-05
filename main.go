package main

import (
	"financial-planner/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	router := gin.Default()

	router.POST("/register", service.Register)

	router.Run(":" + port)
}
