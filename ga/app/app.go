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

func RegisterAction[T action.GoAction](ctor action.GoActionConstructor[T]) func (*App) *action.ActionDefinition {
	return func(app *App) *action.ActionDefinition {
		def, _ := action.NewActionDefinition(ctor)
		return app.actionDefinitionRegistry.acceptDefinition(def)
	}
}

func (a *App) GetActionDef(action action.GoAction) (*action.ActionDefinition, error) {
	actionType := utils.GetValueType(reflect.TypeOf(action))
	return a.actionDefinitionRegistry.getDefinition(actionType)
}

func NewAction[T action.GoAction](a action.GoAction) func(*App) (*action.Action[T], error) {
	return func(app *App) (*action.Action[T], error) {
		def, err := app.GetActionDef(a)
		if err != nil {
			return nil, err
		}
		return action.NewAction[T](def), nil
	}
}
