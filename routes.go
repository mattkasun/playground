package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
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
		state := session.Get("state")
		log.Println(state, session)
		user := session.Get("count")
		log.Println("session.Get('count') returned ", user)
		if user == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		} else {
			c.Next()
		}
	}
}
