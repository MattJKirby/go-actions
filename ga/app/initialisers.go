package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/app/config"
)

type InitialisedAction struct {
	Action              action.GoAction
	InitialisedInstance *action.ActionInstance
}

func InitialiseNewAction(config *config.GlobalConfig, typeDef *definition.ActionTypeDefinition) (*InitialisedAction, error) {
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
