package main

import (
	"fmt"

	"github.com/0xsj/fn-go/pkg/either"
)

func main() {
	leftValue := either.LeftConstructor("Error")
	rightValue := either.RightConstructor(42)

	printEither(leftValue)
	printEither(rightValue)

}

func printEither(eitherValue either.Either[string, int]) {

	switch v := eitherValue.(type) {
	case *either.Left[string]:
		fmt.Printf("Left: %v\n", v.Left)
	case *either.Right[int]:
		fmt.Printf("Right: %v\n", v.Right)
	default:
		fmt.Println("Unknown variant")
	}
}
