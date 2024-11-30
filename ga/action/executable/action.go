package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
)

type Action[T action.GoAction] struct {
	Definition *definition.ActionDefinition
	Instance   *model.ModelInstance
}

func NewAction[T action.GoAction](definition *definition.ActionDefinition) *Action[T] {
	internals := action.NewGoActionInternals(definition.Name)

	return &Action[T]{
		Definition: definition,
		Instance:   internals.ModelInstance,
	}
}

func (a *Action[T]) GetDef() *T {
	def, ok := a.Definition.CtorValue.Interface().(action.GoActionConstructor[T])
	if !ok {
		fmt.Println("ERRRR")
	}

	return def(action.GoActionInternals{
		ModelInstance: a.Instance,
	})
}
