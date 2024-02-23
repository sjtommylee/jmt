package main

import (
	"fmt"

	"github.com/0xsj/fn-go"
)

func main() {
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }
	subtract := func(x, y int) int { return x - y }

	composed := fn.Compose(add, multiply, subtract)
	result := composed.(func(int, int) int)(2, 3)
	fmt.Println("Result:", result)
}
