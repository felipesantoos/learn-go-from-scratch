package main

import (
	"fmt"
)

func getDayOfWeek_1(number int) string {
	// Switch/case
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednessday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func getDayOfWeek_2(number int) string {
	// Switch/case/fallthrough
	var day string
	switch {
	case number == 0:
		fallthrough // Jump to the next
	case number == 1:
		day = "Sunday"
	case number == 2:
		day = "Monday"
	case number == 3:
		day = "Tuesday"
	case number == 4:
		day = "Wednesday"
	case number == 5:
		day = "Thursday"
	case number == 6:
		day = "Friday"
	case number == 7:
		day = "Saturday"
	default:
		day = "Invalid number"
	}

	return day
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	day := getDayOfWeek_1(number)
	fmt.Println(day)
	fmt.Scanf("%d", &number)
	day = getDayOfWeek_2(number)
	fmt.Println(day)
}

// OK
