package random

import (
	"math/rand"
	"time"

	"github.com/0xsj/fn-go/pkg/io"
)

// generates a random float64 within a specified range
var Random io.IO[float64] = func() float64 {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	return randomGenerator.Float64()
}

// generates a random integer
var RandomInt = func(low, high int) io.IO[int] {
	return func() int {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		return low + randomGenerator.Intn(high-low+1)
	}
}

// does this return a tuple??
func RandomRange() {}

// generates a random boolean value.
var RandomBool = func() io.IO[bool] {
	return func() bool {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		return randomGenerator.Intn(2) == 1
	}
}

func RandomElem() {}
