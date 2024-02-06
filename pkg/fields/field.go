package field

// defines the interface for equality comparison
type Eq interface {
	Equals(other interface{}) bool
}

// defines the interface for a field
type Field interface {
	Ring
	Degree(a interface{}) int
	Div(x, y interface{}) interface{}
}

// ring defines  the interface for a ring
type Ring interface {
}

type FieldNumber struct{}

func (f FieldNumber) Zero() interface{} {
	return 0
}

func (f FieldNumber) Add(x, y interface{}) interface{} {
	return x.(int) + y.(int)
}

func (f FieldNumber) Mul(x, y interface{}) interface{} {
	return x.(int) * y.(int)
}

func (f FieldNumber) One() interface{} {
	return 1
}

func Sub(f FieldNumber) {}

func Degree(f FieldNumber) {}

func Div(f FieldNumber) {}

func Mod(f FieldNumber) {}
