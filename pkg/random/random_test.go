package random

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	result := Random()
	fmt.Printf("Random Float64: %f\n", result)
}

func TestRandomInt(t *testing.T) {
	low, high := 1, 10
	result := RandomInt(low, high)()
	// TODO: thnk about this one
	fmt.Printf("Random Int %d is between %d and %d\n", result, low, high)

	// check if the result is really in the range between 1 and 10
	if result < low || result > high {
		t.Errorf("Result %d is not within the expected range [%d, %d]", result, low, high)
	}
}

func TestRandomElem(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	randomElement := RandomElem(elements)
	result := randomElement()

	found := false

	for _, v := range elements {
		if v == result {

			found = true
			break
		}
	}

	if !found {
		t.Errorf("result %d is not one of the elements in the input slice of %v", result, elements)
	}

}
