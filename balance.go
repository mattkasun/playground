package main

import (
	"time"
)

func balance(data *PageData) {
	var expenses []Expense
	balance := 0
	expense := 0
	income := 0
	carryover := 0
	today := time.Now()
	year, week := today.ISOWeek()

	for i := range data.Transactions {
		transDate := data.Transactions[i].Date
		transYear, transWeek := transDate.ISOWeek()
		if transYear == year && transWeek == week {
			if data.Transactions[i].Expense {
				if len(expenses) == 0 {
					expenses = append(expenses, Expense{Cat: data.Transactions[i].Cat, Amount: data.Transactions[i].Amount})
				} else {
					foundata := false
					for j := range expenses {
						if expenses[j].Cat == data.Transactions[i].Cat {
							expenses[j].Amount = expenses[j].Amount + data.Transactions[i].Amount
							foundata = true
						}
					}
					if foundata == false {
						expenses = append(expenses, Expense{Cat: data.Transactions[i].Cat, Amount: data.Transactions[i].Amount})
					}
				}
				balance = balance - data.Transactions[i].Amount
				expense = expense + data.Transactions[i].Amount
			} else {
				balance = balance + data.Transactions[i].Amount
				income = income + data.Transactions[i].Amount
			}
		} else {
			if data.Transactions[i].Expense {
				balance = balance - data.Transactions[i].Amount
				carryover = carryover - data.Transactions[i].Amount
			} else {
				balance = balance + data.Transactions[i].Amount
				carryover = carryover + data.Transactions[i].Amount
			}
		}
	}
	data.Income = income
	data.ExpenseTotal = expense
	data.Balance = balance
	data.Expenses = expenses
	data.CarryOver = carryover
}
