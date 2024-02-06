package eq

// model
type Eq interface {
	Equals(x, y interface{}) bool
}

// structeq - defines a struct for combining multiple Eq instances
type StructEq struct {
	eqs map[string]Eq
}

// Equals implements interface for structEq
func (se StructEq) Equals(x, y interface{}) bool {
	first, ok := x.(map[string]interface{})
	if !ok {
		return false
	}
	second, ok := y.(map[string]interface{})
	if !ok {
		return false
	}

	for key, eq := range se.eqs {
		if !eq.Equals(first[key], second[key]) {
			return false
		}
	}
	return true
}
