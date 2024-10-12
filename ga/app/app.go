package app

import (
	"context"
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

func (a *App) GetActionDef(action action.Action) (*action.ActionDefinition, error) {
	actionType := utils.GetValueType(reflect.TypeOf(action))
	return a.actionDefinitionRegistry.getDefinition(actionType)
}

func NewAction[T action.Action](a action.Action) func(*App) (*action.GoAction[T], error) {
	return func(app *App) (*action.GoAction[T], error) {
		def, err := app.GetActionDef(a)
		if err != nil {
			return nil, err
		}
		inst := action.NewActionInstance(def)
		return action.NewAction[T](def, inst), nil
	}
}
