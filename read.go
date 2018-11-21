package main

import (
	"encoding/json"
	"log"
	"os"
)

func readTrans() []Transaction {

	var transactions []Transaction
	f, err := os.Open("trans.data")
	defer f.Close()
	if err != nil {
		log.Fatal("error opening file ", err)
	}

	decoder := json.NewDecoder(f)
	for decoder.More() {
		var transaction Transaction
		err := decoder.Decode(&transaction)
		if err != nil {
			log.Fatal("decoding transaction", err)
		}
		transactions = append(transactions, transaction)
	}
	return transactions
}

func readCat() []Category {
	var categories []Category
	f, err := os.Open("category.data")
	defer f.Close()
	if err != nil {
		log.Fatal("error opening category file ", err)
	}

	decoder := json.NewDecoder(f)
	for decoder.More() {
		var category Category
		err := decoder.Decode(&category)
		if err != nil {
			log.Fatal("decoding categories ", err)
		}
		categories = append(categories, category)
	}
	return categories
}
