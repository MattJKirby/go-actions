package registration

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"reflect"
)

type ActionRegistry struct {
	actionsByName map[string]*definition.ActionTypeDefinition
	actionsByType map[reflect.Type]*definition.ActionTypeDefinition
}

func NewActionRegistry() *ActionRegistry {
	return &ActionRegistry{
		actionsByName: make(map[string]*definition.ActionTypeDefinition),
		actionsByType: make(map[reflect.Type]*definition.ActionTypeDefinition),
	}
}

func AcceptRegistration[T action.GoAction](reg *action.GoActionRegistration[T]) func(*ActionRegistry) {
	return func(ar *ActionRegistry) {
		definition := definition.TypeDefinitionFromRegistration(reg)
		ar.actionsByName[definition.TypeName] = definition
		ar.actionsByType[definition.ActionType] = definition

	}
}

func GetTypeDefinitionByName(actionName string) func(*ActionRegistry) (*definition.ActionTypeDefinition, error) {
	return func(ar *ActionRegistry) (*definition.ActionTypeDefinition, error) {
		if def, exists := ar.actionsByName[actionName]; exists {
			return def, nil
		}
		return nil, fmt.Errorf("could not retrive action with name '%s'", actionName)
	}
}

func GetTypeDefinitionByType(actionType reflect.Type) func(*ActionRegistry) (*definition.ActionTypeDefinition, error) {
	return func(ar *ActionRegistry) (*definition.ActionTypeDefinition, error) {
		if def, exists := ar.actionsByType[actionType]; exists {
			return def, nil
		}
		return nil, fmt.Errorf("could not retrive action with type '%s'", actionType)
	}
}
