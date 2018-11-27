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
	Today        time.Time
	Start        time.Time
	End          time.Time
	Income       int
	ExpenseTotal int
	Balance      int
	Expenses     []Expense
	Categories   []Category
	Transactions []Transaction
	CarryOver    int
}

var data PageData
var date = time.Now()

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/tabb.gohtml"))
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found\n"+r.URL.Path, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		date = time.Now()
		data := initTemplateData(&date)
		tmpl.Execute(w, data)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		switch r.FormValue("form") {
		case "expense":
			commitTrans(r, true)
			data = initTemplateData(&date)
			tmpl.Execute(w, data)
		case "income":
			commitTrans(r, false)
			data = initTemplateData(&date)
			tmpl.Execute(w, data)
		case "addIncome":
			addCategory(r, false)
			data = initTemplateData(&date)
			tmpl.Execute(w, data)
		case "addExpense":
			addCategory(r, true)
			data = initTemplateData(&date)
			tmpl.Execute(w, data)
		case "back":
			fmt.Println(date)
			date = date.AddDate(0, 0, -7)
			fmt.Println(date)
			data = initTemplateData(&date)
			tmpl.Execute(w, data)
		case "forward":
			date = date.AddDate(0, 0, 7)
			data = initTemplateData(&date)
			tmpl.Execute(w, data)
		default:
			fmt.Println("not yet implemented")
			tmpl.Execute(w, data)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func initTemplateData(date *time.Time) PageData {
	var data PageData
	data.Today = *date
	transactions := readTrans()
	data.Categories = readCat()
	balance(&data, transactions)
	data.Start, data.End = week(data.Today)
	return data
}
func main() {

	fs := http.FileServer(http.Dir("stylesheet"))
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", fs))
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func commitTrans(r *http.Request, expense bool) {

	//r.ParseForm()
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		log.Fatal(err)
	}
	cat := r.FormValue("Category")
	transaction := Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
	writeOne(transaction)
}

func previous() {
	fmt.Println("func previous")
}
func next() {
	fmt.Println("func next")
}
