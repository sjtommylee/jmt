package fn

import "testing"

func TestEither(t *testing.T) {
	// Create a Left instance
	left := LeftConstructor("error")
	if !left.IsLeft() {
		t.Errorf("Expected left to be Left, got Right")
	}

	// Create a Right instance
	right := RightConstructor(42)
	if !right.IsRight() {
		t.Errorf("Expected right to be Right, got Left")
	}

	// Check the values inside Left and Right instances
	if value, ok := left.Left.(string); !ok || value != "error" {
		t.Errorf("Expected left value to be 'error', got %v", left.Left)
	}

	if value, ok := right.Right.(int); !ok || value != 42 {
		t.Errorf("Expected right value to be 42, got %v", right.Right)
	}
}
