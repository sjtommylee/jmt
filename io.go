package fn

// IO =  when executed, produces a value
type IO[A any] func() A

// Map applies a function to the result of an effectful computation.
func Map[A, B any](f func(A) B, io IO[A]) IO[B] {
	return func() B {
		return f(io())
	}
}
