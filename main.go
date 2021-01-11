package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/product/:id", GetProduct)
	router.Run("localhost:8080")
}
