package app

import (
	"context"
	"go-actions/ga/action"
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

func NewAction[T action.Action](actionType reflect.Type) func(*App) (*action.GoAction[T], error) {
	return func(a *App) (*action.GoAction[T], error) {
		def, err := a.GetActionDef(actionType)
		if err != nil {
			return nil, err
		}
		return action.NewAction[T](def), nil
	}
}
