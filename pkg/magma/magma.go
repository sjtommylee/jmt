package magma 

type Magma interface {
  Concat(x, y interface{}) interface{} 
}

type StringSemigroup struct {}
func (s StringSemigroup) Concat(x, y interface{}) interface{} {
  return x.(string) + y.(string)
}
