package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/instance"
)

type Action[T action.GoAction] struct {
	Definition *definition.ActionDefinition
	Instance *instance.ActionInstance
}

func NewAction[T action.GoAction](definition *definition.ActionDefinition) *Action[T] {
	instance := instance.NewActionInstance(definition.Name)
	
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
