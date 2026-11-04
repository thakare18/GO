package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Apple,orange,banana,Mongo"
	parts := strings.Split(data, ",") // what basis ","
	fmt.Println(parts)
	// to count occurance in string
	str := "one two one two"
	count := strings.Count(str, "two")
	fmt.Println("Trimmed parts", count)

	// to remove space in string
	str = "                      hello, everyone    "
	space := strings.TrimSpace(str)
	fmt.Println("trimmed part is", space)

	// to join to strings
	str1 := "Kartik"
	str2 := "Thakare"
	result := strings.Join([]string{str1, "v", str2}, " ")
	fmt.Println("Combined result", result)

}
