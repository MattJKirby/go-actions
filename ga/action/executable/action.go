package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type Action[T action.GoAction, P action.GoActionProps] struct {
	Action     *T
	definition definition.ActionDefinition[T, P]
	Instance   *action.ActionInstance
}

func NewAction[T action.GoAction, P action.GoActionProps](definition definition.ActionDefinition[T, P]) *Action[T, P] {
	instance := action.NewActionInstance(definition.Name)
	action := applyConstructor(definition, instance, definition.DefaultProps)

	return &Action[T, P]{
		Action:     action,
		definition: definition,
		Instance:   instance,
	}
}

func (a *Action[T, P]) ApplyConstructor(props *P) {
	a.Action = applyConstructor(a.definition, a.Instance, props)
}

func applyConstructor[T action.GoAction, P action.GoActionProps](def definition.ActionDefinition[T, P], instance *action.ActionInstance, props *P) *T {
	if props == nil {
		props = def.DefaultProps
	}
	return def.Constructor(instance, *props)
}

// func (a *Action[T, Props]) GetDef() *T {
// 	// def, ok := a.definition.CtorValue.Interface().(action.GoActionConstructor[T, Props])
// 	// if !ok {
// 	// 	fmt.Println("ERRRR")
// 	// }

// 	a.definition.Constructor()

// 	test := a.definition.DefaultProps

// 	return def(&action.ActionInstance{
// 		Model: a.Instance.Model,
// 	}, *test)
// }
