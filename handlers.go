package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func displayMainPage(c *gin.Context) {
	date := time.Now()
	data.init(&date, "Home")
	fmt.Println(data)
	c.HTML(
		http.StatusOK,
		"layout",
		data,
	)
}
