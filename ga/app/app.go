package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app/registration"
	"go-actions/ga/utils"
	"reflect"
)

type App struct {
	ctx                      context.Context
	actionRegistry *registration.ActionRegistry
}

func NewApp() *App {
	return &App{
		ctx:                      context.Background(),
		actionRegistry: registration.NewActionRegistry(),
	}
}

func RegisterAction[T action.GoAction](reg *action.GoActionRegistration[T]) func(*App) *registration.RegisteredAction[T] {
	return func(app *App) *registration.RegisteredAction[T] {
		action, _ := registration.NewRegisteredAction(reg)
		return registration.AcceptAction(action)(app.actionRegistry)
	}
}

func GetActionRegistration[T action.GoAction](action action.GoAction) func(*App) (*registration.RegisteredAction[T], error) {
	return func(app *App) (*registration.RegisteredAction[T], error) {
		actionType := utils.GetValueType(reflect.TypeOf(action))
		return registration.GetAction[T](actionType)(app.actionRegistry)
	}
}

func GetAction[T action.GoAction](a action.GoAction) func(*App) (*executable.Action[T], error) {
	return func(app *App) (*executable.Action[T], error) {
		def, err := GetActionRegistration[T](a)(app)
		if err != nil {
			return nil, err
		}
		return executable.NewAction[T](def.ActionDefinition), nil
	}
}
