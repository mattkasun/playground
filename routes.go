package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Handle main route

	router.POST("/login", processLogin)
	router.GET("/logout", logout)
	router.GET("/login", login)

	private := router.Group("/", AuthRequired())
	{
		private.GET("/", displayMainPage)
		private.POST("date", dateHandler)
		private.POST("expense", expenseHandler)
		private.POST("income", incomeHandler)
		private.POST("newExpense", newExpenseHandler)
		private.POST("edit", editHandler)
		private.POST("update", updateHandler)
		private.POST("cancel", cancelHandler)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthRequired")
		session := sessions.Default(c)
		user := session.Get("user")
		log.Println("session.Get('user') returned ", user)
		cookie, err := c.Cookie("session-cookie")
		log.Println("cookie is ", cookie)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
