package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("What is Your Name?")
	// var name string
	// fmt.Scan(&name) //scan method is used to take input from user
	// fmt.Println("Hello my name is", name)

	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin) // bufio.NewReader is used to create a new reader that reads from standard input (os.Stdin)
	name, _ := reader.ReadString('\n')  // ReadString method is used to read input from user

	fmt.Println("Hello ji my name is", name)

}
