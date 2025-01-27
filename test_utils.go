package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func startProvider(port int) {
	router := gin.Default()
	router.GET("/product/:id", GetProduct)
	router.Run(fmt.Sprintf(":%d", port))
}
