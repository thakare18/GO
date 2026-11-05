package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current Time", currentTime)
	fmt.Printf("Type of that time %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006,3:04 PM")
	fmt.Println("Formated Time", formatted)

	layout_str := "02/01/2006"
	datastr := "25/11/2026"
	formatted_time, _ := time.Parse(layout_str, datastr)
	fmt.Println("Formatted time", formatted_time)
}
