package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type Action[T action.GoAction, Props any] struct {
	Definition *definition.ActionDefinition
	Instance   *action.ActionInstance
	Action     *T
}

func NewAction[T action.GoAction, Props any](definition *definition.ActionDefinition) *Action[T, Props] {
	instance := action.NewActionInstance(definition.Name)

	return &Action[T, Props]{
		Definition: definition,
		Instance:   instance,
	}
}

func (a *Action[T, Props]) GetDef() *T {
	def, ok := a.Definition.CtorValue.Interface().(action.GoActionConstructor[T, Props])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def(&action.ActionInstance{
		Model: a.Instance.Model,
	})
}
