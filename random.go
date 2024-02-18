package fn

import (
	"math/rand"
	"time"
)

// generates a random float64 within a specified group
var Random IO[float64] = func() float64 {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	return randomGenerator.Float64()
}

// generates a random inteer
var RandomInt = func(low, high int) IO[int] {
	return func() int {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		return low + randomGenerator.Intn(high-low+1)
	}
}

// returns a random range
func RandomRange() {}

// generates a random boolena value
var RandomBool = func() IO[bool] {
	return func() bool {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		return randomGenerator.Intn(2) == 1
	}
}

// returns a function that when called, will produce a random elementfrom the provided non empty array
func RandomElem[A any](as []A) func() A {
	return func() A {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		index := randomGenerator.Intn(len(as))
		return as[index]
	}
}
