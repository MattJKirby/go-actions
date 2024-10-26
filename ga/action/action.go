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

func NewAction[T Action](def *ActionDefinition) *GoAction[T] {
	return &GoAction[T]{
		Definition: def,
		Instance: NewActionInstance(def),
	}
}

func (a *GoAction[T]) GetDef() *T {
	def, ok := a.Definition.ctorValue.Interface().(Constructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def()
}
