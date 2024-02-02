package either

import (
	"testing"
)

func TestLeft(t *testing.T) {
	input := 10
	output := LeftConstructor(input)

	t.Logf("Output: %+v", output)

	if output.Left != input {
		t.Errorf("Expected Left to be %v, got %v", input, output.Left)
	}
}

func TestRight(t *testing.T) {
	input := 42
	output := RightConstructor(input)

	t.Logf("Output: %+v", output)

	if output.Tag != "Right" {
		t.Errorf("Expected Tag to be 'Right', got %s", output.Tag)
	}
}

func TestIsLeft(t *testing.T) {
	input := 42
	left := LeftConstructor[int](input)

	if output := left.IsLeft(); output == false {
		t.Errorf("Received %t - expected true", output)
	}

}

func TestIsRight(t *testing.T) {
	input := "foo"
	right := RightConstructor[string](input)

	if output := right.IsRight(); output == false {
		t.Errorf("Received %t - expected true", output)
	}
}
