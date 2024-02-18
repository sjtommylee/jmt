package fn

import "time"

type Eq interface {
	Equals(x, y interface{}) bool
}

// Eq for int
type IntEq struct{}

type StringEq struct{}

type DateEq struct{}

// struct concrete implementation
type StructEq struct {
	Fields map[string]Eq
}

// Equals - implements the Eq interface for int values
func (ie IntEq) Equals(x, y interface{}) bool {
	xInt, ok1 := x.(int)
	yInt, ok2 := y.(int)
	if !ok1 || !ok2 {
		return false
	}
	return xInt == yInt
}

// implements Eq interface for structs
func (se StructEq) Equals(x, y interface{}) bool {
	xStruct, ok1 := x.(map[string]interface{})
	yStruct, ok2 := y.(map[string]interface{})
	if !ok1 || !ok2 {
		return false
	}

	for key, eq := range se.Fields {
		if !eq.Equals(xStruct[key], yStruct[key]) {
			return false
		}
	}

	return true
}

// Date
func (de DateEq) Equals(x, y interface{}) bool {
	xTime, ok1 := x.(time.Time)
	yTime, ok2 := y.(time.Time)
	if !ok1 || !ok2 {
		return false
	}

	return xTime.Equal(yTime)
}

// String
func (se StringEq) Equals(x, y interface{}) bool {
	xStr, ok1 := x.(string)
	yStr, ok2 := y.(string)
	if !ok1 || !ok2 {
		return false
	}

	return xStr == yStr
}
