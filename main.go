package main

import (
	"github.com/GenaroNajera/benkyou/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", handlers.Home)

	router.Run("localhost:8080")
}
