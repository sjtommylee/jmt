package fn

import (
	"fmt"
	"reflect"
)

type Caller func([]reflect.Value) []reflect.Value

// unpacks the inputs into respective types and values
func Unpack(fns ...interface{}) ([]reflect.Type, []reflect.Value, error) {
	ts := make([]reflect.Type, len(fns))
	vs := make([]reflect.Value, len(fns))

	for i, fn := range fns {
		t := reflect.TypeOf(fn)
		if t.Kind() != reflect.Func {
			return nil, nil, fmt.Errorf("argument %d is not a func", i)
		}

		// Create a slice to hold the function's arguments
		args := make([]reflect.Value, t.NumIn())
		for j := 0; j < t.NumIn(); j++ {
			argType := t.In(j)
			arg := reflect.New(argType).Elem()
			args[j] = arg
		}

		ts[i] = t
		vs[i] = reflect.ValueOf(fn)
	}

	return ts, vs, nil
}

// extract
func ExtractIOTypes(fn reflect.Type) ([]reflect.Type, []reflect.Type) {
	in := make([]reflect.Type, fn.NumIn())
	out := make([]reflect.Type, fn.NumOut())
	for i := range in {
		in[i] = fn.In(i)
	}
	for i := range out {
		out[i] = fn.Out(i)
	}
	return in, out
}

// determine
func DetermineCaller(isVariadic bool, value reflect.Value) Caller {
	if isVariadic {
		return value.CallSlice
	}
	return value.Call
}

func CorrectVariadicArgs(in, out []reflect.Type, fnIdx int) (func([]reflect.Value) []reflect.Value, error) {
	ln := len(in) - 1
	if lo := len(out); ln == lo-1 && in[ln] == out[ln] {
		// Exact match, all good, no correction required
		return nil, nil
	} else if ln > lo {
		return nil, fmt.Errorf("too few return values preceding function %d, does not match arguments", fnIdx)
	}

	s := in[ln]
	v := s.Elem()
	for i, t := range out[ln:] {
		if t != v {
			return nil, fmt.Errorf("return value preceding function %d at %d does not match variadic type", fnIdx, i)
		}
	}

	toConvert := len(out) - ln
	correction := func(vs []reflect.Value) []reflect.Value {
		newVs := make([]reflect.Value, ln)
		copy(newVs, vs[:ln])
		sl := reflect.MakeSlice(reflect.SliceOf(v), toConvert, toConvert)
		for i, v := range vs[ln:] {
			sl.Index(i).Set(v)
		}
		var variadicArgs []reflect.Value
		for i := 0; i < sl.Len(); i++ {
			variadicArgs = append(variadicArgs, sl.Index(i))
		}
		return append(newVs, variadicArgs...)
	}
	return correction, nil
}
