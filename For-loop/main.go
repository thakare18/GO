package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Print the value of i: ", i)
	// }

	// counter := 0

	// for {
	// 	fmt.Println("Infinite loop: ", counter)
	// 	counter++
	// 	// if counter == 3 {
	// 	// 	break
	// 	// }
	// }

	//slice and range keyword

	numbers := []int{1, 2, 3, 4, 5} //using range keyword find the index and value of the slice
	for index, value := range numbers {
		fmt.Printf("Index of numbers is %d, Value of numbers is %d\n", index, value)
	}

	data := "Hello world"
	for index, value := range data {
		fmt.Printf("index of the data is %d, value of the data is %c\n", index, value)
	}

}
