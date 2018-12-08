package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	//Handle main route
	router.GET("/", displayMainPage)
	router.POST("/back", backHandler)
	router.POST("today", todayHandler)
	router.POST("forward", forwardHandler)
}
