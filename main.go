package main

import (
	"github.com/gin-gonic/gin"
)

var data PageData

func main() {
	//set router to default
	router := gin.Default()

	//process templates
	router.LoadHTMLGlob("html/*")

	//Initialize routes
	initializeRoutes(router)

	//serve the app
	router.Run()

}
