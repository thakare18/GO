package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current Time", currentTime)
	fmt.Printf("Type of that time %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006,15:04:05")
	fmt.Println("Formated Time", formatted)
}
