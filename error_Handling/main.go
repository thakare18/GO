package main

import "fmt"

func main() {
	name, _ := "Prathamesh", 25 // _ is used to ignore the second return value of the function, which is an error in this case. We are not interested in the error value, so we use _ to ignore it.

	fmt.Println(name)
}
