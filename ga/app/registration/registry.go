package registration

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
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

func AcceptAction[T action.GoAction, P action.GoActionProps](def *definition.ActionDefinition[T, P]) func(*ActionRegistry) {
	return func(ar *ActionRegistry) {
		ar.actionsByName[def.Name] = def
		ar.actionsByType[def.ActionType] = def
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](actionType reflect.Type) func(*ActionRegistry) (*definition.ActionDefinition[T, P], error) {
	return func(ar *ActionRegistry) (*definition.ActionDefinition[T, P], error) {
		if action, exists := ar.actionsByType[actionType]; exists {
			return action.(*definition.ActionDefinition[T, P]), nil
		}
		return nil, fmt.Errorf("could not retrive action '%s'", actionType)
	}
}
