package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

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
	Income       int
	ExpenseTotal int
	Balance      int
	Expenses     []Expense
	Categories   []Category
	Transactions []Transaction
}

var data PageData
var transactions []Transaction
var categories []Category

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/tabb.gohtml"))
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found\n"+r.URL.Path, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		//http.ServeFile(w, r, "html/tab.html")
		transactions = readTrans()
		categories = readCat()
		data = balance(transactions, categories)
		data.Transactions = transactions
		fmt.Println(data)
		tmpl.Execute(w, data)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		switch r.FormValue("form") {
		case "expense":
			transactions = commitTrans(r)
			data = balance(transactions, categories)
			tmpl.Execute(w, data)
		case "cancel":
			http.ServeFile(w, r, "html/main.html")
		case "income":
			http.ServeFile(w, r, "html/income.html")
		case "spending":
			http.ServeFile(w, r, "html/spend.html")
		case "transactions":
			http.ServeFile(w, r, "html/transactions.html")
		case "categories":
			http.ServeFile(w, r, "html/category.html")
		case "previous":
			previous()
			http.ServeFile(w, r, "html/main.html")
		case "next":
			next()
			http.ServeFile(w, r, "html/main.html")
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	fmt.Println(transactions, data)

	fs := http.FileServer(http.Dir("stylesheet"))
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", fs))
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func commitTrans(r *http.Request) []Transaction {

	//r.ParseForm()
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		log.Fatal(err)
	}
	cat := r.FormValue("Category")
	transaction := Transaction{Date: date, Cat: cat, Amount: amount, Expense: true}
	transactions = append(transactions, Transaction{Date: date, Cat: cat, Amount: amount, Expense: true})
	writeOne(transaction)
	return transactions
}

func previous() {
	fmt.Println("func previous")
}
func next() {
	fmt.Println("func next")
}
