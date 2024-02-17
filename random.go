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
