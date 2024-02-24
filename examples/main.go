package main

import (
	"fmt"

	"github.com/0xsj/fn-go"
)

// main.go

func main() {
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }
	subtract := func(x, y int) int { return x - y }

	composed, err := fn.Compose(add, multiply, subtract)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := composed.(func(int, int) int)(2, 3)
	fmt.Println("Result:", result)
}
