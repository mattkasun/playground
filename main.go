package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var data PageData

func main() {
	//set router to default
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//process templates
	router.LoadHTMLGlob("html/*")

	//Initialize routes
	initializeRoutes(router)

	//serve the app
	router.Run()

}
