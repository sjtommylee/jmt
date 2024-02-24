package fn

import (
	"fmt"
	"reflect"
)

type Composable struct {
	Function interface{}
}

func Compose(fns ...interface{}) interface{} {
	if len(fns) == 0 {
		return nil
	}

	var composed interface{}
	composed = fns[len(fns)-1]

	for i := len(fns) - 2; i >= 0; i-- {
		fn := fns[i]
		fnType := reflect.TypeOf(fn)
		composedType := reflect.TypeOf(composed)

		if fnType.NumOut() != 1 || fnType.Out(0) != composedType.In(0) {
			return fmt.Errorf("Output type of function %d does not match input type of composed function", i)
		}

		composed = func(args ...interface{}) interface{} {
			return reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(composed.(func(...interface{}) interface{})(args...))})[0].Interface()
		}
	}

	return composed
}
