package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type Action[T action.GoAction, Props action.GoActionProps] struct {
	definition definition.ActionDefinition[T, Props]
	Instance   *action.ActionInstance
}

func NewAction[T action.GoAction, Props action.GoActionProps](definition definition.ActionDefinition[T, Props]) *Action[T, Props] {
	instance := action.NewActionInstance(definition.Name)

	return &Action[T, Props]{
		definition: definition,
		Instance:   instance,
	}
}

func (a *Action[T, Props]) GetDef() *T {
	def, ok := a.definition.CtorValue.Interface().(action.GoActionConstructor[T, Props])
	if !ok {
		fmt.Println("ERRRR")
	}

	test := a.definition.DefaultProps

	return def(&action.ActionInstance{
		Model: a.Instance.Model,
	}, *test)
}
