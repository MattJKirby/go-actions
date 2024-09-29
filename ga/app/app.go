package app

import (
	"context"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/utils"
	"reflect"
)

type App struct {
	ctx                      context.Context
	actionDefinitionRegistry *ActionDefinitionRegistry
}

func NewApp() *App {
	return &App{
		ctx:                      context.Background(),
		actionDefinitionRegistry: NewActionDefinitionRegistry(),
	}
}

func (a *App) RegisterActionDef(actionDef *action.ActionDefinition) *action.ActionDefinition {
	return a.actionDefinitionRegistry.acceptDefinition(actionDef)
}

func (a *App) GetActionDef(actionType reflect.Type) (*action.ActionDefinition, error) {
	return a.actionDefinitionRegistry.getDefinition(actionType)
}

func (a *App) GetAction(actionType reflect.Type) (*action.Action){
	def, err := a.GetActionDef(actionType)
	if err != nil {
		panic(fmt.Sprintf("could not get action definition '%s'", utils.TypeName(actionType)))
	}
	
	return action.NewAction(def)
}
