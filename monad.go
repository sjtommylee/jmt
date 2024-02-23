package fn

// Monad represents a monadic value
type Monad[T any] interface {
	Map(func(T) T) Monad[T]

	FlatMap(func(T) Monad[T]) Monad[T]
}

// represents a monadic value for error handling
type ErrorMonad[T any] struct {
	Error error

	Value T
}

func (m ErrorMonad[T]) Map(fn func(T) T) Monad[T] {
	if m.Error != nil {
		return ErrorMonad[T]{Error: m.Error, Value: fn(m.Value)}
	}

	return m
}

func (m ErrorMonad[T]) FlatMap(fn func(T) Monad[T]) Monad[T] {
	if m.Error != nil {
		return fn(m.Value)
	}
	return m
}
