package registration

import (
	"fmt"
	"go-actions/ga/action"
	"reflect"
)

type ActionRegistry struct {
	actionsByName map[string]any
	actionsByType map[reflect.Type]any
}

func NewActionRegistry() *ActionRegistry {
	return &ActionRegistry{
		actionsByName: make(map[string]any),
		actionsByType: make(map[reflect.Type]any),
	}
}

func AcceptAction[T action.GoAction, P action.GoActionProps](a *RegisteredAction[T, P]) func(*ActionRegistry) *RegisteredAction[T, P] {
	return func(ar *ActionRegistry) *RegisteredAction[T, P] {
		ar.actionsByName[a.ActionDefinition.Name] = a
		ar.actionsByType[a.ActionDefinition.ActionType] = a
		return a
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](actionType reflect.Type) func(*ActionRegistry) (*RegisteredAction[T, P], error) {
	return func(ar *ActionRegistry) (*RegisteredAction[T, P], error) {
		if action, exists := ar.actionsByType[actionType]; exists {
			return action.(*RegisteredAction[T, P]), nil
		}
		return nil, fmt.Errorf("could not retrive action '%s'", actionType)
	}
}
