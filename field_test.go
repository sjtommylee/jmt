package fn

import (
	"testing"
)

func TestGCD(t *testing.T) {
	// Define a field number
	field := FieldNumber{}

	// Test cases for GCD
	tests := []struct {
		a, b, expected int
	}{
		{8, 12, 4},     // GCD of 8 and 12 is 4
		{17, 23, 1},    // GCD of 17 and 23 is 1
		{15, 25, 5},    // GCD of 15 and 25 is 5
		{105, 140, 35}, // GCD of 105 and 140 is 35
	}

	for _, test := range tests {
		result := gcd(field, field)(test.a, test.b).(int)
		if result != test.expected {
			t.Errorf("GCD(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestLCM(t *testing.T) {
	// Define a field number
	field := FieldNumber{}

	// Test cases for LCM
	tests := []struct {
		a, b, expected int
	}{
		{8, 12, 24},     // LCM of 8 and 12 is 24
		{17, 23, 391},   // LCM of 17 and 23 is 391
		{15, 25, 75},    // LCM of 15 and 25 is 75
		{105, 140, 420}, // LCM of 105 and 140 is 420
	}

	for _, test := range tests {
		result := lcm(field, field)(test.a, test.b).(int)
		if result != test.expected {
			t.Errorf("LCM(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
