package main

import "fmt"

func main() {
	var num int
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println("Value that stored in num", num)
	fmt.Println("Value pointer value means address of that element", ptr)
}
