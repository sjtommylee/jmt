package fn

// Represents the left variant of Either
type Left[E any] struct {
	Tag   string
	Value E
}

// Represents the right variant of Either
type Right[A any] struct {
	Tag   string
	Value A
}

// Either type with left, right variants
type Either[E any, A any] interface {
	IsLeft() bool
	IsRight() bool
	Left() E
	Right() A
}

// Creates instance of Left
func LeftConstructor[E any](value E) *Left[E] {
	return &Left[E]{Tag: "Left", Value: value}
}

// Creates instance of Right
func RightConstructor[A any](value A) *Right[A] {
	return &Right[A]{Tag: "Right", Value: value}
}

// IsLeft returns true for Left instances, false otherwise
func (l *Left[E]) IsLeft() bool {
	return true
}

// IsRight returns true for Right instances, false otherwise.
func (r *Right[A]) IsRight() bool {
	return false
}

// Left returns the value of the Left variant
func (l *Left[E]) Left() E {
	return l.Value
}

// Right returns the value of the Right variant
func (r *Right[A]) Right() A {
	return r.Value
}
