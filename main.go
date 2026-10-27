package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Learning Go language by Hello World")

	myutil.PrintMessage("Hello from myutil package")

	var name string = "praveen"
	fmt.Println("hello", name)

	var version = 76
	fmt.Println("Version:", version)

	var dimension float64 = 87.40
	fmt.Println("Dimension:", dimension)

	var isActive bool = true //using var we can change value but using const we cannot change value
	fmt.Println("Is Active", isActive)

	//const date

	const pi = 3.14
	fmt.Println("Value of pi:", pi)

	//shorthand declaration
	personName := "praveen" //:= is used for shorthand declaration and it can be used only inside the function
	fmt.Println("Person Name:", personName)

	//capitalization of variable name is used to make it public and can be accessed from other packages
	var Public string = "This is not inmportant data "
	fmt.Println("Public", Public)

	// and start with small letter is private and can be accessed only within the package
	var private string = "This is important data"
	fmt.Println("private", private)

	// those same rule applied on function as well

	// thaat is start with capital letter is public and can be accessed from other packages
	// 	func PulicFunction() {
	// 		fmt.Println("This is public function")
	// 	}
	// // and start with small letter is private and can be accessed only within the package that acess n single pkg
	// 	func privateFunction() {
	// 		fmt.Println("This is private function")
	// }
}
