package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/app/registration"
)

type Action[T action.GoAction, Props action.GoActionProps] struct {
	registration registration.RegisteredAction[T, Props]
	Instance     *action.ActionInstance
}

func NewAction[T action.GoAction, Props action.GoActionProps](registration registration.RegisteredAction[T, Props]) *Action[T, Props] {
	instance := action.NewActionInstance(registration.ActionDefinition.Name)

	return &Action[T, Props]{
		registration: registration,
		Instance:     instance,
	}
}

func (a *Action[T, Props]) GetDef() *T {
	def, ok := a.registration.ActionDefinition.CtorValue.Interface().(action.GoActionConstructor[T, Props])
	if !ok {
		fmt.Println("ERRRR")
	}

	test := a.registration.Registration.DefaultProps

	return def(&action.ActionInstance{
		Model: a.Instance.Model,
	}, *test)
}
