package main

import (
  "fmt"
  "time"
  "os"
  "log"
  "encoding/json"
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
  
  var transactions []Transaction

  transactions = append (transactions, Transaction{4, time.Now(), 1, 3, true})
  //balance := 0
  //expense := 0
  //income  := 0
  
  f, err := os.Open("data.file")
  defer f.Close()
  if err != nil {
    log.Fatal ("error opening file " , err)
  }

  dec := json.NewDecoder(f)
  for dec.More() {
    var transaction Transaction
    err := dec.Decode(&transaction)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%v:%v:%v:%v:%v\n", transaction.ID, transaction.Date, transaction.CatID, transaction.Amount, transaction.Expense)
    transactions = append (transactions, transaction)
  }
  fmt.Println (transactions[1].Date)
  year,month,day := transactions[1].Date.Date()
  fmt.Println (year,month,day)
  fmt.Println (transactions[1].Date.Date())
  fmt.Println (transactions[1].Date.String())
}



