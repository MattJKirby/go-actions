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

func DefineAction[T action.GoAction](ctor action.GoActionConstructor[T]) func(*App) *definition.ActionDefinition {
	return func(app *App) *definition.ActionDefinition {
		def, _ := definition.NewActionDefinition(ctor)
		return app.actionDefinitionRegistry.acceptDefinition(def)
	}
}

func GetActionDefinition(action action.GoAction) func(*App) (*definition.ActionDefinition, error) {
	return func(app *App) (*definition.ActionDefinition, error) {
		actionType := utils.GetValueType(reflect.TypeOf(action))
		return app.actionDefinitionRegistry.getDefinition(actionType)
	}
}

func GetAction[T action.GoAction](a action.GoAction) func(*App) (*executable.Action[T], error) {
	return func(app *App) (*executable.Action[T], error) {
		def, err := GetActionDefinition(a)(app)
		if err != nil {
			return nil, err
		}
		return executable.NewAction[T](def), nil
	}
}
