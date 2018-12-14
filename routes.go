package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Handle main route

	router.POST("/auth", processLogin)
	router.GET("/logout", logout)
	router.GET("/login", login)

	private := router.Group("/", authRequired())
	{
		private.GET("/", displayMainPage)
		private.POST("date", dateHandler)
		private.POST("expense", transactionHandler)
		private.POST("income", transactionHandler)
		private.POST("category", newCategoryHandler)
		private.POST("edit", editHandler)
		private.POST("update", updateHandler)
	}
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthRequired")
		cookie, err := c.Cookie("spend")
		if err != nil || cookie != "alldjhaeisislsj" {
			log.Println("unauthorized access, redirect to login")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			log.Println("authorized access, continuing...")
			c.Next()
		}
	}
}
