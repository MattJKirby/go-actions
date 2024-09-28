package app

import (
	"fmt"
	"go-actions/ga/action"
	"reflect"
)

type ActionDefinitionRegistry struct {
	actionsByName map[string]*action.ActionDefinition
	actionsByType map[reflect.Type]*action.ActionDefinition
}

func NewActionDefinitionRegistry() *ActionDefinitionRegistry {
	return &ActionDefinitionRegistry{
		actionsByName: make(map[string]*action.ActionDefinition),
		actionsByType: make(map[reflect.Type]*action.ActionDefinition),
	}
}

func (adr *ActionDefinitionRegistry) acceptDefinition(def *action.ActionDefinition) *action.ActionDefinition {
	adr.actionsByName[def.Name()] = def
	adr.actionsByType[def.ActionType()] = def
	return def
}

func (adr *ActionDefinitionRegistry) getDefinition(actionType reflect.Type) (*action.ActionDefinition, error) {
	if def, exists := adr.actionsByType[actionType]; exists {
		return def, nil
	}
	return nil, fmt.Errorf("could not retrive action '%s'", actionType)
}
