package eq 

type Eq interface {
  Equals(x, y interface{}) bool
}

func fromEquals(equals func(x, y interface{}) bool) Eq {
  return &eqFunc(equals)
}



