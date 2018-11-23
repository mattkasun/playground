package main

import (
	"fmt"
	"time"
)

func balance(t []Transaction, c []Category) PageData {
	var data PageData
	var expenses []Expense
	balance := 0
	expense := 0
	income := 0
	carryover := 0
	today := time.Now()
	year, week := today.ISOWeek()

	fmt.Println(year, week)

	for i := range t {
		transDate := t[i].Date
		transYear, transWeek := transDate.ISOWeek()
		fmt.Println(transYear, transWeek)
		if transYear == year && transWeek == week {
			fmt.Println("use this transaction")
			if t[i].Expense {
				if len(expenses) == 0 {
					expenses = append(expenses, Expense{Cat: t[i].Cat, Amount: t[i].Amount})
				} else {
					found := false
					for j := range expenses {
						if expenses[j].Cat == t[i].Cat {
							expenses[j].Amount = expenses[j].Amount + t[i].Amount
							found = true
						}
					}
					if found == false {
						expenses = append(expenses, Expense{Cat: t[i].Cat, Amount: t[i].Amount})
					}
				}
				balance = balance - t[i].Amount
				expense = expense + t[i].Amount
			} else {
				balance = balance + t[i].Amount
				income = income + t[i].Amount
			}
		} else {
			if t[i].Expense {
				balance = balance - t[i].Amount
				carryover = carryover - t[i].Amount
			} else {
				balance = balance + t[i].Amount
				carryover = carryover + t[i].Amount
			}
		}
		fmt.Println("Income", income, "expense", expense, "balance", balance, "carryover", carryover)
	}
	data.Income = income
	data.ExpenseTotal = expense
	data.Balance = balance
	data.Expenses = expenses
	data.Categories = c
	data.CarryOver = carryover
	return data
}
