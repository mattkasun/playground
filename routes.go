package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	//Handle main route
	router.GET("/", displayMainPage)
}
