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

func AcceptRegistration[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) func(*ActionRegistry) error {
	return func(ar *ActionRegistry) error {
		definition, err := definition.NewActionDefinition(reg)
		if err != nil {
			return err
		}

		ar.actionsByName[definition.Name] = definition
		ar.actionsByType[definition.ActionType] = definition
		return nil
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](actionType reflect.Type) func(*ActionRegistry) (*definition.ActionDefinition[T, P], error) {
	return func(ar *ActionRegistry) (*definition.ActionDefinition[T, P], error) {
		if action, exists := ar.actionsByType[actionType]; exists {
			copy := *action.(*definition.ActionDefinition[T, P])
			return &copy, nil
		}
		return nil, fmt.Errorf("could not retrive action '%s'", actionType)
	}
}
