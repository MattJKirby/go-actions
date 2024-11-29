package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/instance"
	"go-actions/ga/action/internals"
)

type Action[T action.GoAction] struct {
	Definition *definition.ActionDefinition
	Instance   *instance.ActionInstance
}

func NewAction[T action.GoAction](definition *definition.ActionDefinition) *Action[T] {
	internals := internals.NewGoActionInternals(definition.Name)

	return &Action[T]{
		Definition: definition,
		Instance:   internals.ActionInstance,
	}
}

func (a *Action[T]) GetDef() *T {
	def, ok := a.Definition.CtorValue.Interface().(action.GoActionConstructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def(internals.GoActionInternals{
		ActionInstance: a.Instance,
	})
}
