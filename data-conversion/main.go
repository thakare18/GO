package main

import "fmt"

func main() {
	var num int = 42
	fmt.Println("The number is", num)
	fmt.Printf("The type of the num is %T", num)

	var data float64 = float64(num)
	fmt.Println("data is", data)
	fmt.Printf("the type of data is %T", data)
}
