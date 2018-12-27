package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var data PageData

func main() {
	//set router to default
	router := gin.Default()

	//process templates
	router.LoadHTMLGlob("html/*")

	//Redis Connection
	connection, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	//Initialize routes
	initializeRoutes(router, connection)

	//serve the app
	router.Run()

}
