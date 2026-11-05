package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the Program")
	data := add(20, 20)
	defer fmt.Println("The sum of that two numbers is ", data)
	defer fmt.Println("Middle of the program") // defer is used to delay the execution of a function until the surrounding function returns. In this case, the deferred function will be executed after the main function has completed its execution
	fmt.Println("End of the Program")
}
