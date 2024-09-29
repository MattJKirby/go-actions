package action

import (
	"fmt"
	"reflect"
)

type Constructor[T Action] func () T

type Action interface {
	Execute()
}

type GoAction[T Action] struct {
	def *ActionDefinition
}

func NewAction[T Action](def *ActionDefinition) *GoAction[T] {

	return &GoAction[T]{
		def,
	}
}

func (a *GoAction[T]) GetDef() {
	def := a.def.vCtor.Call([]reflect.Value{})[0]
	
	fmt.Println(def)
}
