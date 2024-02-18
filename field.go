package fn

// defines the interface for equality comparison

// defines the interface for a field
type Field interface {
	Ring
	Degree(a interface{}) int
	Div(x, y interface{}) interface{}
	Mod(x, y interface{}) interface{}
}

// ring defines  the interface for a ring
type Ring interface {
	Add(x, y interface{}) interface{}
	Zero() interface{}
	Mul(x, y interface{}) interface{}
	One() interface{}
	Sub(x, y interface{}) interface{}
}

// implements field for type int
type FieldNumber struct{}

// Equals
func (f FieldNumber) Equals(x interface{}, y interface{}) bool {
	// Check if x and y are integers
	xVal, xOK := x.(int)
	yVal, yOK := y.(int)
	if !xOK || !yOK {
		// If x or y is not an integer, return false
		return false
	}
	// Compare x and y for equality
	return xVal == yVal
}

func (f FieldNumber) Zero() interface{} {
	return 0
}

func (f FieldNumber) One() interface{} {
	return 1
}

func (f FieldNumber) Add(x, y interface{}) interface{} {
	return x.(int) + y.(int)
}

func (f FieldNumber) Sub(x, y interface{}) interface{} {
	return x.(int) - y.(int)
}

func (f FieldNumber) Mul(x, y interface{}) interface{} {
	return x.(int) * y.(int)
}

func (f FieldNumber) Div(x, y interface{}) interface{} {
	return x.(int) / y.(int)
}

func (f FieldNumber) Degree(a interface{}) int {
	return 1
}

func (f FieldNumber) Mod(x, y interface{}) interface{} {
	return x.(int) % y.(int)
}

// euclidean algorithm to find the greatest common divisor of two values
func gcd(E Eq, field Field) func(x, y interface{}) interface{} {
	zero := field.Zero()
	var f func(x, y interface{}) interface{}
	f = func(x, y interface{}) interface{} {
		if E.Equals(y, zero) {
			return x
		}
		return f(y, field.Mod(x, y))
	}
	return f
}

// lcm calculates the least common multiple of two values
func lcm(E Eq, F Field) func(x, y interface{}) interface{} {
	zero := F.Zero()
	gcdSF := gcd(E, F)
	return func(x, y interface{}) interface{} {
		if E.Equals(x, zero) || E.Equals(y, zero) {
			return zero
		}
		return F.Div(F.Mul(x, y), gcdSF(x, y))
	}
}
