package registration

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"reflect"
)

type definitionRegistration interface {
	GetTypeDefinition() *definition.ActionTypeDefinition
}

type ActionRegistry struct {
	actionsByName map[string]definitionRegistration
	actionsByType map[reflect.Type]definitionRegistration
}

func NewActionRegistry() *ActionRegistry {
	return &ActionRegistry{
		actionsByName: make(map[string]definitionRegistration),
		actionsByType: make(map[reflect.Type]definitionRegistration),
	}
}

func AcceptRegistration[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) func(*ActionRegistry) error {
	return func(ar *ActionRegistry) error {
		definition, err := definition.NewActionDefinition(reg)
		if err != nil {
			return err
		}

		if defReg, ok := any(definition).(definitionRegistration); ok {
			ar.actionsByName[definition.Name] = defReg
			ar.actionsByType[definition.ActionType] = defReg
			return nil
		}

		return fmt.Errorf("error registering definition for action '%s'", reg.Name)
	}
}

func GetTypedActionDefinition[T action.GoAction, P action.GoActionProps](actionType reflect.Type) func(*ActionRegistry) (*definition.ActionDefinition[T, P], error) {
	return func(ar *ActionRegistry) (*definition.ActionDefinition[T, P], error) {
		if action, exists := ar.actionsByType[actionType]; exists {
			copy := *action.(*definition.ActionDefinition[T, P])
			return &copy, nil
		}
		return nil, fmt.Errorf("could not retrive action '%s'", actionType)
	}
}

func GetRegisteredTypeDefinitionByName(actionName string) func(*ActionRegistry) (*definition.ActionTypeDefinition, error) {
	return func(ar *ActionRegistry) (*definition.ActionTypeDefinition, error) {
		if action, exists := ar.actionsByName[actionName]; exists {
			return action.GetTypeDefinition(), nil
		}
		return nil, fmt.Errorf("could not retrive action with name'%s'", actionName)
	}
}
