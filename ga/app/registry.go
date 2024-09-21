package app

import (
	"go-actions/ga/action"
	"reflect"
)

type ActionDefinitionRegistry struct {
	actionsByName map[string] *action.ActionDefinition
	actionsByType map[reflect.Type] *action.ActionDefinition
}

func NewActionDefinitionRegistry() *ActionDefinitionRegistry {
	return &ActionDefinitionRegistry{
		actionsByName: make(map[string]*action.ActionDefinition),
		actionsByType: make(map[reflect.Type]*action.ActionDefinition),
	}
}

func (adr *ActionDefinitionRegistry) acceptDefinition(def *action.ActionDefinition){
	adr.actionsByName[def.Name()] = def
	adr.actionsByType[def.Type()] = def
}