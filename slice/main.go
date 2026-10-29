package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 2, 3, 4, 5, 6, 67, 7)
	// fmt.Println("The numbers are:", numbers)
	// fmt.Printf("Numbers have this data type: %T\n", numbers)
	// fmt.Println("Length of the numbers", len(numbers))

	// numbers := make([]int, 3, 5) // using make function to create a slice with length 3 and capacity 5
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)
	numbers = append(numbers, 4)

	fmt.Println("Slice", numbers)
	fmt.Println("Length", len(numbers))
	fmt.Println("Capacity", cap(numbers))
}
