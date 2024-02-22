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

func TestIsLeftRight(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"Left", "error", true},
		{"Right", 42, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e Either
			switch tt.name {
			case "Left":
				e = LeftConstructor(tt.value)
			case "Right":
				e = RightConstructor(tt.value)
			}

			if tt.expected != e.IsLeft() {
				t.Errorf("Expected IsLeft() to return %v, got %v", tt.expected, e.IsLeft())
			}

			if !tt.expected != e.IsRight() {
				t.Errorf("Expected IsRight() to return %v, got %v", !tt.expected, e.IsRight())
			}
		})
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
