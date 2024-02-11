package main

import (
	"fmt"

	mf "github.com/0xsj/fn-go/pkg/map"
) // Assuming mf is the alias for fn_map

func main() {
	// Define the numbers slice
	numbers := []int{1, 2, 3, 4, 5}

	// Define the incrementBy function
	incrementBy := func(amount int) func(int) int {
		return func(x int) int {
			return x + amount
		}
	}

	// Apply the transformation using MapInt
	incrementedNumbers := mf.MapInt(numbers, incrementBy(1))
	incrementedNumbers2 := mf.MapInt(numbers, incrementBy(5))

	// Print the original and incremented numbers
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Incremented numbers:", incrementedNumbers)
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Incremented numbers:", incrementedNumbers2)
}
