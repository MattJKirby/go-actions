package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type InitialisedAction struct {
	Action action.GoAction
	InitialisedInstance *action.ActionInstance
}

func NewInitialisedAction(app *App, typeDef *definition.ActionTypeDefinition) (*InitialisedAction, error) {
	instance := action.NewActionInstance(typeDef.TypeName, app.modelConfig)
	defaultProps := typeDef.NewDefaultProps()
	ctor := typeDef.NewConstructor()
	action, err := ctor(instance, defaultProps)
	if err != nil {
		return nil, err
	}
	
	return &InitialisedAction{
		Action: action,
		InitialisedInstance: instance,
	}, nil
}