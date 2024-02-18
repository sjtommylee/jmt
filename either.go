package fn

// Represents the left variant of Either
type Left[E any] struct {
	Tag  string
	Left interface{}
}

// Represents the right variant of Either
type Right[A any] struct {
	Tag   string
	Right interface{}
}

// Either type with left, right variants
type Either[E any, A any] interface{}

// Creates instance of Left
func LeftConstructor[E any](value E) *Left[E] {
	return &Left[E]{Tag: "Left", Left: value}
}

// Creates instance of Right
func RightConstructor[A any](value A) *Right[A] {
	return &Right[A]{Tag: "Right", Right: value}
}

// IsLeft returns true for Left instances, false otherwise
func (l *Left[E]) IsLeft() bool {
	return true
}

// IsRight returns true for Right instances, false otherwise.
func (r *Right[A]) IsRight() bool {
	return false
}
