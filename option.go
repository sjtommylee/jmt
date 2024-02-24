package fn

// Option is a container for an optional value of type T.
// If value exists, Option is of type Some. If the value is absent, Option is of type None.
type Option[T any] struct {
	isPresent bool
	value     T
}

// Some builds an Option when value is present.
func Some[T any](value T) Option[T] {
	return Option[T]{
		isPresent: true,
		value:     value,
	}
}

// None builds an Option when value is absent.
func None[T any]() Option[T] {
	return Option[T]{
		isPresent: false,
	}
}

// Map executes the mapper function if value is present or returns None if absent.
func (o Option[T]) Map(mapper func(value T) (T, bool)) Option[T] {
	if o.isPresent {
		return TupleToOption(mapper(o.value))
	}

	return None[T]()
}

// FlatMap executes the mapper function if value is present or returns None if absent.
func (o Option[T]) FlatMap(mapper func(value T) Option[T]) Option[T] {
	if o.isPresent {
		return mapper(o.value)
	}

	return None[T]()
}

// Match executes the first function if value is present and second function if absent.
// It returns a new Option.
func (o Option[T]) Match(onValue func(value T) (T, bool), onNone func() (T, bool)) Option[T] {
	if o.isPresent {
		return TupleToOption(onValue(o.value))
	}
	return TupleToOption(onNone())
}

// TupleToOption builds a Some Option when second argument is true, or None.
func TupleToOption[T any](value T, ok bool) Option[T] {
	if ok {
		return Some(value)
	}
	return None[T]()
}
