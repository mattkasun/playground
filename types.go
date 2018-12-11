package main

import "time"

// Category -- expense and income types
type Category struct {
	ExpenseCat bool
	Name       string
}

// Transaction -- contains information on single transaction
type Transaction struct {
	Date    time.Time
	Cat     string
	Amount  int
	Expense bool
}

//Expense -- type to expense data
type Expense struct {
	Cat    string
	Amount int
}

//PageData - contains data for html template
type PageData struct {
	Page         string
	Today        time.Time
	Start        time.Time
	End          time.Time
	Income       int
	ExpenseTotal int
	Balance      int
	Expenses     []Expense
	Categories   []Category
	Transactions []Transaction
	Transaction  Transaction
	CarryOver    int
}

//EditData - contains data to edit a transaction
type EditData struct {
	Old        Transaction
	Categories []Category
}