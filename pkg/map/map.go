package fn_map

import "github.com/0xsj/fn-go/pkg/either"

func Map[E any, A any, B any](e either.Either[E, A], f func(A) B) either.Either[E, B] {
	switch v := e.(type) {
	case *either.Left[E]:
		return v
	case *either.Right[A]:
		return either.RightConstructor[B](f(v.Right))
	default:
		return nil
	}
}
