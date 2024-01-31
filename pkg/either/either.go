package either

// Represents the left variant of Either
type Left[E any] struct {
	_Tag string
	Left E
}

// Reprsents the right variant of Either
type Right[A any] struct {
	_Tag  string
	Right A
}

// Either type with left, right variants
type Either[E any, A any] interface{}

// Creates instance of Left
func LeftConstructor[E any](value E) *Left[E] {
	return &Left[E]{_Tag: "Left", Left: value}
}

// Creates instanece of Right
func RightConstructor[A any](value A) *Right[A] {
	return &Right[A]{_Tag: "Right", Right: value}
}

// isLeft
func (l *Left[E]) IsLeft() bool {
	return true
}

// isRight
func (r *Right[A]) IsRight() bool {
	return true
}
