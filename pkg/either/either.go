package either

// Represents the left variant of Either
type Left[E any] struct {
	Tag  string
	Left E
}

// Reprsents the right variant of Either
type Right[A any] struct {
	Tag   string
	Right A
}

// Either type with left, right variants
type Either[E any, A any] interface{}

// Creates instance of Left
func LeftConstructor[E any](value E) *Left[E] {
	return &Left[E]{Tag: "Left", Left: value}
}

// Creates instanece of Right
func RightConstructor[A any](value A) *Right[A] {
	return &Right[A]{Tag: "Right", Right: value}
}

// isLeft returns true for Left instances, false otherwise
func (l *Left[E]) IsLeft() bool {
	return true
}

// isRight - returns true for Right instances, false otherwise.
func (r *Right[A]) IsRight() bool {
	return false
}
