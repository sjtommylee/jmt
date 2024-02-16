package fn_test

import (
	"testing"
)

type MockMagma struct{}

func (m *MockMagma) Concat(x, y interface{}) interface{} {
	return x.(int) + y.(int)
}

func TestReverse(t *testing.T) {
	m := &MockMagma{}
	result := m.Concat(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Concat(3, 4) returned %d, expected %d", result, expected)
	}

}

// func TestFilterFirst() {}

// func TestFilterSecond() {}
