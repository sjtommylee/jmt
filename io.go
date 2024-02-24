package fn

// IO =  when executed, produces a value
type IO[A any] func() A
