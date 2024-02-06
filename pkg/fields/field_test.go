package field

import (
	"testing"
)

func TestFieldNumber(t *testing.T) {
	// instance of fieldnumber
	fieldNum := FieldNumber{}

	result := fieldNum.Add(3, 4)
	if result != 7 {
		t.Errorf("Additiona failed, expected %d, got %d", 7, result)
	}

	result = fieldNum.Sub(5, 4)
	if result != 1 {
		t.Errorf("Subtraction failed, expected %d, got %d", 1, result)
	}

	// Test multiplication
	result = fieldNum.Mul(2, 3)
	if result != 6 {
		t.Errorf("Multiplication failed. Expected %d, got %d", 6, result)
	}

	// Test division
	result = fieldNum.Div(8, 2)
	if result != 4 {
		t.Errorf("Division failed. Expected %d, got %d", 4, result)
	}

	// Test degree
	degree := fieldNum.Degree(5)
	if degree != 1 {
		t.Errorf("Degree calculation failed. Expected %d, got %d", 1, degree)
	}

	// Test modulo
	result = fieldNum.Mod(10, 3)
	if result != 1 {
		t.Errorf("Modulo operation failed. Expected %d, got %d", 1, result)
	}
}

func TestGCDAndLCM(t *testing.T) {
	// Create an instance of FieldNumber
	fieldNum := FieldNumber{}

	// Test gcd
	gcd := gcd(fieldNum, fieldNum)
	result := gcd(12, 8)
	if result != 4 {
		t.Errorf("GCD calculation failed. Expected %d, got %d", 4, result)
	}

	// Test lcm
	lcm := lcm(fieldNum, fieldNum)
	result = lcm(12, 8)
	if result != 24 {
		t.Errorf("LCM calculation failed. Expected %d, got %d", 24, result)
	}
}

func TestInterfaces(t *testing.T) {
	// Check if FieldNumber implements Eq interface
	var _ Eq = FieldNumber{}

	// Check if FieldNumber implements Ring interface
	var _ Ring = FieldNumber{}

	// Check if FieldNumber implements Field interface
	var _ Field = FieldNumber{}
}
