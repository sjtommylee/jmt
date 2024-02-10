package monoid

import "github.com/0xsj/fn-go/pkg/semigroup"

type Monoid interface {
	semigroup.Semigroup
	Empty() interface{}
}

type StringMonoid struct{}

// empty
func (s StringMonoid) Empty() interface{} {
	return ""
}
