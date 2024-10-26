package executable

import (
	"fmt"
	"go-actions/ga/action"
)

type Action[T action.GoAction] struct {
	Definition *action.ActionDefinition
	Instance *action.ActionInstance
}

func NewAction[T action.GoAction](definition *action.ActionDefinition) *Action[T] {
	instance := action.NewActionInstance(definition)
	
	return &Action[T]{
		Definition: definition,
		Instance: instance,
	}
}

func (a *Action[T]) GetDef() *T {
	def, ok := a.Definition.CtorValue.Interface().(action.GoActionConstructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def()
}
