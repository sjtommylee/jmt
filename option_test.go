package fn 

import (
	"testing"
)

func TestOptionMapSome(t *testing.T) {
	opt := Some[int]{Value: 5}

	newOpt := opt.Map(func(val int) int {
		return val + 1 
	})

	if newOpt.(Some[int]).Value != 6 {
		t.Errorf("Expected value 6, got %d", newOpt.(Some[int]).Value)
	}
}


func TestOptionMapNone(t *testing.T) {
	opt := None[int]{}

	newOpt := opt.Map(func(val int) int {
		return val + 1
	})

	if _, ok := newOpt.(None[int]); !ok {
		t.Errorf("Expected None, got Some")
	}
}

