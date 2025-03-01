package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
)

type InitialisedAction struct {
	Action              action.GoAction
	InitialisedInstance *action.ActionInstance
}

func InitialiseNewAction(config model.ActionModelConfig, typeDef *definition.ActionTypeDefinition) (*InitialisedAction, error) {
	instance := action.NewActionInstance(typeDef.TypeName, config)
	defaultProps := typeDef.NewDefaultProps()
	ctor := typeDef.NewConstructor()
	action, err := ctor(instance, defaultProps)
	if err != nil {
		return nil, err
	}

	return &InitialisedAction{
		Action:              action,
		InitialisedInstance: instance,
	}, nil
}

type InitialisedTypedAction[T action.GoAction] struct {
	Action              *T
	InitialisedInstance *action.ActionInstance
}

func InitialiseNewTypedAction[T action.GoAction, P action.GoActionProps](config model.ActionModelConfig, def *definition.ActionDefinition[T, P]) (*InitialisedTypedAction[T], error) {
	instance := action.NewActionInstance(def.TypeName, config)
	action := def.Constructor(instance, *def.DefaultProps)

	// executableAction := executable.NewExecutableAction(app.modelConfig, reg.GetTypeDefinition())
	// act, ok := any(executableAction.Action).(*T)
	// if !ok {
	// 	return nil, fmt.Errorf("could nt ")
	// }

	return &InitialisedTypedAction[T]{
		Action:              action,
		InitialisedInstance: instance,
	}, nil
}
