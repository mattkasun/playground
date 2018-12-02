package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func writeAll(transactions []Transaction) {

	f, err := os.OpenFile("trans.data", os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error opening file", err)
	}
	total := 0
	for i := range transactions {
		b, err := json.Marshal(transactions[i])
		if err != nil {
			log.Fatal("encoding error: ", err)
		}
		n, err := f.Write(b)
		if err != nil {
			log.Fatal("write err:", err)
		}
		total = total + n
	}
	fmt.Println("wrote ", total, " bytes")
}
func writeOne(t Transaction) {
	f, err := os.OpenFile("trans.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error creating file", err)
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal("encoding err", err)
	}
	n, err := f.Write(b)
	if err != nil {
		log.Fatal("write error", err)
	}

	fmt.Println("wrote ", n, " bytes")
}

func addCategory(r *http.Request, expense bool) {
	f, err := os.OpenFile("category.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal("error creating file", err)
	}
	name := r.FormValue("category")
	category := Category{ExpenseCat: expense, Name: name}
	b, err := json.Marshal(category)
	if err != nil {
		log.Fatal("encoding error", err)
	}
	n, err := f.Write(b)
	if err != nil {
		log.Fatal("write error", err)
	}

	fmt.Println("wrote ", n, " bytes")

}
