package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func writeCookie(u, c string, valid time.Time) {
	f, err := os.OpenFile("data/user.data", os.O_RDWR, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error opening file", err)
	}
	decoder := json.NewDecoder(f)
	//get all users
	var users []User
	for decoder.More() {
		var user User
		err = decoder.Decode(&user)
		if err != nil {
			log.Println("decoding error")
			break
		}
		//find user whose cookie is to be updated
		if user.UserName == u {
			user.Cookie = c
			user.ValidTo = valid
		}
		users = append(users, user)
	}
	//clear file
	f.Truncate(0)
	f.Seek(0, 0)
	for i := range users {
		b, err := json.Marshal(users[i])
		if err != nil {
			log.Fatal("error encoding user", err)
		}
		_, err = f.Write(b)
		if err != nil {
			log.Fatal("error writing to user file ", err)
		}
	}
}

func writeAll(transactions []Transaction) {

	f, err := os.OpenFile("data/trans.data", os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error opening file", err)
	}
	for i := range transactions {
		b, err := json.Marshal(transactions[i])
		if err != nil {
			log.Fatal("encoding error: ", err)
		}
		_, err = f.Write(b)
		if err != nil {
			log.Fatal("write err:", err)
		}
	}
}
func writeOne(t Transaction) {
	f, err := os.OpenFile("data/trans.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error creating file", err)
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal("encoding err", err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Fatal("write error", err)
	}

}

func addCategory(c *gin.Context, expense bool) {
	f, err := os.OpenFile("data/category.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error creating file", err)
	}
	name := c.PostForm("category")
	category := Category{ExpenseCat: expense, Name: name}
	b, err := json.Marshal(category)
	if err != nil {
		log.Fatal("encoding error", err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Fatal("write error", err)
	}
}
