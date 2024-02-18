package fn

import "testing"

func TestRandom(t *testing.T) {
	randomValue := Random()
	if randomValue < 0 || randomValue >= 1 {
		t.Errorf("Random() returned out-of-range value: %f", randomValue)
	}
}

func TestRandomInt(t *testing.T) {
	low, high := 5, 10
	result := RandomInt(low, high)()
	if result < low || result > high {
		t.Errorf("RandomInt(%d, %d) = %d; want result between %d and %d", low, high, result, low, high)
	}
}

func TestRandomBool(t *testing.T) {
	result := RandomBool()()
	if result != true && result != false {
		t.Errorf("RandomBool() = %v; want true or false", result)
	}
}

func TestRandomElem(t *testing.T) {
	// Test the RandomElem function
	as := []int{1, 2, 3, 4, 5}
	result := RandomElem(as)()
	if result < 1 || result > 5 {
		t.Errorf("RandomElem(%v) = %d; want element in %v", as, result, as)
	}
}
