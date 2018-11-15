package main

import (
  "fmt"
  "time"
)

type Transaction struct {
  ID int
  Date time.Time
  CatID int
  Amount int
  Expense bool
}

type Category struct {
  ID int
  name string
}

func main () {
  expenseCats := []Category {
    Category {
      ID: 1,
      name: "coffee",
    },
    Category {
      ID: 2,
      name: "lunch",
    },
  }

  incomeCats := Category {
    ID: 1,
    name: "salary",
  }
  fmt.Println(expenseCats, incomeCats)
  
  transactions := []Transaction {
    Transaction {
      ID: 1,
      Date: time.Now(),
      CatID: 1,
      Amount: 200,
      Expense: false,
    },
    Transaction {
      ID: 2,
      Date: time.Now(),
      CatID: 1,
      Amount: 3,
      Expense: true,
    },
    Transaction {
      ID: 3,
      Date: time.Now(),
      CatID: 2,
      Amount: 9,
      Expense: true,
    },
  }

  transactions = append (transactions, Transaction{4, time.Now(), 1, 3, true})
  balance := 0
  expense := 0
  income  := 0
  for i := range transactions {
    if transactions[i].Expense {
      balance = balance - transactions[i].Amount
      expense = expense + transactions[i].Amount
    } else {
        balance = balance + transactions[i].Amount
        income = income + transactions[i].Amount
    }
  fmt.Println ("New balances: ", balance, income, expense)
}
  fmt.Println (transactions[1].Date)
  year,month,day := transactions[1].Date.Date()
  fmt.Println (year,month,day)
  fmt.Println (transactions[1].Date.Date())
  fmt.Println (transactions[1].Date.String())
}



