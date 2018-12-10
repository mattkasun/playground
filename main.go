package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var data PageData

func main() {
	//set router to default
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//process templates
	router.LoadHTMLGlob("html/*")

	//Initialize routes
	initializeRoutes(router)

	//serve the app
	router.Run()

}
