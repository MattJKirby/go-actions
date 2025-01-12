package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app/registration"
	"go-actions/ga/utils"
	"reflect"
)

type App struct {
	ctx            context.Context
	actionRegistry *registration.ActionRegistry
}

func NewApp() *App {
	return &App{
		ctx:            context.Background(),
		actionRegistry: registration.NewActionRegistry(),
	}
}

func RegisterAction[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) func(*App) {
	return func(app *App) {
		registration.AcceptRegistration(reg)(app.actionRegistry)
	}
}

func GetActionRegistration[T action.GoAction, P action.GoActionProps]() func(*App) (*definition.ActionDefinition[T, P], error) {
	return func(app *App) (*definition.ActionDefinition[T, P], error) {
		action := new(T)
		actionType := utils.GetValueType(reflect.TypeOf(action))
		return registration.GetAction[T, P](actionType)(app.actionRegistry)
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*executable.Action[T, P], error) {
	return func(app *App) (*executable.Action[T, P], error) {
		reg, err := GetActionRegistration[T, P]()(app)
		if err != nil {
			return nil, err
		}

		action := executable.NewAction(*reg)
		if props != nil {
			action.PopulateActionInstance(props)
		}

		return action, nil
	}
}
