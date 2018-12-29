package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Handle main routed

	router.Static("/stylesheet", "./stylesheet")
	router.StaticFile("favicon.ico", "./resources/favicon.ico")
	router.POST("/auth", processLogin)
	router.GET("/logout", logout)
	router.GET("/login", displayLogin)

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
		if err != nil {
			log.Println("unauthorized access, redirect to login")
			displayLogin(c)
			c.Abort()
			return
		}

		if validateCookie(cookie) {
			log.Println("authorized access, continuing ....")
			c.Next()
		} else {
			displayLogin(c)
			c.Abort()
		}
	}
}
