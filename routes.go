package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, connection redis.Conn) {
	//Redis Connection
	connection, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	router.Use(redisSetup(connection))

	//Handle main route

	router.POST("/auth", processLogin)
	router.GET("/logout", logout)
	router.GET("/login", login)

	private := router.Group("/", authRequired(connection))
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

func redisSetup(connection redis.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", connection)
		c.Next()
	}
}

func authRequired(connection redis.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthRequired")

		cookie, err := c.Cookie("spend")
		if err != nil {
			log.Println("unauthorized access, redirect to login")
			login(c)
			c.Abort()
			return
		}
		userid, err := redis.Int(connection.Do("HGET", "auths", cookie))
		if err == redis.ErrNil {
			login(c)
			c.Abort()
			return
		}
		key := fmt.Sprintf("user:%d", userid)
		valid, err := redis.String(connection.Do("HGET", key, "cookie"))
		if cookie != valid {
			login(c)
			c.Abort()
			return
		}
		log.Println("authorized access, continuing...")
		c.Next()
	}
}
