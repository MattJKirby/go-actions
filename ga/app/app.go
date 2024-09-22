package app

import (
	"context"
	"go-actions/ga/action"
	"reflect"
)

type App struct {
	ctx context.Context
	actionDefinitionRegistry *ActionDefinitionRegistry
}

func NewApp() *App{
	return &App{
		ctx: context.Background(),
		actionDefinitionRegistry: NewActionDefinitionRegistry(),
	}
}

func (a *App) RegisterActionDef(actionDef *action.ActionDefinition) *action.ActionDefinition {
	return a.actionDefinitionRegistry.acceptDefinition(actionDef)
}

func (a *App) GetAction(v reflect.Value) (*action.ActionDefinition, error) {
	return a.actionDefinitionRegistry.getDefinition(v)
}