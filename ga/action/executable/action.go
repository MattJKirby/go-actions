package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type Action[T action.GoAction] struct {
	Definition *definition.ActionDefinition
	Instance   *action.ActionInstance
}

func NewAction[T action.GoAction](definition *definition.ActionDefinition) *Action[T] {
	instance := action.NewActionInstance(definition.Name)

	return &Action[T]{
		Definition: definition,
		Instance:   instance,
	}
}

func (a *Action[T]) GetDef() *T {
	def, ok := a.Definition.CtorValue.Interface().(action.GoActionConstructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def(&action.ActionInstance{
		Model: a.Instance.Model,
	})
}
