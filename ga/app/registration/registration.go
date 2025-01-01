package registration

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type RegisteredAction[T action.GoAction] struct {
	registration *action.GoActionRegistration[T]
	ActionDefinition *definition.ActionDefinition
}

func NewRegisteredAction[T action.GoAction](registration *action.GoActionRegistration[T]) (*RegisteredAction[T], error) {
	ActionDefinition, err := definition.NewActionDefinition(registration.Constructor)
	if err != nil {
		return nil, err
	}

	return &RegisteredAction[T]{
		registration,
		ActionDefinition,
	}, nil
}