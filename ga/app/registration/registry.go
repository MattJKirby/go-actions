package registration

import (
	"fmt"
	"go-actions/ga/action"
	"reflect"
)

type ActionRegistry struct {
	actionsByName map[string]*action.ActionTypeDefinition
	actionsByType map[reflect.Type]*action.ActionTypeDefinition
}

func NewActionRegistry() *ActionRegistry {
	return &ActionRegistry{
		actionsByName: make(map[string]*action.ActionTypeDefinition),
		actionsByType: make(map[reflect.Type]*action.ActionTypeDefinition),
	}
}

func AcceptRegistration[T action.GoAction](reg *action.ActionRegistration[T]) func(*ActionRegistry) {
	return func(ar *ActionRegistry) {
		definition := action.TypeDefinitionFromRegistration(reg)
		ar.actionsByName[definition.TypeName] = definition
		ar.actionsByType[definition.ActionType] = definition

	}
}

func GetTypeDefinitionByName(actionName string) func(*ActionRegistry) (*action.ActionTypeDefinition, error) {
	return func(ar *ActionRegistry) (*action.ActionTypeDefinition, error) {
		if def, exists := ar.actionsByName[actionName]; exists {
			return def, nil
		}
		return nil, fmt.Errorf("could not retrive action with name '%s'", actionName)
	}
}

func GetTypeDefinitionByType(actionType reflect.Type) func(*ActionRegistry) (*action.ActionTypeDefinition, error) {
	return func(ar *ActionRegistry) (*action.ActionTypeDefinition, error) {
		if def, exists := ar.actionsByType[actionType]; exists {
			return def, nil
		}
		return nil, fmt.Errorf("could not retrive action with type '%s'", actionType)
	}
}
