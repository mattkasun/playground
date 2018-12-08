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

func backHandler(c *gin.Context) {
	date = date.AddDate(0, 0, -7)
	data.init(&date, "Home")
	c.HTML(
		http.StatusOK,
		"layout",
		data,
	)
}

func todayHandler(c *gin.Context) {
	date := time.Now()
	data.init(&date, "Home")
	c.HTML(
		http.StatusOK,
		"layout",
		data,
	)
}

func forwardHandler(c *gin.Context) {
	date = date.AddDate(0, 0, 7)
	data.init(&date, "Home")
	c.HTML(
		http.StatusOK,
		"layout",
		data,
	)
}
