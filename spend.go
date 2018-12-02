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

var data PageData
var date = time.Now()

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"html/layout.gohtml", "html/buttonbar.gohtml", "html/mainpage.gohtml"))
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found\n"+r.URL.Path, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		date = time.Now()
		initTemplateData(&date, &data)
		data.Page = "Home"
		tmpl.ExecuteTemplate(w, "layout", data)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		switch r.FormValue("form") {
		case "expense":
			commitTrans(r, true)
			initTemplateData(&date, &data)
			data.Page = "Expense"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "income":
			commitTrans(r, false)
			initTemplateData(&date, &data)
			data.Page = "Income"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "addIncome":
			addCategory(r, false)
			initTemplateData(&date, &data)
			data.Page = "Income"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "addExpense":
			addCategory(r, true)
			initTemplateData(&date, &data)
			data.Page = "Expense"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "back":
			date = date.AddDate(0, 0, -7)
			initTemplateData(&date, &data)
			data.Page = "Home"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "forward":
			date = date.AddDate(0, 0, 7)
			initTemplateData(&date, &data)
			data.Page = "Home"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "today":
			date = time.Now()
			initTemplateData(&date, &data)
			data.Page = "Home"
			tmpl.ExecuteTemplate(w, "layout", data)
		case "edit":
			var data EditData
			fmt.Println(r.Form, r.FormValue("form"), r.FormValue("name"), r.FormValue("transaction"))
			data.Old = edit(r)
			data.Categories = readCat()

			tmpl = template.Must(template.ParseFiles("html/edit.gohtml"))
			fmt.Println(data)
			tmpl.ExecuteTemplate(w, "layout", data)
		case "update":
			fmt.Println("Update: ", r.Form)
			//read old values
			date, err := time.Parse("2006-01-02", r.FormValue("OldDate"))
			if err != nil {
				log.Fatal(err)
			}
			amount, _ := strconv.Atoi(r.FormValue("OldAmount"))
			cat := r.FormValue("OldCat")
			expense, _ := strconv.ParseBool(r.FormValue("OldExpense"))
			old := Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
			date, err = time.Parse("2006-01-02", r.FormValue("date"))
			if err != nil {
				log.Fatal(err)
			}
			amount, _ = strconv.Atoi(r.FormValue("Amount"))
			cat = r.FormValue("Category")
			new := Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
			updateTrans(old, new)
		default:
			fmt.Println(w, "not yet implemented", r.FormValue("form"))
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func initTemplateData(date *time.Time, data *PageData) {
	data.Today = *date
	data.Start, data.End = week(data.Today)
	transactions := readTrans()
	data.Categories = readCat()
	balance(data, transactions)
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

func edit(r *http.Request) Transaction {
	var transaction Transaction
	amount, _ := strconv.Atoi(r.FormValue("Amount"))
	expense, _ := strconv.ParseBool(r.FormValue("Expense"))
	date, err := time.Parse("2006-01-02", r.FormValue("Date"))
	if err != nil {
		log.Fatal(err)
	}
	cat := r.FormValue("Cat")
	transaction = Transaction{Date: date, Cat: cat, Amount: amount, Expense: expense}
	return transaction
}
func previous() {
	fmt.Println("func previous")
}
func next() {
	fmt.Println("func next")
}
