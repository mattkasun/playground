package main

func balance(t []Transaction, c []Category) PageData {
	var data PageData
	var expenses []Expense
	balance := 0
	expense := 0
	income := 0

	for i := range t {
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
	}
	data.Income = income
	data.ExpenseTotal = expense
	data.Balance = balance
	data.Expenses = expenses
	data.Categories = c
	return data
}
