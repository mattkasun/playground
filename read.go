package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func read() []Transaction {

	var transactions []Transaction
	f, err := os.Open("data.file")
	defer f.Close()
	if err != nil {
		log.Fatal("error opening file ", err)
	}

	decoder := json.NewDecoder(f)
	for decoder.More() {
		var transaction Transaction
		err := decoder.Decode(&transaction)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v:%v:%v:%v:\n", transaction.Date, transaction.Cat, transaction.Amount, transaction.Expense)
		transactions = append(transactions, transaction)
	}
	return transactions
}
