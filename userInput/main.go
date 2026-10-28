package main

import "fmt"

func main() {
	fmt.Println("What is Your Name?")
	var name string
	fmt.Scan(&name) //scan method is used to take input from user
	fmt.Println("Hello my name is", name)
}
