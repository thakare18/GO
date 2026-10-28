package main

import "fmt"

func main() {
	fmt.Println("Am learing Array in Go lang")

	var name [5]string //5 is the size of the array and string is the type of the array
	name[0] = "Anand"
	name[2] = "Pratham"

	fmt.Println("My name is ", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is ", numbers)

	//to find the length of the array we can use len() function

	fmt.Println("Length of the array is", len(numbers))

	//Accessing the array/
	fmt.Println("Accessing the element at 2 index is ", numbers[2])
}
