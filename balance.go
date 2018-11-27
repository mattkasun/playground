package main

func balance(data *PageData, transactions []Transaction) {
	var expenses []Expense
	balance := 0
	expense := 0
	income := 0
	carryover := 0
	year, week := data.Today.ISOWeek()
	data.Transactions = nil

	for i := range transactions {
		transDate := transactions[i].Date
		transYear, transWeek := transDate.ISOWeek()
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
				balance = balance + transactions[i].Amount
				income = income + transactions[i].Amount
			}
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
	data.CarryOver = carryover
}
