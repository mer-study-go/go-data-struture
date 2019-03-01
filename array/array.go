package main

import "fmt"

func main() {
	numbers := []int{1, 1, 2, 3, 5, 8, 13}
	// Get one element from the array
	fmt.Println(numbers[3])
	// Get the length of the array
	fmt.Println(len(numbers))
	// Iterate the array
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}
	// Iterate the array without using index
	for _, value := range numbers {
		fmt.Println("Value:", value)
	}
	// Add one item
	fmt.Println(append(numbers, 100))
	// Add multiple items
	fmt.Println(append(numbers, 100, 9, 50, 13))
}
