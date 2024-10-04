package action

import (
	"fmt"
)

type Constructor[T Action] func () *T

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

func (a *GoAction[T]) GetDef() *T {
	def, ok := a.def.vCtor.Interface().(Constructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}


	return def()
}
