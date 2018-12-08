package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

var data PageData
var date = time.Now()

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
