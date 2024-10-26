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

func NewAction[T Action](definition *ActionDefinition) *GoAction[T] {
	instance := NewActionInstance(definition)
	
	return &GoAction[T]{
		Definition: definition,
		Instance: instance,
	}
}

func (a *GoAction[T]) GetDef() *T {
	def, ok := a.Definition.ctorValue.Interface().(Constructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def()
}
