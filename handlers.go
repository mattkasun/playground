package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
)

func displayMainPage(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("displayMainPage")
	}
	date := time.Now()
	data.init(&date, "Home")
	c.HTML(http.StatusOK, "layout", data)
}

func dateHandler(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("dateHandler")
	}
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
	if gin.IsDebugging() {
		log.Println("transactionHandler")
	}
	action := c.PostForm("action")
	date, err := time.Parse("2006-01-02", c.PostForm("date"))
	if err != nil {
		log.Fatal("transaction handler", err)
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
	if gin.IsDebugging() {
		log.Println("newCategoryHandler")
	}
	date := time.Now()
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
	if gin.IsDebugging() {
		log.Println("editHandler")
	}
	var data EditData
	data.Old = edit(c)
	data.Categories = readCat()
	c.HTML(http.StatusOK, "edit", data)
}

func cancelHandler(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("cancleHandler")
	}
	data.init(&data.Today, "Transaction")
	c.HTML(http.StatusOK, "layout", data)
}

func updateHandler(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("updateHandler")
	}
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
func displayLogin(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("displayLogin")
	}
	c.HTML(http.StatusOK, "login", "")
}

func logout(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("logout")
	}
	c.SetCookie("spend", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func addNewUser(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("addNewUser")
	}
	var user User
	username := c.PostForm("user")
	password := c.PostForm("pass")
	f, err := os.Create("data/user.data")
	if err != nil {
		log.Panic("unable to create user file ", err)
	}
	user.UserName = username
	user.Password = password
	user.ID = 1
	user.Cookie = ""
	user.ValidTo = time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
	b, err := json.Marshal(user)
	if err != nil {
		log.Fatal("error encoding new user ", err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Fatal("error writing user file ", err)
	}
	c.HTML(http.StatusOK, "login", "")
}

func processLogin(c *gin.Context) {
	if gin.IsDebugging() {
		log.Println("processLogin")
	}
	username := c.PostForm("user")
	password := c.PostForm("pass")
	valid, user := validateUser(username, password)
	if valid {
		log.Println("user ", username, " logged in")
		//is saved cookie still valid
		if user.ValidTo.After(time.Now()) {
			c.SetCookie("spend", user.Cookie, 604800, "/", "", false, true)
			//set new cookie
		} else {
			s := uniuri.New()
			c.SetCookie("spend", s, 604800, "/", "", false, true)
			writeCookie(username, s, time.Now().Add(time.Hour*7*24))
		}
		date := time.Now()
		data.init(&date, "Home")
		c.HTML(http.StatusOK, "layout", data)
	} else {
		log.Println("invalid login")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Login"})
	}
}
