package main

import "fmt"

func main() {
	day := 30

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("Truesday")
	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("This is invalid day")
	}

	month := "august"

	switch month {
	case "january", "february", "march":
		fmt.Println("winter season")
	case "april", "may", "june":
		fmt.Println("summer season")
	default:
		fmt.Println("This is other season")
	}

	//using expression in switch case

	temperature := 30
	switch {
	case temperature < 20:
		fmt.Println("It's cold")
	case temperature >= 23:
		fmt.Println("Its warm")
	}
}
