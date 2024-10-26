package action

import (
	"fmt"
)

type GoActionConstructor[T GoAction] func() *T

type GoAction interface {
	Execute()
}

type Action[T GoAction] struct {
	Definition *ActionDefinition
	Instance *ActionInstance
}

func NewAction[T GoAction](definition *ActionDefinition) *Action[T] {
	instance := NewActionInstance(definition)
	
	return &Action[T]{
		Definition: definition,
		Instance: instance,
	}
}

func (a *Action[T]) GetDef() *T {
	def, ok := a.Definition.ctorValue.Interface().(GoActionConstructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def()
}
