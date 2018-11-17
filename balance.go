package main

func balance(t []Transaction) PageData {
	var data PageData
	var categories []Expense
	balance := 0
	expense := 0
	income := 0
	for i := range t {
		if t[i].Expense {
			if len(categories) == 0 {
				categories = append(categories, Expense{Cat: t[i].Cat, Amount: t[i].Amount})
			} else {
				found := false
				for j := range categories {
					if categories[j].Cat == t[i].Cat {
						categories[j].Amount = categories[j].Amount + t[i].Amount
						found = true
					}
				}
				if found == false {
					categories = append(categories, Expense{Cat: t[i].Cat, Amount: t[i].Amount})
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
	data.Expenses = categories

	return data
}
