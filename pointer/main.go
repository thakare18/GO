package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	// var num int
	// num = 2
	num := 2

	// var ptr *int
	// ptr = &num

	// short HAND to write pointer
	ptr := &num

	// without num we can also print the value of num
	fmt.Println("Value of num using pointer", *ptr)

	fmt.Println("Value that stored in num", num)
	fmt.Println("Value pointer value means address of that element", ptr)

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value Contains", value)
}
