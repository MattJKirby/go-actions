package app

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"reflect"
)

type definitionRegistration interface {
	GetTypeDefinition() *definition.ActionTypeDefinition
}

type actionRegistry struct {
	actionsByName map[string]definitionRegistration
	actionsByType map[reflect.Type]definitionRegistration
}

func newActionRegistry() *actionRegistry {
	return &actionRegistry{
		actionsByName: make(map[string]definitionRegistration),
		actionsByType: make(map[reflect.Type]definitionRegistration),
	}
}

func acceptRegistration[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) func(*actionRegistry) error {
	return func(ar *actionRegistry) error {
		definition := definition.NewActionDefinition(reg)

		if defReg, ok := any(definition).(definitionRegistration); ok {
			ar.actionsByName[definition.TypeName] = defReg
			ar.actionsByType[definition.ActionType] = defReg
			return nil
		}

		return fmt.Errorf("error registering definition for action '%s'", reg.Name)
	}
}

func getTypedActionDefinition[T action.GoAction, P action.GoActionProps](actionType reflect.Type) func(*actionRegistry) (*definition.ActionDefinition[T, P], error) {
	return func(ar *actionRegistry) (*definition.ActionDefinition[T, P], error) {
		if action, exists := ar.actionsByType[actionType]; exists {
			copy := *action.(*definition.ActionDefinition[T, P])
			return &copy, nil
		}
		return nil, fmt.Errorf("could not retrive action '%s'", actionType)
	}
}

func getRegisteredTypeDefinitionByName(actionName string) func(*actionRegistry) (*definition.ActionTypeDefinition, error) {
	return func(ar *actionRegistry) (*definition.ActionTypeDefinition, error) {
		if action, exists := ar.actionsByName[actionName]; exists {
			return action.GetTypeDefinition(), nil
		}
		return nil, fmt.Errorf("could not retrive action with name '%s'", actionName)
	}
}
