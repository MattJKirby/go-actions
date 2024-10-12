package action

import (
	"fmt"
)

type Constructor[T Action] func() *T

type Action interface {
	Execute()
}

type GoAction[T Action] struct {
	Definition *ActionDefinition
	Instance *ActionInstance
}

func NewAction[T Action](def *ActionDefinition, inst *ActionInstance) *GoAction[T] {
	return &GoAction[T]{
		Definition: def,
		Instance: inst,
	}
}

func (a *GoAction[T]) GetDef() *T {
	def, ok := a.Definition.vCtor.Interface().(Constructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def()
}
