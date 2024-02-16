package fn

// Magma - a basic algebraic structure consisting of a set equipped with one binary operation
type Magma[A any] interface {
	Concat(x, y A) A
}

// Predicate - represents a function that takes an argument of type A, and returns a boolean
type Predicate[A any] func(A) bool

// Endomorphism - reprsents a function that takes an argument of type A and returns a value of the same type A
type Endomorphism[A any] func(A) A

// Magmafunc - helper type
type MagmaFunc[A any] func(A, A) A

type MagmaInt struct{}

// reverse - reverse the order of concats
func Reverse[A any](M Magma[A]) Magma[A] {
	return &reversedMagma[A]{M}
}

type reversedMagma[A any] struct {
	M Magma[A]
}

func (rm *reversedMagma[A]) Concat(first, second A) A {
	return rm.M.Concat(second, first)
}

// FilterFirst returns a new Magma instance with the first arg filtered
func FilterFirst[A any](predicate Predicate[A]) func(Magma[A]) Magma[A] {
	return func(M Magma[A]) Magma[A] {
		return &filteredFirst[A]{predicate, M}
	}
}

type filteredFirst[A any] struct {
	predicate Predicate[A]
	M         Magma[A]
}

func (ff *filteredFirst[A]) Concat(first, second A) A {
	if ff.predicate(first) {
		return ff.M.Concat(first, second)
	}
	return second
}

// FilterSecond returns a new Magma instance with the second argument filtered
func FilterSecond[A any](predicate Predicate[A]) func(Magma[A]) Magma[A] {
	return func(M Magma[A]) Magma[A] {
		return &filteredSecond[A]{predicate, M}
	}
}

type filteredSecond[A any] struct {
	predicate Predicate[A]
	M         Magma[A]
}

func (fs *filteredSecond[A]) Concat(first, second A) A {
	if fs.predicate(first) {
		return fs.M.Concat(first, second)
	}
	return second
}

func Endo() {}
