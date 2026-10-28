package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(a, b int) int {
	return a + b
	//result = a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return

}

func main() {
	fmt.Println("We are learing func in go lang")
	simpleFunction() // calling simpleFunction in main function becouse main function is executable function
	fmt.Println("Adding the two numbers is ", add(200, 9))
	data := multiply(1, 20)
	fmt.Println("Result of multiplication is ", data)

}
