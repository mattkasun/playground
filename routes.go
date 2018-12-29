package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Handle main routed

	router.Static("/stylesheet", "./stylesheet")
	router.StaticFile("favicon.ico", "./resources/favicon.ico")
	router.POST("/auth", processLogin)
	router.GET("/logout", logout)
	router.GET("/login", displayLogin)
	router.POST("newuser", addNewUser)

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
		//if user file doesn't exist, this is the first time
		//so we have to create a user
		_, err := os.Open("data/user.data")
		if err != nil {
			if gin.IsDebugging() {
				log.Println("create new user")
			}
			c.HTML(http.StatusOK, "new", "")
			c.Abort()
			return
		}

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
			log.Println("invalid cookie, redirect to login")
			displayLogin(c)
			c.Abort()
		}
	}
}
