package fn

import "testing"

func TestLeftConstructor(t *testing.T) {
	value := "error"
	left := LeftConstructor(value)
	if left.Tag != "Left" {
		t.Errorf("Expected Tag to be 'Left', got %s", left.Tag)
	}
	if left.Value != value {
		t.Errorf("Expected Value to be %s, got %s", value, left.Value)
	}
}

func TestRightConstructor(t *testing.T) {
	value := 42
	right := RightConstructor(value)
	if right.Tag != "Right" {
		t.Errorf("Expected Tag to be 'Right', got %s", right.Tag)
	}
	if right.Value != value {
		t.Errorf("Expected Value to be %d, got %d", value, right.Value)
	}
}

func TestIsLeft(t *testing.T) {
	left := LeftConstructor("error")
	if !left.IsLeft() {
		t.Error("Expected IsLeft to return true")
	}

	right := RightConstructor(42)
	if right.IsLeft() {
		t.Error("Expected IsLeft to return false")
	}
}

func TestIsRight(t *testing.T) {
	left := LeftConstructor("error")
	if left.IsRight() {
		t.Error("Expected IsRight to return false")
	}

	right := RightConstructor(42)
	if !right.IsRight() {
		t.Error("Expected IsRight to return true")
	}
}

func TestLeft(t *testing.T) {
	value := "error"
	left := LeftConstructor(value)
	if left.Left() != value {
		t.Errorf("Expected Left() to return %s, got %s", value, left.Left())
	}
}

func TestRight(t *testing.T) {
	value := 42
	right := RightConstructor(value)
	if right.Right() != value {
		t.Errorf("Expected Right() to return %d, got %d", value, right.Right())
	}
}
