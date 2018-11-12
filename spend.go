package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Expense struct {
	Cat    string
	Amount int
}
type PageData struct {
	Income       int
	ExpenseTotal int
	Expenses     []Expense
}

var data = PageData{
	Income:       200,
	ExpenseTotal: 15,
	Expenses: []Expense{
		Expense{
			Cat:    "lunch",
			Amount: 12,
		},
		Expense{
			Cat:    "coffee",
			Amount: 3,
		},
	},
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/tabb.gohtml"))
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found\n"+r.URL.Path, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		//http.ServeFile(w, r, "html/tab.html")
		tmpl.Execute(w, data)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Println("Post from website! r.PostFrom = %v\n", r.PostForm)
		fmt.Println("Request", r)
		switch r.FormValue("form") {
		case "expense":
			CommitTrans(r)
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

		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	fs := http.FileServer(http.Dir("stylesheet"))
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", fs))
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CommitTrans(r *http.Request) {

	fmt.Println("Post from website! r.PostFrom = \n", r.PostForm)
	fmt.Println("current data", data.Income, data.ExpenseTotal, data.Expenses[0].Amount, data.Expenses[1].Amount)
	fmt.Println("commiting transaction")
	r.ParseForm()
	amount, _ := strconv.Atoi(strings.Join(r.Form["amount"], " "))
	fmt.Println("amount: ", amount)

	if strings.Join(r.Form["Category"], " ") == "lunch" {
		data.Expenses[0].Amount += amount
	} else {
		data.Expenses[1].Amount += amount
	}
	data.ExpenseTotal += amount
	fmt.Println ("ExpenseTotal: ", data.ExpenseTotal)
}

func previous() {
	fmt.Println("func previous")
}
func next() {
	fmt.Println("func next")
}
