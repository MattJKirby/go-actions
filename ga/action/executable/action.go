package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type Action[T action.GoAction, P action.GoActionProps] struct {
	definition definition.ActionDefinition[T, P]
	Instance   *action.ActionInstance
	Action     *T
}

func NewAction[T action.GoAction, P action.GoActionProps](definition definition.ActionDefinition[T, P]) *Action[T, P] {
	instance, constructed := newPopulatedInstance(definition, definition.DefaultProps)
	return &Action[T, P]{
		definition: definition,
		Instance:   instance,
		Action:     constructed,
	}
}

func (a *Action[T, P]) PopulateActionInstance(props *P) {
	a.Instance, a.Action = newPopulatedInstance(a.definition, props)
}

func newPopulatedInstance[T action.GoAction, P action.GoActionProps](def definition.ActionDefinition[T, P], props *P) (*action.ActionInstance, *T) {
	if props == nil {
		props = def.DefaultProps
	}
	instance := action.NewActionInstance(def.Name)
	constructed := def.Constructor(instance, *props)
	return instance, constructed
}
