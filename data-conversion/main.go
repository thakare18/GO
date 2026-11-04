package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("The number is", num)
	fmt.Printf("The type of the num is %T", num)
	// num = num + 1.2
	var data float64 = float64(num)
	data = data + 3.3
	fmt.Println("data is", data)
	fmt.Printf("the type of data is %T", data)
	//int to string
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("The str is", str)
	fmt.Printf("the type of str is %T\n", str)

	//string to int
	number_string := "123"
	number_int, _ := strconv.Atoi(number_string) // _ is used to ignore the second return value of the function, which is an error in this case. We are not interested in the error value, so we use _ to ignore it
	fmt.Println("number_int is", number_int)
	fmt.Printf("type of number_int is %T", number_int)
	// string to float conversion
	number_string = "123.45"
	number_float, _ := strconv.ParseFloat(number_string, 64)
	fmt.Println("number_float is", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}
