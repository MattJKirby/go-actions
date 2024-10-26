package app

import (
	"fmt"
	"go-actions/ga/action/definition"
	"reflect"
)

type ActionDefinitionRegistry struct {
	actionsByName map[string]*definition.ActionDefinition
	actionsByType map[reflect.Type]*definition.ActionDefinition
}

func NewActionDefinitionRegistry() *ActionDefinitionRegistry {
	return &ActionDefinitionRegistry{
		actionsByName: make(map[string]*definition.ActionDefinition),
		actionsByType: make(map[reflect.Type]*definition.ActionDefinition),
	}
}

func (adr *ActionDefinitionRegistry) acceptDefinition(def *definition.ActionDefinition) *definition.ActionDefinition {
	adr.actionsByName[def.Name] = def
	adr.actionsByType[def.ActionType] = def
	return def
}

func (adr *ActionDefinitionRegistry) getDefinition(actionType reflect.Type) (*definition.ActionDefinition, error) {
	if def, exists := adr.actionsByType[actionType]; exists {
		return def, nil
	}
	return nil, fmt.Errorf("could not retrive action '%s'", actionType)
}
