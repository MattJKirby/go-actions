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

func (a *App) GetActionDef(actionType reflect.Type) (*action.ActionDefinition, error) {
	return a.actionDefinitionRegistry.getDefinition(actionType)
}

func (a *App) getActionFromType(action action.Action) reflect.Type {
	return utils.GetValueType(reflect.TypeOf(action))
}

func NewAction[T action.Action](a action.Action) func(*App) (*action.GoAction[T], error) {
	return func(app *App) (*action.GoAction[T], error) {
		actionType := app.getActionFromType(a)
		def, err := app.GetActionDef(actionType)
		if err != nil {
			return nil, err
		}
		inst := action.NewActionInstance(def)
		return action.NewAction[T](def, inst), nil
	}
}
