// Option represents a value that may or may not be present
// It is used to avid null pointer exceptions, and encourages
// handling both cases where a value is present (Some), and
// where it is absent (None)
package fn

// Option represents a value that may or may not be present.

type Option[T any] interface {
	// Map applies the given function the value inside the Option
	// and returns a new Optino with the result
	Map(func(T) T) Option[T]
}

// some represents a present value
type Some[T any] struct {
	Value T
}

func (s Some[T]) Map(fn func(T) T) Option[T] {
	return Some[T]{Value: fn(s.Value)}
}

// None represents an absent value
type None[T any] struct{}

func (n None[T]) Map(fn func(T) T) Option[T] {
	return n
}
