package main

import "fmt"

type Person struct { //using a struct we can group multiple data types  together

	FirstName string
	LastName  string
	Age       int
}

func main() {
	var prince Person
	prince.FirstName = "Prince"
	prince.LastName = "Thakre"
	prince.Age = 22
	fmt.Println("Prince Person", prince)

	// using shorterhand

	person1 := Person{
		FirstName: "ABC",
		LastName:  "xyz",
		Age:       33,
	}
	fmt.Println("Person 1", person1)

	//using new keyword
	var person2 = new(Person)
	person2.FirstName = "Suraj"
	person2.LastName = "Poddar"
	person2.Age = 23

	fmt.Println("Person 2 all details", person2.Age) // dot operator
}
