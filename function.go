package fn

// idea 1

type PipeFunc func(interface{}) interface{}

// // pipe
func Pipe(input interface{}, funcs ...PipeFunc) interface{} {
	result := input
	for _, f := range funcs {
		result = f(result)
	}
	return result
}

// func add(a int) PipeFunc {
// 	return func(b interface{}) interface{} {
// 		return a + b.(int)
// 	}
// }

// func multiply(a int) PipeFunc {
// 	return func(b interface{}) interface{} {
// 		return a * b.(int)
// 	}
// }

// // actual usage might look like
// func baz() {
// 	result := Pipe(1, add(2), multiply(3))
// }
