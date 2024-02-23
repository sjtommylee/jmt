package fn

import "fmt"

var (
	ErrInvalidArgumentId = fmt.Errorf("invalid argument id")
	ErrMissingArg        = fmt.Errorf("missing argument")
)

// Either represents a value that can be one of many types.
type Either struct {
	argId int
	args  []interface{}
}

// NewEither creates a new Either with the specified argument set.
func NewEither(argId int, args ...interface{}) Either {
	return Either{
		argId: argId,
		args:  args,
	}
}

// GetArg returns the argument at the specified index.
func (e Either) GetArg(index int) (interface{}, error) {
	if index < 0 || index >= len(e.args) {
		return nil, ErrMissingArg
	}
	return e.args[index], nil
}

// IsArg returns true if the argument at the specified index is set.
func (e Either) IsArg(argId int) bool {
	return e.argId == argId
}

// Match executes the given function for the argument at the specified index.
func (e Either) Match(argId int, fn func(interface{}) interface{}) (Either, error) {
	if !e.IsArg(argId) {
		return e, ErrInvalidArgumentId
	}
	arg, err := e.GetArg(argId)
	if err != nil {
		return e, err
	}
	result := fn(arg)
	return NewEither(argId, result), nil
}
