package fn

import (
	"reflect"
)
// represents a function with defind input and output types
type Composable struct {
	InputType  reflect.Type
	OutputType reflect.Type
}

// input returns the input type of the function via reflect. 
func (f Composable) Input() reflect.Type {
	return f.InputType
}

// output returns the output type of the function via reflect
func (f Composable) Output() reflect.Type {
	return f.OutputType
}

// Compose 
func Mega(fns ...Composable) (Composable, error) {

}
