package io

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	addOneIO := func() int {
		return 1
	}

	square := func(x int) int {
		return x * x
	}

	produceInt := func() int {
		return addOneIO()
	}

	// Using Map to apply square to the result of effectfulComputation
	squareAddOneIO := Map(square, produceInt)
	fmt.Printf("Type of squareAddOneIO: %T\n", squareAddOneIO)

	// Execute the composed effectful computation
	result := squareAddOneIO()
	fmt.Printf("Result of squareAddOneIO(): %d\n", result)

	// Adjust the expected value to match the result when addOneIO returns 1
	expected := 1 // Square of 1 is 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
