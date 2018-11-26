package main

import (
	"fmt"
	"time"
)

func week(date time.Time) (time.Time, time.Time) {

	var start, end time.Time

	//date := time.Date(2010, 12, 2, 12, 30, 0, 0, time.UTC)
	day := date.Weekday()
	year, week := date.ISOWeek()
	fmt.Println(date, day, week, year)
	switch day {
	case 0:
		start = date.AddDate(0, 0, -6)
		end = date.AddDate(0, 0, 0)
	case 1:
		start = date.AddDate(0, 0, 0)
		end = date.AddDate(0, 0, 6)
	case 2:
		start = date.AddDate(0, 0, -1)
		end = date.AddDate(0, 0, 5)
	case 3:
		start = date.AddDate(0, 0, -2)
		end = date.AddDate(0, 0, 4)
	case 4:
		start = date.AddDate(0, 0, -3)
		end = date.AddDate(0, 0, 3)
	case 5:
		start = date.AddDate(0, 0, -4)
		end = date.AddDate(0, 0, 2)
	case 6:
		start = date.AddDate(0, 0, -5)
		end = date.AddDate(0, 0, 1)
	default:
		fmt.Println("switch not working")
	}
	return start, end
}
