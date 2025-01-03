package registration

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type RegisteredAction[T action.GoAction, Props action.GoActionProps] struct {
	Registration     *action.GoActionRegistration[T, Props]
	ActionDefinition *definition.ActionDefinition
}

func NewRegisteredAction[T action.GoAction, Props action.GoActionProps](registration *action.GoActionRegistration[T, Props]) (*RegisteredAction[T, Props], error) {
	ActionDefinition, err := definition.NewActionDefinition(registration.Constructor)
	if err != nil {
		return nil, err
	}

	return &RegisteredAction[T, Props]{
		registration,
		ActionDefinition,
	}, nil
}
