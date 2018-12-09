package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func displayMainPage(c *gin.Context) {
	date := time.Now()
	data.init(&date, "Home")
	fmt.Println(data)
	c.HTML(http.StatusOK, "layout", data)
}

func backHandler(c *gin.Context) {
	date := data.Today.AddDate(0, 0, -7)
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func todayHandler(c *gin.Context) {
	date := time.Now()
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func forwardHandler(c *gin.Context) {
	date := data.Today.AddDate(0, 0, 7)
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func expenseHandler(c *gin.Context) {
	commitTrans(c, true)
	data.init(&data.Today, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func incomeHandler(c *gin.Context) {
	commitTrans(c, false)
	data.init(&data.Today, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func newExpenseHandler(c *gin.Context) {
	addCategory(c, true)
	data.init(&data.Today, "Expense")
	c.HTML(http.StatusOK, "layout", data)
}

func editHandler(c *gin.Context) {
	var data EditData
	data.Old = edit(c)
	data.Categories = readCat()
	c.HTML(http.StatusOK, "edit", data)
}

func cancelHandler(c *gin.Context) {
	data.init(&data.Today, "Transaction")
	c.HTML(http.StatusOK, "layout", data)
}

func updateHandler(c *gin.Context) {
	date, err := time.Parse("2006-01-02", c.PostForm("OldDate"))
	if err != nil {
		log.Fatal(err)
	}
	amount, _ := strconv.Atoi(c.PostForm("OldAmount"))
	cat := c.PostForm("OldCat")
	expense, _ := strconv.ParseBool(c.PostForm("OldExpense"))
	old := Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
	date, err = time.Parse("2006-01-02", c.PostForm("date"))
	if err != nil {
		log.Fatal(err)
	}
	amount, _ = strconv.Atoi(c.PostForm("Amount"))
	cat = c.PostForm("Category")
	new := Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
	updateTrans(old, new)
	data.init(&data.Today, "Transaction")
	c.HTML(http.StatusOK, "layout", data)
}
