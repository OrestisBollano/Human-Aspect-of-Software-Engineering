package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// 1. Create a slice (list) of integers
	numbers := []int{1, 2, 3}

	// 2. Print the whole slice
	fmt.Println("Here is the slice:", numbers)

	// 3. Or, loop through the slice and print each item
	fmt.Println("Here are the items one by one:")
	for _, num := range numbers {
		fmt.Println(num)
	}
}