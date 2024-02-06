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

// returns a function that when called, will produce a random element from the provided no empty arr
func RandomElem[A any](as []A) func() A {
	return func() A {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		index := randomGenerator.Intn(len(as))
		return as[index]
	}
}
