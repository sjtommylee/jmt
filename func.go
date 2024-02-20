package fn

// idea 2
import (
	"fmt"
	"testing"
)

type Pipeline struct {
	value interface{}
}

func NewPipe(value interface{}) *Pipeline {
	return &Pipeline{value: value}
}

func (p *Pipeline) Add(a int) *Pipeline {
	p.value = p.value.(int) + a
	return p
}

func (p *Pipeline) Multiply(a int) *Pipeline {
	p.value = p.value.(int) * a
	return p
}

func (p *Pipeline) Result() interface{} {
	return p.value
}

func TestPipe2(t *testing.T) {
	result := NewPipe(1).Add(2).Multiply(3).Result()
	fmt.Println(result) // result is 9
}

// actual usage in this case might look like

// func foo() {
// 	result := fn.NewPipeline(1).
// 		Add(2).
// 		Multiply(3).
// 		Result()

// 	// Ok()
// }
