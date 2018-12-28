package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (data *PageData) init(date *time.Time, page string) {
	data.Page = page
	data.Today = *date
	data.Start, data.End = week(data.Today)
	transactions := readTrans()
	data.Categories = readCat()
	balance(data, transactions)
}

func week(date time.Time) (time.Time, time.Time) {

	var start, end time.Time

	//date := time.Date(2010, 12, 2, 12, 30, 0, 0, time.UTC)
	day := date.Weekday()
	switch day {
	case 0:
		start = date.AddDate(0, 0, -6)
		end = date.AddDate(0, 0, 0)
	case 1:
		start = date.AddDate(0, 0, 0)
		end = date.AddDate(0, 0, 6)
	case 2:
		start = date.AddDate(0, 0, -1)
		end = date.AddDate(0, 0, 5)
	case 3:
		start = date.AddDate(0, 0, -2)
		end = date.AddDate(0, 0, 4)
	case 4:
		start = date.AddDate(0, 0, -3)
		end = date.AddDate(0, 0, 3)
	case 5:
		start = date.AddDate(0, 0, -4)
		end = date.AddDate(0, 0, 2)
	case 6:
		start = date.AddDate(0, 0, -5)
		end = date.AddDate(0, 0, 1)
	}
	return start, end
}

func balance(data *PageData, transactions []Transaction) {
	var expenses []Expense
	var incomes []Expense
	balance := 0
	expense := 0
	income := 0
	carryover := 0
	year, week := data.Today.ISOWeek()
	data.Transactions = nil

	for i := range transactions {
		transDate := transactions[i].Date
		// ignore transactions in the future
		if transDate.After(data.End) {
			continue
		}
		transYear, transWeek := transDate.ISOWeek()
		//handle current time period transactions
		if transYear == year && transWeek == week {
			data.Transactions = append(data.Transactions, transactions[i])
			if transactions[i].Expense {
				if len(expenses) == 0 {
					expenses = append(expenses, Expense{Cat: transactions[i].Cat, Amount: transactions[i].Amount})
				} else {
					foundata := false
					for j := range expenses {
						if expenses[j].Cat == transactions[i].Cat {
							expenses[j].Amount = expenses[j].Amount + transactions[i].Amount
							foundata = true
						}
					}
					if foundata == false {
						expenses = append(expenses, Expense{Cat: transactions[i].Cat, Amount: transactions[i].Amount})
					}
				}
				balance = balance - transactions[i].Amount
				expense = expense + transactions[i].Amount
			} else {
				if len(incomes) == 0 {
					incomes = append(incomes, Expense{Cat: transactions[i].Cat, Amount: transactions[i].Amount})
				} else {
					foundata := false
					for j := range incomes {
						if incomes[j].Cat == transactions[i].Cat {
							incomes[j].Amount = incomes[j].Amount + transactions[i].Amount
							foundata = true
						}
					}
					if foundata == false {
						incomes = append(incomes, Expense{Cat: transactions[i].Cat, Amount: transactions[i].Amount})
					}
				}

				balance = balance + transactions[i].Amount
				income = income + transactions[i].Amount
			}
			//update balance, carryover for transaction before current period.
		} else {
			if transactions[i].Expense {
				balance = balance - transactions[i].Amount
				carryover = carryover - transactions[i].Amount
			} else {
				balance = balance + transactions[i].Amount
				carryover = carryover + transactions[i].Amount
			}
		}
	}
	data.Income = income
	data.ExpenseTotal = expense
	data.Balance = balance
	data.Expenses = expenses
	data.Incomes = incomes
	data.CarryOver = carryover
}

func commitTrans(c *gin.Context, expense bool) {
	amount, _ := strconv.Atoi(c.PostForm("amount"))
	date, err := time.Parse("2006-01-02", c.PostForm("date"))
	if err != nil {
		log.Fatal(err)
	}
	cat := c.PostForm("Category")
	transaction := Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
	writeOne(transaction)
}

func edit(c *gin.Context) Transaction {
	var transaction Transaction
	amount, _ := strconv.Atoi(c.PostForm("Amount"))
	expense, _ := strconv.ParseBool(c.PostForm("Expense"))
	date, err := time.Parse("2006-01-02", c.PostForm("Date"))
	if err != nil {
		log.Fatal(err)
	}
	cat := c.PostForm("Cat")
	transaction = Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
	return transaction
}
