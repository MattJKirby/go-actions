package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
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

func RegisterAction[T action.GoAction](ctor action.GoActionConstructor[T]) func(*App) *definition.ActionDefinition {
	return func(app *App) *definition.ActionDefinition {
		def, _ := definition.NewActionDefinition(ctor)
		return app.actionDefinitionRegistry.acceptDefinition(def)
	}
}

func (a *App) GetActionDef(action action.GoAction) (*definition.ActionDefinition, error) {
	actionType := utils.GetValueType(reflect.TypeOf(action))
	return a.actionDefinitionRegistry.getDefinition(actionType)
}

func NewAction[T action.GoAction](a action.GoAction) func(*App) (*executable.Action[T], error) {
	return func(app *App) (*executable.Action[T], error) {
		def, err := app.GetActionDef(a)
		if err != nil {
			return nil, err
		}
		return executable.NewAction[T](def), nil
	}
}
