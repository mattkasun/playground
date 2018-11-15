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
  //balance := 0
  //expense := 0
  //income  := 0
  
  f, err := os.OpenFile("data.file", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
  defer f.Close()
  if err != nil {
    log.Fatal ("error creating file" , err)
  }

  for i := range transactions {
    b, err := json.Marshal(transactions[i])
    if err != nil {
      log.Fatal("encoding error: ", err)
    }
    n, err := f.Write(b)
    if  err != nil {
      log.Fatal("write err:", err)
    }
    fmt.Println("wrote ", n, " bytes")
  }
  fmt.Println (transactions[1].Date)
  year,month,day := transactions[1].Date.Date()
  fmt.Println (year,month,day)
  fmt.Println (transactions[1].Date.Date())
  fmt.Println (transactions[1].Date.String())
}



