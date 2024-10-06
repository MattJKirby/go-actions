package action

import (
	"fmt"
)

type Constructor[T Action] func() *T

type Action interface {
	Execute()
}

type GoAction[T Action] struct {
	definition *ActionDefinition
	instance *ActionInstance
}

func NewAction[T Action](def *ActionDefinition, inst *ActionInstance) *GoAction[T] {
	return &GoAction[T]{
		definition: def,
		instance: inst,
	}
}

func (a *GoAction[T]) GetDef() *T {
	def, ok := a.definition.vCtor.Interface().(Constructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def()
}
