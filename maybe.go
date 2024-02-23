// 

package fn 

// Maybe represents a value that may or may not be present

type Maybe[T any] interface {
	Map(func(T) T) Maybe[T]
}

// just - present value
type Just[T any] struct {
	Value T
}

func (j Just[T]) Map(fn func(T) T) Maybe[T] {
	return Just[T]{Value: fn(j.Value)}
}

// nothing - represents an absent value 
type Nothing[T any] struct {}

func (n Nothing[T]) Map(fn func(T) T) Maybe[T] {
	return n
}
