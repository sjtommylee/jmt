package fn

type Result[T any] struct {
	isErr bool
	value T
	err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		isErr: false,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		err:   err,
		isErr: true,
	}
}

func Map[T any, U any](r Result[T], mapper func(T) U) Result[U] {
	if r.isErr {
		return Err[U](r.err)
	}
	return Ok[U](mapper(r.value))
}

func FlatMap[T any, U any](r Result[T], mapper func(T) Result[U]) Result[U] {
	if r.isErr {
		return Err[U](r.err)
	}
	return mapper(r.value)
}

func Match[T any](r Result[T], onSuccess func(T) Result[T], onError func(error) Result[T]) Result[T] {
	if r.isErr {
		return onError(r.err)
	}
	return onSuccess(r.value)
}
