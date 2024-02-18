package fn

type PipeFunc func(interface{}) interface{}

// // pipe
func Pipe(input interface{}, funcs ...PipeFunc) interface{} {
	result := input
	for _, f := range funcs {
		result = f(result)
	}
	return result
}
