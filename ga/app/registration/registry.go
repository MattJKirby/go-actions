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

func AcceptAction[T action.GoAction](a *RegisteredAction[T]) func (*ActionRegistry) *RegisteredAction[T] {
	return func(ar *ActionRegistry) *RegisteredAction[T] {
    ar.actionsByName[a.ActionDefinition.Name] = a
    ar.actionsByType[a.ActionDefinition.ActionType] = a
    return a
	}
}

func GetAction[T action.GoAction](actionType reflect.Type) func (*ActionRegistry) (*RegisteredAction[T], error) {
	return func(ar *ActionRegistry) (*RegisteredAction[T], error) {
		if action, exists := ar.actionsByType[actionType]; exists {
			return action.(*RegisteredAction[T]), nil
		}
		return nil, fmt.Errorf("could not retrive action '%s'", actionType)
	}
}
