package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type RegisteredAction[T action.GoAction] struct {
	registration *action.GoActionRegistration[T]
	actionDefinition *definition.ActionDefinition
}

func NewRegisteredAction[T action.GoAction](registration *action.GoActionRegistration[T]) (*RegisteredAction[T], error) {
	actionDefinition, err := definition.NewActionDefinition(registration.Constructor)
	if err != nil {
		return nil, err
	}
	
	return &RegisteredAction[T]{
		registration,
		actionDefinition,
	}, nil
}