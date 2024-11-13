package main

import (
	"fmt"
	"time"
)

func compare(a, b int) (z string) {
	if a > b {
		return "a > b"
	} else if a < b {
		return "a < b"
	}
	return
}

func isEven(number int) bool {
	if num := number % 2; num == 0 {
		return true
	}
	return false
}

// switch case
func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	}
	return "Invalid day"
}

// switch case with condition
func showGreeter() string {
	h := time.Now().Unix()
	switch {
	case h < 12:
		return "Good Morning"
	case h < 17:
		return "Good Afternoon"
	default:
		return "Good Evening"
	}
}

func main() {
	t := compare(1, 2)
	fmt.Println(t)
	fmt.Println(isEven(8))
	fmt.Println(dayOfWeek(3))
	fmt.Println(showGreeter())
}
