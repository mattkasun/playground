package main

import (
	"fmt"
	"time"
)

func main() {
	//funcMap := template.FuncMap{
	//	"after": time.After,
	//	"now":   time.Now,
	//}
	type Transaction struct {
		Date   time.Time
		Amount int
	}

	type PageData struct {
		Today        time.Time
		Transactions []Transaction
	}
	var start, end time.Time

	date := time.Date(2010, 12, 2, 12, 30, 0, 0, time.UTC)
	day := date.Weekday()
	year, week := date.ISOWeek()
	fmt.Println(date, day, week, year)
	switch day {
	case 0:
		start = date.AddDate(0, 0, -6)
		end = date.AddDate(0, 0, 0)
		fmt.Println("Sunday")

	case 1:
		start = date.AddDate(0, 0, 0)
		end = date.AddDate(0, 0, 6)
		fmt.Println("monday")
	case 2:
		start = date.AddDate(0, 0, -1)
		end = date.AddDate(0, 0, 5)
		fmt.Println("tues")
	case 3:
		start = date.AddDate(0, 0, -2)
		end = date.AddDate(0, 0, 4)
		fmt.Println("wed")
	case 4:
		start = date.AddDate(0, 0, -3)
		end = date.AddDate(0, 0, 3)
		fmt.Println("thurs")
	case 5:
		start = date.AddDate(0, 0, -4)
		end = date.AddDate(0, 0, 2)
		fmt.Println("friday")
	case 6:
		start = date.AddDate(0, 0, -5)
		end = date.AddDate(0, 0, 1)
		fmt.Println("sat")
	default:
		fmt.Println("switch not working")
	}
	year, week = start.ISOWeek()
	fmt.Println(start, start.Weekday(), week)
	year, week = end.ISOWeek()
	fmt.Println(end, end.Weekday(), week)
}
