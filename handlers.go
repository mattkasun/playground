package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func displayMainPage(c *gin.Context) {
	date := time.Now()
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func dateHandler(c *gin.Context) {
	where := c.PostForm("action")
	date := time.Now()
	switch where {
	case "back":
		date = data.Today.AddDate(0, 0, -7)
	case "today":
		date = time.Now()
	case "forward":
		date = data.Today.AddDate(0, 0, 7)
	}
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func transactionHandler(c *gin.Context) {
	action := c.PostForm("action")
	date, err := time.Parse("2006-01-02", c.PostForm("date"))
	if err != nil {
		log.Fatal(err)
	}
	if action == "expense" {
		commitTrans(c, true)
	} else {
		commitTrans(c, false)
	}
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func newCategoryHandler(c *gin.Context) {
	date, err := time.Parse("2006-01-02", c.PostForm("date"))
	if err != nil {
		log.Fatal(err)
	}
	action := c.PostForm("action")
	if action == "expense" {
		addCategory(c, true)
		data.init(&date, "Expense")
	} else {
		addCategory(c, false)
		data.init(&date, "Income")
	}
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
	data.init(&date, "Transaction")
	c.HTML(http.StatusOK, "layout", data)
}

//loginHandler
func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", "")
}

func logout(c *gin.Context) {
	c.SetCookie("spend", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func processLogin(c *gin.Context) {
	user := c.PostForm("user")
	pass := c.PostForm("pass")
	if validateUser(user, pass) {
		log.Println("user", user, "password", pass)
		c.SetCookie("spend", "alldjhaeisislsj", 0, "/", "localhost", false, true)
		date := time.Now()
		data.init(&date, "Home")
		c.HTML(http.StatusOK, "layout", data)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication Failed"})
	}
}
