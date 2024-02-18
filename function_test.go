package fn

import "testing"

func add(a int) PipeFunc {
	return func(b interface{}) interface{} {
		return a + b.(int)
	}
}

func multiply(a int) PipeFunc {
	return func(b interface{}) interface{} {
		return a * b.(int)
	}
}

func TestPipe(t *testing.T) {
	result := Pipe(1, add(2), multiply(3))

	expected := 9 // (1 + 2) * 3 = 9
	if result != expected {
		t.Errorf("Pipe(1, add(2), multiply(3)) = %d; want %d", result, expected)
	}
}
